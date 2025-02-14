package main

import (
	"fmt"
)

type ProductOfNumbers struct {
	products []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		products: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.products = []int{}
	} else {
		for i := 0; i < len(this.products); i++ {
			this.products[i] = this.products[i] * num
		}
		this.products = append(this.products, num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > len(this.products) {
		return 0
	}
	return this.products[len(this.products)-k]
}

func main() {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)
	if val := obj.GetProduct(2); val != 20 {
		fmt.Printf("1. Test failed, expected 20, got %d", val)
		return
	}
	if val := obj.GetProduct(3); val != 40 {
		fmt.Printf("2. Test failed, expected 40, got %d", val)
		return
	}
	if val := obj.GetProduct(4); val != 0 {
		fmt.Printf("3. Test failed, expected 0, got %d", val)
		return
	}
	obj.Add(8)
	if val := obj.GetProduct(2); val != 32 {
		fmt.Printf("4. Test failed, expected 32, got %d", val)
		return
	}
	fmt.Println("Passed")
}
