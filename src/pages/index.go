package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/sattfatt/website/src/templs"
	"io"
)

func init() {
	New(Page{
		Path: "/",
		Render: func(c echo.Context, w io.Writer) error {
			return templs.Index().Render(c.Request().Context(), w)
		},
	})
}
