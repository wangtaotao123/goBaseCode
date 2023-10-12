package main

import (
	"fmt"
)

func main() {
	/*
		1.回调函数：可以将一个函数作为另一个函数的参数
		将fun1函数作为fun2函数的参数，那么fun2函数，就叫做高阶函数
		而fun1函数，就叫做回调函数，作为另一个函数的参数
	*/
}

func main29() {
	//匿名函数：就是没有名字的函数
	//引用类型赋值
	f05 := f04 //两者同一个地址
	//调用
	f05()
	//匿名函数 的1种调用方法，就是函数名 + （），那么直接再定义的时候就可以 + （），进行调用
	f06 := func() {
		fmt.Println("匿名函数赋值给f06")
	}
	f06()
	//只执行一次，自己调用了自己
	func() {
		fmt.Println("匿名函数赋值给f07，直接+ （）")
	}()
	//自己调用自己，带参数
	func(a, b int) {
		//打印形参的两个值
		fmt.Println(a, b)
		fmt.Println("匿名函数赋值给f08，直接+ （）")
	}(10, 20) //直接再这里将实参进行传递
	//自己调用自己，带参数 ，带返回值
	num := func(a, b int) int {
		//打印形参的两个值
		fmt.Println(a, b)
		fmt.Println("匿名函数赋值给f09，直接+ （）")
		return a + b
	}(10, 20) //直接再这里将实参进行传递
	fmt.Println("来自匿名函数的返回值num = ", num)
}
func f04() {
	fmt.Println("i am f04()")
}
func main28() {
	/*
		1.函数的本质



	*/
	//	打印数据类型
	var a int = 10
	fmt.Printf("%T\n", a)
	//这里f02不带（），因为带了（），那就是函数调用了
	//如果不加（），那函数名就是一个变量
	fmt.Printf("%T\n", f02)
	//	函数也是一种类型，定义一个函数类型
	var f03 func(int)
	//将函数f02的功能赋给f03,f03就可以进行操作了
	//而且，其是引用数据类型，
	f03 = f02
	f03(300)
}
func main27() {
	//延迟函数 ：**先将参数进行传递**，然后保持再栈中，等待最后执行
	//后续还会使用defer来进行关闭操作
	a := 10
	fmt.Println("a=", a)
	defer f02(a)
	a++
	fmt.Println("end a= ", a)
	//
}
func f02(a int) {
	fmt.Println("f02 func a= ", a)
}
func main26() {
	//	延迟函数，延迟到块的最后执行；而且  defer会将其压入栈中，
	//  可以看见789被先压入栈中，然后再是1，所以再进行出栈
	fPrint("123")
	fmt.Println("456")
	//将会延迟本行，在块之后执行
	defer fPrint("789")
	fPrint("101112")
	defer fPrint("1")
	fPrint("2")
}
func fPrint(s string) {
	fmt.Println(s)
}

func main25() {
	//	递归函数，自己调用自己，要有出口
	//	累加递归
	fmt.Println(getSum02(10))

}
func getSum02(num int) int {
	if num == 1 {
		return 1
	}
	return getSum02(num-1) + num
}
func main24() {
	//	变量作用域：局部变量 ---- 就近原则
	//	全局变量在  func外面，全部的函数都可以调用

}
func main23() {
	//	切片,可以扩容的数组   也就是说切片是引用类型的传递
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("默认 ", s1)
	update2(s1)
	fmt.Println("调用后", s1)
}

// 传递进来切片类型
func update2(s2 []int) {
	s2[0] = 100
}

func main22() {
	//	参数传递
	/*
		1.值传递：操作的是数据本身 int,string,bool,float,array
		2.引用类型传递 :操作的是数据的地址 slice（切片类型） ,map,chan ...

	*/
	//	1.数组值传递
	//定义数组 [个数]类型
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr[1])
	fmt.Println("函数调用，数组是值传递")
	updateArry(arr)
	//	从这里可以看出，数组在形式参数时是值传递，
	fmt.Println(arr[1])
	//在外面就可以进行改变
	arr[1] = 200
	fmt.Println(arr[1])
}

// 注意的是，形式参数的数据的大小必须固定
func updateArry(arr2 [4]int) {
	arr2[1] = 100
}

func main21() {
	getSum(1, 2, 3, 4)
}

// ...可变参数  ,可变参数在最后，一个函数列表中，只能有一个可变参数
func getSum(nums ...int) {
	//得到参数个数
	fmt.Println(len(nums))
	//	遍历里面的值
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d,", nums[i])
	}

}

