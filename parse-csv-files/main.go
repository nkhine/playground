package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Person ...
type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Email     *Email   `json:"email,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//Email ...
type Email struct {
	User   string `json:"name"`
	Domain string `json:"domain"`
}

//Address ...
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	csvFile, _ := os.Open("my.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person // we get an emplty list to store people
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// if not end of file and no error
		splitEmail, err := strings.Split(line[2], "@")
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Email: &Email{
				User:   splitEmail[0],
				Domain: splitEmail[1],
			},
			Address: &Address{
				City:  line[3],
				State: line[4],
			},
		})
	}
	peopleJSON, _ := json.Marshal(people)
	fmt.Println(string(peopleJSON))
}

// The ciphertext is encoded in base64 format
// func splitEmail(email string) (string, error) {
// 	e := strings.Split(email, "@")
// 	return e, nil
// }
