package handlers

import (
	"athleticclub/models"
	"athleticclub/database"
	"html/template"
	"net/http"
)

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db := database.GetDB()
		name := r.FormValue("name")
		password := r.FormValue("password")

		_, err := db.Exec("INSERT INTO employees(name, password) VALUES (?, ?);", name, password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	}
}

func ShowEmployees(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, employee)
	}

	tmpl := template.Must(template.ParseFiles("templates/employees.html"))
	if err := tmpl.Execute(w, employees); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}