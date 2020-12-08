package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/miraclew/tao/pkg/ce"
)

func JSON(c echo.Context, res interface{}, err error) error {
	if err != nil {
		if e, ok := err.(*ce.Error); ok {
			if e.Code >= 600 {
				return c.JSON(200, echo.Map{"message": e.Message, "code": e.Code})
			}
			return c.JSON(e.Code, echo.Map{"message": e.Message, "code": e.Code})
		}
		return err
	}

	var response = echo.Map{
		"code":    0,
		"message": "",
		"data":    res,
	}

	return c.JSON(200, response)
}
