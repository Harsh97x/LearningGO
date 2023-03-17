// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("go" + "lang")
// 	fmt.Println("1+1=", 1+1)
// 	fmt.Println("7/3=", 7.0/3.0)
// }

// package main

// import "fmt"

// func main() {
//   fmt.Println("This" + "is" + "a" + "tutorial")
// }

// package main

// import "fmt"

// func main() {
// 	var a = "initial"
// 	fmt.Println(a)

//   var b,c int= 1,2
//   fmt.Println(b,c)

//   f := "apple"
//   fmt.Println(f)
// }

package main

import (
	"fmt"
	"math"
)

const s string = "Constant"

func main (){
	fmt.Println(s)
	
	const a = 500000

	const n = 5000000

	fmt.Print(a)

	const c = 3e20/n 

	fmt.Println(c)

	fmt.Println(int64(c))

	fmt.Println(math.Sin(n))
}
