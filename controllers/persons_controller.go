package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stephenalexbrowne/zoom"
	"github.com/stephenalexbrowne/zoom_example/models"
	"net/http"
	"strconv"
)

type PersonsController struct{}

func (*PersonsController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	persons, err := models.FindAllPersons()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	personsJson, err := json.Marshal(persons)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(personsJson))
}

func (*PersonsController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

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
		return
	}

	p := models.NewPerson(name, ageInt)
	err = zoom.Save(p)
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

func (*PersonsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, "persons.update")
	// TODO: implement this
}

func (*PersonsController) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

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
	w.Header().Set("Content-Type", "application/json")

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
	err = zoom.Delete(p)
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
