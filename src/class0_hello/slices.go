package main

import "fmt"

func changeSlices(arr []int) {
	arr[1] = 0
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr1", arr[2:])
	fmt.Println("arr2", arr[:])

	arr1 := arr[2:]
	arr2 := arr[:]
	//changeSlices(arr[2:7])//不能用
	changeSlices(arr1)
	fmt.Println(arr1)

	changeSlices(arr2)
	fmt.Println(arr2)

	//切片是对原数组的view
	var arr0 = [...]int{1, 2, 3, 4, 5, 6, 7}
	var s1 = arr0[2:6]
	var s2 = s1[3:5] //后面可以向后到6,因为cap大于5,哪怕len的超限也没超cap的
	s2[0] = 0        //注意,此时直接改s2也会影响s1
	fmt.Println("s1")
	fmt.Println(s1)
	fmt.Println(len(s1)) //切片长度
	fmt.Println(cap(s1)) //总计长度(包括可向后延伸部分)

	s2 = append(s2, 3) //此后再修改s2与s1无关
	s2[1] = 22
	fmt.Println("s2")
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	var s3 = append(s1, 7)

	fmt.Println(cap(s1), cap(s3)) //s3和arr已经无关,是新的数组
	s3 = append(s3, 1)
	fmt.Println(cap(s1), cap(s3)) //map会自动扩容,而且避免多次扩容提前扩好

}
