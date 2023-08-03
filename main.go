package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type employee struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Experience   int    `json:"experience"`
	Gender       string `json:"gender"`
	PrevEmployer string `json:"prevEmployer"`
}

var employees = []employee{
	{Id: 1, Name: "Mounika", Email: "mounika@gmail.com", Experience: 2, Gender: "female", PrevEmployer: "Infosys"},
	{Id: 2, Name: "Sandhya", Email: "sandhya@gmail.com", Experience: 5, Gender: "female", PrevEmployer: "Walmart"},
	{Id: 3, Name: "Poojitha", Email: "poojitha@gmail.com", Experience: 2, Gender: "female", PrevEmployer: "MindTree"},
	{Id: 4, Name: "Shivani", Email: "shivani@gmail.com", Experience: 3, Gender: "female", PrevEmployer: "Wipro"},
}

func getEmployees(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, employees)
}

func addEmployee(context *gin.Context) {
	var newEmployee employee

	if err := context.BindJSON(&newEmployee); err != nil {
		return
	}

	employees = append(employees, newEmployee)

	context.IndentedJSON(http.StatusCreated, newEmployee)
}

func getEmployeeById(id int) (*employee, error) {
	for i, t := range employees {
		if t.Id == id {
			return &employees[i], nil
		}
	}

	return nil, errors.New("Employee is not Found")
}

func getEmployeeRecordByName(name string) (*[]employee, error) {
	var filteredNames = []employee{}
	for i := 0; i < len(employees); i++ {
		if strings.Contains(strings.ToLower(employees[i].Name), strings.ToLower(name)) {
			filteredNames = append(filteredNames, employees[i])
		}
	}
	if len(filteredNames) != 0 {
		return &filteredNames, nil
	}
	return nil, errors.New("No Employee data found with the given Employee name")
}

func getByEmployeeName(context *gin.Context) {
	employeeName := context.Param("name")
	filteredNames, err := getEmployeeRecordByName(employeeName)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Emplotees Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, filteredNames)

}

func getEmployee(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	employee, err := getEmployeeById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, employee)
}

func updateEmployeeInfo(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	employee, err := getEmployeeById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not Found"})
		return
	}

	employee.Email = "myworld@gmail.com"
	employee.Experience = 5
	employee.PrevEmployer = "Google"

	context.IndentedJSON(http.StatusOK, employee)
}

func deleteEmployeeById(id int) (string, error) {
	flag := false
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			employees[i] = employees[len(employees)-1]
			employees = employees[:len(employees)-1]
			flag = true
		}
	}
	if flag {
		return "Employee record Deleted Successfully", nil
	}
	return "", errors.New("No Employee details found with the provided Employee Id")
}

func deleteById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	msg, err := deleteEmployeeById(int(id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Employees Found to Delete"})
		return
	}
	context.IndentedJSON(http.StatusOK, msg)

}

func main() {
	router := gin.Default()
	router.GET("/getEmployees", getEmployees)
	router.GET("/getEmployee/:id", getEmployee)
	router.GET("/getByEmployeeName/:name", getByEmployeeName)
	router.PATCH("/updateEmployeeInfo/:id", updateEmployeeInfo)
	router.POST("/addEmployee", addEmployee)
	router.DELETE("/deleteEmployeeById/:id", deleteById)
	router.Run("localhost:9090")
}
