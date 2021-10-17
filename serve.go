package main

/*We import 4 important libraries
1. “net/http” to access the core go http functionality
2. “fmt” for formatting our text
3. “html/template” a library that allows us to interact with our html file.
4. "time" - a library for working with date and time.*/
import (
	"fmt"
	"go_server/controller"
	"net/http"
	"text/template"
)

//Go application entrypoint
func main() {
	type welcome struct {
		Name string
	}

	var person1 welcome
	person1.Name = "Sangeeth"

	var person2 welcome
	person2.Name = "Joseph"

	a := 1234
	sum := controller.Add(a)
	fmt.Print(sum)
	//Instantiate a Welcome struct object and pass in some random information.
	//We shall get the name of the user as a query parameter from the URL
	//welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	// the relative path). We wrap it in a call to template.Must() which handles any errors and halts if there are fatal errors

	templates := template.Must(template.ParseFiles("home.html"))

	//Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	// a handle that looks in the static directory, go then uses the "/static/" as a url that our
	//html can refer to when looking for our css and other files.

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
	//url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
	//once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
	//It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.

		if err := templates.ExecuteTemplate(w, "home.html", person1); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	templates2 := template.Must(template.ParseFiles("new.html"))
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {

		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.

		if err := templates2.ExecuteTemplate(w, "new.html", person2); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println("http://127.0.0.1:8080/")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
