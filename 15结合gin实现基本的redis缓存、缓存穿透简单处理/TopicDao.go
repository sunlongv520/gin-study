package src

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"net/http"
)
func MustLogin() gin.HandlerFunc {  //必须登录
	return func(c *gin.Context) {
		if _,status:=c.GetQuery("token");!status{
			c.String(http.StatusUnauthorized,"缺少token参数")
			c.Abort()
		}else{
			c.Next()
		}
	}
}

func GetTopicDetail(c *gin.Context)  {
	tid:=c.Param("topic_id")
	topics:=Topics{}
   // DBHelper.Find(&topics,tid)
	//c.JSON(200,topics)
	conn:=RedisDefaultPool.Get()
	redisKey:="topic_"+tid
	defer conn.Close()
	ret,err:=redis.Bytes(conn.Do("get",redisKey))
	if err!=nil{ //缓存里没有
		DBHelper.Find(&topics,tid)
		retData,_:=ffjson.Marshal(topics)
		if topics.TopicID==0{ //代表从数据库没有匹配到
			conn.Do("setex",redisKey,20,retData)
		}else{//正常数据 50秒缓存
			conn.Do("setex",redisKey,50,retData)
		}


		c.JSON(200,topics)
		log.Println("从数据库读取")
	}else{//代表有值
		log.Println("从 redis读取")
		ffjson.Unmarshal(ret,&topics)
		c.JSON(200,topics)
	}


}
func NewTopic(c *gin.Context)  { //单帖子新增
	topic:=Topics{}
	err:=c.BindJSON(&topic)
	if err!=nil{
		c.String(400,"参数错误:%s",err.Error())
	}else{
		c.JSON(200,topic)
	}
}
func NewTopics(c *gin.Context)  { //多帖子批量新增
	topics:=TopicArray{}
	err:=c.BindJSON(&topics)
	if err!=nil{
		c.String(400,"参数错误:%s",err.Error())
	}else{
		c.JSON(200,topics)
	}
}

func DelTopic(c *gin.Context)  {
	//判断登录
	c.String(200,"删除帖子")
}
func GetTopicList(c *gin.Context)  {
	 query:=TopicQuery{}
	 err:=c.BindQuery(&query)
	 if err!=nil{
	 	c.String(400,"参数错误:%s",err.Error())
	 }else{
	 	c.JSON(200,query)
	 }
}