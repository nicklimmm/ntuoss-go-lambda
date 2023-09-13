package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nicklimmm/ntuoss-go-lambda/fx"
)

// registerFx registers routes under "/fx"
func registerFx(e *echo.Echo) {
	grp := e.Group("/fx")
	grp.GET("/rate", getRate)
	grp.POST("/convert", convert)
}

type getRateReq struct {
	From string `query:"from" validate:"required"`
	To   string `query:"to" validate:"required"`
}

type getRateResp struct {
	Rate float64 `json:"rate"`
}

func getRate(c echo.Context) error {
	req := new(getRateReq)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	rate, err := fx.GetRate(req.From, req.To)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &getRateResp{rate})
}

type convertReq struct {
	From   string  `json:"from" validate:"required"`
	To     string  `json:"to" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type convertResp struct {
	Amount float64 `json:"amount"`
}

func convert(c echo.Context) error {
	req := new(convertReq)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	amount, err := fx.Convert(req.From, req.To, req.Amount)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &convertResp{amount})
}
