package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func handleExpressions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		InfoLogger.Println("http get expressions")
		getExpressions(w, r)
	case http.MethodPost:
		InfoLogger.Println("http post expression")
		addExpression(w, r)
	default:
		ErrorLogger.Println("http method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func getExpressions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expressions)
	InfoLogger.Println("http get expressions completed")
}

func addExpression(w http.ResponseWriter, r *http.Request) {
	var expression Expression
	var result Expression_with_result
	expr := r.URL.RawQuery
	expr2 := strings.Split(expr, "=")
	if expr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid expression")
		ErrorLogger.Println("add expression error: invalid expression")
		return
	}
	if len(expr2) == 2 {
		if expr2[0] != "expression" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid expression")
			ErrorLogger.Println("add expression error: invalid expression")
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid expression")
		ErrorLogger.Println("add expression error: invalid expression")
		return
	}
	expression.Expression = expr2[1]
	id := generateID()
	expression.ID = id
	expression.Status = "Processing"
	expression.CreatedAt = time.Now()
	result.ID = id
	result.Result = "Не посчитано"
	results = append(results, result)
	expressions = append(expressions, expression)
	tasks = append(tasks, expression)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, id)
	InfoLogger.Println("expression with id " + id + " added")
}

func handleExpressionByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/expressions/"):]
	switch r.Method {
	case http.MethodGet:
		InfoLogger.Println("http get expression by id")
		getExpressionByID(w, r, id)
	default:
		ErrorLogger.Println("http get expression by id error: method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func getExpressionByID(w http.ResponseWriter, r *http.Request, id string) {
	for _, expression := range expressions {
		if expression.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(expression)
			InfoLogger.Println("expression by id returned")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	ErrorLogger.Println("Expression not found")
	fmt.Fprintf(w, "Expression not found")
}
