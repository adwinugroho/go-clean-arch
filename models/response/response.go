package response

type (
	GetByIDOrderResponse struct {
		ID     string      `json:"id"`
		Number string      `json:"number"`
		Menus  interface{} `json:"menus"`
	}

	GeneralResponse struct {
		Status bool        `json:"status"`
		Code   int         `json:"code"`
		Data   interface{} `json:"data,omitempty"`
	}

	SuccesListResponse struct {
		Status bool        `json:"status"`
		Code   int         `json:"code"`
		Data   interface{} `json:"data,omitempty"`
		Total  int         `json:"total"`
	}
)

func Success(code int) *GeneralResponse {
	return &GeneralResponse{Status: true, Code: code}
}

func (r *GeneralResponse) SetData(data interface{}) *GeneralResponse {
	r.Data = data
	return r
}

func SuccessList(code int) *SuccesListResponse {
	return &SuccesListResponse{Status: true, Code: code}
}

func (r *SuccesListResponse) SetData(data interface{}) *SuccesListResponse {
	r.Data = data
	return r
}

func (r *SuccesListResponse) SetTotal(total int) *SuccesListResponse {
	r.Total = total
	return r
}
