package main

import (
	"log"
	"pricing_details/pkg/handler"
	"pricing_details/pkg/postgres_client"
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
