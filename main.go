package main

import (
	"fmt"
	"os"

	"Goproject/leetcode/List"
	"Goproject/leetcode/array"
	"Goproject/leetcode/backtrace"
	"Goproject/leetcode/binary"
	cache2 "Goproject/leetcode/cache"
	"Goproject/leetcode/commonstring"
	"Goproject/leetcode/daily"
	"Goproject/leetcode/dp"
	"Goproject/leetcode/packagetest"
	"Goproject/leetcode/queue"
	"Goproject/leetcode/redundent"
	"Goproject/leetcode/reverse"
	"Goproject/leetcode/roman"
	"Goproject/leetcode/stack"
	"Goproject/leetcode/stock"
	"Goproject/leetcode/twosum"
)

func main() {
	//a := make([]int, 5, 10)
	//fmt.Println(a)
	//set(a)
	//fmt.Println(a)


	//fmt.Println("hello world!")
	//Twosum()
	//Reverse()
	//LongSubString()
	//LongPalindrome()
	//romanToInt()
	//longgestCommon()
	//maxArea()
	//isValidBrackets()
	//threeSum()
	//removeRedun()
	//removeElement()
	//threeSumClosest()
	//numberOfLIS()

	//subsets()
	//subsetsUnique()
	//appendTrick()
	//generateParenthesis()
	//nextPermutation()
	//binarySearch()
	//search()
	//searchRange()
	//palindrome()
	//poolTest()
	//isValidSudoku()
	//combinationSum()
	//combinationSumOnce()
	//maxQueue()
	//evalRPN()
	//reverseList()
	//findPeakElement()
	//permute()
	//permuteUnique()
	//combine()
	//coinChange()
	//maxProfit()
	//minDistance()
	//gcdOfStrings()
	//maxSlidingWindow()
	//lruCache()
	//frequency()
	//drink()
	//arraySum()
	//decode()
	//countCharacters()
	//maxPackageValue()
	//longestPalindrome()
	//getLeastNumbers()
	//canMeasureWater()
	compressString()
}

func set(a []int) {
	a[0] = 1
	a[4] = 2
	a = append(a, 1,2)
}

func Twosum() {
	//twosum
	nums := []int{3,2,4}
	target := 6

	twosum.Sum(nums, target)
}

func Reverse() {
	x := 12121

	//fmt.Println(reverse.IntReverse(x))
	fmt.Println(reverse.IsPalindrome(x))
}

func reverseList() {
	p5 := &List.ListNode{Val: 5, Next: nil}
	p4 := &List.ListNode{Val: 4, Next: p5}
	p3 := &List.ListNode{Val: 3, Next: p4}
	p2 := &List.ListNode{Val: 2, Next: p3}
	p1 := &List.ListNode{Val: 1, Next: p2}

	head := List.ReversePart(p1, 2,4)
	q := head
	for {
		if q == nil {
			break
		}

		fmt.Println(q.Val)
		q = q.Next
	}
}

func LongSubString() {
	s := "acqacqcc"

	fmt.Println(redundent.LengthOfLongestSubstring(s))
}

func LongPalindrome() {
	//s := "babad"
	s := "aaba"

	fmt.Println(dp.LongestPalindrome(s))
}

func romanToInt() {
	s := "MCMXCIV"
	fmt.Println(roman.RomanToInt(s))
}

func longgestCommon() {
	s := []string{"flower","flow","flight"}
	fmt.Println(commonstring.LongestCommonPrefix(s))
}

func maxArea() {
	s := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(array.MaxArea(s))
}

func isValidBrackets() {
	s := "([)]"
	fmt.Println(stack.IsValid(s))
}

func threeSum() {
	n := []int{0,0,0}
	fmt.Println(array.ThreeSum(n))
}

func removeRedun() {
	n := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(array.RemoveDuplicates(n))
}

func removeElement() {
	n := []int{0,1,2,2,3,0,4,2}
	val := 2
	fmt.Println(array.RemoveElement(n, val))
}

func threeSumClosest() {
	n := []int{1,1,1,1}
	val := 0
	fmt.Println(array.ThreeSumClosest(n, val))
}

func numberOfLIS() {
	n := []int{1,3,5,4,7}
	//n := []int{2,2,2,2,2}
	fmt.Println(dp.NumberOfLIS(n))
}

func subsets() {
	n := []int{1,2,3}
	//fmt.Println(dp.Subsets(n))
	fmt.Println(backtrace.Backtrace(n))
}

