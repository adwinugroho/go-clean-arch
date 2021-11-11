package controller

import (
	"bytes"
	"encoding/json"
	reqModel "go-clean-arch/models/request"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

// validation, ex:user session etc

func (t *OrderRoute) middlewareOrder(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t.user = new(reqModel.User)
		body := new(reqModel.General)
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		}
		json.Unmarshal(bodyBytes, &body)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		if body.User == nil {
			// if body.Child == "" {
			// 	body.Child = "0"
			// }
			var token = c.Request().Header.Get("authorization")
			var customID = c.Request().Header.Get("x-consumer-custom-id")
			c.Set("customId", customID)

			t.user.Child = "0"
			t.user.LangCode = body.LangCode
			t.user.Token = token
			t.user.ID = customID
		}
		return next(c)
	}
}
