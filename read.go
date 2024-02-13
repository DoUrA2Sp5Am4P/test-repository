package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadAll() {
	readExpressions()
	readResults()
	readTasks()
	readOperations()
	readSleep()
}

func readExpressions() {
	a, _ := exists("expressions")
	if !a {
		InfoLogger.Println("Не найдено предыдущих сессий")
	} else {
		dir, err := os.Open("./expressions")
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		for _, elem := range files {
			filename := "expressions/" + elem.Name()
			f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			ID_arr := strings.Split(elem.Name(), ".")
			ID_ := ID_arr[0]
			scanner.Scan()
			expr_ := scanner.Text()
			scanner.Scan()
			stat_ := scanner.Text()
			scanner.Scan()
			created_at_, _ := time.Parse(time.RFC3339, scanner.Text())
			scanner.Scan()
			completed_at_, _ := time.Parse(time.RFC3339, scanner.Text())
			expressions = append(expressions, Expression{ID: ID_, Status: stat_, Expression: expr_, CreatedAt: created_at_, CompletedAt: completed_at_})
		}
	}
}
func readTasks() {
	a, _ := exists("tasks")
	if !a {
		InfoLogger.Println("Не найдено предыдущих сессий")
	} else {
		dir, err := os.Open("./tasks")
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		for _, elem := range files {
			filename := "tasks/" + elem.Name()
			f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			ID_arr := strings.Split(elem.Name(), ".")
			ID_ := ID_arr[0]
			scanner.Scan()
			expr_ := scanner.Text()
			scanner.Scan()
			stat_ := scanner.Text()
			scanner.Scan()
			created_at_, _ := time.Parse(time.RFC3339, scanner.Text())
			scanner.Scan()
			completed_at_, _ := time.Parse(time.RFC3339, scanner.Text())
			tasks = append(tasks, Expression{ID: ID_, Status: stat_, Expression: expr_, CreatedAt: created_at_, CompletedAt: completed_at_})
		}
	}
}
func readResults() {
	a, _ := exists("expressions_with_result")
	if !a {
		InfoLogger.Println("Не найдено предыдущих сессий")
	} else {
		dir, err := os.Open("./expressions_with_result")
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		for _, elem := range files {
			filename := "expressions_with_result/" + elem.Name()
			f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			ID_arr := strings.Split(elem.Name(), ".")
			ID_ := ID_arr[0]
			scanner.Scan()
			res_ := scanner.Text()
			results = append(results, Expression_with_result{ID: ID_, Result: res_})
		}
	}
}
func readOperations() {
	a, _ := exists("operations")
	if !a {
		InfoLogger.Println("Не найдено предыдущих сессий")
	} else {
		dir, err := os.Open("./operations")
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		for _, elem := range files {
			filename := "operations/" + elem.Name()
			f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			Name_of_oper := strings.Split(elem.Name(), ".")
			scanner.Scan()
			execution_time_ := scanner.Text()
			n, _ := strconv.Atoi(execution_time_)
			if Name_of_oper[0] == "Сложение" {
				operations[0].ExecutionTime = n
			}
			if Name_of_oper[0] == "Вычитание" {
				operations[1].ExecutionTime = n
			}
			if Name_of_oper[0] == "Умножение" {
				operations[2].ExecutionTime = n
			}
			if Name_of_oper[0] == "Деление" {
				operations[3].ExecutionTime = n
			}
		}
	}
}
func readSleep() {
	a, _ := exists("sleep")
	if !a {
		InfoLogger.Println("Не найдено предыдущих сессий")
	} else {
		dir, err := os.Open("./sleep")
		if err != nil {
			ErrorLogger.Println(err)
			return
		}
		defer dir.Close()
		filename := "sleep/sleep.txt"
		f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		sl, _ = strconv.Atoi(scanner.Text())
	}
}
