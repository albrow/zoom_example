package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stephenalexbrowne/models-example/models"
	"net/http"
	"strconv"
)

type PersonsController struct{}

func (*PersonsController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("PersonsController#Create() was called.")

	// get the Name and Age from the form data
	name := r.FormValue("Name")
	if (name == "") || (name == "null") {
		msg := "Missing required paramater: Name"
		http.Error(w, msg, 400)
		return
	}

	age := r.FormValue("Age")
	if (age == "") || (age == "null") {
		msg := "Missing required paramater: Age"
		http.Error(w, msg, 400)
		return
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	p := models.NewPerson(name, ageInt)
	result, err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	personJson, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(personJson))
}

func (*PersonsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("PersonsController#Update() was called.")

	fmt.Fprint(w, "persons.update")
	// TODO: implement this
}

func (*PersonsController) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("PersonsController#Show() was called.")

	// get the Id from the url muxer
	vars := mux.Vars(r)
	personId := vars["Id"]
	if personId == "" {
		http.Error(w, "Missing required paramater: Id", 400)
		return
	}

	p, err := models.FindPersonById(personId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	personJson, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(personJson))
}

func (*PersonsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("PersonsController#Update() was called.")

	fmt.Fprint(w, "persons.delete")
	// TODO: implement this
}
