package loops

import (
	_ "fmt"
)

//1 check daigonal sum is equal from both sides in matrix
type Matrix [][]int

func MultidimenssionalArray() {
	givenMatix := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := checkdaigonalSum(givenMatix)
	if result {
		println("matrix is daigonal sum is equal")
	} else {
		println("matrix is  daigonal sum is not equal")
	}
	givenMatix2 := Matrix{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}
	result2 := checkdaigonal(givenMatix2)
	if result2 {
		println("Martix is Daigonal")
	} else {
		println("matrix is not daigonal")
	}

}
func checkdaigonal(matrix Matrix) bool {
	var counter int = 0
	for row, i := range matrix {
		for column, j := range i {
			if row == column && j != 0 {
				continue
			} else if j == 0 {
				counter++
			}

		}
	}
	if counter == 6 {
		return true
	} else {
		return false
	}
}

func checkdaigonalSum(givenMatix Matrix) bool {
	var leftDaigonalSum = 0
	var rightDaigonalSum = 0
	for row, i := range givenMatix {
		for column, j := range i {
			if row == column {
				rightDaigonalSum = rightDaigonalSum + j
			}
			if row+column == 2 {
				leftDaigonalSum = leftDaigonalSum + j
			}
		}

	}
	if leftDaigonalSum == rightDaigonalSum {
		return true
	} else {
		return false
	}

}
