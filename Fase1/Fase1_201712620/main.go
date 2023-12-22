package main

import (
	"Fase1_201712620/estructuras/ArbolAVL"
	"Fase1_201712620/estructuras/ColaPrioridad"
	"Fase1_201712620/estructuras/Listas"
	"Fase1_201712620/estructuras/MatrizDispersa"
	"fmt"
	"strconv"
)

var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var colaPrioridad *ColaPrioridad.Cola = &ColaPrioridad.Cola{Primero: nil, Longitud: 0}
var matrizDispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var arbolCursos *ArbolAVL.ArbolAVL = &ArbolAVL.ArbolAVL{Raiz: nil}

var loggeado_estudiante string = ""

func main() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Inicio de Sesion")
		fmt.Println("2. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			MenuLogin()
		case 2:
			salir = true
		}
	}
}

func MenuLogin() {
	fmt.Print("\033[H\033[2J")
	usuario := ""
	password := ""
	fmt.Print("Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_201712620" && password == "Admin" {
		fmt.Println("Administrador Inicio Sesion")
		MenuAdmin()
	} else if listaDoble.Buscar(usuario, password) {
		loggeado_estudiante = usuario
		MenuEstudiantes()
	} else {
		fmt.Println("ERROR EN CREDENCIALES!!!!")
	}
}

func MenuAdmin() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar de Cursos")
		fmt.Println("4. Control de Estudiantes")
		fmt.Println("5. Reportes")
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			CargaTutores()
		case 2:
			CargaEstudiantes()
		case 3:
			CargaCursos()
		case 4:
			ControlEstudiantes()
		case 5:
			Reportes()
		case 6:
			salir = true
		}

	}
}

func Reportes() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Reporte de Alumnos")
		fmt.Println("2. Reporte de Tutores aceptados")
		fmt.Println("3. Reporte de asignaciones")
		fmt.Println("4. Reporte de Cursos")
		fmt.Println("5. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			listaDoble.Reporte()
		case 2:
			listaDobleCircular.Reportev2()
		case 3:
			matrizDispersa.Reporte("Matriz.jpg")
		case 4:
			arbolCursos.Graficar()
		case 5:
			salir = true
		}

	}
}

func MenuEstudiantes() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("1. Ver Tutores Disponibles")
		fmt.Println("2. Asignarse Tutores")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Print("\033[H\033[2J")
			listaDobleCircular.Mostrar()
		case 2:
			AsignarCurso()
		case 3:
			salir = true
		}
	}
}

func AsignarCurso() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Println("Teclee el codigo del curso: ")
		fmt.Scanln(&opcion)
		if arbolCursos.Busqueda(opcion) {
			if listaDobleCircular.Buscar(opcion) {
				TutorBuscado := listaDobleCircular.BuscarTutor(opcion)
				estudiante, err := strconv.Atoi(loggeado_estudiante)
				if err != nil {
					break
				}
				matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.Carnet, opcion)
				matrizDispersa.Reporte("Matriz.jpg")
				break
			} else {
				fmt.Println("No hay tutores para ese curso....")
				break
			}
		} else {
			fmt.Println("El curso no existe en el sistema")
			break
		}

	}
}

func CargaTutores() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerCSV(ruta)
	fmt.Println("Se cargo a la Cola los tutores")
}

func CargaEstudiantes() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	listaDoble.LeerCSV(ruta)
	fmt.Println("Se cargo los estudiantes")
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Print("\033[H\033[2J")
		colaPrioridad.Primero_Cola()
		fmt.Println("════════════════════")
		if colaPrioridad.Primero != nil {
			fmt.Println("1. Aceptar")
		}
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1 && colaPrioridad.Primero != nil {
			curso := colaPrioridad.Primero.Tutor.Curso
			tutorCarnet := colaPrioridad.Primero.Tutor.Carnet
			tutorNombre := colaPrioridad.Primero.Tutor.Nombre
			tutorNota := colaPrioridad.Primero.Tutor.Nota

			if tutorExistente := listaDobleCircular.BuscarTutor(curso); tutorExistente != nil {
				if tutorExistente.Tutor.Nota < tutorNota {
					listaDobleCircular.SustituirTutor(tutorCarnet, tutorNombre, curso, tutorNota)
					fmt.Println("Se acepto al tutor de curso actual")

					fmt.Scanln()
					colaPrioridad.Descolar()
					listaDobleCircular.Reportev2()
				} else {
					fmt.Println("Se rechazo al tutor de curso actual")
					colaPrioridad.Descolar()
				}
			} else {
				listaDobleCircular.Agregar(tutorCarnet, tutorNombre, curso, tutorNota)
				colaPrioridad.Descolar()
				fmt.Println("Se registro tutor con exito")
				listaDobleCircular.Reportev2()
			}
		} else if opcion == 2 {
			colaPrioridad.Descolar()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}

func CargaCursos() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	arbolCursos.LeerJson(ruta)
	fmt.Println("Se cargaron los cursos")
}
