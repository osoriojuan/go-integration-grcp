package structs

type DefaultRequest struct {
	Operation string `json:"operation"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
}

type DefaultResponse struct {
	Result  string `json:"operation"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
