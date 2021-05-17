package request

type (
	CreateOrderLRequest struct {
		Number string `json:"number" validate:"required"`
		Menus  []Menu `json:"menus"`
	}

	GetByIDorderRequest struct {
		ID string `json:"id" validate:"required"`
	}

	Menu struct {
		Name                  string `json:"name"`
		Qty                   int    `json:"qty"`
		AdditionalInformation string `json:"additionalInformation"`
	}
)
