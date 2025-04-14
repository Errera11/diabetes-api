package domain

type Prediction struct {
	CholLevel     int32     `json:"cholLevel"`
	DiffWalk      bool      `json:"diffWalk"`
	PhysHealth    int32     `json:"physHealth"`
	Birthdate     string    `json:"birthdate"`
	BloodPressure float32   `json:"bloodPressure"`
	Weight        float32   `json:"weight"`
	Height        float32   `json:"height"`
	HeartDisease  bool      `json:"heartDisease"`
	GenHealth     int32     `json:"genHealth"`
	PhysActivity  bool      `json:"physActivity"`
	Result        []float32 `json:"result"`
	UserId        int32     `json:"userId"`
}

type PredictionRecord struct {
	Prediction
	Id        int32
	CreatedAt string
}
