package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shinobe179/api/internal/controller/system"
	"github.com/shinobe179/api/internal/controller/user"
)

const latest = "/v1"

type Server struct {}

func(s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// 死活監視用
	// これはなんで{}で囲まれてるの？
	// TODO: 意味のまとまりでしかない可能性がある。括弧外してテスト
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	// ユーザー管理機能
	// これもなんで{}で囲まれてるの？
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUserById)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}