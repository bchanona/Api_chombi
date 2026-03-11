package main

import (
	"fmt"

	"github.com/bchanona/Api_chombi.git/src/helper/config"
	userDependencies "github.com/bchanona/Api_chombi.git/src/users/infrastructure/dependencies"
	userRoutes "github.com/bchanona/Api_chombi.git/src/users/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	//Inicializar dependencias
	userDependencies.Init()

	//Configurar router
	r := gin.Default()

	//Inicializar CORS
	config.InitCORS(r)
	
	//Configurar rutas
	api := r.Group("/api")
	v2 := api.Group("/v2")

	userRoutes.UserRoutes(v2)
	
	//Iniciar servidor
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}

}