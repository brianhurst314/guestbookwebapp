package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

/*
Test ...
This is a comment
*/
type Test struct {
	Title       string
	Instruction string
	Names       []string
}

func (t *Test) save() error {
	filename := "testsave.txt"
	input, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, []byte(input), 0600)
}

func load() (*Test, error) {
	filename := "testsave.txt"
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var t Test
	err = json.Unmarshal(fileData, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	test, _ := load()

	//test.Title = "Hello World"

	//test.save()

	t, _ := template.ParseFiles("./static/index.html")

	fmt.Println(test, "This is a Test *********************************")

	t.Execute(w, test)
}

func addNameHandler(w http.ResponseWriter, r *http.Request) {

	//read the form data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//assign value from form to a variable
	newName := r.Form.Get("newName")

	//load current data from file
	test, _ := load()

	//test.Title = "This is Working!!!"

	//add newName to the list
	test.Names = append(test.Names, newName)

	//save the new data to the file
	test.save()

	//reload the webpage to reflect the changes
	t, _ := template.ParseFiles("./static/index.html")

	t.Execute(w, test)
}

func main() {

	name := []string{"Test", "Test2"}
	test := Test{Title: "Guest Book", Instruction: "Enter a new name below", Names: name}
	test.save()

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/addName/", addNameHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
