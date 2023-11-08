// проверка массива на сортированность, два способа
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	//создадим сортированный массив
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = /*10*i +*/ rand.Intn(10) //закоментируем, чтобы разупорядочить массив
	}
	fmt.Println(checkSorted(a))
}

//проверим массив на сортированность, если поиск необходимо провести один раз,
//проверять на сортированность не имеет смысла, так сложность такого поиска О(n)
//и сложность сортировки тоже О(n). Но если необходимо сделать поиск несколько раз,
//проверив массив на сортированность, в случае его сортированности можно применить
//алгоритм бинарного поиска, сложность которого значительно меньше чем О(n),
//и будет О(logn)
/*func checkSorted(a [n]int) (result bool) {
	result = true
	for i := 0; i < n-1; i++ {
		if a[i+1] < a[i] {
			result = false
			break
		}
	}
	return
}*/
//проверка на сортированность массива методом монте-карло, в больших массивах данный
//метод позволяет ускорить время сортировки, но он не дает 100%-ю гарантию точности результата
func checkSorted(a [n]int) (result bool) {
	result = true
	firstIndex := rand.Intn(n / 3)
	secondIndex := rand.Intn(n/3) + n/3
	thirdIndex := rand.Intn(n/3) + 2*n/3
	if a[firstIndex] > a[secondIndex] || a[secondIndex] > a[thirdIndex] || a[firstIndex] > a[thirdIndex] {
		result = false
		fmt.Println("Определено методом Монте-карло")
		return
	}
	for i := 0; i < n-1; i++ {
		if a[i+1] < a[i] {
			result = false
			fmt.Println("Определено методом перебора")
			break
		}
	}
	return
}
