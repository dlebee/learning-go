package main

import "fmt"

type ILogger interface {
	Log(msg string)
}

type Logger struct {
}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

type Service struct {
	Logger
}

type Other struct {
	log Logger
}

func IsLogger(x any) bool {
	if _, ok := x.(ILogger); ok {
		return true
	} else {
		return false
	}
}

func main() {
	// service A is automatically implementing ILogger
	serviceA := Service{}
	serviceA.Log("Composition implicit Acccess!")
	serviceA.Logger.Log("Explicit Access!")
	fmt.Println("Service a is ILogger", IsLogger(serviceA))

	// other is not, because its not using composition.
	other := Other{}
	other.log.Log("Can only do it explicitely because we named the field.")
	fmt.Println("Other is ILogger", IsLogger(other))
}
