package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)

type client struct {
  ID      string //`json:"id"`
  Name    string //`json:"name"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
  apiVersion := "v0.1"
  fmt.Fprintf(w, "API %s", apiVersion)
}

func getClient(w http.ResponseWriter, r *http.Request) {
  clients := map[string]client{
    "1": {
        ID: "1",
        Name: "Blair Trump",
      },
    "2": {
        ID: "2",
        Name: "Mina Hu",
      },
    }

  vars := mux.Vars(r)
  payload, _ := json.Marshal(clients[vars["id"]])
  w.Header().Set("Content-Type", "application/json")
  w.Write(payload)
  fmt.Println(clients)
}

func getClients(w http.ResponseWriter, r *http.Request) {
  clients := map[string]client{
    "1": {
        ID: "1",
        Name: "Blair Trump",
      },
    "2": {
        ID: "2",
        Name: "Mina Hu",
      },
    }

  payload, _ := json.Marshal(clients)
  w.Header().Set("Content-Type", "application/json")
  w.Write(payload)
  fmt.Println(clients)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", homeLink)
  router.HandleFunc("/client", getClients)
  router.HandleFunc("/client/{id}", getClient)
  port := 8081
  msg := fmt.Sprintf("Listening on port %d", port)
  fmt.Println(msg)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
