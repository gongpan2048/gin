package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	//"strconv"
	mo "test2/models"

	"github.com/gin-gonic/gin"
)

/*type Blogs struct {
	 string `json:"content" form:"content"`
}
*/
//初始页面
func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! this is a dome about blog")
}

//博客列表页面
func ListHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{
		"title": "GIN: 博客列表页面",
	})
}

//获取博客列表
func GetDataList(c *gin.Context) {
	//得到参数用户ID，得到该用户的博客
	userId := c.PostForm("userid")

	//将userid(string)类型转换成(int)
	fmt.Println(userId)
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatalln(err)
	}
	//调用models里的GetBlogList(uid)获取博客数据列表
	datalist := mo.GetBlogList(uid)
	//count：=mo.GetBlogSum(uid)

	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"datalist": datalist,
	})

}

//新增页面
func AddHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "GIN: 新增博客页面",
	})
}

//新增记录API
func AddBlogApi(c *gin.Context) {
	//得到参数userid
	userId := c.PostForm("userid")

	//将userid(string)类型转换成(int)
	fmt.Println(userId)
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatalln(err)
	}
	//得到参数blogid
	bid := mo.GetBlogSum(uid)

	if err != nil {
		log.Fatalln(err)
	}

	//得到参数title，content
	tilte := c.PostForm("title")
	content := c.PostForm("content")

	//赋值
	b := mo.Blogs{UserID: uid, BlogID: bid, Title: tilte, Content: content}
	//调用models中的AddBlogs(),增加博客
	r1 := b.AddBlogs()

	c.JSON(http.StatusOK, gin.H{
		"success": r1,
	})
	c.Redirect(http.StatusOK, "/home/list")
}

//删除记录API
func DeleteBlogApi(c *gin.Context) {

	//得到参数userid
	userId := c.PostForm("userid")
	//得到参数blogid
	blogId := c.PostForm("blogid")
	//将userid(string)类型转换成(int)
	fmt.Println(userId)
	uid, err := strconv.Atoi(userId)
	//将blogid(string)类型转换成(int)
	fmt.Println(blogId)
	bid, err := strconv.Atoi(blogId)

	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}

	/*uid := 201801
	bid := 1*/

	b := mo.Blogs{UserID: uid, BlogID: bid, Title: "", Content: ""}

	//调用models中的DeleteBlogs(),删除博客
	r2 := b.DeleteBlogs()
	if r2 == false {
		log.Fatalln("删除失败")
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": r2,
	})

}

//编辑页面
func EditHtml(c *gin.Context) {
	//得到URL请求的参数
	uid_t := c.Query("userid")
	bid_t := c.Query("blogid")

	uid, err1 := strconv.Atoi(uid_t)
	bid, err2 := strconv.Atoi(bid_t)

	if err1 != nil {
		log.Fatalln(err1)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}

	b := mo.GetBlogData(uid, bid)
	if b == nil {
		fmt.Println("得到数据错误")
	} else {
		fmt.Println(b)
		fmt.Println("得到数据正确")
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"userid": b.UserID,
		"blogid": b.BlogID,
		"title":  b.Title,
		//"content":	 b.Content
	})
}

//编辑记录API
func EditBlogApi(c *gin.Context) {

	//得到参数
	//得到参数userid
	userId := c.PostForm("userid")

	//将userid(string)类型转换成(int)
	fmt.Println(userId)
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatalln(err)
	}
	//得到参数blogid
	blogId := c.PostForm("blogid")

	//将blogid(string)类型转换成(int)
	fmt.Println(blogId)
	bid, err := strconv.Atoi(blogId)
	if err != nil {
		log.Fatalln(err)
	}

	//得到参数title，content
	tilte := c.PostForm("title")
	content := c.PostForm("content")

	//赋值
	b := mo.Blogs{UserID: uid, BlogID: bid, Title: tilte, Content: content}

	//调用models中的EditBlogs()，编辑函数
	r3 := b.EditBlogs()
	if r3 == false {
		log.Fatalln("编辑失败")
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": r3,
	})
	c.Redirect(http.StatusOK, "/home/list")

}
