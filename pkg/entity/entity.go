package entity

// Results global api return value
type Results struct {
	Message interface{} `json:"message"`
	Error   string      `json:"error"`
}
