package response

type (
	GeneralResponse struct {
		Code    int         `json:"code"`
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message,omitempty"`
		Status  bool        `json:"status"`
	}

	GetByIDOrderResponse struct {
		ID     string      `json:"id"`
		Number string      `json:"number"`
		Menus  interface{} `json:"menus"`
	}
)

func Success(code int, data interface{}) *GeneralResponse {
	return &GeneralResponse{Status: true, Code: code, Data: data}
}

func Error(code int, message string) *GeneralResponse {
	return &GeneralResponse{Status: false, Code: code, Message: message}
}
