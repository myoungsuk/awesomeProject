package backend

import (
	"github.com/labstack/echo"
	"net/http"
)

func EchoRequestHandler(c echo.Context) error {
	reshandler := c.FormValue("Handler")
	shandle := static(reshandler)

	switch shandle {
	case "1":
		return GetArticleList(c)
		//case "2":
		//	return EchoNoticeboardContentView(c)
		//case "3":
		//	return EchoNoticeboardWriteView(c)
	}

	return c.Render(http.StatusOK, "error.html", 0)
}
