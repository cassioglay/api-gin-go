package controlles

import (
	"github.com/cassioglay/api-gin-go/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(ctx *gin.Context) {
	ctx.JSON(200, models.Alunos)
}

func Saudacao(ctx *gin.Context) {
	nome := ctx.Params.ByName("nome")
	ctx.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})

}
