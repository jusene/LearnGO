package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Formatter interface {
	Format(f fmt.State, c rune)
}

type Person struct {
	Name string
	Age int
	Sex int
}

func (p *Person) String() string {
	buffer := bytes.NewBufferString("这是")
	buffer.WriteString(p.Name + ",")
	if p.Sex == 0 {
		buffer.WriteString("他")
	} else {
		buffer.WriteString("她")
	}
	buffer.WriteString("今年")
	buffer.WriteString(strconv.Itoa(p.Age))
	buffer.WriteString("岁。 ")
	return buffer.String()
}

func(p *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(p.String()))
		f.Write([]byte("Person有三个字段。"))
	} else {
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}

func main() {
	person := &Person{
		Name: "jusene",
		Age:  27,
		Sex:  0,
	}

	Formatter(person).Format('L', 33)
}
