package main

import (
	"log"
	"rest-api-server/pkg/handler"
	"rest-api-server/pkg/postgres_client"
)

func main() {
	client := postgres_client.GetPostgresClient()
	err := client.Ping()
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	routerHandler := &handler.RouteHandler{
		Engine: client,
	}

	handler.GetRouterEngine(routerHandler)
}
