package loops

import (
	"fmt"
	_ "fmt"
)

func SumOfDigits() {
	var num int
	fmt.Scanf("%d",&num)
  var sum int =0
	var  copynum int = num
  var lastDigit int
	var sum2 int
	for num>9 {
		for num !=0{
			lastDigit = num%10
			sum = lastDigit + sum
			num = num/10

		}
		sum2=sum
		num=sum
		sum=0
	}

println("Given digit ",copynum,"Sum = ",sum2)
}
