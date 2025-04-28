package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Errera1/prediction/internal/prediction/domain"
	"io"
	"net/http"
)

type PredictionResponse struct {
	Prediction [][]float32 `json:"prediction"`
}

type ApiService struct {
	ApiAddr *string
}

type PredictionRequest struct {
	Data *domain.Prediction `json:"data"`
}

func NewApiService(addr *string) *ApiService {
	return &ApiService{ApiAddr: addr}
}

func (a *ApiService) MakePrediction(ctx context.Context, payload *domain.Prediction) (*PredictionResponse, error) {
	apiPayload := &PredictionRequest{
		Data: payload,
	}

	serializedPayload, err := json.Marshal(apiPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize payload: %w", err)
	}

	fmt.Println("a.ApiAddr", *a.ApiAddr)
	fmt.Printf("Make req for prediction to: %d", *a.ApiAddr+"/predict/")
	resp, err := http.Post(*a.ApiAddr+"/predict/", "application/json", bytes.NewBuffer(serializedPayload))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed make req for prediction: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed reading body: %v", err)
	}

	var unmarshalledBody PredictionResponse
	if err = json.Unmarshal(body, &unmarshalledBody); err != nil {
		return nil, fmt.Errorf("failed unmarshall body: %w", err)
	}

	return &unmarshalledBody, nil
}
