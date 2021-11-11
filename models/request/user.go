package request

type (
	User struct {
		ID       string                 `json:"_key"`
		Email    string                 `json:"email"`
		Child    string                 `json:"child"`
		ChildKey []string               `json:"childKey"`
		Device   string                 `json:"device"`
		Fbid     string                 `json:"fbid"`
		BCAdd    string                 `json:"bcadd"`
		CTAdd    string                 `json:"ctadd"`
		IsActive int                    `json:"isActive"`
		LangCode string                 `json:"langCode"`
		Phone    string                 `json:"phone"`
		RespUser map[string]interface{} `json:"user,omitempty"`
		Token    string                 `json:"token,omitempty"`
		FullName string                 `json:"fname"`
	}
)
