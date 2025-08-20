package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Bark() string {
	return fmt.Sprintf("%s: Лает", d.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Meow() string {
	return fmt.Sprintf("%s: Мяукает", c.Name)
}

type Parrot struct {
	Name string
}

// Попугай и так говорит, поэтому реализует интерфейс
func (p *Parrot) Speak() string {
	return fmt.Sprintf("%s: Говорит", p.Name)
}

// ---- Адаптеры ----
// делают Dog и Cat совместимыми с Animal
type DogAdapter struct {
	dog *Dog
}

func NewDogAdapter(d *Dog) *DogAdapter {
	return &DogAdapter{dog: d}
}

func (a *DogAdapter) Speak() string {
	return a.dog.Bark()
}

type CatAdapter struct {
	cat *Cat
}

func NewCatAdapter(c *Cat) *CatAdapter {
	return &CatAdapter{cat: c}
}

func (a *CatAdapter) Speak() string {
	return a.cat.Meow()
}

func announce(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := &Dog{Name: "Бетховен"}
	cat := &Cat{Name: "Гарфилд"}

	dogAdapter := NewDogAdapter(dog)
	catAdapter := NewCatAdapter(cat)

	parrot := &Parrot{Name: "Кеша"}

	animals := []Animal{dogAdapter, catAdapter, parrot}

	for _, a := range animals {
		announce(a)
	}
}

//Данный паттерн стоит применять когда
//1. Интеграция кода (к примеру легаси который нельзя менять)
//2. Моки в тестах
//Реальные примеры:
//1. Замена логгера в проекте
//2. Подключение дополнительной платежной системы
//3. Обертка драйверов бд
//4. Привязка к внешним библиотекам
//Плюсы
//1. Удобство в изменении кода без дальнейших изменений в клиентском коде
//2. Код становится гибким
//Минусы
//1. Код становится менее читаемым и слишком абстрактным
