package controllers

import (
	"fmt"
	"net/http"
)

type PersonsController struct{}

func (*PersonsController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("PersonsController#Create() was called.")

	fmt.Fprint(w, "persons.create")
	// TODO: implement this
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

	fmt.Fprint(w, "persons.show")
	// TODO: implement this
}

func (*PersonsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("PersonsController#Update() was called.")

	fmt.Fprint(w, "persons.delete")
	// TODO: implement this
}
