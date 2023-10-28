package main

import "fmt"

// Мои размеры дома
type Dimensions struct {
	Width  float64
	Height float64
	Depth  float64
}

// Вся мебель присутствующая в доме

type Furniture struct {
	Name   string
	Width  float64
	Height float64
	Depth  float64
}

// Вся бытовая техника присутствующая в доме

type Appliance struct {
	Name   string
	Width  float64
	Height float64
	Depth  float64
}

// Состав семьи, которая живет в данном доме

type FamilyMember struct {
	Name string
	Age  int
}

// Сам дом
type House struct {
	Size       Dimensions
	Furnitures []Furniture
	Appliances []Appliance
	Family     []FamilyMember
}

func main() {
	myHouse := House{
		Size: Dimensions{Width: 10, Height: 6, Depth: 12},
		Furnitures: []Furniture{
			{Name: "Диван", Width: 2.3, Height: 1.3, Depth: 0.6},
			{Name: "Стол", Width: 1.9, Height: 0.9, Depth: 1.5},
			{Name: "Стул", Width: 0.62, Height: 0.8, Depth: 0.4},
			{Name: "Кровать", Width: 2.2, Height: 0.7, Depth: 1.4},
			{Name: "Шкаф", Width: 2.4, Height: 2.5, Depth: 0.8},
		},
		Appliances: []Appliance{
			{Name: "Телевизор", Width: 1.3, Height: 0.85, Depth: 0.13},
			{Name: "Холодильник", Width: 0.9, Height: 2.5, Depth: 0.8},
			{Name: "Стиральная машина", Width: 0.62, Height: 0.9, Depth: 0.7},
			{Name: "Микроволновка", Width: 0.6, Height: 0.3, Depth: 0.8},
			{Name: "Кондиционер", Width: 1, Height: 0.33, Depth: 0.23},
			{Name: "Приставка", Width: 0.01, Height: 0.2, Depth: 0.5},
			{Name: "Сушилка", Width: 0.62, Height: 0.9, Depth: 0.7},
			{Name: "Робот пылесос", Width: 0.03, Height: 0.01, Depth: 0.01},
		},
		Family: []FamilyMember{
			{Name: "Эдвард", Age: 20},
			{Name: "Эльдар", Age: 16},
			{Name: "Дарья", Age: 18},
			{Name: "Ярик", Age: 20},
			{Name: "Сергей", Age: 19},
			{Name: "Семен", Age: 18},
			{Name: "Никита", Age: 20},
		},
	}

	DescribeHouse(myHouse)
}

func DescribeHouse(h House) {
	fmt.Println("Описание дома:")
	fmt.Printf("Размеры: Ширина - %v м, Высота - %v м, Глубина - %v м\n", h.Size.Width, h.Size.Height, h.Size.Depth)

	fmt.Println("\nМебель:")
	for _, f := range h.Furnitures {
		fmt.Printf("%v: Ширина - %v м, Высота - %v м, Глубина - %v м\n", f.Name, f.Width, f.Height, f.Depth)
	}

	fmt.Println("\nБытовая техника:")
	for _, a := range h.Appliances {
		fmt.Printf("%v: Ширина - %v м, Высота - %v м, Глубина - %v м\n", a.Name, a.Width, a.Height, a.Depth)
	}

	fmt.Println("\nСостав семьи:")
	for _, fm := range h.Family {
		fmt.Printf("Имя: %v, Возраст: %v лет\n", fm.Name, fm.Age)
	}
}
