package backend

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

//func GetArticleList(c echo.Context) error {
//	c.Request().ParseForm()
//
//	respage := c.FormValue("Page") //페이지 번호 값을 가져온다
//
//	if respage != "" {
//		intrespage, err := strconv.Atoi(respage)
//		if err != nil {
//			log.Fatal(err)
//		}
//		intrespage = (intrespage * 10) - 10 //페이지 번호계산
//		page := strconv.Itoa(intrespage)    //페이지 번호를 문자열로 변환
//
//		//var noticeviewstring = "SELECT title, created_by, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s')  FROM article WHERE id > 0"
//
//		var noticeviewstring = "SELECT title, created_by, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s'), (SELECT MAX(id) FROM article) FROM article WHERE id > 0 ORDER BY id DESC LIMIT 10 OFFSET ?"
//		result, _ := ArticleSelectQuery(db1, noticeviewstring, "index", page)
//
//		return c.Render(http.StatusOK, "index.html", result)
//	}
//	return c.Render(http.StatusOK, "error.html", 0)
//}

func GetArticleList(c echo.Context) error {
	c.Request().ParseForm()

	respage := c.FormValue("Page") //페이지 번호 값을 가져온다

	pageNumber := 1 // 기본 페이지 번호 설정
	if respage != "" {
		var err error
		pageNumber, err = strconv.Atoi(respage)
		if err != nil {
			log.Fatal(err)
		}
		// 이 부분은 더 이상 필요하지 않습니다.
		// intrespage = (intrespage * 10) - 10 //페이지 번호계산
		// page := strconv.Itoa(intrespage)    //페이지 번호를 문자열로 변환
	}

	var noticeviewstring = "SELECT title, created_by, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s'), (SELECT MAX(id) FROM article) FROM article WHERE id > 0 ORDER BY id DESC"
	result, err := ArticleSelectQuery(db1, noticeviewstring, "index", pageNumber)
	if err != nil {
		// 데이터베이스 에러 처리를 추가하였습니다.
		log.Println("Database Error:", err)
		return c.Render(http.StatusInternalServerError, "error.html", err.Error())
	}

	return c.Render(http.StatusOK, "index.html", result)
}
