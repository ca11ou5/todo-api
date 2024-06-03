package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"todo-list/configs"
	"todo-list/internal/handler"
	_ "todo-list/internal/handler/http/docs"
)

func StartListening(cfg *configs.Config, h *handler.Handler) {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.POST("task", h.CreateTask)
	r.GET("task/:id", h.GetTask)
	r.PUT("task/:id", h.UpdateTask)
	r.DELETE("task/:id", h.DeleteTask)
	r.GET("task", h.GetTaskList)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
