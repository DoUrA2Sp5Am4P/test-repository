package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

type worker_info struct {
	Worker_number string `json:"worker_number"`
	Status        string `json:"status"`
}

var workers []worker_info

type Expression struct {
	ID          string    `json:"id"`
	Expression  string    `json:"expression"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}
type Expression_with_result struct {
	ID     string `json:"id"`
	Result string `json:"value"`
}
type Operation struct {
	Name          string `json:"name"`
	ExecutionTime int    `json:"execution_time"`
}

var sl int = 1
var expressions []Expression
var tasks []Expression
var results []Expression_with_result
var operations = []Operation{
	{Name: "Сложение", ExecutionTime: 1},
	{Name: "Вычитание", ExecutionTime: 2},
	{Name: "Умножение", ExecutionTime: 3},
	{Name: "Деление", ExecutionTime: 4},
}
var InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
var WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime)

func main() {
	InfoLogger.Println("start program")
	ReadAll()
	go saveAll_r()
	InfoLogger.Println("started save opertions")
	go start()
	InfoLogger.Println("started backend")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./html/")))
	mux.HandleFunc("/expressions", handleExpressions)
	mux.HandleFunc("/expressions/", handleExpressionByID)
	mux.HandleFunc("/operations", handleOperations)
	mux.HandleFunc("/operations/addition", handleAddition)
	mux.HandleFunc("/operations/subtraction", handleSubtraction)
	mux.HandleFunc("/operations/multiplication", handleMultiplication)
	mux.HandleFunc("/operations/division", handleDivision)
	mux.HandleFunc("/tasks", handleTasks)
	mux.HandleFunc("/results", handleResults)
	mux.HandleFunc("/results/", handleResultsByID)
	mux.HandleFunc("/database", handleDatabaseTime)
	mux.HandleFunc("/workers", handleWorkers)
	handler := cors.Default().Handler(mux)
	ErrorLogger.Println(http.ListenAndServe(":8080", handler))
}

type Task struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
}

type Result struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
