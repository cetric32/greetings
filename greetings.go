package greetings

// A Greeting is a message to be printed.
type Greeting struct {
	// Name is the name of the person to greet.
	Name string
	// Message is the message to print.
	Message string
}

// NewGreeting returns a new Greeting with the given name and message.
func NewGreeting(name, message string) *Greeting {
	return &Greeting{name, message}
}

// Greet prints the greeting to the standard output.
func (g *Greeting) Greet() {
	println(g.Name, "says:", g.Message)
}

// Greet prints the greeting to the standard output.
func Greet(name, message string) {
	println(name, "says:", message)
}
