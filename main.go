package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ussii39/go_rest_api/controller"
	"github.com/ussii39/go_rest_api/db"
)

func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load() // envファイルのパスを渡す。何も渡さないと、どうディレクトリにある、.envファイルを探す
		if err != nil {
			panic("Error loading .env file")
		}
	}

	dbConn, err := db.Init()
	if err != nil {
		log.Printf("db init failed: %v", err)
		os.Exit(1)
	}
	tc := &controller.TaskController{dbConn}
	tc2 := &controller.HouseholdController{dbConn}

	router := mux.NewRouter()
	router.HandleFunc("/tasks", tc.CreateTask).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/tasks", tc.GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/houseHold", tc2.CreateHousehold).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/houseHold", tc2.GetHouseholds).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", tc.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", tc.PutTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{uuid}", tc.DeleteTask).Methods(http.MethodDelete, http.MethodOptions)
	log.Print(http.ListenAndServe("0.0.0.0:80", router))
	os.Exit(1)
}
