package main

import "fmt"

type AliveThing interface {
	IsDead() bool
	Age() int
}
type ImmortalCat struct {
	age int
}

func (ic ImmortalCat) IsDead() bool {
	return true
}

func (ic ImmortalCat) Age() int {
	return ic.age
}

type Animal interface {
	Lick() error
}

type Cat struct {
	AliveThing
}

// if you comment and uncomment this
// you can see how c.Age() will either use, the cats age or ic age
// func (c Cat) Age() int {
// 	return 3
//}

func (c Cat) Lick() error {
	fmt.Println("cat licked")

	return nil
}

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


