package main

import (
	"fmt"
)

func main() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[len(names)])
}


package main

import (
        "fmt"
)

type Shark struct {
	Name string
}

func (s * Shark) SayHello() {
	fmt.Println("Hi! My name is", s.Name)
}

func main() {
	s := &Shark{"Sammy"}
	s = nil
	s.SayHello()
}




package main

import (
	"fmt"

)

func main() {
	defer func ()  {
		fmt.Println("hello from the deferred function!")
		
	}()

	panic("oh no!")
}


package main

func main() {
	foo()
}

func foo() {
	panic("oh no!")
}



package main

import (
	"fmt"
	"log"
)

func main() {
	divideByzero()
	fmt.Println("We survived division by zero")
}

func divideByzero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occured:", err)
		}
	}()
	fmt.Println(divide(1, 0))
}
func divide(a, b int) int {
	return a / b
}		
