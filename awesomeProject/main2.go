package main
import (
	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {

	router:=gin.Default()
	router.GET("/haha",Test)
    router.POST("/testPost",Test2)
	router.Run(":8899")
}

func Test(c *gin.Context){
	fmt.Println("in")
	username:=c.Query("username")
	password:=c.Query("password")
	fmt.Println("data",username,password)
	c.JSON(http.StatusOK,gin.H{"message":"hello"})
}

func Test2(c *gin.Context){
	fmt.Println("in")
	var user Data
    c.BindQuery(user)
	fmt.Println("data:",user[0],user[1])
	c.JSON(http.StatusOK,gin.H{"message":"hello"})
}

