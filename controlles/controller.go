package controlles

import (
	"net/http"

	"github.com/cassioglay/api-gin-go/database"
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

func CriaNovoAluno(ctx *gin.Context) {
	var aluno models.Aluno
	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	ctx.JSON(http.StatusOK, aluno)
}
