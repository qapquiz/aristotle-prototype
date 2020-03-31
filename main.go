package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/sendquestion", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "sendquestion.html")
	})

	r.GET("/channel/:name", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "chan.html")
	})

	r.GET("/channel/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		chanAndMessage := strings.Split(string(msg), ":")
		m.BroadcastFilter([]byte(chanAndMessage[1]), func(q *melody.Session) bool {
			paths := strings.Split(q.Request.URL.Path, "/")
			return paths[len(paths)-2] == chanAndMessage[0]
		})
	})

	r.Run(":5000")
}
