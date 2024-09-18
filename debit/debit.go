package main 
import(
	"fmt"

)
func debit(num int) int{
	var amount int
	fmt.Println("enter the amount:")
	fmt.Scanln(&amount)
	
	return num + amount
}
func main(){
	var name string
	var accountNo int
	fmt.Println("enter your name:")
	fmt.Scanln(&name)
	fmt.Println("enter your accountNo:")
	fmt.Scanln(&accountNo)
	
	initialAmount:= 1000
	result :=debit(initialAmount)
	fmt.Printf("the balance is %d\n",result)

	


}