//模型配置文件，对数据库的增删改查
package models

import (
	"fmt"
	"log"
	db "test2/mysql"
)

/*用户ID，博客ID，标题，正文*/
type Blogs struct {
	UserID  int    `json:"userid" form:"userid"`
	BlogID  int    `json:"blogid" form:"blogid"`
	Title   string `json:title" form:"title"`
	Content string `json:"content" form:"content"`
}

//新增博客
func (b *Blogs) AddBlogs() bool {
	rs, err := db.SqlDB.Exec("insert into blog values(?,?,?,?)", b.UserID, b.BlogID, b.Title, b.Content)
	if err != nil {
		return false
	}

	id, err := rs.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return false
	} else {
		return true
	}

}

//删除博客
func (b *Blogs) DeleteBlogs() bool {
	rs, err := db.SqlDB.Exec("delete from blog where userid=? and blogid=?", b.UserID, b.BlogID)
	if err != nil {
		return false
	}

	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err != nil {
		return false
	} else {
		return true
	}
}

//编辑博客
func (b *Blogs) EditBlogs() bool {
	rs, err := db.SqlDB.Exec("update blog set title=?,content=? where userid=? and blogid=?", b.Title, b.Content, b.UserID, b.BlogID)
	if err != nil {
		return false
	}

	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err != nil {
		return false
	} else {
		return true
	}
}

//根据用户ID得到该用户博客序号
func GetBlogSum(uid int) int {
	//计数器，记录博客序号(数量）
	num := 0
	rows, err := db.SqlDB.Query("select * from blog where userid=?", uid)
	if err != nil {
		return 0
	}

	defer rows.Close()

	for rows.Next() {
		num++
	}

	return num + 1
}

//根据用户ID得到博客记录列表
func GetBlogList(uid int) (blog []Blogs) {
	rows, err := db.SqlDB.Query("select * from blog where userid=?", uid)
	if err != nil {
		return nil
	}
	defer rows.Close()

	Blogs_t := make([]Blogs, 0)

	for rows.Next() {
		var blog Blogs
		rows.Scan(&blog.UserID, &blog.BlogID, &blog.Title, &blog.Content)
		Blogs_t = append(Blogs_t, blog)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	//返回记录列表
	return Blogs_t
}

//根据用户ID，博客ID得到博客内容
func GetBlogData(uid int, bid int) (b *Blogs) {
	var blog Blogs

	//将查询的值赋给变量blog
	err := db.SqlDB.QueryRow("select userid,blogid,title,content from blog where userid=? and blogid=?", uid, bid).Scan(
		&blog.UserID, &blog.BlogID, &blog.Title, &blog.Content,
	)

	if err != nil {
		log.Println(err)
	}

	return &blog
}
