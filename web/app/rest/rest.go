package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/misraak/app/api/add"
)

func Register() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var numbers []int
			var error string
			query := r.URL.Query()
			for key, value := range query {
				if value == nil || (len(value) > 0 && value[0] == "") {
					error = "Value for " + key + " is not provided"
					w.WriteHeader(http.StatusBadRequest)
				} else if len(value) > 1 {
					error = "Value for " + key + " is provided more than once."
					w.WriteHeader(http.StatusBadRequest)
				} else {
					fmt.Println("Value for key: " + key + ", Value: " + value[0])
					number, err := strconv.ParseInt(value[0], 10, 0)
					if err != nil {
						error = "Invalid number: " + value[0]
						w.WriteHeader(http.StatusBadRequest)
					} else {
						numbers = append(numbers, int(number))
					}
				}

				if error != "" {
					json.NewEncoder(w).Encode(error)
					return
				}
			}

			if len(numbers) > 0 {
				sum := add.AddAll(numbers)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(sum)
			}
		default:
			w.WriteHeader(http.StatusNotImplemented)
			message := "Method not supported"
			json.NewEncoder(w).Encode(message)
		}
	})

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server is ready in port 3000")
}
