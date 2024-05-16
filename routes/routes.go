package routes

import (
	"github.com/cassioglay/api-gin-go/controlles"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controlles.ExibeTodosAlunos)
	r.GET("/:nome", controlles.Saudacao)
	r.POST("/alunos", controlles.CriaNovoAluno)
	r.Run() // listen and serve on 0.0.0.0:8080
}
