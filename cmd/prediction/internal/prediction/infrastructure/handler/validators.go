package handler

type SavePredictionValidator struct {
	CholLevel     float32 `validate:"required"`
	DiffWalk      bool    `validate:"boolean"`
	PhysHealth    int32   `validate:"required"`
	Birthdate     string  `validate:"required"`
	BloodPressure float32 `validate:"required"`
	Weight        float32 `validate:"required"`
	Height        float32 `validate:"required"`
	HeartDisease  bool    `validate:"boolean"`
	GenHealth     int32   `validate:"required"`
	PhysActivity  bool    `validate:"boolean"`
}
