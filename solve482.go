/*

Problem
Chef is the financial incharge of Chefland and one of his duties is identifying if any company has gained a monopolistic advantage in the market.

There are exactly 33 companies in the market each of whose revenues are denoted by R_1R
1
​
 , R_2R
2
​
  and R_3R
3
​
  respectively. A company is said to have a monopolistic advantage if its revenue is strictly greater than the sum of the revenues of its competitors.

Given the revenue of the 33 companies, help Chef determine if any of them has a monopolistic advantage.

*/

package main

import (
	"fmt"
)

func main() {
	var cases int
	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		var val int
		fmt.Scan(&val)
		var val2 int
		fmt.Scan(&val2)
		var val3 int
		fmt.Scan(&val3)
		li := []int{val, val2, val3}
		max := li[0]
		idx := 0
		for i := 0; i < len(li); i++ {
			if li[i] >= max {
				max = li[i]
				idx = i
			}
		}
		sum := 0
		for i := 0; i < len(li); i++ {
			if i == idx {
				sum = sum + 0
			} else {
				sum = sum + li[i]
			}
		}

		if sum < max {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
