package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
 *	Start app with `go run .`
 *	For files, you can use 'file=@./matrix.csv', 'file=@./matrix-invalid.csv', or 'file=@./matrix-empty.csv'
 *	To run tests, use `go test`
 */

func main() {
	// Let handler handle the routes
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Divert to the appropriate method based on URL path
	if r.URL.Path == "/echo" {
		echo(w, r)
	} else if r.URL.Path == "/invert" {
		invert(w, r)
	} else if r.URL.Path == "/flatten" {
		flatten(w, r)
	} else if r.URL.Path == "/sum" {
		sum(w, r)
	} else if r.URL.Path == "/multiply" {
		multiply(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write(([]byte(fmt.Sprintln("error endpoint not found"))))
	}
}

// Extract method for reading file
func readFile(w http.ResponseWriter, r *http.Request) [][]string {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil
	}
	validateData(records, w)
	return records
}

/*
 *	curl -F 'file=@./matrix.csv' "localhost:8080/echo"
 */
func echo(w http.ResponseWriter, r *http.Request) {
	records := readFile(w, r)
	var response string

	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

/*
 *	curl -F 'file=@./matrix.csv' "localhost:8080/invert"
 */
func invert(w http.ResponseWriter, r *http.Request) {
	records := readFile(w, r)
	var response string

	// Create new slices to populate our inverted matrix
	inverted := make([][]string, len(records))
	for i := range inverted {
		// Initialize each row
		inverted[i] = make([]string, len(records))
	}

	for i, row := range records {
		for j := range row {
			// Invert by taking the opposing positions
			inverted[i][j] = records[j][i]
			inverted[j][i] = records[i][j]
		}
	}

	for _, row := range inverted {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

/*
 *	curl -F 'file=@./matrix.csv' "localhost:8080/flatten"
 */
func flatten(w http.ResponseWriter, r *http.Request) {
	records := readFile(w, r)
	slice := returnAllNums(records)

	fmt.Fprint(w, strings.Join(slice, ","), "\n")
}

/*
 *	curl -F 'file=@./matrix.csv' "localhost:8080/sum"
 */
func sum(w http.ResponseWriter, r *http.Request) {
	records := readFile(w, r)
	slice := returnAllNums(records)

	total := 0
	for _, num := range slice {
		// Convert string to integer first
		current, err := strconv.Atoi(num)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		}
		total += current
	}
	fmt.Fprint(w, total, "\n")
}

/*
 *	curl -F 'file=@./matrix.csv' "localhost:8080/multiply"
 */
func multiply(w http.ResponseWriter, r *http.Request) {
	records := readFile(w, r)
	slice := returnAllNums(records)

	total := 1
	for _, num := range slice {
		// Convert string to integer first
		current, err := strconv.Atoi(num)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		}
		total *= current
	}
	fmt.Fprint(w, total, "\n")
}

func validateData(records [][]string, w http.ResponseWriter) {
	rows := len(records)

	if rows == 0 {
		// Empty csv provided
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln("error no data provided")))
	}

	for _, row := range records {
		if len(row) != rows {
			// Row and column lengths should be equal
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintln("error unequal rows and column lengths")))
		}
	}
}

func returnAllNums(records [][]string) []string {
	// Create a single slice with length equal to the total number of elements
	// This will make it easier to go through each element one by one
	slice := make([]string, len(records)*len(records))
	i := 0

	for _, row := range records {
		for _, num := range row {
			slice[i] = num
			i++
		}
	}
	return slice
}
