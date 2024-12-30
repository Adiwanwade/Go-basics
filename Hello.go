// It is complied language
package main
import ("fmt")

// func main(){
	// var num=2
	// var num int
	// var num ,num2 int 
	// var num int =2  =>another way of declaring variable
	// fmt.Print(num)
	// num=2
    // num2=3
// num ,num2=2,3
// 	var result=num+num2

// result:=9 // var result=9
// const num=2 // constant variable
// num2:=2
// result,result2:=cal(num,num2)
// 	fmt.Println(result,result2)
	// fmt.Print("\n")
// }
// func cal (num int,num2 int)(int ,int){
// 	var out =num+num2
// 	var out2  =num-num2
// 	return out,out2
// }

// func cal (num int,num2 int)(out,out2 int){
// 	 out =num+num2
// 	 out2  =num-num2
// 	return 
// }


// for loop

// func main(){
// 	for i := 0; i < 3; i++ {
		
// 		fmt.Println("Ganesha")
// 	}
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i)
// 	}
// }

// exported names
// use capital letter to export the function

// math package
// import ("math"
// "fmt"
// )
// func main()  {
// 	var num float64=12
// 	// fmt.Println(math.Sqrt(num))
// 	fmt.Printf("The output is %.2v",math.Sqrt(num))
// 	// fmt.Printf("The output is %.2g",math.Sqrt(num)) // Same result as %.2v
// 	fmt.Printf("The output is %.2f",math.Sqrt(num)) // Don't round off the value


// }

func main(){
	// if math.IsNaN(2){
	// 	fmt.Println("The number is not a number")
	// } else {
	// 	fmt.Println("The number is a number")
	// }

	// switch statement

	// switch num:=2; num {
	// case 1:
	// 	if num > 2 {
	// 		fmt.Println("The number is greater than 2")
	// 	}
	// case 2:
	// 	fmt.Println("The number is 2")
	// default:
	// 	fmt.Println("The number is not 1 or 2")
	// }


	// defer statement
	
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i) // Will work after a function completes its execution
	// }

	// fmt.Println("Bye..")


	// Structure in Go

	type person struct {
		name string
		age int
		salary int
	}
	var p person=person{name:"Ganesha",age:22,salary:10000}
	fmt.Println(p)
}