package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"algogrit.com/emp-server/employee/entities"
	"algogrit.com/emp-server/employee/repository"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type jsonService struct {
	repo repository.EmployeeRepository
}

func (es *jsonService) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	employees := es.repo.RetrieveAll()
	json.NewEncoder(w).Encode(employees)
}

func (es *jsonService) Create(w http.ResponseWriter, r *http.Request) {
	var e entities.Employee
	json.NewDecoder(r.Body).Decode(&e)

	validate := validator.New()
	errs := validate.Struct(e)

	if errs != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unable to process: %v", errs)
		return
	}

	savedEmployee, errs := es.repo.Save(e)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(savedEmployee)
}

func (es *jsonService) Show(w http.ResponseWriter, r *http.Request) {
	empID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(empID)

	emp := es.repo.FindBy(id)

	json.NewEncoder(w).Encode(emp)
}

// NewJSON returns an instance of new EmployeeService
func NewJSON(repo repository.EmployeeRepository) EmployeeService {
	return &jsonService{repo}
}
