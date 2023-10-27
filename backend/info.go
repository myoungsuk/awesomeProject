package backend

// DB 정보나 구조체를 정의
var db1 = dbInfo{"root", "myoung1249!", "localhost:3306", "mysql", "board"}
var id string = "kang212"
var pw string = "asdf1234"
var header string = "frontend/header.html"
var footer string = "frontend/footer.html"

/*
데이터베이스 접속에 필요한 정보를 구조체로 정의
*/

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

type Article struct {
	ID         int64
	CreatedAt  string
	CreatedBy  string
	ModifiedAt string
	ModifiedBy string
	Content    string
	Title      string
	UserID     string
}

func static(handle string) string {
	var reshandle string

	switch handle {
	case "n_main":
		reshandle = "1"
	case "n_content":
		reshandle = "2"
	case "n_write":
		reshandle = "3"
	}

	return reshandle
}
