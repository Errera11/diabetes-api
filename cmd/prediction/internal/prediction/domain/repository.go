package domain

import "context"

type PredictionRepo interface {
	SavePrediction(ctx context.Context, data *Prediction) (*PredictionRecord, error)
	GetAllPredictionsByUserId(ctx context.Context, userId int32) ([]*PredictionRecord, error)
	//GetAllPredictionsByUserEmail(ctx context.Context, email string) ([]*PredictionRecord, error)
	GetPredictionById(ctx context.Context, id int32) (*PredictionRecord, error)
	DeletePredictionById(ctx context.Context, id int32) (*PredictionRecord, error)
}