func main20() {
	//	1.无参无返回值
	printInfo01()
	//2.有参无返回值的
	//3.无参有返回值--有1个返回值,有多个返回值的
	//4.有参有返回值的
	fmt.Println(printInfo02())
}
func printInfo01() {
	fmt.Println("无参无返回值")
}
func printInfo02() (int, int) {
	a, b := 10, 12
	return a, b
}
func main19() {
	/*
		函数

	*/
	//	函数调用
	var b int
	var m float64 = 12.1

	b = 2
	//c = add(a, b)
	//fmt.Println("c = ", c)

	m = add(m, b)
	fmt.Println("c = ", m)
}

/*
//add函数

	   通过func来定义函数 函数名，变量，变量类型，返回值类型
			func add(a, b int) int {
				c := a + b
				return c
			}
*/
func add(a float64, b int) float64 {
	var c float64 = 0.0
	c = a + float64(b)
	return c
}
func main18() {
	//go 中没有while
	//	遍历string
	str := "hello go"
	fmt.Println("字符串的长度为: ", len(str))
	// str[0] 将会输出ascii编码的数字，
	fmt.Printf("0号字符: %c\n", str[0])
	//	遍历string
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
	//	for range循环遍历　返回两个下标ｉ，以及下标下的字符v
	for i, v := range str {
		fmt.Printf("str[%d] = %c ", i, v)
	}

}

func main17() {
	//	break 结束循环体;continue 结束当此循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			//break
			continue
		}
		fmt.Println(i)
	}

	//fmt.Println("调用")
	//main16()
}
func main16() {
	/*	打印5*5方阵
		for i := 1; i <= 5; i++ {
			for j := 1; j <= 5; j++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}

	*/
	//	打印9*9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
		}
		fmt.Println()
	}

}
func main15() {
	//for循环
	//下面是快速定义
	//i := 10
	//fmt.Println(i)
	//下面循环输出了9次
	//起始值；循环条件；控制变量
	//for i := 1; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//	计算1-10的累加
	sum := 0
	for i := 1; i <= 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

}
func main14() {
	//	switch
	/*
		这里不同与c，只要运行完了对应case就直接跳出了，不在往下执行下
		然后都没有case的情况下，进入default
		可以使用fallthrough来进行case穿透：即匹配到case后，仍然向下进行。
		只穿透一次
	*/
	var a int = 30
	switch a {
	case 10:
		fmt.Println("10")

	case 20:
		fmt.Println("20")
		//fallthrough
	case 100:
		fmt.Println("100")
	case 30, 40, 50:
		fmt.Println("30,40,50,多个中的一个")
		//fallthrough
	default:
		fmt.Println("default!")
	}
}

func main13() {
	/*	循环控制
		1.顺序
		2.选择
		if
		switch
		select,后面的channel再讲
		3.循环结构：
		for
	*/
	/*
			var a int = 20

			if a < 12 {
				fmt.Println("<12")
			} else if a >= 12 && a < 20 {
				fmt.Println("12<=a<20")
			} else {
				fmt.Println("a>=20")
			}


		//	if嵌套
		var a int = 100
		var b int = 200

		if a > 20 {
			if b > 100 {
				fmt.Println("a and b is true!")
			}
		}
	*/
	//	用户名密码登陆验证
	var name string = "wangtaotao"
	var pwd string = "123"
	var inputA string
	fmt.Println("input name:")
	fmt.Scanln(&inputA)
	if inputA == name {
		fmt.Println("name is true and please input pwd:")
		fmt.Scanln(&inputA)
		if inputA == pwd {
			fmt.Println("login is success!")
		}
	}

}

func main12() {
	//键盘输入输出
	var x int
	var y float64
	fmt.Println("请输入x,y:")
	//这样输入是对的
	//fmt.Scanln(&x)
	//fmt.Scanln(&y)
	//注意下面这个：输入的中间用空格，而不是，
	//要是用，这样后面那个y就接受不到，
	fmt.Scanln(&x, &y)
	//下面这个是对的
	//fmt.Scanf("%d,%f", &x, &y)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func main11() {
	//	位运算
	var a1 uint = 60
	var a2 uint = 13
	fmt.Printf("a1    = %.8b\n", a1)
	fmt.Printf("a2    = %.8b\n", a2)
	//fmt.Printf("a1&a2 = %.8b\n", a1&a2)
	//	^ 异或运算
	//	移动 --下面这个出现问题了
	//右移动没有问题，高出补0
	fmt.Printf("a1<<3 = %.8b\n", a1>>4)
	//但是左移出现问题:即左移动只到左边第一位为1
	fmt.Printf("a2<<3 = %.8b\n", a2<<5)
	//fmt.Printf("a1>>2 = %.8b\n", a1>>2)

}
func main10() {
	//关系运算符结果位布尔类型
	a := 20
	b := 10
	fmt.Println(a > b)
	/*	逻辑运算符
		    &&   与
			||   或
			!    非

	*/
}

func main09() {
	var num int = 10
	//10%3 --> 1   求模，取余
	fmt.Println(num % 3)
}

