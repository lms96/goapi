package main

import (
	"bytes"
	"encoding/json"
	"goapi/controllers"
	"goapi/database"
	"goapi/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Test Student", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetAllStudentsHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindStudentByIdHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/student/:id", controllers.GetStudent)
	path := "/student/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Test Student", studentMock.Name)
	assert.Equal(t, "12345678910", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
}

func TestEditStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.PATCH("/student/:id", controllers.EditStudent)
	student := models.Student{Name: "Test Student Edited", CPF: "98765432100", RG: "987654321"}
	data, _ := json.Marshal(student)

	path := "/student/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(data))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var editedData models.Student
	json.Unmarshal(res.Body.Bytes(), &editedData)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, student.Name, editedData.Name)
	assert.Equal(t, student.CPF, editedData.CPF)
	assert.Equal(t, student.RG, editedData.RG)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.DELETE("/student/:id", controllers.DeleteStudent)
	path := "/student/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindStudentByCPFHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/student/cpf/:cpf", controllers.GetStudentByCPF)

	req, _ := http.NewRequest("GET", "/student/cpf/12345678910", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
