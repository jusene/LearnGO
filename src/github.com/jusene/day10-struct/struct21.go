package main

import "fmt"

type Person struct {
	name string
	age int
	dreams []string
}

func (p *Person) SetDream(dreams []string) {
	p.dreams = dreams
}

func (p *Person) SetDream2(dreams []string) {
	p.dreams = make([]string, len(dreams))
	fmt.Println(p.dreams)
	fmt.Println(dreams)
	copy(p.dreams, dreams)
	fmt.Println(p.dreams)
}

func main() {
	p1 := Person{
		name:  "jusene",
		age:   27,
	}

	data := []string{"eat", "sleep", "play"}
	//p1.SetDream(data)

	data[1] = "不睡觉"
	p1.SetDream2(data)


	fmt.Println(p1.dreams)
}
