package controlles

import (
	"net/http"

	"github.com/cassioglay/api-gin-go/database"
	"github.com/cassioglay/api-gin-go/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(ctx *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	ctx.JSON(200, &alunos)
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

func BuscarAlunoPorId(ctx *gin.Context) {
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não econtrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, &aluno)

}
