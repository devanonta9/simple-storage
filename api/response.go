package api

// struct for response
type Response struct {
	HttpCode int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}
