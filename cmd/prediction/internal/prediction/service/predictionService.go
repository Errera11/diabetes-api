package service

import (
	"context"
	"fmt"
	"github.com/Errera1/prediction/internal/prediction/domain"
	diabetesproto "github.com/Errera1/prediction/internal/protogen"
	"github.com/Errera1/prediction/internal/utils/mappers"
)

type PredictionService struct {
	predictionRepo domain.PredictionRepo
	apiService     *ApiService
}

func (p PredictionService) SavePrediction(ctx context.Context, request *diabetesproto.SavePredictionRequest) (*diabetesproto.PredictionResponse, error) {
	payload := mappers.MapToDomain(request, *request.UserId)
	apiPrediction, err := p.apiService.MakePrediction(ctx, payload)

	if err != nil {
		return nil, fmt.Errorf("Error MakePrediction: %w", err)
	}

	if request.UserId == nil {
		return &diabetesproto.PredictionResponse{
			CholLevel:     request.CholLevel,
			DiffWalk:      request.DiffWalk,
			PhysHealth:    request.PhysHealth,
			Birthdate:     request.Birthdate,
			BloodPressure: request.BloodPressure,
			Weight:        request.Weight,
			Height:        request.Height,
			HeartDisease:  request.HeartDisease,
			GenHealth:     request.GenHealth,
			PhysActivity:  request.PhysActivity,
			Result:        apiPrediction.Prediction[0],
			UserId:        -1,
			Id:            -1,
			CreatedAt:     "-1",
		}, nil
	}

	withResultPayload := mappers.MapToDomain(&diabetesproto.SavePredictionRequest{
		CholLevel:     payload.CholLevel,
		DiffWalk:      payload.DiffWalk,
		PhysHealth:    payload.PhysHealth,
		Birthdate:     payload.Birthdate,
		BloodPressure: payload.BloodPressure,
		Weight:        payload.Weight,
		Height:        payload.Height,
		HeartDisease:  payload.HeartDisease,
		GenHealth:     payload.GenHealth,
		PhysActivity:  payload.PhysActivity,
		Result:        apiPrediction.Prediction[0],
	}, *request.UserId)
	newRecord, err := p.predictionRepo.SavePrediction(ctx, withResultPayload)
	if err != nil {
		return nil, err
	}

	return mappers.MapToResponse(newRecord, *request.UserId), nil
}

func (p PredictionService) GetAllPredictionsByUserId(ctx context.Context, request *diabetesproto.GetAllPredictionsByUserIdRequest) (*diabetesproto.GetAllPredictionsByUserIdResponse, error) {
	records, err := p.predictionRepo.GetAllPredictionsByUserId(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	var response []*diabetesproto.PredictionResponse
	for _, record := range records {
		serializedResponse := mappers.MapToResponse(record, request.UserId)
		response = append(response, serializedResponse)
	}

	return &diabetesproto.GetAllPredictionsByUserIdResponse{Predictions: response}, nil
}

//func (p PredictionService) GetAllPredictionsByUserEmail(ctx context.Context, request *diabetesproto.GetAllPredictionsByUserEmailRequest) (*diabetesproto.GetAllPredictionsByUserIdResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (p PredictionService) GetPredictionById(ctx context.Context, request *diabetesproto.GetPredictionByIdRequest) (*diabetesproto.PredictionResponse, error) {
	record, err := p.predictionRepo.GetPredictionById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return mappers.MapToResponse(record, record.UserId), nil
}

func (p PredictionService) DeletePredictionById(ctx context.Context, request *diabetesproto.DeletePredictionByIdRequest) (*diabetesproto.DeletePredictionByIdResponse, error) {
	record, err := p.predictionRepo.DeletePredictionById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &diabetesproto.DeletePredictionByIdResponse{
		Id: record.Id,
	}, nil
}

func New(predictionRepo domain.PredictionRepo, service *ApiService) *PredictionService {
	return &PredictionService{
		predictionRepo: predictionRepo,
		apiService:     service,
	}
}
