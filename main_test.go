package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/cassioglay/api-gin-go/controlles"
	"github.com/cassioglay/api-gin-go/database"
	"github.com/cassioglay/api-gin-go/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // Simplifica a visualização dos testes...
	rotas := gin.Default()
	return rotas
}

var ID int

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "1234567898", RG: "12345678901"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controlles.Saudacao)            // ROTA PARA TESTE
	req, _ := http.NewRequest("GET", "/Maria", nil) // REALIZA A REQUISIÇÃO PARA A ROTA
	reposta := httptest.NewRecorder()               // IMPLEMENTEA A INTERFACE DE QUEM VAI REALIZAR A REQUISIÇÃO (GUARDA OS DADOS DA REQUISIÇÃO)
	r.ServeHTTP(reposta, req)                       // REALIZA E GUARDA OS DADOS DA REQUISIÇÃO

	assert.Equal(t, http.StatusOK, reposta.Code)
	mockDaResposta := `{"API diz":"E ai Maria, tudo beleza?"}`
	repostaBody, _ := ioutil.ReadAll(reposta.Body)

	assert.Equal(t, mockDaResposta, string(repostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controlles.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controlles.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/123456", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controlles.BuscarAlunoPorId)
	pathBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "1234567898", alunoMock.CPF)
	assert.Equal(t, "12345678901", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
