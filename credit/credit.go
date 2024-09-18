package main 
import(
	"fmt"
)


func credit(num int)int{
	var amount int
	fmt.Println("enter the amount:")
	fmt.Scanln(&amount)
	if num <= amount{
		return num
		
	}else{
		return num - amount
	}
	
	}

func main(){
	var name string
	var accountNo int
	fmt.Println("enter your name:")
	fmt.Scanln(&name)
	fmt.Println("enter your accountNo:")
	fmt.Scanln(&accountNo)
	initialAmount:= 1000
	result :=credit(initialAmount)
	fmt.Printf("the balance is %d\n",result)


	


}