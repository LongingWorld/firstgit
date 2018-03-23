package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 { //类型自有的函数用于实现类型自有的行为;func (r ReceiverType) funcName(parameters) (results)
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Color byte

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	withd, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.depth * b.height * b.withd
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (BList BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range BList {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (BList BoxList) PaintItBlack() {
	for i, _ := range BList {
		BList[i].SetColor(BLACK)
	}
	/*for _, b := range BList {//b 只是个临时变量，只存储BList的copy,b.SetColor(BLACK)只改变了b本身的值，并不能改变BList里的值
		b.SetColor(BLACK)
	}*/
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	r1 := Rectangle{12, 9}
	r2 := Rectangle{10, 8}
	c1 := Circle{10}
	c2 := Circle{5}

	fmt.Println("Area of r1 is ", r1.area())
	fmt.Println("Area of r2 is ", r2.area())
	fmt.Println("Area of c1 is ", c1.area())
	fmt.Println("Area of c2 is ", c2.area())

	fmt.Println("~~~~~~~~~~~~~~Next Knowledge Point~~~~~~~~~~~~~~~~")

	Boxes := BoxList{
		Box{2, 4, 6, RED},
		Box{4, 4, 4, BLUE},
		Box{1, 2, 10, BLACK},
		Box{6, 6, 6, WHITE},
		Box{8, 8, 8, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(Boxes))
	fmt.Println("The volume of the first box is ", Boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last box is ", Boxes[len(Boxes)-1].color.String())
	fmt.Println("The biggest box is ", Boxes.BiggestsColor().String())

	fmt.Println("Let's paint them all black!")
	Boxes.PaintItBlack()
	for i, box := range Boxes {
		fmt.Printf("Now the color of the %d box is %s !\n", i, box.color.String())
	}
	fmt.Println("Obviously, now,the biggest box is ", Boxes.BiggestsColor().String())

}
