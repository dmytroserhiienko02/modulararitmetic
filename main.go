package main

import "fmt"

func mod(a int, m int) int {
	res := a % m
	if res < 0 {
		res += m
	}
	return res
}
func modn(a int, b int, m int) int {
	if b == 1 {
		return a
	}
	if b%2 == 0 {
		var t int = modn(a, b/2, m)
		return t * t % m
	} else {
		return modn(a, b-1, m) * a % m
	}
}
func main() {
	var a, b, m int
	for {
		var menu int
		fmt.Println("menu:")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			return
		}
		if menu == 0 {
			break
		}
		switch menu {
		case 1:
			fmt.Println("Введіть число а:")
			_, err := fmt.Scanln(&a)
			if err != nil {
				return
			}
			fmt.Println("Введіть модуль m:")
			_, err = fmt.Scanln(&m)
			if err != nil {
				return
			}
			fmt.Println(mod(a, m))
		case 2:
			fmt.Println("Введіть число а:")
			_, err := fmt.Scanln(&a)
			if err != nil {
				return
			}
			fmt.Println("Введіть степінь b:")
			_, err = fmt.Scanln(&b)
			if err != nil {
				return
			}
			fmt.Println("Введіть модуль m:")
			_, err = fmt.Scanln(&m)
			if err != nil {
				return
			}
			fmt.Println(modn(a, b, m))
		}
	}
}
