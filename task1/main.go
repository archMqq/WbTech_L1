package main

import "fmt"

type Human struct {
	age    int
	height int
}

func (h *Human) SetAge(newAge int) error {
	if newAge < 0 {
		return fmt.Errorf("error age can't be less zero: %d", newAge)
	}

	if newAge > 120 {
		return fmt.Errorf("can't believe you're a centerian")
	}

	h.age = newAge
	return nil
}

func (h *Human) SetHeight(newHeight int) error {
	if newHeight < 20 {
		return fmt.Errorf("wow, are you an embryo?...")
	}

	if newHeight > 272 {
		return fmt.Errorf("you're not the highest man in the world")
	}

	h.height = newHeight
	return nil
}

func (h Human) Speak() {
	fmt.Println("Hahahah... Hello...")
}

type Action struct {
	Human
}

func main() {
	a := Action{}

	a.SetAge(20)
	a.SetHeight(170)
	a.Speak()

	fmt.Println(a)
}
