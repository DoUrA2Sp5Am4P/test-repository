package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-calculator"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func handleWorkers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		InfoLogger.Println("get workers info")
		getWorkers(w, r)
	default:
		ErrorLogger.Println("error: method not allowed (workers info)")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func getWorkers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workers)
	InfoLogger.Println("get workers info OK")
}

var serverURL = "http://localhost:8080" // Адрес сервера оркестратора

func start() {
	concurrency := getConcurrency()
	for i := 0; i < concurrency; i++ {
		go startWorker()
		InfoLogger.Println("started worker with number" + strconv.Itoa(i+1))
	}

	InfoLogger.Println("Agent started with concurrency: ", concurrency)

	select {}
}

func startWorker() {
	numb := len(workers) + 1
	numb_str := strconv.Itoa(numb)
	workers = append(workers, worker_info{Worker_number: numb_str, Status: "Free"})
	for {
		workers[numb-1].Status = "Free"
		task, err := getTask_()
		if err != nil {
			WarningLogger.Println(err)
			time.Sleep(time.Second)
			continue
		}
		workers[numb-1].Status = "Busy"
		result := processTask(task)
		err = sendResult(result)
		if err != nil {
			ErrorLogger.Println(err)
		}
	}
}

func getTask_() (*Task, error) {
	resp, err := http.Get(serverURL + "/tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("no available tasks")
	}

	var task Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("failed to decode task: %v", err)
	}

	return &task, nil
}

func processTask(task *Task) *Result {
	result := &Result{
		ID:    task.ID,
		Value: evaluateExpression(task.Expression),
	}

	return result
}

func calculate(input string) (float64, error) {
	containsSymbols := func(str string, symbols string) bool {
		for _, symbol := range symbols {
			if strings.Contains(str, string(symbol)) {
				return true
			}
		}
		return false
	}
	if containsSymbols(input, "0123456789/*-+ ,.()") {
		a, err := calculator.Calculate(input)
		if err != nil {
			return 0, err
		}
		return a, nil
	} else {
		return 0, fmt.Errorf("Недопустимые символы")
	}
}

func evaluateExpression(expression string) string {
	input := strings.ReplaceAll(expression, " ", "")
	var t int = 0
	var q int = 0
	for _, char := range input {
		if char == 42 {
			t += operations[2].ExecutionTime
			q = 1
		}
		if char == 47 {
			t += operations[3].ExecutionTime
			q = 1
		}
		if char == 43 {
			t += operations[0].ExecutionTime
			q = 1
		}
		if char == 45 {
			t += operations[1].ExecutionTime
			q = 1
		}
	}
	if q == 0 {
		return "error, when calculating"
	}
	time.Sleep(time.Duration(t) * time.Second)
	res, err := calculate(input)
	if err != nil {
		return "error, when calculating"
	}
	fmt.Println(res)
	return fmt.Sprintf("%f", res)
}

func sendResult(result *Result) error {
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal result: %v", err)
	}

	resp, err := http.Post(serverURL+"/results", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send result: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send result, status code: %d", resp.StatusCode)
	}

	return nil
}
