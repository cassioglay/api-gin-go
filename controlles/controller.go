package controlles

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"id":   "1",
		"nome": "Paulo Santos",
	})
}
