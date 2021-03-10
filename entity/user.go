package entity

type (
	User struct {
		ID    string   `json:"_key"`
		Data  DataUser `json:"data"`
		Audit Audit    `json:"audit"`
	}

	DataUser struct {
		Email    string `json:"email"`
		Gender   string `json:"gender"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	BindVarsUser struct {
		ID     string
		Limit  int    `json:"limit"`
		Offset int    `json:"offset"`
		Page   int    `json:"page"`
		Filter Filter `json:"filter"`
		Search Search `json:"search,omitempty"`
	}

	Filter struct {
		Gender string `json:"gender,omitempty"`
	}

	Search struct {
		Text string `json:"text"`
	}
)
