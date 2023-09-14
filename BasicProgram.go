//Program to find Factorial of number
package main
import "fmt"

/* Variable Declaration */
var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
                       // Range: 0 through 18446744073709551615. 
var i int = 1
var n int

/*     function declaration        */
func factorial(n int) uint64 {   
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
            factVal *= uint64(i)  // mismatched types int64 and int
        }
        
    }    
    return factVal  /* return from function*/
}

func check_prime(n int) bool {   
    for i:=2; i<=n/2; i++ {
		if(n%i==0){ return false}
	}
	return true
}

func sum_of_fib(n int) uint64 {

	var ans uint64= 1

	var arr = make([]int, n+1)
	arr[1] = 1;
	for i:=2;i<=n;i++ {
		arr[i] = arr[i-1] + arr[i-2]
		ans+= uint64(arr[i])
	}
	return ans
}

func main(){    
    // fmt.Print("Enter a positive integer between 0 - 50 : ")
    // fmt.Scan(&n)   
    // fmt.Print("Factorial is: ",factorial(n))

	
    // fmt.Print("Enter a positive integer to check for prime : ")
    // fmt.Scan(&n)   
	// if (check_prime(n)) {
	// 	fmt.Println("The no is Prime")
	// } else {
	// 	fmt.Println("The no is not Prime")
	// }

	// fmt.Print("Enter a positive integer : ")
    // fmt.Scan(&n)  
	// fmt.Println("The Sum of fib up to the given no is ", sum_of_fib(n))

	fmt.Print("Enter size of array: ")
    fmt.Scan(&n)  
	fmt.Println("Enter array elements")
	var arr = make([]int, n)
	even:=0
	odd:=0
	zero:=0
	for i:=0;i<n;i++ {
		fmt.Scan(&arr[i])
		if(arr[i] == 0) {
			zero++
		} else if(arr[i]%2 == 0) {
			even++
		} else {
			odd++
		}
	}



	fmt.Println("Count of Evens : ", even)
	fmt.Println("Count of Odds : ", odd)
	fmt.Println("Count of Zeros : ", zero)

	


}