package main

import "fmt"

type Topic interface {
	register(ob Observer)
	broadcast()
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func (ec *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from client %s\n", value, ec.id)
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	obs       []Observer
	name      string
	available bool
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(ob Observer) {
	i.obs = append(i.obs, ob)
}

func (i *Item) broadcast() {
	for _, observer := range i.obs {
		observer.updateValue(i.name)
	}
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func main() {

	nvidiaItem := NewItem("RTX 500")

	firstObserver := &EmailClient{"12b"}
	secondObserver := &EmailClient{"34dc"}

        nvidiaItem.register(firstObserver)
        nvidiaItem.register(secondObserver)

        nvidiaItem.UpdateAvailable()

        nvidiaItem.UpdateAvailable()





}
