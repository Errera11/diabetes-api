package db

import (
	"context"
	"fmt"
	"github.com/Errera1/prediction/internal/prediction/domain"
	"github.com/jackc/pgx/v5"
)

type PredictionRepo struct {
	conn *pgx.Conn
}

func (p PredictionRepo) SavePrediction(ctx context.Context, data *domain.Prediction) (*domain.PredictionRecord, error) {
	createPredictionQuery := `INSERT INTO prediction (chol_level, diff_walk, phys_health, birthdate,
							  blood_pressure, weight, height, heart_disease, gen_health, 
							  phys_activity, result) 
							  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
							  RETURNING id, chol_level, diff_walk, phys_health, birthdate, blood_pressure, weight, height, 
							  heart_disease, gen_health, phys_activity, result, 
							  created_at::text AS created_at`

	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return &domain.PredictionRecord{}, err
	}
	defer tx.Rollback(ctx)

	var predictionRecord domain.PredictionRecord
	err = tx.QueryRow(ctx, createPredictionQuery,
		data.CholLevel,
		data.DiffWalk,
		data.PhysHealth,
		data.Birthdate,
		data.BloodPressure,
		data.Weight,
		data.Height,
		data.HeartDisease,
		data.GenHealth,
		data.PhysActivity,
		data.Result,
	).Scan(&predictionRecord.Id,
		&predictionRecord.CholLevel,
		&predictionRecord.DiffWalk,
		&predictionRecord.PhysHealth,
		&predictionRecord.Birthdate,
		&predictionRecord.BloodPressure,
		&predictionRecord.Weight,
		&predictionRecord.Height,
		&predictionRecord.HeartDisease,
		&predictionRecord.GenHealth,
		&predictionRecord.PhysActivity,
		&predictionRecord.Result,
		&predictionRecord.CreatedAt)

	if err != nil {
		return &domain.PredictionRecord{}, err
	}

	createUserPredictionQuery := `INSERT INTO user_prediction (user_id, prediction_id) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, createUserPredictionQuery, data.UserId, predictionRecord.Id)

	if err != nil {
		return &domain.PredictionRecord{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return &domain.PredictionRecord{}, err
	}

	return &predictionRecord, nil
}

func (p PredictionRepo) GetAllPredictionsByUserId(ctx context.Context, userId int32) ([]*domain.PredictionRecord, error) {
	query := `SELECT DISTINCT p.id, chol_level, diff_walk, phys_health, birthdate, blood_pressure, weight, height, 
							  heart_disease, gen_health, phys_activity, result, 
							  created_at::text AS created_at, up.user_id FROM user_prediction up 
							  INNER JOIN prediction p ON up.prediction_id=p.id WHERE up.user_id=$1;`

	rows, err := p.conn.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("could not fetch predictions: %v", err)
	}
	defer rows.Close()

	var predictions []*domain.PredictionRecord

	for rows.Next() {
		var pred domain.PredictionRecord
		err := rows.Scan(&pred.Id, &pred.CholLevel, &pred.DiffWalk, &pred.PhysHealth, &pred.Birthdate,
			&pred.BloodPressure, &pred.Weight, &pred.Height, &pred.HeartDisease, &pred.GenHealth, &pred.PhysActivity, &pred.Result,
			&pred.CreatedAt, &pred.UserId)
		if err != nil {
			return nil, fmt.Errorf("could not scan prediction: %v", err)
		}
		predictions = append(predictions, &pred)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over rows: %v", err)
	}

	return predictions, nil
}

//func (p PredictionRepo) GetAllPredictionsByUserEmail(ctx context.Context, email string) ([]*domain.PredictionRecord, error) {
//	query := `SELECT p.* up.user_id FROM (SELECT * FROM pr) p INNER JOIN user_prediction up ON p.id=up.prediction_id;`
//
//	rows, err := p.conn.Query(ctx, query)
//	if err != nil {
//		return nil, fmt.Errorf("could not fetch predictions: %v", err)
//	}
//	defer rows.Close()
//
//	var predictions []*domain.PredictionRecord
//
//	for rows.Next() {
//		var pred domain.PredictionRecord
//		err := rows.Scan(&pred.Id, &pred.CholLevel, &pred.DiffWalk, &pred.PhysHealth, &pred.Birthdate,
//			&pred.BloodPressure, &pred.Weight, &pred.Height, &pred.HeartDisease, &pred.GenHealth, &pred.PhysActivity, &pred.Result,
//			&pred.CreatedAt, &pred.UserId)
//		if err != nil {
//			return nil, fmt.Errorf("could not scan prediction: %v", err)
//		}
//		predictions = append(predictions, &pred)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("could not iterate over rows: %v", err)
//	}
//
//	return predictions, nil
//}

func (p PredictionRepo) GetPredictionById(ctx context.Context, id int32) (*domain.PredictionRecord, error) {
	query := `SELECT * FROM prediction WHERE id=$1`
	var pred domain.PredictionRecord

	err := p.conn.QueryRow(ctx, query, id).Scan(&pred.Id, &pred.CholLevel, &pred.DiffWalk, &pred.PhysHealth, &pred.Birthdate,
		&pred.BloodPressure, &pred.Weight, &pred.Height, &pred.HeartDisease, &pred.GenHealth, &pred.PhysActivity, &pred.Result,
		&pred.CreatedAt)
	if err != nil {
		return &pred, fmt.Errorf("could not fetch predictions by id: %v", err)
	}

	return &pred, nil
}

func (p PredictionRepo) DeletePredictionById(ctx context.Context, id int32) (*domain.PredictionRecord, error) {
	query := `DELETE FROM prediction WHERE id=$1 RETURNING id, chol_level, diff_walk, phys_health, birthdate, blood_pressure, weight, height, 
							  heart_disease, gen_health, phys_activity, result, 
							  created_at::text AS created_at;`

	var pred domain.PredictionRecord

	err := p.conn.QueryRow(ctx, query, id).Scan(&pred.Id, &pred.CholLevel, &pred.DiffWalk, &pred.PhysHealth, &pred.Birthdate,
		&pred.BloodPressure, &pred.Weight, &pred.Height, &pred.HeartDisease, &pred.GenHealth, &pred.PhysActivity, &pred.Result,
		&pred.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not fetch predictions by id: %v", err)
	}

	return &pred, nil
}

func New(conn *pgx.Conn) domain.PredictionRepo {
	return &PredictionRepo{
		conn: conn,
	}
}
