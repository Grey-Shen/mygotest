package main

import "fmt"

func main() {

	fmt.Println("vim-go")
}

type Person struct {
	name string
	age  int
}

type PersonHandler interface {
	Batch(origs <-chan Person) <-chan Person
	Handle(orig Person)
}

type PersonHandlerImpl struct{}

func (handler *PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	dests := make(chan Person, 100)
	go func() {
		for {
			p, ok := <-origs
			if ok {
				close(dests)
				break
			}

			handler.Handle(p)
			dests <- p
		}

	}()

	return dests
}

func (handler *PersonHandlerImpl) Handle(orig Person) {
	if orig.name != "shen" {
		orig.name = "shen"
	}
}

func getPersonHandler() PersonHandler {
	return PersonHandlerImpl{}
}

func feathPerson(origs chan<- Person) {

}

func savePerson(dest <-chan Person) <-chan byte {

}
