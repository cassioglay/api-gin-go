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
	r.GET("/alunos/:id", controlles.BuscarAlunoPorId)
	r.DELETE("/alunos/:id", controlles.DeletaAluno)
	r.PATCH("/alunos/:id", controlles.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controlles.BuscaAlunoPorCPF)
	r.Run() // listen and serve on 0.0.0.0:8080
}
