package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// InsertQuery is mysql Insert Query
func InsertQuery(db dbInfo, query string) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	result, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	nRow, err := result.RowsAffected()
	fmt.Println("insert count : ", nRow)
	conn.Close()
}

func UpdateQuery(db dbInfo, query string) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec(query)
	conn.Close()
}

func ArticleSelectQuery(db dbInfo, query string, sel string, pageNumber int) ([]Article, error) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 페이지 크기를 정의 (예: 10개의 게시물)
	pageSize := 10
	offset := (pageNumber - 1) * pageSize

	// 쿼리에 LIMIT과 OFFSET 추가
	// 주의: 직접 문자열 연결을 통한 쿼리 구성은 SQL 인젝션의 위험이 있으므로, 대안으로 쿼리 인자를 사용하는 것이 좋습니다.
	queryWithPagination := query + " LIMIT ? OFFSET ?"

	rows, err := conn.Query(queryWithPagination, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		var maxID int
		if sel == "index" {
			err := rows.Scan(&article.Title, &article.CreatedBy, &article.CreatedAt, &maxID)
			if err != nil {
				return nil, err
			}
		}

		articles = append(articles, article)
	}

	// 에러 핸들링: rows.Next() 이후에 발생할 수 있는 에러들
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

/* 게시글의 첫 화면, 선택된 게시글 내용을 출력하기 위한 select 문*/

//func ArticleSelectQuery(db dbInfo, query string, sel string, page string) ([]Article, error) {
//	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
//	conn, err := sql.Open(db.engine, dataSource)
//	if err != nil {
//		return nil, err
//	}
//	defer conn.Close()
//
//	rows, err := conn.Query(query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var articles []Article
//	for rows.Next() {
//		var article Article
//
//		if sel == "index" {
//			err := rows.Scan(&article.Title, &article.CreatedBy, &article.CreatedAt)
//			if err != nil {
//				return nil, err
//			}
//		}
//		//else {
//		//	err := rows.Scan(&article.ID, &article.CreatedAt, &article.CreatedBy, &article.ModifiedAt, &article.ModifiedBy, &article.Content, &article.Title, &article.UserID)
//		//	if err != nil {
//		//		return nil, err
//		//	}
//		//}
//
//		articles = append(articles, article)
//	}
//
//	// 에러 핸들링: rows.Next() 이후에 발생할 수 있는 에러들
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return articles, nil
//}
