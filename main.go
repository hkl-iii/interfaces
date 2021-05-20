package main

import "fmt"

type bot interface {
	// define a set of methods/functions (behavior)
	getGreeting() string
}

// design algorithm

// Generics Code
// Dependency Injection

// frontend
// web apps
// HTML, CSS, Bootstrap
// Angular, Vue.js, React.js(Redux) -> Components
// programming language -> JavaScrip /TypeScript(OOP)
// API REST -> client -> customers data JSON/XML/CSV
// JSON (JavaScript object notations)

// backend
// -> Desing API RESTful (Desing)/ Code
// -> Methods, Endpoints, Status codes(errors, oks),Representations
// -> Protocol Communcatyions -> HTTP, gRPC, MQTT, AQMT, WebSockets.io
// -> Database -> Desing/Model, Relational databases RDBMS(MySQL, Postgres, SQlSever), NoSQL(MongoDB), GraphQL

type englishBot struct {
	name string
}

func (e englishBot) getGreeting() string {
	//Very custom logic for generating an english greeting
	return "hi There!" + e.name
}

type books []int

func (b books) getGreeting() string {
	return fmt.Sprintf("hello I am a book, and I have %v books", len(b))
}

type spanishBot struct{}

func main() {

	eb := englishBot{
		name: "golang",
	}

	myBooks := books{1}
	myBook1 := books{1, 2}
	myBook2 := books{1, 2, 3}
	myBook4 := books{1, 2, 3, 5}

	printGreeting(myBooks, myBook1, myBook2, myBook4, eb)
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(bots ...bot) {
	for _, myBot := range bots {
		fmt.Println(myBot.getGreeting())
	}
}
