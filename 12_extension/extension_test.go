package _2_extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

//func (d *Dog) Speak() {
//	//d.p.Speak()
//	fmt.Println("Wang!")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	//d.p.SpeakTo(host)
//	d.Speak()
//	fmt.Println(" ", host)
//}

func TestDog(t *testing.T) {
	//var dog Pet := *Pet(new(Dog)) // error cannot use new(Dog) (type *Dog) as type Pet in assignment
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
