package main
import "fmt"

func takeInteger(X int, Y int) int{
   return X + Y
}


func takeString(X string, Y string) string{

	return X + Y
}

func main(){

	
  fmt.Println(takeInteger(4,100))

  fmt.Println(takeString("obinna ", "franklin"))


}

