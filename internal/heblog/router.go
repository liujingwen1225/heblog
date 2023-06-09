package heblog

import (
	"github.com/gin-gonic/gin"
	"heblog/internal/heblog/controller/user"
	"heblog/internal/heblog/store"
	"heblog/internal/pkg/core"
	"heblog/internal/pkg/errno"
	"heblog/internal/pkg/log"
)

func initRouter(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function calle")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	userController := user.New(store.S)
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", userController.Create)
		}
	}

	return nil
}
