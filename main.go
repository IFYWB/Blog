package main

import (
	"Blog/blog"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

var db *sql.DB

const (
	url string = "127.0.0.1:3310" //数据库连接端口
)

func initdb() {
	db,_:=sql.Open("mysql","root:123456@(127.0.0.1:3310)/test") // 设置连接数据库的参数
	defer db.Close()	//关闭数据库
	err:=db.Ping()		//连接数据库
	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("连接成功")
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {
	initdb()		//连接数据库

	blog.Register()			//注册
	//http://localhost:2975
	//开启服务器
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/login",  blog.Login)
	}

	r.Run(":2975")
}