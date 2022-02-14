package main

import "fmt"

type IProduct interface {
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

func newLaptop() IProduct {
	return &Laptop{
		Computer{"Laptop 234", 45},
	}
}

func newDesktop() IProduct {
	return &Desktop{
		Computer{"Desktop 23444", 445},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid Computer Type")
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Product Name: %s, with Stock: %d. \n", p.getName(), p.getStock())
}

func main() {
	myLaptop, _ := GetComputerFactory("laptop")
        myDesktop, _ := GetComputerFactory("desktop")

        printNameAndStock(myDesktop)
        printNameAndStock(myLaptop)
}
