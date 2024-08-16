package webserver

import (
	"conversion-app/proceed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func StartServer() {
	http.HandleFunc("/", serveMain)
	http.HandleFunc("/weight", weightConversionHandler)
	http.HandleFunc("/length", lengthConversionHandler)
	http.HandleFunc("/temperature", temperatureConversionHandler)
	http.HandleFunc("/resault", serveResult)

	fmt.Println("Server starting at port 8080...")
	http.ListenAndServe(":8080", nil)
}

func serveMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("html", "main.html"))
}

func serveResult(w http.ResponseWriter, r *http.Request) {
	refrense := r.Header.Get("Referer")
	//fmt.Println(strings.Contains(refrense, "weight"))
	fromUnit := r.FormValue("fromUnit")
	toUnit := r.FormValue("toUnit")
	amountStr := r.FormValue("amount")

	// Anonymous function to handle conversion based on the reference
	convert := func(refrense, amountStr, fromUnit, toUnit string) (string, error) {
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			return "", fmt.Errorf("invalid amount: %v", err)
		}
		switch {
		case strings.Contains(refrense, "weight"):
			return proceed.ConvertWeight(amount, fromUnit, toUnit)
		case strings.Contains(refrense, "length"):
			return proceed.ConvertLength(amount, fromUnit, toUnit)
		case strings.Contains(refrense, "temperature"):
			return proceed.ConvertTemperature(amount, fromUnit, toUnit)
		default:
			return "", fmt.Errorf("unsupported conversion type: %s", refrense)
		}
	}

	// Perform the conversion
	result, err := convert(refrense, amountStr, fromUnit, toUnit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Render the result in the template
	tmpl := template.Must(template.ParseFiles(filepath.Join("html", "resault.html")))
	tmpl.Execute(w, map[string]string{"resault": result})
}

func weightConversionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("html", "weight.html")))
	tmpl.Execute(w, nil)
}

func temperatureConversionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("html", "temperature.html")))
	tmpl.Execute(w, nil)

}
func lengthConversionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("html", "length.html")))
	tmpl.Execute(w, nil)

}
