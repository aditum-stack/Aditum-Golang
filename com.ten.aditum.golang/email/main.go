package main

import (
	"./controller"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

// 根路径
func handleConnections(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Aditum Email!")
}

type health struct {
	Status string `json:"status"`
}

// eureka心跳监测
func healthCheck(w http.ResponseWriter, r *http.Request) {
	health := health{
		Status: "UP"}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	fmt.Println("health check ", health)
	json.NewEncoder(w).Encode(health)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleConnections)
	router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/email/{emailId}", controller.GetEmailInfoById).Methods("GET")
	router.HandleFunc("/email", controller.GetAllEmailInfo).Methods("GET")
	router.HandleFunc("/email", controller.SendEmail).Methods("POST")
	//router.HandleFunc("/email", controller.UpdateEmailInfoById).Methods("PUT")
	//router.HandleFunc("/email/{emailId}", controller.DeleteEmailInfo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12001", router))
}
