package mappers

import (
	"github.com/Errera1/prediction/internal/prediction/domain"
	diabetesproto "github.com/Errera1/prediction/internal/protogen"
)

func MapToDomain(request *diabetesproto.SavePredictionRequest, userId int32) *domain.Prediction {
	return &domain.Prediction{
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
		Result:        request.Result,
		UserId:        userId,
	}
}

func MapToResponse(request *domain.PredictionRecord, userId int32) *diabetesproto.PredictionResponse {
	return &diabetesproto.PredictionResponse{
		Id:            request.Id,
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
		Result:        request.Result,
		CreatedAt:     request.CreatedAt,
		UserId:        userId,
	}
}
