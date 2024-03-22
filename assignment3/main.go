package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

type Status struct {
	WaterStatus uint `json:"water"`
	WindStatus  uint `json:"wind"`
}

type Response struct {
	Status []Status `json:"status"`
}

var PORT = ":8080"
var JSONFilePath = "status.json"

func main() {
	http.HandleFunc("/", getStatus)
	fmt.Println("Applications is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := Response{
			Status: []Status{{
				WaterStatus: uint(rand.Intn(20)),
				WindStatus:  uint(rand.Intn(20)),
			}},
		}
		if err := saveToJSON(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		return

	}
	http.Error(w, "Invalid method", http.StatusBadRequest)

}

func saveToJSON(data Response) error {
	jsonData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(JSONFilePath, jsonData, 0644)
}
