//路由配置，根据路由决定需要调用那个apis函数
package routers

import (
	ap "test2/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//初始
	router.GET("/", ap.IndexApi)

	//新增
	router.GET("/home/add", ap.AddHtml)
	router.POST("/home/saveadd", ap.AddBlogApi)

	//删除

	router.POST("/home/delete", ap.DeleteBlogApi)

	//编辑
	router.GET("/home/edit", ap.EditHtml)
	router.POST("/home/saveedit", ap.EditBlogApi)

	//显示博客列表
	router.GET("/home/list", ap.ListHtml)
	router.POST("/home/PageData", ap.GetDataList)

	return router
}
