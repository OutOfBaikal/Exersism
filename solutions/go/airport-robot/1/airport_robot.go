package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(a string) string 
}

// Реализация Greeter для немецкого языка
type GermanGreeter struct{}

func (g GermanGreeter) LanguageName() string {
    return "German"
}

func (g GermanGreeter) Greet(name string) string {
    return "Hallo " + name + "!"
}

// Функция SayHello
func SayHello(name string, greeter Greeter) string {
    return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

type Italian struct{}

func (g Italian) LanguageName() string {
    return "Italian"
}

func (g Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

type Portuguese struct{}

func (g Portuguese) LanguageName() string {
    return "Portuguese"
}

func (g Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
