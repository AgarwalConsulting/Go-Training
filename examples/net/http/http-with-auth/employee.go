package main

type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProjectID   int    `json:"project" validate:"gte=1000"`
	Designation string `json:"title"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf("{\"name\": \"%s\", \"title\": \"%s\", \"project\": %d}", e.Name, e.Designation, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var sequence = 0

func nextID() int {
	sequence++
	return sequence
}

var employees = []Employee{
	Employee{nextID(), "Pramod", 1002, "Sr. Java Dev"},
	Employee{nextID(), "Nandan", 1004, "Sr. Go Dev"},
}

func FindEmployeeBy(id int) *Employee {
	for _, emp := range employees {
		if emp.ID == id {
			return &emp
		}
	}

	return nil
}
