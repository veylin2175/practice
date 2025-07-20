package main

import "fmt"

type dog struct {
	name string
}

func (d *dog) Bark() {
	fmt.Println(d.name, "гавкает.")
}

type cat struct {
	name string
}

func (c *cat) Meow(period string) {
	fmt.Printf("%v мяукает %v.\n", c.name, period)
}

type adapter interface {
	reaction()
}

type dogAdapter struct {
	*dog
}

func (d *dogAdapter) reaction() {
	d.Bark()
}

func newDogAdapter(d *dog) adapter {
	return &dogAdapter{d}
}

type catAdapter struct {
	*cat
}

func (c *catAdapter) reaction() {
	c.Meow("постоянно")
}

func newCatAdapter(c *cat) adapter {
	return &catAdapter{c}
}

func main() {
	animals := [...]adapter{newDogAdapter(&dog{"Песик"}), newCatAdapter(&cat{"Китик"}), &cow{}}

	for _, animal := range animals {
		animal.reaction()
	}
}

// Спокойно можем добавить еще животных:

type cow struct{}

func (c *cow) reaction() {
	fmt.Println("Кто-то мычит...")
}

// Применимость паттерна:
// 1) Когда нужно использовать старый класс с новым интерфейсом
// 2) Для интеграции сторонних библиотек
// 3) Для унификации разных интерфейсов

// Плюсы:
// 1) Гибкость - позволяет работать с несовместимыми интерфейсами
// 2) Инкапсуляция
// 3) Расширяемость - легко добавить новые адаптеры (пример - cow)

// Минусы:
// 1) Усложнение кода
// 2) Доп ресурсы

// Примеры использования:
// 1) Адаптеры БД
// 2) Логгеры
