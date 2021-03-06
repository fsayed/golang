package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the data.txt file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Create a new Scanner for file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//Read a line
		//Print that line
		fmt.Println(scanner.Text())
	}

	//close file function has only err return value, report it and exit
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//If there was an error scanning the file, report it and exit
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
