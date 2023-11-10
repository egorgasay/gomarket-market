package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(c echo.Context, err error) error {
	var unmarshalTypeError *json.UnmarshalTypeError
	if err != nil {
		if errors.Is(err, ErrEnabledData) {
			err := fmt.Sprintf("enabled data %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		if errors.As(err, &unmarshalTypeError) {
			err := fmt.Sprintf("bad json %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	c.Response().Status = http.StatusOK
	return nil
}
