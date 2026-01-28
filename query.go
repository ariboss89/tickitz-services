package main

import (
	"fmt"
	"strings"
)

func main() {
	title := "har"
	author := "dan"
	sql, val := GetBooks(title, author)
	fmt.Println(sql)
	fmt.Println(val)
}

func GetBooks(title string, author string) (string, []any) {
	var sql strings.Builder
	values := []any{}
	sql.WriteString("SELECT * FROM books")
	if title != "" || author != "" {
		sql.WriteString(" WHERE")
	}
	if title != "" {
		fmt.Fprintf(&sql, " title=$%d", len(values)+1)
		values = append(values, title)
	}
	if author != "" {
		if len(values) > 0 {
			sql.WriteString(" AND")
		}
		fmt.Fprintf(&sql, " author=$%d", len(values)+1)
		values = append(values, author)
	}
	return sql.String(), values
}
