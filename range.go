package main

import "fmt"

func main() {
	nums := []int{2, 4, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums{
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apples", "b": "bananas"}
	for k, v := range kvs{
		fmt.Println("%s -> %s\n", k, v)
	}
}
