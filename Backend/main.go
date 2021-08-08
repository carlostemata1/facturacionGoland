package main

import (
	"net/http"
	//"log"
	"fmt"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/medicamentos/crear", CrearMedicamento)
	fmt.Println("servidor corriendo")
	http.ListenAndServe(":8000", nil)
}

func Inicio (w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"hola mundito")
	plantillas.ExecuteTemplate(w,"inicio", nil)
}

func CrearMedicamento (w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"hola mundito")
	plantillas.ExecuteTemplate(w,"crear", nil)
}