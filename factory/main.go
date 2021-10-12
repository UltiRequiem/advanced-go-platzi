package main

import "fmt"

type Iproduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func newLaptop() Iproduct {
	return &Laptop{
		Computer{"Laptop 234", 45},
	}
}

func newDesktop() Iproduct {
	return &Laptop{
		Computer{"Desktop 23444", 445},
	}
}

func GetComputerFactory(computerType string) (Iproduct, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid Computer Type")
}

func printNameAndStock(p Iproduct) {
	fmt.Printf("Product Name: %s, with Stock: %d. \n", p.getName(), p.getStock())
}

func main() {
	myLaptop, _ := GetComputerFactory("laptop")
        myDesktop, _ := GetComputerFactory("desktop")

        printNameAndStock(myDesktop)
        printNameAndStock(myLaptop)

}
