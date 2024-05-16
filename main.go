package main

import (
	"github.com/cassioglay/api-gin-go/database"
	"github.com/cassioglay/api-gin-go/models"
	"github.com/cassioglay/api-gin-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	models.Alunos = []models.Aluno{
		{Nome: "Aluno 0", CPF: "000000000", RG: "000000000"},
		{Nome: "Aluno 1", CPF: "111111111", RG: "111111111"},
		{Nome: "Aluno 2", CPF: "222222222", RG: "222222222"},
	}

	routes.HandleRequest()
}
