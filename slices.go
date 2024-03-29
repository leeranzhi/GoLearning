package main

import "fmt"

//参数数组  不加长度  即为slices（切片）
func updateSlice(s []int) {
	s[0] = 100
}

//数组的切片
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//slices本身是没有数据的，是对底层array的view（视图）

	//左包括但右不包括
	//[2,3,4,5]
	fmt.Println("arr[2:6]=", arr[2:6])
	//[0,1,2,3,4,5]
	fmt.Println("arr[:6]=", arr[:6])
	//[2,3,4,5,6,7]
	s1 := arr[2:]
	fmt.Println("s1=", s1)
	//[0,1,2,3,4,5,6,7]
	s2 := arr[:]
	fmt.Println("s2=", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("s2=", s2)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	//slices之上再次进行slice
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//slice的扩展
	fmt.Println("Extending slice")
	fmt.Println("---arr---", arr)
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	//slice的扩展
	s2 = s1[3:5]

	//可以向后扩展，不可向前扩展
	//s[i]不可以超越len[s]，向后扩展可超越，但不可超过底层数组cap[s]

	//s1[4]提示出错，访问不到
	//fmt.Println(s1[4])
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))

	fmt.Println("------")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5", s3, s4, s5)
	//s4 and s5 no longer view arr.
	//超越cap时，系统会重新拷贝到更大的底层数组
	//原来的数组将会被回收
	//由于值传递的关系，必须接收append的返回值
	fmt.Println("arr=", arr)
}
