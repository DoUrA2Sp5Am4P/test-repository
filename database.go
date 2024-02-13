package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handleDatabaseTime(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		InfoLogger.Println("http get database time")
		getDatabaseTime(w, r)
	case http.MethodPost:
		InfoLogger.Println("http post database time")
		changeDatabaseTime(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorLogger.Println("error: method not allowed (database time)")
		fmt.Fprintf(w, "Method not allowed")
	}
}

func getDatabaseTime(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(sl)
	InfoLogger.Println("get database time OK")
}

func changeDatabaseTime(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("time")
	sl_, err := strconv.Atoi(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println("error (post databse time)")
		fmt.Fprint(w, "Error")
		return
	}
	InfoLogger.Println("post database time OK")
	w.WriteHeader(http.StatusOK)
	sl = sl_
}
