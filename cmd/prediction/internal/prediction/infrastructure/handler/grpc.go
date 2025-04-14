package handler

import (
	"context"
	"github.com/Errera1/prediction/internal/prediction/service"
	diabetesproto "github.com/Errera1/prediction/internal/protogen"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

type PredictionGrpcHandler struct {
	predictionService *service.PredictionService
	validate          *validator.Validate
	diabetesproto.UnimplementedPredictionServiceServer
}

func (p PredictionGrpcHandler) SavePrediction(ctx context.Context, request *diabetesproto.SavePredictionRequest) (*diabetesproto.PredictionResponse, error) {
	parsedReq := &SavePredictionValidator{
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
	}
	err := p.validate.Struct(parsedReq)
	if err != nil {
		return nil, err
	}

	user := ctx.Value("user").(*diabetesproto.AuthResponse)

	return p.predictionService.SavePrediction(ctx, &diabetesproto.SavePredictionRequest{
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
		UserId:        &user.Id,
	})
}

func (p PredictionGrpcHandler) GetAllPredictionsByUserId(ctx context.Context, request *diabetesproto.GetAllPredictionsByUserIdRequest) (*diabetesproto.GetAllPredictionsByUserIdResponse, error) {
	return p.predictionService.GetAllPredictionsByUserId(ctx, request)
}

//func (p PredictionGrpcHandler) GetAllPredictionsByUserEmail(ctx context.Context, request *diabetesproto.GetAllPredictionsByUserEmailRequest) (*diabetesproto.GetAllPredictionsByUserIdResponse, error) {
//	return p.predictionService.SavePrediction(ctx, request)
//}

func (p PredictionGrpcHandler) GetPredictionById(ctx context.Context, request *diabetesproto.GetPredictionByIdRequest) (*diabetesproto.PredictionResponse, error) {
	return p.predictionService.GetPredictionById(ctx, request)
}

func (p PredictionGrpcHandler) DeletePredictionById(ctx context.Context, request *diabetesproto.DeletePredictionByIdRequest) (*diabetesproto.DeletePredictionByIdResponse, error) {
	return p.predictionService.DeletePredictionById(ctx, request)
}

func New(grpc *grpc.Server, predictionService *service.PredictionService) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	gRPCHandler := &PredictionGrpcHandler{
		predictionService: predictionService,
		validate:          validate,
	}
	diabetesproto.RegisterPredictionServiceServer(grpc, gRPCHandler)
}
