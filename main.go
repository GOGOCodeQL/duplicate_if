package main

func main() {
	one := 1
	if one == 1 {
		println("A") //!
	} else if one == 1 {
		println("B")
	} else if one == 1 {
		println("C")
	} else if one = 2; one == 2 {
		println("D")
	}
	two := 2
	if two = 1; two == two {
		println("E") //!
	} else if two = 1; two == two {
		println("F") //!
	}
	if two = 1; two == two {
		println("F")
	}

	println("ByeBye")
}
