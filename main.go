package main

import (
	"tidb-gin-demo/config"
	"tidb-gin-demo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 创建Gin实例
	r := gin.Default()

	// 创建用户控制器
	userController := &controllers.UserController{}

	// 用户相关路由
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.POST("", userController.CreateUser)      // 创建用户
		userRoutes.GET("", userController.GetUsers)         // 获取所有用户
		userRoutes.GET("/:id", userController.GetUser)       // 获取单个用户
		userRoutes.PUT("/:id", userController.UpdateUser)    // 更新用户
		userRoutes.DELETE("/:id", userController.DeleteUser) // 删除用户
	}

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "TiDB Gin Demo is running!",
		})
	})

	// 启动服务器
	r.Run(":8080")
}