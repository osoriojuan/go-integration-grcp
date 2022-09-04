package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	structs "github.com/osoriojuan/go-integration-grcp/structs"
)

var router *gin.Engine

func StartDefaultRouter() {
	router = gin.Default()
}

func ConfigureRoutesAvailables() {

	defResponse := structs.DefaultResponse{
		Result:  "SUCCESS!",
		Message: "ENDPOINT SUCCESS",
		Code:    200,
	}

	log.Println("Setting endpoints...")

	router.GET("/test", func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, defResponse)

	})

}

func Serve() {
	log.Println("Runing router...")
	router.Run()
}
