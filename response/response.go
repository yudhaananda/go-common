package response

type Response struct {
	Meta   Meta        `json:"meta"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}, errors interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta:   meta,
		Data:   data,
		Errors: errors,
	}
	return jsonResponse
}
