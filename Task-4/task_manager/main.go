package main

// import "github.com/gin-gonic/gin"
import "fmt"

func main() {
	// r := gin.Default()

	// r.GET("/")

	// r.Run(":8080")

	var i interface{}

	i = 42
	fmt.Printf("%v, %T\n", i, i)

	i = "hello"
	fmt.Printf("%v, %T\n", i, i)

	i = true
	fmt.Printf("%v, %T\n", i, i)
}
