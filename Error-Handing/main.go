package main

import  ( "fmt"
		"errors"
)

func main() {
	err := errors.New("Barcles")
	fmt.Println("The error occurred", err)

}


package main

import (
	"time"
	"fmt"
)

func main() {
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)
}
	


package main

import (
	"errors"
	"fmt"

)

func boom() error {
	return errors.New("barnacles")
}

func main() {
	err := boom()

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("Anchors away!")

}


	package main

	import (
		"errors"
		"fmt"
	)
	
	func boom() error {
		return errors.New("barnacles")
	}
	
	func main() {
		if err := boom(); err != nil {
			fmt.Println("An error occurred:", err)
			return
		}
		fmt.Println("Anchors away!")
	}
		

package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func main() {
	name, err := capitalize("")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}

	fmt.Println("Capitalized name:", name)
} 


package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func main() {
	_, err := capitalize("")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}
	fmt.Println("Success!")
}



package main

import (
	"fmt"
	"os"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}

func sayHello() (string, error) {
	return "", &MyError{}
}

func main() {
	s, err := sayHello()
	if err != nil {
		fmt.Println("unexpected error: err:", err)
		os.Exit(1)
	}
	fmt.Println("The string:", s)
}
