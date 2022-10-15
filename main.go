package main

import (
	"fmt"
	"math/big"
)

// a mod m = x
func mod(a int, m int) int {
	res := a % m
	if res < 0 {
		res += m
	}
	return res
}

//   ab mod m = x
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

// a*x ≡ b mod m
func solve(a int, b int, m int) int {
	if gcd(a, m) == 1 {
		return mod(b*modn(a, phi(m)-1, m), m)
	} else {
		fmt.Println("wrong argument")
		return 0
	}
}

// НСД
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// функція Ейлера
func phi(n int) int {
	var result = n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func generaterandom() *big.Int {
	//p, _ := rand.Prime(rand.Reader, 64)
	//return p
	return big.NewInt(0)
}
func main() {
	var a, b, m int
	for {
		var menu int
		fmt.Println("\nmenu:\n0: вихід\n1: a mod m = x\n2: ab mod m = x\n3: a * x ≡ b mod m\n4: згенерувати просте число в проміжку між a b\n")
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
			fmt.Printf("%d mod %d = %d\n", a, m, mod(a, m))
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
			fmt.Printf("%d pow %d mod %d = %d\n", a, b, m, modn(a, b, m))
		case 3:
			fmt.Println("Введіть число а:")
			_, err := fmt.Scanln(&a)
			if err != nil {
				return
			}
			fmt.Println("Введіть число b:")
			_, err = fmt.Scanln(&b)
			if err != nil {
				return
			}
			fmt.Println("Введіть модуль m:")
			_, err = fmt.Scanln(&m)
			if err != nil {
				return
			}
			fmt.Println(solve(a, b, m))
			fmt.Printf("x = %d ", solve(a, b, m))
			//case 4:
			//	fmt.Println("Введіть нижню межу а:")
			//	_, err := fmt.Scanln(&a)
			//	if err != nil {
			//		return
			//	}
			//	fmt.Println("Введіть верхню межу b:")
			//	_, err = fmt.Scanln(&b)
			//	if err != nil {
			//		return
			//	}
			//	fmt.Println(generaterandom())
		}
	}
}
