package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func soma(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func subtracao(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func multiplicacao(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func divisao(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func calculo(w http.ResponseWriter, r *http.Request) {
	numero1, err := strconv.ParseFloat(r.FormValue("num1"), 32)
	if err != nil {
		fmt.Fprintf(w, "ERROR, %v", err)
	}
	numero2, err := strconv.ParseFloat(r.FormValue("num2"), 32)
	if err != nil {
		fmt.Fprintf(w, "ERROR %v", err)
	}
	operacao := r.Form.Get("op")

	var result string
	switch operacao {
	case "soma":
		result = fmt.Sprintf("%f", soma(numero1, numero2))
		w.Write([]byte(result))
	case "sub":
		result = fmt.Sprintf("%f", subtracao(numero1, numero2))
		w.Write([]byte(result))
	case "mult":
		result = fmt.Sprintf("%f", multiplicacao(numero1, numero2))
		w.Write([]byte(result))
	case "div":
		result = fmt.Sprintf("%f", divisao(numero1, numero2))
		w.Write([]byte(result))
	}

}

func main() {
	server := &http.Server{
		Addr: ":8000",
	}

	http.HandleFunc("/calcular", calculo)

	server.ListenAndServe()
}
