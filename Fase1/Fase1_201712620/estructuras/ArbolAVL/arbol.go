package ArbolAVL

import (
	"Fase1_201712620/estructuras/GenerarArchivos"
	"encoding/json"
	"log"
	"math"
	"os"
	"strconv"
)

type ArbolAVL struct {
	Raiz *NodoArbol
}

/***************************************/
type Curso struct {
	Codigo string `json:"Codigo"`
	Nombre string `json:"Nombre"`
}

type DatosCursos struct {
	Cursos []Curso `json:"Cursos"`
}

/***************************************/

func (a *ArbolAVL) altura(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (a *ArbolAVL) equilibrio(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return (a.altura(raiz.Derecho) - a.altura(raiz.Izquierdo))
}

func (a *ArbolAVL) rotacionI(raiz *NodoArbol) *NodoArbol {
	raiz_derecho := raiz.Derecho
	hijo_izquierdo := raiz_derecho.Izquierdo
	raiz_derecho.Izquierdo = raiz
	raiz.Derecho = hijo_izquierdo
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	numeroMax = math.Max(float64(a.altura(raiz_derecho.Izquierdo)), float64(a.altura(raiz_derecho.Derecho)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz_derecho.Factor_Equilibrio = a.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (a *ArbolAVL) rotacionD(raiz *NodoArbol) *NodoArbol {
	raiz_izquierdo := raiz.Izquierdo
	hijo_derecho := raiz_izquierdo.Derecho
	raiz_izquierdo.Derecho = raiz
	raiz.Izquierdo = hijo_derecho
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	numeroMax = math.Max(float64(a.altura(raiz_izquierdo.Izquierdo)), float64(a.altura(raiz_izquierdo.Derecho)))
	raiz_izquierdo.Altura = 1 + int(numeroMax)
	raiz_izquierdo.Factor_Equilibrio = a.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (a *ArbolAVL) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor > nuevoNodo.Valor {
			raiz.Izquierdo = a.insertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.insertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	balanceo := a.equilibrio(raiz)
	raiz.Factor_Equilibrio = balanceo
	if balanceo > 1 && nuevoNodo.Valor > raiz.Derecho.Valor {
		return a.rotacionI(raiz)
	} else if balanceo < -1 && nuevoNodo.Valor < raiz.Izquierdo.Valor {
		return a.rotacionD(raiz)
	} else if balanceo > 1 && nuevoNodo.Valor < raiz.Derecho.Valor {
		raiz.Derecho = a.rotacionD(raiz.Derecho)
		return a.rotacionI(raiz)
	} else if balanceo < -1 && nuevoNodo.Valor > raiz.Izquierdo.Valor {
		raiz.Izquierdo = a.rotacionI(raiz.Izquierdo)
		return a.rotacionD(raiz)
	}
	return raiz
}

func (a *ArbolAVL) InsertarElemento(valor string) {
	if a.Busqueda(valor) {
		return
	}
	nuevoNodo := &NodoArbol{Valor: valor}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *ArbolAVL) busqueda_arbol(valor string, raiz *NodoArbol) *NodoArbol {
	var valorEncontro *NodoArbol
	if raiz != nil {
		if raiz.Valor == valor {
			valorEncontro = raiz
		} else {
			if raiz.Valor > valor {
				valorEncontro = a.busqueda_arbol(valor, raiz.Izquierdo)
			} else {
				valorEncontro = a.busqueda_arbol(valor, raiz.Derecho)
			}
		}
	}
	return valorEncontro
}

func (a *ArbolAVL) Busqueda(valor string) bool {
	buscarElemento := a.busqueda_arbol(valor, a.Raiz)
	if buscarElemento != nil {
		return true
	}
	return false
}

func (a *ArbolAVL) LeerJson(ruta string) {
	data, err := os.ReadFile(ruta)
	if err != nil {
		log.Fatal("Error al leer el archivo:", err)
	}

	var datos DatosCursos
	err = json.Unmarshal(data, &datos)
	if err != nil {
		log.Fatal("Error al decodificar el JSON:", err)
	}

	for _, curso := range datos.Cursos {
		a.InsertarElemento(curso.Codigo)
	}
}

func (a *ArbolAVL) Graficar() {
	cadena := ""
	nombre_archivo := "./ArbolAVL.dot"
	nombre_imagen := "Arbol AVL de cursos.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol{ "
		cadena += a.retornarValoresArbol(a.Raiz, 0)
		cadena += "}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolAVL) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor
		cadena += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "\"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Izquierdo.Valor) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Valor
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Derecho.Valor) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}
