package main

import (
	"fmt"
	"time"
)

// ============================================================================
// 1. PENGECEKAN/VERIFIKASI TIPE VARIABLE
// ============================================================================
func checkType(v any) {
	switch v.(type) {
	case int:
		fmt.Println("Tipe data: int")
	case string:
		fmt.Println("Tipe data: string")
	case float64:
		fmt.Println("Tipe data: float64")
	default:
		fmt.Println("Tipe data tidak dikenal")
	}
}

func typeVerification() {
	checkType(10)
	checkType("Hello")
	checkType(3.14)
	checkType(true)
}

// ============================================================================
// 2. FUNCTION CLOSURE
// ============================================================================
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func closure() {
	next := counter()
	fmt.Println(next()) // Output: 1
	fmt.Println(next()) // Output: 2
	fmt.Println(next()) // Output: 3
}

// ============================================================================
// 3. CLASS INHERITANCE (menggunakan struct embedding)
// ============================================================================
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Halo, nama saya", p.Name)
}

type Student struct {
	Person
	City string
}

func inheritance() {
	s := Student{
		Person: Person{Name: "Andri"},
		City:   "Surabaya",
	}
	s.Greet()
	fmt.Println("Kota:", s.City)
}

// ============================================================================
// 4. PERMUTASI
// ============================================================================
func permute(arr []int, l int, r int) {
	if l == r {
		fmt.Println(arr)
	} else {
		for i := l; i <= r; i++ {
			arr[l], arr[i] = arr[i], arr[l] // swap
			permute(arr, l+1, r)
			arr[l], arr[i] = arr[i], arr[l] // backtrack
		}
	}
}

func permutation() {
	arr := []int{1, 2, 3}
	permute(arr, 0, len(arr)-1)
}

// ============================================================================
// 5. GOROUTINE
// ============================================================================
func worker(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Goroutine dihentikan.")
			return
		default:
			fmt.Println("Goroutine berjalan...")
			time.Sleep(1 * time.Second)
		}
	}
}

func goroutine() {
	stop := make(chan bool)
	go worker(stop)
	time.Sleep(5 * time.Second)
	stop <- true
	time.Sleep(1 * time.Second)
}

func main() {
	typeVerification()
	fmt.Println("===============================================================")
	closure()
	fmt.Println("===============================================================")
	inheritance()
	fmt.Println("===============================================================")
	permutation()
	fmt.Println("===============================================================")
	goroutine()
}
