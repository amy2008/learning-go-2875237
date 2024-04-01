package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from meli!"
	file, err := os.Create("./fromString.txt")
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Println("Wrote a file with %v characters\n", length)
	defer file.Close()
	fmt.Println(file.Name())
	defer readFromFile(file.Name()) // file.Name equals "./fromString.txt"
}

func readFromFile(fileName string) {
	data, err := os.ReadFile(fileName) //ioutil.ReadFile() is obsoleted
	checkErr(err)
	fmt.Println("Text read from file: ", string(data))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
