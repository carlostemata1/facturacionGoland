package main

import (
	"database/sql"
	"net/http"
	"fmt"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "xQ47YG6Gtv"
	Contrasenia := "you9gXSOqw"
	Nombre := "xQ47YG6Gtv"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(remotemysql.com)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/medicamentos", Inicio)
	http.HandleFunc("/crearMedicamento", CrearMedicamento)
	http.HandleFunc("/insertarMedicamento", InsertarMedicamento)
	http.HandleFunc("/promociones", ListarPromocion)
	http.HandleFunc("/crearPromocion", CrearPromocion)
	http.HandleFunc("/insertarPromocion", InsertarPromocion)
	http.HandleFunc("/facturas", ListarFacturas)
	http.HandleFunc("/crearFactura", CrearFactura)
	http.HandleFunc("/insertarFactura", InsertarFactura)
	http.HandleFunc("/simulador", Simulador)
	http.HandleFunc("/reporte", Reporte)
	http.HandleFunc("/generarReporte", GenerarReporte)

	fmt.Println("servidor corriendo")
	http.ListenAndServe(":8000", nil)
}

//---------------------------------------parte del medicamento --------------------------------------
type Medicamento struct {
	Id        int
	Nombre    string
	Precio    float64
	Ubicacion string
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM Medicamento")

	if err != nil {
		panic(err.Error())
	}
	medicamento := Medicamento{}
	arregloMedicameto := []Medicamento{}

	for obtenerRegistros.Next() {
		var id int
		var nombre string
		var precio float64
		var ubicacion string
		err = obtenerRegistros.Scan(&id, &nombre, &precio, &ubicacion)
		if err != nil {
			panic(err.Error())
		}
		medicamento.Id = id
		medicamento.Nombre = nombre
		medicamento.Precio = precio
		medicamento.Ubicacion = ubicacion

		arregloMedicameto = append(arregloMedicameto, medicamento)
	}
	plantillas.ExecuteTemplate(w, "medicamentos", arregloMedicameto)
}

func CrearMedicamento(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crearMedicamento", nil)
}

func InsertarMedicamento(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		precio := r.FormValue("precio")
		ubicacion := r.FormValue("ubicacion")

		conexionEstablecida := conexionBD()
		insertarMedicamento, err := conexionEstablecida.Prepare("INSERT INTO Medicamento (nombre, precio, ubicacion) VALUES (?, ?, ?);")
		if err != nil {
			panic(err.Error())
		}
		insertarMedicamento.Exec(nombre, precio, ubicacion)
		http.Redirect(w, r, "/medicamentos", 301)
	}
}

//---------------------------------------fin del medicamento --------------------------------------

//-------------------------------------- parte las promociones  --------------------------------------
type Promocion struct {
	Id           int
	Descripcion  string
	Porcentaje   float64
	Fecha_inicio string
	Fecha_fin    string
}
type PromocionV struct {
	IdV           int
	DescripcionV  string
	PorcentajeV   float64
	Fecha_inicioV string
	Fecha_finV    string
}

func ListarPromocion(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM promocion")

	if err != nil {
		panic(err.Error())
	}
	promocion := Promocion{}
	arregloPromocion := []Promocion{}

	for obtenerRegistros.Next() {
		var id int
		var descripcion string
		var porcentaje float64
		var fecha_inicio string
		var fecha_fin string
		err = obtenerRegistros.Scan(&id, &descripcion, &porcentaje, &fecha_inicio, &fecha_fin)
		if err != nil {
			panic(err.Error())
		}
		promocion.Id = id
		promocion.Descripcion = descripcion
		promocion.Porcentaje = porcentaje
		promocion.Fecha_inicio = fecha_inicio
		promocion.Fecha_fin = fecha_fin

		arregloPromocion = append(arregloPromocion, promocion)
	}
	plantillas.ExecuteTemplate(w, "promociones", arregloPromocion)
}

func CrearPromocion(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crearPromocion", nil)
}

func InsertarPromocion(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sePuede := false
		descripcion := r.FormValue("descripcion")
		porcentaje := r.FormValue("porcentaje")
		fecha_inicio := r.FormValue("fecha_inicio")
		fecha_fin := r.FormValue("fecha_fin")
		conexionEstablecida := conexionBD()
		i1, err := strconv.Atoi(porcentaje)
		if err == nil {
			fmt.Println(i1)
		}
		if i1 < 70 {
			sePuede = true
		}
		insertarMedicamento, err := conexionEstablecida.Prepare("INSERT INTO promocion (descipcion, porcentaje, fecha_inicio, fecha_fin) VALUES (?, ?, ?,?);")
		if err != nil {
			panic(err.Error())
		}
		if sePuede {
			insertarMedicamento.Exec(descripcion, porcentaje, fecha_inicio, fecha_fin)
			http.Redirect(w, r, "/promociones", 301)
		} else {
			plantillas.ExecuteTemplate(w, "crearPromocion", nil)
		}
	}
}

//-------------------------------------- parte las Facturas  --------------------------------------
type Factura struct {
	Id          int
	Fecha_crear string
	Pago_total  float64
}

func ListarFacturas(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM Factura")

	if err != nil {
		panic(err.Error())
	}
	factura := Factura{}
	arregloFacturas := []Factura{}

	for obtenerRegistros.Next() {
		var id int
		var fecha_crear string
		var pago_total float64
		err = obtenerRegistros.Scan(&id, &fecha_crear, &pago_total)
		if err != nil {
			panic(err.Error())
		}
		factura.Id = id
		factura.Fecha_crear = fecha_crear
		factura.Pago_total = pago_total

		arregloFacturas = append(arregloFacturas, factura)
	}
	plantillas.ExecuteTemplate(w, "facturas", arregloFacturas)
}

func CrearFactura(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crearFactura", nil)
}

func InsertarFactura(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fecha_crear := r.FormValue("fecha_crear")
		pago_total := r.FormValue("pago_total")
		conexionEstablecida := conexionBD()
		insertarFactura, err := conexionEstablecida.Prepare("INSERT INTO Factura (fecha_crear, pago_total) VALUES (?, ?);")
		if err != nil {
			panic(err.Error())
		}
		insertarFactura.Exec(fecha_crear, pago_total)
		http.Redirect(w, r, "/facturas", 301)
	}
}

//------------------------------

func Simulador(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "simulador", nil)
}

//----------------------------------------------- parte reporte ---------------
func Reporte(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "reporte", nil)
}


func GenerarReporte(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fecha_inicio := r.FormValue("fecha_inicio")
		fecha_fin := r.FormValue("pago_final")
		//conexionEstablecida := conexionBD()
		//insertarFactura, err := conexionEstablecida.Prepare("INSERT INTO Factura (fecha_crear, pago_total) VALUES (?, ?);")
		//if err != nil {
		//	panic(err.Error())
		//}
		//insertarFactura.Exec(fecha_crear, pago_total)
		fmt.Print(fecha_inicio)
		fmt.Print(fecha_fin)
		plantillas.ExecuteTemplate(w, "generarReporte", nil)

	}
}
