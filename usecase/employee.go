package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jkrish1011/projectx-go/model"
	"github.com/Jkrish1011/projectx-go/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	// Assign new employee id
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	insertId, err := repo.InsertEmployee(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee inserted with id : ", insertId, emp)

}

func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("employee id , ", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeByID(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emps, err := repo.FindAllEmployee()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emps
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var emp model.Employee
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("employee id , ", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id : ", empID)
		res.Error = "Invalid Employee ID"
		return
	}

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	// If this is not set, then the id is set as empty string
	emp.EmployeeID = empID
	count, err := repo.UpdateEmployeeById(empID, &emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("employee id , ", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id : ", empID)
		res.Error = "Invalid Employee ID"
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	// If this is not set, then the id is set as empty string

	count, err := repo.DeleteEmployeeByID(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	// If this is not set, then the id is set as empty string

	count, err := repo.DeleteAllEmployee()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
