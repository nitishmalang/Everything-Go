package main

import "fmt"

func main() {
	grade := 70
	if (grade >= 65) {
		fmt.Println("Passed")
	
	} 
}
	
  package main

	import "fmt"
	
	func main() {
		grade := 60
	
		if grade >= 65 {
			fmt.Println("Passing grade")
		} else {
			fmt.Println("Failing grade")
		}
	}

  package main

	import "fmt"

	func main() {
		grade := 60

		if grade >= 90 {
			fmt.Println("A grade")

		} else if grade >= 80 {
			fmt.Println("B grade")
		} else if grade >= 70 {
			fmt.Println("C grade")
		} else if grade >= 60 {
			fmt.Println("D grade")
		} else {
			fmt.Println("Failing grade")
		}

	}
   package main

	import "fmt"

	func main() {
		grade := 92
		if grade > 65 {
			fmt.Println("Its passing grade")
			if grade >= 90 {
				fmt.Println("A grade")
			}else if grade >= 80 {
				fmt.Println("B grade")
			}else if grade >= 70 {
				fmt.Println("C grade")
			}else if grade >= 60 {
				fmt.Println("D grade")
			}
		}else {
			fmt.Println("Failing grade")
		}	

		
	}




    package main
	    import "fmt"

	func main() {
		grade := 98
		if grade >= 65 {
			fmt.Println("Passing grade")
			if grade > 90 {
				if grade >= 96 {
					fmt.Println("A+grade")
				}else if grade <= 96 && grade >= 93 {
					fmt.Println("A grade")

				}else {
					fmt.Println("A- grade")
				}
			}
		}
	}

	
	


      
