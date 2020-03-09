package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
)

func query() {
	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", "percona1/ppercona1234@10.16.16.26:1521/svdp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select name from t1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("00000")

	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Printf("Name = %s, len=%d", name, len(name))
		fmt.Printf("1111")

		log.Printf("Name = %s, len=%d", name, len(name))
	}
	rows.Close()
}

func update() {
	// 为log添加短文件名,方便查看行数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 用户名/密码@实例名  跟sqlplus的conn命令类似
	db, err := sql.Open("oci8", "percona1/ppercona1234@ASMSDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare(`UPDATE FUB_B set name ='cnm'`)
	result, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	count, _ := result.RowsAffected()
	log.Printf("result count:%d", count)
}

func main() {
	fmt.Println("开始执行查询")
	query()
	fmt.Println("结束执行查询")

	//update()
}
