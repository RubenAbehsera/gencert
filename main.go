package main

import (
	cert "example/gencert/cert"
	"example/gencert/html"
	"fmt"
	"os"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during certificate creation : %v", err)
		os.Exit(1)
	}
	var saver cert.Saver

	// Generate PDF
	//saver, err = pdf.New("output")

	// Generate HTML
	saver, err = html.New("output")

	if err != nil {
		fmt.Printf("Error during pdf generation : '%v'", err)
		os.Exit(1)
	}
	err = saver.Save(*c)
	if err != nil {
		return
	}

}
