package main

import "fmt"

type MaySay interface {
	Say()
}

type Human struct {
}

func (h Human) Say() {
	fmt.Println("Ola!")
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("Woof woof")
}

type Cat struct {
}

func (c *Cat) Meow() {
	fmt.Println("Meow")
}

type DogAdapter struct {
	*Dog
}

func (a *DogAdapter) Say() {
	a.Bark()
}

type CatAdapter struct {
	*Cat
}

func (a *CatAdapter) Say() {
	a.Meow()
}

func main() {
	dog, cat, human := DogAdapter{&Dog{}}, CatAdapter{&Cat{}}, Human{}

	dog.Say()
	cat.Say()
	human.Say()
}

/*
Паттерн адаптер используется для приведения одного интерфейса к другому
Его часто применяют когда необходимо использовать композицию, иметь возможность взаимозамещения объектов
	или интегрировать внешнюю библиотеку, имеющую другие интерфейсы

Плюсы:
1. Повторное использование кода
2. Изолирование от деталей
3. Повышение гибкости системы

Минусы:
1. Увеличение количества классов и сложности архитектуры
2. В случае если интерфейсы сильно отличаются друг от друга, легко повысить сложность кода
3. Чрезмерное использование приводит к усложнению чтения кода


Реальное применение:
1. Интеграция сторонних библиотек, например я использую данные в XML, а библиотека в JSON, тогда мой адаптер будет
	преобразовывать XML в JSON и передавать библиотеке
2. Пример в прямом смысле из реальной жизни - переходники например из USB в USB-C
*/
