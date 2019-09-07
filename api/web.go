package api

import (
	"github.com/gin-gonic/gin"
)

func RunWebServer(exit chan string) {
	router := PrepareRoutes()

	/* ******************************** */
	/* ***		PORTA DA API 		***	*/
	/* ********************************	*/
	router.Run(":8080")
	exit <- "Exit"
}

func PrepareRoutes() *gin.Engine{
	r := gin.Default()

	/* ******************************** */
	/* ***		ROTAS DA APLICAÇÃO	***	*/
	/* ********************************	*/
	r.GET("/ping", Pong)
	return r
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}