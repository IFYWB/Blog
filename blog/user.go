package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	_"github.com/jmoiron/sqlx"
)

type user struct {
	Username string	  `db:"username"`
	Sex string  `db:"sex"`
	Password string		`db:"password"`
}

var db *sql.DB

//用户登录
func Login(c *gin.Context) {
	router := gin.Default()
	// 匹配的url格式:  http://localhost:2975/getMes/mes
	router.GET("/user/login", func(c *gin.Context) {
		id := c.DefaultQuery("id", "Guest")
		password := c.Query("password") // 是 c.Request.URL.Query().Get("lastname") 的简写
		//id, password, ok := Check(id, password)
		ok := Check(id, password)
		if !ok {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "验证失败",
			})
			return
		} else{
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "验证成功,正在跳转",

			})
		}
		c.String(http.StatusOK, "Hello %s", id)
	})
	router.Run(":2975")
}

func Check(id string,password string)(bool){
	if id=="LiuC"&&password=="123456"{
		return true
	}else{
		return false
	}
}

func Register(){
	db,_:=sql.Open("mysql","root:123456@(127.0.0.1:3310)/test") // 设置连接数据库的参数
	defer db.Close()	//关闭数据库
	err:=db.Ping()		//连接数据库
	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("连接成功")
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
	r, err := db.Exec("insert into user(username, sex, password)values(?, ?, ?)", "LiuC", "man", "123456")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