func subsetsUnique() {
	n := []int{1,2,2}
	fmt.Println(backtrace.Backtrace2(n))
}

func generateParenthesis() {
	n := 3
	//fmt.Println(dp.GenerateParenthesis(n))
	fmt.Println(backtrace.GenerateParenthesis(n))
}

func nextPermutation() {
	n := []int{3,2,1}
	array.NextPermutation(n)
}

func binarySearch() {
	n := []int{-1,0,3,5,9,12}
	target := 9
	fmt.Println(array.BinarySearch(n, target))
}

func search() {
	n := []int{5,1,3}
	target := 3
	fmt.Println(array.Search(n, target))
}

func searchRange() {
	n := []int{2}
	target := 2
	fmt.Println(array.SearchRange(n, target))
}

func palindrome() {
	s := "dbbacc"
	fmt.Println(array.Palindrome(s))
}

func poolTest() {
	packagetest.PoolTest(os.Stdout, "dablelv", "monkey")
}

func isValidSudoku() {
	array.IsValidSudoku(nil)
}

func combinationSum() {
	n := []int{2,3,5}
	target := 8
	fmt.Println(array.CombinationSum(n, target))
}

func combinationSumOnce() {
	n := []int{2,3,5}
	target := 8
	fmt.Println(array.CombinationSumOnce(n, target))
}

func maxQueue() {
	mq := queue.Constructor()

	mq.Push_back(10)
	mq.Push_back(11)
	mq.Push_back(12)
	mq.Push_back(11)
	mq.Push_back(10)
	mq.Push_back(9)
	mq.Push_back(13)
	fmt.Println(mq.Max_value())
}

func evalRPN() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(stack.EvalRPN(tokens))
}

func findPeakElement() {
	n := []int{1,2,1,3,5,6,4}
	fmt.Println(binary.FindPeakElementFor(n))
}

func permute() {
	n := []int{1,2,3}
	fmt.Println(backtrace.Permute(n))
}

func permuteUnique() {
	n := []int{1,1,2}
	fmt.Println(backtrace.PermuteUnique(n))
}

func combine() {
	n := 4
	k := 2
	fmt.Println(backtrace.Combine(n, k))
}

func coinChange() {
	//n := []int{186,419,83,408}
	//target := 6249
	n := []int{2}
	amount := 3
	fmt.Println(dp.CoinChange(n, amount))
}

func maxProfit() {
	//n := []int{7,1,5,3,6,4}
	//n := []int{1,2,3,4,5}
	n := []int{1,2,3,0,2}
	//fmt.Println(stock.MaxProfitOnce(n))
	//fmt.Println(stock.MaxProfitUnlimit(n))
	//fmt.Println(stock.MaxProfixWithK(n))
	fmt.Println(stock.MaxProfitWithCooldown(n))
}

func minDistance() {
	word1 := ""
	word2 := "a"

	fmt.Println(dp.MinDistance(word1, word2))
}

func gcdOfStrings() {
	str1 := "ABABAB"
	str2 := "ABAB"
	fmt.Println(daily.GcdOfStringsOptimize(str1, str2))
}

func maxSlidingWindow() {
	//n := []int{1,3,-1,-3,5,3,6,7}
	//k := 3
	n := []int{1,-1}
	k := 1
	fmt.Println(dp.MaxSlidingWindowOptimiz(n, k))
}

func lruCache() {
	cache := cache2.Constructor(2)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Put(2, 6))
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Put(1,5))
	//cache2.Print()
	fmt.Println(cache.Put(1,2))
	//cache2.Print()
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	//cache2.Print()
}

func countCharacters() {
	//words := []string{"cat","bt","hat","tree"}
	//chars := "atach"
	words := []string{"hello","world","leetcode"}
	chars := "welldonehoneyr"
	fmt.Println(array.CountCharacters(words, chars))
}

func maxPackageValue() {
	w := 4
	weights := []int{2, 1, 3}
	values := []int{4, 2, 3}
	fmt.Println(dp.MaxPackageValue(w, weights, values))
}

func longestPalindrome() {
	s := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"
	fmt.Println(daily.LongestPalindromeOptimize(s))
}

func getLeastNumbers() {
	arr := []int{4,5,1,6,2,7,3,8}
	k := 4
	fmt.Println(array.QuickSort(arr, k))
}

func canMeasureWater() {
	x := 2
	y := 6
	z := 5
	fmt.Println(daily.CanMeasureWater(x, y, z))
}

func compressString() {
	S := "abbccdKKKK"
	fmt.Println(daily.CompressString(S))
}