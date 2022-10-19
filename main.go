package main

import "fmt"

type AliveThing interface {
	IsDead() bool
	Age() int
}
type Animal interface {
	Lick() error
}

// ----------------------------------------------------------------------------------------

type ImmortalCat struct { // is type AliveThing
	age int
}

func (ic ImmortalCat) IsDead() bool {
	return true
}

func (ic ImmortalCat) Age() int {
	return ic.age
}

// ----------------------------------------------------------------------------------------

type Cat struct { // is type Animal AND AliveThing
	// hidden from view, its also implmements Animal
	AliveThing // explicit, I am telling people now I am an alive thing
}

func (c Cat) Lick() error {
	fmt.Println("cat licked")
	return nil
}

//if you comment and uncomment this
//you can see how c.Age() will either use, the cats age or ic age
//func (c Cat) Age() int {
//	return 3
//}

func main() {
	ic := ImmortalCat{
		age: 999999,
	}
	c := Cat{
		ic,
	}
	c.Lick()
	fmt.Printf("Cat is alive: %v\n", c.IsDead())
	fmt.Printf("cat's age: %d\n", c.Age())
}


