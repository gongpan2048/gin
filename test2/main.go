// test2 project main.go

package main

import (
	//"log"
	//"net/http"
	//mo "test2/models"
	//ap "test2/apis"
	db "test2/mysql"
	ro "test2/routers"

	//"github.com/gin-gonic/gin"
)

func main() {
	//数据库
	defer db.SqlDB.Close()
	/*------------------------------------------
	测试models中函数是否正确
	--测试AddBlogs()
	bo1 := mo.Blogs{201801, 3, "123", "111111111111111111111111111111"}
	bo1.AddBlogs()
	--测试DeleteBlogs()
	bo2 := mo.Blogs{201801, 3, "", ""}
	bo2.DeleteBlogs()
	--测试EditBlogs()
	bo3 := mo.Blogs{201801, 1, "second blog2", "test blog"}
	bo3.EditBlogs()*/

	//路由
	router := ro.InitRouter()

	//运行
	router.Run(":8002")

}
