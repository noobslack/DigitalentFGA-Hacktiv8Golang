package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	panjang, lebar float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (p rectangle) area() float64 {
	return p.panjang * p.lebar
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (p rectangle) perimeter() float64 {
	return 2 * (p.panjang + p.lebar)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
func main() {

	var c1 shape = circle{radius: 5}
	var p1 shape = rectangle{panjang: 2, lebar: 10}

	//empty interface
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
		fmt.Println(v)
	}
	////////////////////////////////////
	//empty interface bisa diterapkan dislice dan map, dimana nantinya akan memuat data dengan tipe data berbeda beda
	rs := []interface{}{"faris", 1, 2, true, "as'ad"}

	_ = rs

	// value, ok := c1.(circle)
	// if ok == true {
	// 	fmt.Printf("%+v\n", value)
	// 	fmt.Printf("circle volume: %f \n", value.volume())
	// }
	fmt.Println("luas lingkaran:", c1.area())
	fmt.Println("luas persegi panjang:", p1.area())
	fmt.Println("keliling lingkaran:", c1.perimeter())
	fmt.Println("keliling persegi panjang:", p1.perimeter())

}
