package entity

type (
	// create struct modelling data in DB
	Order struct {
		ID    string `json:"_key,omitempty"`
		Data  *Data  `json:"data"`
		Owner string `json:"owner"`
		Audit *Audit `json:"audit,omitempty"`
	}

	// add struct Data
	Data struct {
		Number string `json:"number"`
		Menus  []Menu `json:"menus"`
	}

	Menu struct {
		Name                  string `json:"name"`
		Qty                   int    `json:"qty"`
		AdditionalInformation string `json:"additionalInformation"`
	}
)
