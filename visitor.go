package main

import (
	"fmt"
)

type Node interface {
	Accept(Visitor)
}

type ConcreteNodeX struct{}
func (n ConcreteNodeX) Accept(visitor Visitor) {
	visitor.Visit(n)
}

type ConcreteNodeY struct{}
func (n ConcreteNodeY) Accept(visitor Visitor) { 
	fmt.Println("ConcreteNodeY being visited !")
	visitor.Visit(n)
}


type Visitor interface {
	Visit(Node)
}
 
type ConcreteVisitor struct{}
func (v ConcreteVisitor) Visit(node Node) {
	fmt.Println("doing something concrete")
            
	switch node.(type) {
	case ConcreteNodeX:
		fmt.Println("on Node X")
	case ConcreteNodeY:
		fmt.Println("on Node Y")
	}
}


func main() {
	aggregate := []Node {ConcreteNodeX{}, ConcreteNodeY{},}
	
	visitor := new(ConcreteVisitor)
	for _, node := range(aggregate){
		node.Accept(visitor)
	}

}
