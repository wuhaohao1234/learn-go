# 从0开始

* Go没有char，所有的字符都用rune（跟int是一个）代替，输出需要手动转string

* Go语言的大括号必须放在一句话的后面

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World");
}
```

Go语言只有一个入口，main包下的main函数

* Go语言函数内不能声明不用的变量和包名，编译器会报错，但是在函数外面是可以的

```
package main

import (
	"fmt"
	"net/http" //Error，因为下面没用到
)

const a = 3 //虽然没用但是可以
var b = 4 //同上

func main() {
	var str string //Error，因为下面没用到
	fmt.Println("Hello World")
}
```

* Go语言不用写分号除非你想在一行里面写多条语句

* Go不支持静态变量，并没有static

* Go语言的变量和函数的公有私有靠首字母大小写区分，首字母大写是共有的，小写是私有的

```
fmt.Println(str); //这是一个公有的函数，因为P大写了，变量也是一样的
```

* Go语言的定义变量，变量类型写变量后面

```
package main

import (
	"fmt"
)
func main() {
	var str string = "Hello World"
	fmt.Println(str);
}
```

* Go没有try/catch，finally的功能靠defer关键字实现 （稍后解惑）

* Go函数可以有多个返回值 （稍后解惑）

## 变量定义

```
const IP = "127.0.0.1" //常量用const
var str1 string = "small"
str2 := "solo" //这里go会自动解析出类型
var str3 = "smallSohoSolo" //同上
```

这些定义的方法都是等价的

```
str2 := "solo"
```

只能在函数内部定义，不能在函数外部

```
package main

import (
	"fmt"
)

global := "small" //这个不行！！！！！！！

func main() {
	fmt.Println("Hello World")
}
```

还有更爽的方式

```
str1,str2,str3 := 1,2,3
var str1,str2,str3 = "1","2","3"
var str1,str2,str3 string = "1","2","3"
```

以后很地方会用到，函数也可以一次返回多个返回值

PS:int在32位机器当前都是32位，未来可能在64位机器上变成64位，如果要指定准确的字长，那就用int8，int16等手动指定

另外不同类型的值是不能相加的（吐槽一下Java，Java会自己默默地转然后error = =）非常强的类型比对，非常适合大型企业项目

## 流程控制

if条件句

```
a := 1
if a == 1 {
  //true
} else if a == 2{
  
} else {
  
}
```

另外if支持在条件句中申明一个变量

```
if a:= 1; a == 1 {
  //true
}
```

for循环

```
for i := 0; i < 10; i++ {
  //循环体
}
```

另外并没有while，用for解决

```
i := 0
for i != 10 {
  i++;
}
```
for的遍历写法，下面是遍历一个map和array

```
for k,v := range map { //循环体 } list := []int{1,2,3,4,5} for index,val := range list { fmt.Println(index,val) }
```

goto走一波

```
i := 1
Here:
if i == 1 {
  i = 2
  goto Here
}
```
switch不用break自己就中断了，想要接着执行就
```
var i = 3
switch i {
  case 1,3:
     fmt.Println("123")
     fallthrough
  case 2:
     fmt.Println("789")
  default:
     fmt.Println("456")
}
```

函数

函数可以有多个返回值，支持闭包，func后面跟的第一个括号是入参，第二个括号是返回值，可以有多个哦

```
//第一种
c,d := func (a int,b int)(c int,d int) {
		c = a + b
		d = 4
		return
	}(1,2)
fmt.Println(c,d)
//第二种
c := func (a int,b int)(c int) {
  return a +b
}(1,2)
fmt.Println(c)
//第三种
c := func (a int,b int)(int) {
  return a + b
}(1,2)
fmt.Println(c)
//第四种
c,d := func (a int,b int)(int,int) {
  return a,b
}(1,2)
fmt.Println(c,d)
```

*和&

两者同c一样

* *是代表指针，也可以从地址中获取内容

* &获取一个对象的地址

```
i := 1
point := &i // point等于一个地址
data = *point //data == 1
```

另外函数中要分清楚值传递和指针传递，这是一个交换函数，很清晰

```
func change(a *int,b *int) {
  c := *a
  *a = *b
  *b = c
}
a := 3
b := 4
func(&a,&b)
```

## 结构体

```
type SmallSoho struct {
	Name string //公有变量大写
	Sex string 
	kg int //私有变量小写
}
```

## 如何申请一个结构体

```
meStruct := new(SmallSoho) //可以申请一个空的结构体，返回一个指针
meStruct := &SmallSoho{} //同上
meStruct := SmallSoho{} //返回一个结构体，是类型的
meStruct := &SmallSoho{"small","男"，1} //可以初始化
meStruct := &SmallSoho{Name:"SmallSoho"} //也可以键值对这样来初始化
```

```
meStruct := &SmallSoho{"small","男"，1} //可以初始化
meStruct := &SmallSoho{Name:"SmallSoho"} //也可以键值对这样来初始化
```

## 数据结构

数组，Slice（划分），map，channel（并发编程应用）

数组就是跟其他语言一样的数组，没有什么特别的，并且长度只能写死，不能通过变量来定义

```
list := [10]string{}
list := [10]int{1,2,3,4,5,6,7,8,9,10}
list := []int{1,2,3}
list := [3]int{1} //后面两个默认是0
```

slice是划分，也就是可以通过数组得到slice，其中每一个元素都指向数组的一个元素，也可以通过make来为划分申请内存

```
list := []int{1,2,3}
slice := list[0:3] //左闭右开，不支持负数索引
```

上述情况slice的每一个元素都指向了list，修改任意其一的元素，两者均变化

```
slice := make([]int,4)
```

也可以为slice申请一块内存，这样就可以做到动态数组

map的使用方法同上，是一个映射，具体更多的参数请参考文档

摘要:
https://smallsoho.com/backend/2016/11/20/Go%E7%AE%80%E6%98%8E%E6%95%99%E7%A8%8B/