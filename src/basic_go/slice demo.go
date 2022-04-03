package main
import "fmt"
func main() {
	// array demo
	var a [2]string
	a[0] = "Hello"
	a[1] = "Array"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println("-------------------------------------")

	// slice demo
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println("-------------------------------------")

	var cheeses = make([]string, 2)
	cheeses[0] = "Cheddar"
	cheeses[1] = "Edam"
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))
	fmt.Println(cap(cheeses))
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println("-------------------------------------")
	// append to slice
	cheeses = append(cheeses, "Gouda")
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))
	fmt.Println(cap(cheeses))
	fmt.Println(cheeses[2])
	fmt.Println("-------------------------------------")
	cheeses = append(cheeses, "Munster", "Stilton", "Stracchino")
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))
	fmt.Println(cap(cheeses))
	fmt.Println("-------------------------------------")
	// delete from slice
	// delete Gouda
	cheeses = append(cheeses[:2], cheeses[3:]...)
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))
	fmt.Println(cap(cheeses))
	fmt.Println("-------------------------------------")
	// copy slice
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numbers2 = make([]int, 10)
	fmt.Println(numbers)
	fmt.Println(numbers2)
	copy(numbers2, numbers)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))
	fmt.Println("-------------------------------------")

	//map demo
	var players = make(map[string]int)
	players["Ronaldo"] = 10
	players["Messi"] = 11
	players["Neymar"] = 12
	fmt.Println(players)
	fmt.Println(players["Ronaldo"])
	fmt.Println(players["Messi"])
	fmt.Println(players["Neymar"])
	fmt.Println("-------------------------------------")
	// delete from map
	delete(players, "Neymar")
	fmt.Println(players)

}