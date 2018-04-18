package main

import (
	"fmt"
)

func main() {

	//createArr()
	//editArr()
	//pointArr()
	//rangeArr()
	//newarr()

	var testArr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(testArr)

	modifyarr(testArr)

	fmt.Println("In main", testArr)

}

func modifyarr(arr [5]int) {
	arr[0] = 10
	fmt.Println("In modify", arr)
}

func createArr() {

	var a [5]byte //长度为5的数组，每个元素为一个字节
	fmt.Println(a)
	var c [5]*int // 指针数组
	fmt.Println(c)
	var d [2][3]int //二维数组
	fmt.Println(d)
	// 指定数组长度
	// 创建数组如果给定数组长度，数组个数就不能超过这个长度
	var arr1 = [5]int{1, 3, 5, 2, 9}
	fmt.Println(arr1)

	//数组的长度一定是固定的
	//如果不指定数组长度
	//Go 语言会根据元素的个数来设置数组的大小
	var arr2 = []int{1, 2}       // 此种方式会变成切片
	var arr3 = [...]int{2, 3, 4} // 此种方式数组长度由系统自动计算
	fmt.Println(arr2)
	fmt.Println(arr3)

	//创建变量或者常量 可以简写
	arr4 := []int{5}
	fmt.Println(arr4)
	//初始化对指定元素赋值
	//第三个元素 初始化为1
	var arr5 = [5]int{3: 1}
	fmt.Println(arr5)

	// 获取数组的长度和容量
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

}

func editArr() {

	var arr1 = []int{1, 2}
	fmt.Print(len(arr1), cap(arr1))
	fmt.Println(arr1)

	// 给arr1 添加元素 3，4
	var arr2 = append(arr1, 3, 4)
	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)
	fmt.Print(len(arr1), cap(arr1))

	//改变数组中元素的值
	arr1[0] = 9
	fmt.Println(arr1)
}

func pointArr() {
	//数组指针   它是一个指针，指向数组的地址
	a := []int{5: 1}
	var p *[]int = &a
	fmt.Println(a)
	fmt.Println(p) //比a 多了一个取地址符

	// 指针数组 数组里存放的是指针地址，不是实际的值
	var x, y = 4, 5
	arr := []*int{&x, &y}
	fmt.Println(arr)
}

func rangeArr() {

	var arr = []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//使用new 关键字，返回的是数组的指针
func newarr() {

	var a = [5]int{}
	a[1] = 2
	fmt.Println(a)

	p := new([5]int) //可以通过下标赋值
	p[1] = 2
	fmt.Println(p)

}
