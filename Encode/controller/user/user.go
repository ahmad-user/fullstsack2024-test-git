package user

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"Encode/helper"
	model "Encode/model/user"
)

type ClientController struct {
	Model model.MyClientModel
}

func (cc *ClientController) AddClient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.MyClient
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupported") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim tidak sesuai", nil))
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirim kurang sesuai", nil))
		}

		err = cc.Model.AddClient(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan pada sistem", nil))
		}

		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "client berhasil ditambahkan", input))
	}
}
