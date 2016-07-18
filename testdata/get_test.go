package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"h12.me/html-query"
	"h12.me/html-query/expr"
)

func TestGetS1(*testing.T) {
	db, err := sql.Open("mysql", "root:admin@tcp(server-01:3306)/bi_data")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT device_ua FROM request")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	m := make(map[string]int)
	for rows.Next() {
		var ua string
		err = rows.Scan(&ua)
		if err != nil {
			panic(err)
		}
		m[ua]++
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	for ua, cnt := range m {
		fmt.Println(cnt, ua)
	}
}

func TestGetIOS(*testing.T) {
	resp, err := http.Get(`http://www.webapps-online.com/online-tools/user-agent-strings/dv/operatingsystem51849/ios`)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	root, err := query.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	var (
		table = expr.Table
		class = expr.Class
	)

	for _, t := range root.Div(class("CmsDataPage_table-cCPPages")).Children(table, class("uas_item")).All() {
		fmt.Println(*t.Td(class("uas_useragent")).Text())
	}
}
