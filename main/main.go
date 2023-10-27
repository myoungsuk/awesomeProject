package main

import (
	"awesomeProject/backend"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// template 기능을 사용하기 위한 구조체
type TemplateRenderer struct {
	templates *template.Template
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"subtract": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
}

func GetTempFilesFromFolders(folders []string) []string {
	var filepaths []string
	for _, folder := range folders {
		files, _ := ioutil.ReadDir(folder)

		for _, file := range files {
			if strings.Contains(file.Name(), ".html") {
				filepaths = append(filepaths, folder+file.Name())
			}
		}
	}
	return filepaths
}

// HTML 템플릿 렌더링
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	dirs := []string{
		"./frontend/",
		"./frontend/static/",
		"./frontend/templates/",
		"./frontend/templates/articles/"}

	tempfiles := GetTempFilesFromFolders(dirs)

	//load template from folder
	t := &TemplateRenderer{
		templates: template.Must(template.ParseFiles(tempfiles...)),
	}

	//fs는 정적 파일을 제공하는 HTTP 파일 서버이다.
	fs := http.FileServer(http.Dir("./frontend/static"))

	//새로운 Echo인스턴스 생성
	e := echo.New()

	//e.AutoTLSManager.Cache = autocert.DirCache("/static/ssl/")

	//now := time.Now()
	//custom := now.Format("2023-10-23")
	//fileName := custom + "_log.txt"
	//
	//f, file := os.OpenFile("./logs/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if file != nil {
	//	panic(file)
	//}
	//defer f.Close()
	//
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: `{"time":"${time_rfc3339}", "remote_ip":"${remote_ip}", ` +
	//		`"host":"${host}", "method":"${method}", "uri":"${uri}", "user_agent":"${user_agent}",` +
	//		`"status":${status}, ` + "\n",
	//	Output: f,
	//}))
	//e.Use(middleware.Recover())

	e.Static("/static", "public")
	e.Renderer = t

	//라우터 설정
	//각 경로에 대한 핸들러 함수 설정
	e.GET("/", backend.GetArticleList)
	e.GET("/articleDetail", getArticleDetail) //상세페이지

	e.GET("/articleUpdate", getArticleUpdate)   //수정페이지
	e.POST("/articleUpdate", postArticleUpdate) //수정하기

	e.GET("/articleRegisteration", getArticleRegisteration)   //게시물 등록페이지
	e.POST("/articleRegisteration", postArticleRegisteration) //게시물 등록

	e.POST("/articleDetail", postArticleDelete) //삭제하기

	e.GET("/sign-up", getSignUp)   //회원가입페이지
	e.POST("/sign-up", postSignUp) //회원가입

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))
	e.POST("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))

	e.Logger.Fatal(e.Start(":1323"))
}

func getArticleDetail(c echo.Context) error {
	return c.Render(http.StatusOK, "detail.html", "")
}

func getArticleUpdate(c echo.Context) error {
	return c.Render(http.StatusOK, "form.html", "")
}

func postArticleUpdate(c echo.Context) error {
	return c.Render(http.StatusOK, "form.html", "")
}

func getArticleRegisteration(c echo.Context) error {
	return c.Render(http.StatusOK, "form.html", "")
}

func postArticleRegisteration(c echo.Context) error {
	return c.Render(http.StatusOK, "form.html", "")
}

func postArticleDelete(c echo.Context) error {
	return c.Render(http.StatusOK, "detail.html", "")
}

func getSignUp(c echo.Context) error {
	return c.Render(http.StatusOK, "sign-up.html", "")
}

func postSignUp(c echo.Context) error {
	return c.Render(http.StatusOK, "sign-up.html", "")
}
