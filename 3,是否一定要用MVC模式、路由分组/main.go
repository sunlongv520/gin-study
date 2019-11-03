package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {


	router:=gin.Default()

	//router.GET("/v1/topics", func(c *gin.Context) {
	//	if c.Query("username")==""{
	//		c.String(200,"获取帖子列表")
	//	}else {
	//		c.String(200,"获取用户名=%s的帖子列表",c.Query("username"))
	//	}
	//})
	//router.GET("/v1/topics/:topic_id", func(c *gin.Context) {
	//	c.String(200,"获取topicid=%s的帖子",c.Param("topic_id"))
	//})

	v1:=router.Group("/v1/topics")
	{
		v1.GET("", func(c *gin.Context) {
			if c.Query("username")==""{
				c.String(200,"获取帖子列表")
			}else {
				c.String(200,"获取用户名=%s的帖子列表",c.Query("username"))
			}
		})

		v1.GET("/:topic_id", func(c *gin.Context) {
			c.String(200,"获取topicid=%s的帖子",c.Param("topic_id"))
		})
	}

	router.Run()

}
