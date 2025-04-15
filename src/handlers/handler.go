package handlers

import (
    "net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the Home Page!"))
}