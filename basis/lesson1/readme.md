# basis(基础)

## packages,variable and function(包，变量，函数)

### packages(包)

每个 Go 程序都是由包组成的。

程序运行的入口是包 `main`。

这个程序使用并导入了包 "fmt" 和 `"math/rand"`。

按照惯例，包名与导入路径的最后一个目录一致。例如，`"math/rand"` 包由 package rand 语句开始。

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

```

该程序会输出一个非负伪随机数
[rand.Intn](https://golang.org/pkg/math/rand/#Intn)

### import(导入)

```
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}
```

import () 为导入语句

当然也可以写为

```
import "fmt"
import "math"
```
该程序返回一个2到3的小数

### export-names(导出名)

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(10))
}
```

### function(函数)

函数可以没有参数或接受多个参数。

```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

这个函数接收两个整形,返回值也是个整形


#### function Shorthand(函数简写)

```
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

x int , y int 可以替换为 x, y int

### multiple-results(多值返回)

函数可以返回任意数量的返回值

```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

a,b := swap("hello","world") 等价于 var a = "world" , b = "hello"

### named-results(命名返回)

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

```

go中的返回值可以被命名,并且可以直接使用,上述代码返回的7与10被直接打印

### variables(变量)

通过var关键字定义变量,后面写类型

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

以上定义了4个变量但是都没有赋值
c,python,java 是布尔值，默认为false
i 是整形,默认是0

变量定义可以包含初始值，每个变量对应一个。

如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。

### short-varabel-declaractions(短声明变量)

```
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

```

在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。

函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。

### basic-types(基本类型)

1. bool布尔值:true&&false
2. string字符串
3. int int8 int16 int64
4. uint uint8 uint16 uint64 uintprt

5. byte unit8的别名

6. rune int32的别名代表一个Unicode码

7. float32 float64

8. complex64 complex128

```
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
```

这个例子演示了具有不同类型的变量。 同时与导入语句一样，变量的定义“打包”在一个语法块中。

### zero(零值)

```
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

分别打印出0 0 false "" 这些定义了没有初始化的打印出零值

### type-conversions(类型转换)

demo

```
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}
```
example
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
and

```
i := 42
f := float64(i)
u := uint(f)
```
与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换。

### type-inference(类型推导)

在定义一个变量但不指定其类型时（使用没有类型的 var 或 := 语句）， 变量的类型由右值推导得出。

当右值定义了类型时，新变量的类型与其相同：

```
var i int
j := i // j 也是一个 int
```

但是当右边包含了未指名类型的数字常量时，新的变量就可能是 int 、 float64 或 `complex128`。 这取决于常量的精度：

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```
example
```
package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
```

### constants(常量)

常量的定义与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用 := 语法定义。

```
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

