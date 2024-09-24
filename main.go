package main

func main() {
	/*result := Ft_coin([]int{1, 2, 5}, 11) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(result)
	result = Ft_coin([]int{2}, 3) // resultat : -1
	fmt.Println(result)
	result = Ft_coin([]int{1}, 0) // resultat : 0
	fmt.Println(result)*/

	/*result := Ft_missing([]int{3, 1, 2}) // resultat : 0
	fmt.Println(result)
	result = Ft_missing([]int{0, 1}) // resultat : 2
	fmt.Println(result)
	result = Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) // resultat : 8
	fmt.Println(result)*/

	/*result := Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	fmt.Println(result)
	result = Ft_non_overlap([][]int{{1, 2}, {2, 3}}) // resultat : 0
	fmt.Println(result)
	result = Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}) // resultat : 2
	fmt.Println(result)*/

	/*result := Ft_profit([]int{7, 1, 5, 3, 6, 4}) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(result)
	result = Ft_profit([]int{7, 6, 4, 3, 1}) // resultat : 0
	fmt.Println(result)*/

	/*result := Ft_max_substring("abcabcbb") // resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(result)
	result = Ft_max_substring("bbbbb") // resultat : 1
	// "b" est la plus grande sous chaine
	fmt.Println(result)*/

	/*result := Ft_min_window("ADOBECODEBANC", "ABC") // resultat : "BANC"
	fmt.Println(result)
	result = Ft_min_window("a", "aa") // resultat : ""
	fmt.Println(result)*/
}

/*func Ft_min_window(s string, t string) string {
	...
}*/

func Ft_max_substring(s string) int {
	var checked []rune
	var combo []int
	maxLength := 0
	for i, char := range s {
		checked = append(checked, char)
		for x := i + 1; x < len(s); x++ {
			isAlreadyChecked := false
			for _, c := range checked {
				if c == rune(s[x]) {
					isAlreadyChecked = true
					break
				}
			}
			if isAlreadyChecked {
				combo = append(combo, len(checked))
				if len(checked) > maxLength {
					maxLength = len(checked)
				}
				checked = []rune{}
				break
			}
			checked = append(checked, rune(s[x]))
		}
		if len(checked) > maxLength {
			maxLength = len(checked)
		}
	}
	return maxLength
}

func Ft_profit(prices []int) int {
	var profits []int
	for i, j := range prices {
		for x := i; x < len(prices); x++ {
			if j >= prices[x] {
				continue
			}
			profits = append(profits, prices[x]-j)
		}
	}
	if len(profits) < 1 {
		return 0
	}
	bestProfit := 0
	for i := bestProfit; i < len(profits); i++ {
		if profits[i] > bestProfit {
			bestProfit = profits[i]
		}
	}
	return bestProfit
}

func Ft_non_overlap(intervals [][]int) int {
	result := 0
	var checked []int
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		badInterval := false
		for y := start; y < end; y++ {
			for _, j := range checked {
				if j == y {
					badInterval = true
					break
				}
			}
			if badInterval {
				break
			}
			checked = append(checked, y)
		}
		if badInterval {
			result++
		}
	}
	return result
}

func Ft_missing(nums []int) int {
	result := 0
	has := true
	for has {
		y := result
		for i := 0; i < len(nums); i++ {
			if result == nums[i] {
				result++
			}
		}
		if result == y {
			has = false
		}
	}
	return result
}

func Ft_coin(coins []int, amount int) int {
	if len(coins) < 1 {
		if amount == 0 {
			return 0
		} else {
			return -1
		}
	}
	result := 0
	coinsNum := 0
	i := len(coins) - 1
	for y := i; y > -1; y-- {
		for result+coins[y] <= amount {
			result = result + coins[y]
			coinsNum++
		}
	}
	if result != amount {
		return -1
	}
	return coinsNum
}
