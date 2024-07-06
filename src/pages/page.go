package pages

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var AllPages Pages

type Pages []Page

type Page struct {
	Path   string
	Render func(c echo.Context, w io.Writer) error
}

func New(p Page) {
	fmt.Println("registering page: " + p.Path)
	if AllPages == nil {
		AllPages = make(Pages, 0)
	}
	AllPages = append(AllPages, p)
}

func InitializePages(ps Pages, e *echo.Echo) *echo.Echo {
	for _, p := range ps {
		e.Add(
			http.MethodGet,
			p.Path,
			func(c echo.Context) error {
				sb := strings.Builder{}
				err := p.Render(c, &sb)
				if err != nil {
					return err
				}
				return c.HTML(200, sb.String())
			},
		)
	}

	return e
}
