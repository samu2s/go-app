package main

import "fmt"



func main() {
	x := addizione(10, 5)
	println(x)

	y := sottrazione(10, 5)
	println(y)

	z := moltiplicazione(10, 2)
	println(z)

	v := divisione(10, 2)
	println(v)

	fmt.Println("ciao")

	r := pariodisp(5)
	println(r)

	j := max (5,3)
	println(j)

	k := add5(6)
	println (k)

	fmt.Printf("%v", marche())

}

func addizione(a, b int) int {
	sum := a + b
	return sum
}

func sottrazione(a, b int) int {
	diff := a - b
	return diff
}

func moltiplicazione(a, b int) int {
	molt := a * b
	return molt
}

func divisione(a, b int) int {
	div := a / b
	return div
}


func pariodisp(x int) bool{
	if x%2==0 {
		return true
	} else {
		return false
	}
}

func max(a,b int) bool{
	if a>b {
		return true
	} else {
		return false
	}

}

func add5(x int) int {
  	return 5+x
}

func marche() []string {
	telefoni := []string{"iphone","xiaomi","samsung","oppo"}
	return telefoni
}