// 全局变量：在函数体之外的变量；局部变量在函数体内部的变量，就近原则（有局部就是用局部）
// var a int = 100
func boolFunc() bool {

	var isFlag bool = true
	return isFlag
}
func main08() {
	//数据类型的转换，强显示的
	a := 5
	b := 10.23

	//将低精确度给高精确度
	c := float64(a)
	d := int(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T,%d\n", d, d)
}

func main07() {
	//	定义变量类型
	/*
				一、基本数据类型
					1.布尔类型
					2.数字类型
						1.整型：int,int8,int16,int32,int64,uint,uintptr,uint8,byte,uint16,uint32,uint64
						//byte 类似uint8
		        		//rune 类似int32
						//int  类似int64
						2.浮点型:float32,float64,complex64,complex128
					3.字符串类型


				二、派生数据类型
					1.指针
			        2.数组
					3.结构体
					4.通道（channel）
					5.切片（slice）
					6.函数
					7.接口（interface）
					8.Map

	*/

	//1.bool类型 :默认类型是false
	var isFlag = boolFunc()
	//%T,类型，%t是bool类型的值，%p是地址
	fmt.Printf("bool's Type: %T, bool's 值:%t\n", isFlag, isFlag)
	var age int = 32
	var money float64 = 100.12
	fmt.Printf("age's Type: %T, age's 值:%d\n", age, age)
	//Printf是进行格式化输出的符号%f默认打印6位小数，  %.2f打印2位小数
	//丢失去精度，四舍五入  a = 3.16  %.1f  -->3.2
	fmt.Printf("money's Type: %T, money's 值:%.2f\n", money, money)
	//使用双引号是字符串
	var str1 string = "A"
	//使用单引号  int32类型的。
	/*中国的编码表是gbk
	全世界的是unicode
	*/
	//A -65
	//a -97
	str2 := 'a'
	fmt.Printf("%T,%s\n", str1, str1) //string A
	fmt.Printf("%T,%d", str2, str2)   //int32 a
	//	字符串的拼接 + ,使用反斜杠\来进行转义字符
	fmt.Println("123" + "456 \"这里 \"")

}

func main06() {
	//	特殊的自增常量iota 在同一蹙里面
	const (
		//iota从0开始计数
		a = iota   //0
		b          //1
		c          //2
		d = "haha" //haha iota = 3
		e          //haha iota = 4
		f = 100    //100 iota = 5
		g          //100 iota = 6
		//这里哪怕上述的e等改变了，但是这个iota仍然在计数
		h = iota //7
		i        //8
	)
	//新的族里iota将重新从0计数
	const (
		j = iota
		k
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
	//同一组中只定义了前一个，后一个不定义就沿用上一个的定义
	const (
		a1 = "adsf"
		b1
		c1
		d1
	)
	fmt.Println(a1, b1, c1, d1)

}

// 函数返回两个int类型的变量，注意写在（）外面还有一个（）
func test() (int, int) {
	return 10, 20

}

func main05() {
	//接受函数返回值
	var a, b int = test()
	fmt.Println(a, b)
	//	匿名变量丢弃  _ 这是匿名变量将丢弃获得的值
	num, _ := test()
	fmt.Println(num)
	//	匿名变量不会占用空间，不用有重复命名这一说
	_, age := test()
	fmt.Println(age)
	//函数调用
	//main04()

}
func main04() {
	//	常量，定义后不能被修改。只能是布尔型，数字型（整型，小数，复数）和字符串
	const PAI float32 = 3.1425
	const UTF string = "UTF-8"
	//	隐式定义
	const UTFS = "UTF-8"
	fmt.Println(PAI, UTF, UTFS)
}

func main03() {
	//交换ab
	var a, b int = 100, 200
	fmt.Println(a, b)
	fmt.Println("交换之后")
	//这里进行交换
	a, b = b, a
	fmt.Println(a, b)
}
func main02() {
	//var name string = "wangtaotao"
	var age int = 23
	//格式化打印地址空间 %p  , &age 用到取地址符
	fmt.Printf("age:%d,addAge:%p", age, &age)
}

func main01() {
	//name 变量
	/*
		var name string = "wangtaotao"
		var age int = 23
		//还可以这样打印东西。
		fmt.Println(name, age)
		//后面没有;
		fmt.Println("hello world from go!")
		name = "zhuhuahua"
		fmt.Println("changed :" + name)
		var (
			name string
			age  int
			addr string
		)

	*/

	//没有进行初始化，就是默认值
	//string   : null
	//int      :0
	//fmt.Println(name, age, addr)

	//	偷懒的方式定于变量  自动推到格式   短变量定义
	name := "wangtaotao"
	age := 12
	fmt.Println(name, age)
	//格式化打印他们的格式
	fmt.Printf("%T, %T", name, age)
}
