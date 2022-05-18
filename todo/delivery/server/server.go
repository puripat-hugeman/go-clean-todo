package server

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func NewHttpServer() *Server {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return &Server{
		engine: r,
	}
}
func (g *Server) Serve(addr string) error {
	return g.engine.Run(addr)
}
