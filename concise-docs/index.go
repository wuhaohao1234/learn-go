// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello World");
// }\


// package main

// import (
// 	"fmt"
// 	// "net/http" //Error，因为下面没用到
// )

// const a = 3 //虽然没用但是可以
// var b = 4 //同上

// func main() {
// 	var str string //Error，因为下面没用到
// 	fmt.Println(str)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 3
// 	if a == 1 {
// 	//true
// 		fmt.Println("成功")
// 	} else if a == 2{
// 		fmt.Println("失败")	
// 	} else {
// 		fmt.Println("成功")
// 	}
// }

//第一种
package main

import (
	"fmt"
)

func main() {
	c,d := func (a int,b int)(c int,d int) {
		c = a + b
		d = 4
		return
	}(1,2)
	fmt.Println(c,d)
}
