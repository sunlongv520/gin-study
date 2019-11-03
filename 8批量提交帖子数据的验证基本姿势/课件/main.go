package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	. "topic.jtthink.com/src"
)


func main()  {
	router:=gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", TopicUrl)
		v.RegisterValidation("topics", TopicsValidate) //验证长度
	}

	v1:=router.Group("/v1/topics")//单条帖子
	{
		v1.GET("", GetTopicList)
		v1.GET("/:topic_id",GetTopicDetail)

		v1.Use(MustLogin())
		{
			v1.POST("",NewTopic)
			v1.DELETE("/:topic_id",DelTopic)
		}
	}
	v2:=router.Group("/v1/mtopics")//多条帖子
	{
		v2.Use(MustLogin())
		{
			v2.POST("",NewTopics)

		}
	}

	router.Run()

}
