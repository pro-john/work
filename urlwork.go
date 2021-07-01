package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var Summa int
var Success bool
var ErrCode string

func JSONError(w http.ResponseWriter, err error, bcode string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(struct {
		Code    string
		Message string
	}{
		Code:    bcode,
		Message: err.Error(),
	})
}

func main() {

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {

		Success = true
		ErrCode = "ОК"

		name := r.URL.Query().Get("a")
		name2, _ := strconv.Atoi(name)
		age := r.URL.Query().Get("b")
		age2, _ := strconv.Atoi(age)
		met := r.URL.String()
		//fmt.Println("a = ", name, ",", "b= ", age, ",", "met = ", met[5:8], ",", "name2 = ", name2, ",", "age2 = ", age2)

		switch met[5:8] {
		case "add":
			Summa = name2 + age2
		case "sub":
			Summa = name2 - age2
		case "mul":
			Summa = name2 * age2
		case "div":
			Summa = name2 / age2
		default:

			Success = false
			ErrCode = "Not OK"

		}

		response := struct {
			Success bool
			ErrCode string
			Value   int
		}{
			Success,
			ErrCode,
			Summa,
		}

		err := json.NewEncoder(w).Encode(&response)
		if err != nil {

			JSONError(w, fmt.Errorf("cannot unmarshal response: %w", err), "custom_code2", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusCreated)

	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":10443", nil)
}
