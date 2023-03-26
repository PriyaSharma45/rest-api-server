package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type RouteHandler struct {
	Engine *sql.DB
}

func GetRouterEngine(rh *RouteHandler) {

	http.HandleFunc("/add_card_details", rh.AddCardDetails)
	http.HandleFunc("/get_card_details", rh.GetCardDetails)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("error listenandserve", err)
		log.Fatal("error creating port ", err)
	}
}
