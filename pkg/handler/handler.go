package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	//services *service.Service
}

func NewHandler() *Handler {
	return &Handler{}
}

// routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// sockets
	sockets := router.Group("/ws")
	{
		// voice room
		sockets.GET("/voice/:id", func(c *gin.Context) {
			voiceSocket(c.Writer, c.Request, c)
		})

		// video room
		sockets.GET("/video/:id", func(c *gin.Context) {
			videoSocket(c.Writer, c.Request, c)
		})
	}

	return router
}
