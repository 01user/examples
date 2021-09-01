package main

import (
	"github.com/gin-gonic/examples/onboarding/po"
	"github.com/gin-gonic/examples/onboarding/service"
	"github.com/gin-gonic/examples/onboarding/vo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// list all users fixme  将访问数据抽为接口，实现可为内存，可为数据访问
	r.GET("/users", func(c *gin.Context) {
		// 调用业务层
		userService := new(service.UserService)
		users := userService.QueryAll()
		// 处理vo
		var userVos []vo.UserVo
		for _, user := range users {
			userVo := user2Vo(user)
			userVos = append(userVos, userVo)
		}
		c.JSON(http.StatusOK, userVos)
		//c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	})
	// create a user fixme post请求 todo
	r.POST("/users", func(c *gin.Context) {
		// 接收vo
		var voSource vo.UserVo
		err := c.ShouldBind(&voSource)
		if err != nil {
			panic(err)
		}
		// 转为业务po
		user := vo2User(voSource)
		// 调用业务层
		userService1 := new(service.UserService)
		addUser := userService1.AddUser(user)
		// 处理结果
		userVo := user2Vo(addUser)
		c.JSON(http.StatusOK, userVo)
	})
	// list a user's all relationships fixme  将访问数据抽为接口，实现可为内存，可为数据访问
	r.GET("/users/:user_id/relationships", func(c *gin.Context) {
		userId := c.Params.ByName("user_id")
		userIdNum, err := strconv.Atoi(userId)
		if err != nil {
			panic(err)
		}
		// 调用业务层
		relationService := new(service.RelationshipService)
		relationships := relationService.QueryAll(userIdNum)
		var relationshipVos []vo.RelationshipVo
		// 处理结果
		for _, relationship := range relationships {
			relationshipVo := relationship2Vo(relationship)
			relationshipVos = append(relationshipVos, relationshipVo)
		}
		c.JSON(http.StatusOK, relationshipVos)
		//c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	})
	// create/update relationship state to another user fixme put请求 todo
	r.PUT("/users/:user_id/relationships/:other_user_id", func(c *gin.Context) {
		// 接收vo
		var voSource vo.RelationshipInVo
		//err := c.ShouldBind(&voSource)
		err := c.Bind(&voSource)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		userId := c.Params.ByName("user_id")
		userIdNum, err := strconv.Atoi(userId)
		if err != nil {
			panic(err)
		}
		otherUserId := c.Params.ByName("other_user_id")
		otherUserIdNum, err := strconv.Atoi(otherUserId)
		if err != nil {
			panic(err)
		}
		// 转为业务po
		relationship := vo2Relationship(voSource)
		relationship.UserId = userIdNum
		relationship.OtherUserId = otherUserIdNum
		// 调用业务层
		relationshipService1 := new(service.RelationshipService)
		addRelationship := relationshipService1.AddRelationship(relationship)
		// 处理结果
		relationshipVo := relationship2Vo(addRelationship)
		c.JSON(http.StatusOK, relationshipVo)
	})
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func user2Vo(user po.User) vo.UserVo {
	userVo := new(vo.UserVo)
	userVo.Id = user.Id
	userVo.Name = user.Name
	userVo.Type = user.Type
	return *userVo
}

func vo2User(userVo vo.UserVo) po.User {
	user := new(po.User)
	user.Name = userVo.Name
	return *user
}

func relationship2Vo(relationship po.Relationship) vo.RelationshipVo {
	relationshipVo := new(vo.RelationshipVo)
	relationshipVo.OtherUserId = relationship.OtherUserId
	relationshipVo.State = relationship.State
	relationshipVo.Type = relationship.Type
	return *relationshipVo
}

func vo2Relationship(relationshipVo vo.RelationshipInVo) po.Relationship {
	relationship := new(po.Relationship)
	//relationship.UserId = relationshipVo.UserId
	//relationship.OtherUserId = relationshipVo.OtherUserId
	relationship.State = relationshipVo.State
	return *relationship
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
