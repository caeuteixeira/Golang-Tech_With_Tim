package main

import "fmt"

func main() {
	fmt.Printf("Number: %d", 3435)
	/*print out in decimals*/
}

func main() {
	fmt.Printf("Number: %e", 2.37363762)
	/*pirnt out 2.373638e+00 (scientific notation)*/
}

func main() {
	fmt.Printf("Number: %s", "Caeu")
	/*print out Caeu*/
}

func main() {
	fmt.Printf("Number: %q", "Caeu")
	/*print out Caeu but includes the double quotation marks*/
}

func main() {
	fmt.Printf("Number: %-9q is cool", "Caeu")
	/*adds padding*/
}

func main() {
	fmt.Printf("Number: %.2f", 3.4563442)
	/*round it to two decimal points*/
}

func main() {
	fmt.Printf("Number: %.f", 3.4563442)
	/*round it to the whole number*/
}

func main() {
	fmt.Printf("Number: %07d", 3.4563442)
	/*do the padding but it fills with zeros in seven spaces*/
}

func main() {
	var out string = fmt.Sprintf("Number: \n %07d is cool", 45)
	fmt.Println(out)
	/*\n works like an Enter*/
}

func main() {
	var out string = fmt.Sprintf("Number: \t %07d is cool", 45)
	fmt.Println(out)
	/*\t works like a Tab*/
}
