package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func saveAll() {
	time.Sleep(time.Duration(sl) * time.Second)
	saveExpressions()
	saveTasks()
	saveExpressions_with_result()
	saveOperations()
	saveSleep()
}
func saveSleep() {
	a, _ := exists("sleep")
	if !a {
		os.Mkdir("sleep", os.ModeAppend)
	}
	filename := "sleep/sleep.txt"
	b, _ := exists(filename)
	if !b {
		f, _ := os.Create(filename)
		_ = f.Close
	}
	f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	v := 0
	if scanner.Text() == fmt.Sprint(sl) {
		v = 1
	}
	if v == 0 {
		f.Close()
		os.Remove(filename)
		f2, _ := os.Create(filename)
		f2.WriteString(fmt.Sprint(sl) + "\n")
	}
}

func saveExpressions() {
	a, _ := exists("expressions")
	if !a {
		os.Mkdir("expressions", os.ModeAppend)
	}
	for _, elem := range expressions {
		filename := "expressions/" + elem.ID + ".txt"
		a, _ := exists(filename)
		if !a {
			f, _ := os.Create(filename)
			_ = f.Close
		}
		f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		v := 0
		if scanner.Text() == elem.Expression {
			scanner.Scan()
			if scanner.Text() == elem.Status {
				scanner.Scan()
				if scanner.Text() == elem.CreatedAt.Format(time.RFC3339) {
					scanner.Scan()
					if scanner.Text() == elem.CreatedAt.Format(time.RFC3339) {
						v = 1
					}
				}
			}
		}
		if v == 0 {
			f.Close()
			os.Remove(filename)
			f2, _ := os.Create(filename)
			f2.WriteString(elem.Expression + "\n")
			f2.WriteString(elem.Status + "\n")
			f2.WriteString(elem.CreatedAt.Format(time.RFC3339) + "\n")
			f2.WriteString(elem.CreatedAt.Format(time.RFC3339) + "\n")
		}
	}
}

func saveTasks() {
	a, _ := exists("tasks")
	if !a {
		os.Mkdir("tasks", os.ModeAppend)
	}
	for _, elem := range tasks {
		filename := "tasks/" + elem.ID + ".txt"
		a, _ := exists(filename)
		if !a {
			f, _ := os.Create(filename)
			_ = f.Close
		}
		f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		v := 0
		if scanner.Text() == elem.Expression {
			scanner.Scan()
			if scanner.Text() == elem.Status {
				scanner.Scan()
				if scanner.Text() == elem.CreatedAt.Format(time.RFC3339) {
					scanner.Scan()
					if scanner.Text() == elem.CreatedAt.Format(time.RFC3339) {
						v = 1
					}
				}
			}
		}
		if v == 0 {
			f.Close()
			os.Remove(filename)
			f2, _ := os.Create(filename)
			f2.WriteString(elem.Expression + "\n")
			f2.WriteString(elem.Status + "\n")
			f2.WriteString(elem.CreatedAt.Format(time.RFC3339) + "\n")
			f2.WriteString(elem.CreatedAt.Format(time.RFC3339) + "\n")
		}
	}
	dir, err := os.Open("./tasks")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, elem := range files {
		y := 0
		for _, filfilename2_ := range tasks {
			filfilename2_2 := filfilename2_.ID + ".txt"
			if elem.Name() == filfilename2_2 {
				y = 1
			}
		}
		if y == 0 {
			os.Remove("tasks/" + elem.Name())
		}
	}
}

func saveExpressions_with_result() {
	a, _ := exists("expressions_with_result")
	if !a {
		os.Mkdir("expressions_with_result", os.ModeAppend)
	}
	for _, elem := range results {
		filename := "expressions_with_result/" + elem.ID + ".txt"
		a, _ := exists(filename)
		if !a {
			f, _ := os.Create(filename)
			_ = f.Close
		}
		f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		v := 0
		if scanner.Text() == elem.Result {
			v = 1
		}
		if v == 0 {
			f.Close()
			os.Remove(filename)
			f2, _ := os.Create(filename)
			f2.WriteString(elem.Result + "\n")
		}
	}
}
func saveOperations() {
	a, _ := exists("operations")
	if !a {
		os.Mkdir("operations", os.ModeAppend)
	}
	for _, elem := range operations {
		filename := "operations/" + elem.Name + ".txt"
		a, _ := exists(filename)
		if !a {
			f, _ := os.Create(filename)
			_ = f.Close
		}
		f, _ := os.OpenFile(filename, os.O_RDWR, 0666)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		v := 0
		if scanner.Text() == fmt.Sprint(elem.ExecutionTime) {
			v = 1
		}
		if v == 0 {
			f.Close()
			os.Remove(filename)
			f2, _ := os.Create(filename)
			f2.WriteString(fmt.Sprint(elem.ExecutionTime) + "\n")
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func saveAll_r() {
	for {
		saveAll()
		InfoLogger.Println("saved")
		time.Sleep(time.Duration(sl) * time.Second)
	}
}
