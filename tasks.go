package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		InfoLogger.Println("http get tasks")
		getTask(w, r)
	default:
		ErrorLogger.Println("http get task error: method is not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	if len(tasks) > 0 {
		task := tasks[0]
		tasks = tasks[1:]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
		InfoLogger.Println("get completed")
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "No available tasks")
}
