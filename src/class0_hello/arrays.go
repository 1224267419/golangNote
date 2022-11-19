package main

import "fmt"

func printArr(arr [5]int) {
	for i, j := range arr { //i得到下标,j得到元素

		fmt.Println(i, j)
	}
}

func printarrPoint(arr *[5]int) {
	for i, j := range arr { //i得到下标,j得到元素

		fmt.Println(i, j)
	}
}

func main() {
	var arr0 = new([5]int) //先定义再赋值1,注意,在GO中就是它返回一个指针，该指针指向新分配的，类型为T的零值。
	arr0[0] = 1
	fmt.Println(len(arr0))
	var arr1 [5]int //先定义再赋值2
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10} //懒得自己数数组长度
	var grid [4][5]bool              //二维数组
	fmt.Println(arr0)
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	for i, j := range arr3 { //i得到下标,j得到元素

		fmt.Println(i, j)
	}
	for j, _ := range arr3 { //只要下标,不要元素

		fmt.Println(j)
	}
	printArr(arr3)
	//printArr(arr2)//报错,因为 [5]int是有5个元素的数组,
	//arr2只有三个,显然不行
	printarrPoint(&arr3)

}
