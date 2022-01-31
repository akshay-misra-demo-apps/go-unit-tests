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

			numbers, err := getNumbersToAdd(r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())

				return
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

	http.HandleFunc("/sumpositive", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			numbers, err := getNumbersToAdd(r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())

				return
			}

			if len(numbers) > 0 {
				for _, number := range numbers {
					if number < 0 {
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(-1)

						return
					}
				}

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

func getNumbersToAdd(request *http.Request) ([]int, error) {
	var numbers []int
	var errr string
	query := request.URL.Query()
	for key, value := range query {
		if value == nil || (len(value) > 0 && value[0] == "") {
			errr = "Value for " + key + " is not provided"
		} else if len(value) > 1 {
			errr = "Value for " + key + " is provided more than once."
		} else {
			fmt.Println("Value for key: " + key + ", Value: " + value[0])
			number, err := strconv.ParseInt(value[0], 10, 0)
			if err != nil {
				errr = "Invalid number: " + value[0]
			} else {
				numbers = append(numbers, int(number))
			}
		}

		if errr != "" {
			return nil, fmt.Errorf(errr)
		}
	}

	fmt.Println("getNumbersToAdd, number: ", numbers)
	return numbers, nil
}
