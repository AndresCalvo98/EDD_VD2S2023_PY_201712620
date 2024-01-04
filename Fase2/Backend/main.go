package main

import (
	"Fase2/estructuras/Peticiones"
	"Fase2/estructuras/tablaHash"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tablaAlumnos *tablaHash.TablaHash

func main() {
	tablaAlumnos = &tablaHash.TablaHash{Tabla: make(map[int]tablaHash.NodoHash), Capacidad: 7, Utilizacion: 0}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/login", Validar)
	app.Listen(":4000")

}
func Validar(c *fiber.Ctx) error {
	var usuario Peticiones.PeticionLogin
	c.BodyParser(&usuario)
	if usuario.UserName == "ADMIN_201712620" {
		if usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status":  200,
				"message": "Credenciales correctas",
				"rol":     1,
			})
		} //else {
		//return c.JSON(&fiber.Map{
		//	"status": 400,
		//	"message": "Credenciales incorrectas",
		//})
		//}
		//FLUJO DE Admin
	} else {
		if usuario.Tutor {
			//buscar en arbol B
		} else {
			//buscar en tabla hash
			//if tablaAlumnos.Buscar(usuario.UserName, SHA256(usuario.Password)) {
			//	return c.JSON(&fiber.Map{
			//		"status":  200,
			//		"message": "Credenciales correctas",
			//		"rol":     3,
			//	})
			//}

		}
	}
	return c.JSON(&fiber.Map{
		"status":  400,
		"message": "Credenciales incorrectas",
		"rol":     0,
	})
}
