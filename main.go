package main

import (
	"github.com/cassioglay/api-gin-go/database"
	"github.com/cassioglay/api-gin-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
