package main

import (
	"Goproject/alibaba"
	"Goproject/leetcode/cache/LFU"
	"Goproject/leetcode/daily/april"
	"Goproject/leetcode/evenweekly"
	"Goproject/leetcode/index"
	"Goproject/leetcode/kmp"
	"Goproject/leetcode/stdsort"
	"Goproject/leetcode/weekly"
	"Goproject/tecent"
	"context"
	"fmt"
	"os"
	"time"

	"Goproject/leetcode/List"
	"Goproject/leetcode/array"
	"Goproject/leetcode/backtrace"
	"Goproject/leetcode/binary"
	cache2 "Goproject/leetcode/cache/LRU"
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
	//compressString()
	//massage()
	//SurfaceArea()
	//numRookCaptures()
	//hasGroupsSizeX()
	//trieTree()
	//minimumLengthEncoding()
	//minIncrementForUnique()
	//findLucky()
	//numTeams()
	//undergroundSystem()
	//escapeZero()
	//stdSort()
	//kmpTest()
	//oneCount()
	//pathNum()
	//maxDepthAfterSplit()
	//gameOfLife()
	//myAtoi()
	//trap()
	//countLargestGroup()
	//canConstruct()
	//checkOverlap()
	//maxSatisfaction()
	//minSubsequence()
	//NumSteps()
	//stoneGame()
	//lfu()
	//rotate()
	movingCount()
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

func massage() {
	nums := []int{2,1,4,5,3,1,1,3}
	fmt.Println(dp.Massage(nums))
}

func SurfaceArea() {
	input := [][]int{{1,1,1}, {1,0,1},{1,1,1}}
	fmt.Println(daily.SurfaceArea(input))
}

func numRookCaptures() {
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	select{
	case <-ctx.Done():
		//doSometing
	}

	ch := make(chan int)
	t := time.After(5 * time.Second)
	select {
	case <-t:
		close(ch)
	}
	input := [][]byte{{'.','.','.','.','.','.','.','.'},
					{'.','.','.','p','.','.','.','.'},
					{'.','.','.','p','.','.','.','.'},
					{'p','p','.','R','.','p','B','.'},
					{'.','.','.','.','.','.','.','.'},
					{'.','.','.','B','.','.','.','.'},
					{'.','.','.','p','.','.','.','.'},
					{'.','.','.','.','.','.','.','.'}}
	fmt.Println(daily.NumRookCaptures(input))
}

func hasGroupsSizeX() {
	input := []int{1,2,3}
	fmt.Println(daily.HasGroupsSizeX(input))
}

func trieTree() {
	trie := index.Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple")) // 返回 true
	fmt.Println(trie.Search("app"))   // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app"))     // 返回 true
	fmt.Println(trie.ExactlyStartsWith("apple")) // 返回 true
}

func minimumLengthEncoding() {
	words := []string{"time", "me", "bell"}
	fmt.Println(daily.MinimumLengthEncoding(words))
}

func minIncrementForUnique() {
	input := []int{3,2,1,2,1,7}
	fmt.Println(daily.MinIncrementForUnique(input))
}

func findLucky() {
	arr := []int{7,7,7,7,7,7,7}
	fmt.Println(weekly.FindLucky(arr))
}

func numTeams() {
	arr := []int{1,2,3,4}
	fmt.Println(weekly.NumTeams(arr))
}

func undergroundSystem() {
	undergroundSystem := weekly.Constructor();
	undergroundSystem.CheckIn(334530, "40R0G6EM", 27)
	undergroundSystem.CheckOut(334530, "40R0G6EM", 30)
	undergroundSystem.CheckIn(568697, "IESH9C88", 99)
	undergroundSystem.CheckOut(568697, "YFGL6BD9", 185)
	undergroundSystem.CheckIn(228445, "TFU233RE", 200)
	undergroundSystem.CheckOut(228445, "03RA2AIV", 245)
	undergroundSystem.CheckIn(931420, "E5DR624I", 254)
	undergroundSystem.CheckOut(931420, "AKXE2R5P", 255)
	fmt.Println(undergroundSystem.GetAverageTime("IESH9C88", "YFGL6BD9"))
	undergroundSystem.CheckIn(579295, "3V2KP8CY", 355)
	undergroundSystem.CheckIn(350549, "IESH9C88", 428)
	fmt.Println(undergroundSystem.GetAverageTime("40R0G6EM", "D4MEVFON"))
	fmt.Println(undergroundSystem.GetAverageTime("TFU233RE", "03RA2AIV"))
	undergroundSystem.CheckOut(579295, "16W1KJWQ", 488)
	undergroundSystem.CheckOut(350549, "YFGL6BD9", 491)
	fmt.Println(undergroundSystem.GetAverageTime("TFU233RE", "03RA2AIV"))
	undergroundSystem.CheckIn(670740, "IUTXOOZE", 576)
	undergroundSystem.CheckIn(721153, "CEIKSL9H", 643)
	undergroundSystem.CheckIn(590057, "HSI0KVF2", 664)
	fmt.Println(undergroundSystem.GetAverageTime("TFU233RE", "03RA2AIV"))
	fmt.Println(undergroundSystem.GetAverageTime("3V2KP8CY", "16W1KJWQ"))
	fmt.Println(undergroundSystem.GetAverageTime("E5DR624I", "AKXE2R5P"))
	undergroundSystem.CheckOut(590057, "IG3AUY9H", 673)
	undergroundSystem.CheckIn(722875, "IESH9C88", 767)
	fmt.Println(undergroundSystem.GetAverageTime("E5DR624I", "AKXE2R5P"))
	fmt.Println(undergroundSystem.GetAverageTime("E5DR624I", "AKXE2R5P"))
	fmt.Println(undergroundSystem.GetAverageTime("3V2KP8CY", "16W1KJWQ"))
	undergroundSystem.CheckIn(480231, "M2VQ1GJG", 813)
	fmt.Println(undergroundSystem.GetAverageTime("IESH9C88", "YFGL6BD9"))
	fmt.Println(undergroundSystem.GetAverageTime("3V2KP8CY", "16W1KJWQ"))
	fmt.Println(undergroundSystem.GetAverageTime("TFU233RE", "03RA2AIV"))
}

func escapeZero() {
	//input := [][]int{{0,0,0,0,0},{0,1,1,1,0},{1,0,0,0,1},{0,1,1,0,0},{0,0,0,0,0}}
	//input := [][]int{{1,1,1,1,1},{0,0,0,0,1},{0,0,0,0,1}}
	//input := [][]int{{1,1,1,1,0},{0,0,1,1,0},{0,0,0,1,1}}
	input := [][]int{{1,1},{1,1}}
	//input := [][]int{{0,1,0},{1,0,1},{0,1,0}}
	fmt.Println(alibaba.EscapeZero(input))
}

func stdSort() {
	input := []int{5,1,1,2,0,0}
	fmt.Println(stdsort.QuickSort(input))
}

func kmpTest() {
	s := "BBC ABCDAB ABCDABCDABDE"
	p := "ABCDABD"
	fmt.Println(kmp.KmpSearch(s, p))
}

func oneCount() {
	m := 111
	n := 112
	fmt.Println(tecent.Count(m,n))
}

func pathNum() {
	input := [][]byte{{'S','o', 'o', 'o', 'o'},
		{'o','x', 'o', 'x', 'o'},
		{'o','o', 'o', 'o', 'x'},
		{'o','o', 'o', 'o', 'E'}}
	fmt.Println(alibaba.PathNumDP(input))
}

func maxDepthAfterSplit() {
	s := "()(())()"
	fmt.Println(daily.MaxDepthAfterSplit(s))
}

func gameOfLife() {
	input := [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
	april.GameOfLife(input)
	fmt.Println(input)
}

func myAtoi() {
	//s := "-91283472332"
	s := "9223372036854775808"
	//s := "-2147483648"
	fmt.Println(april.MyAtoi(s))
}

func trap() {
	input := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(april.TrapOptimize(input))
}

func countLargestGroup() {
	n := 24
	fmt.Println(evenweekly.CountLargestGroup(n))
}

func canConstruct() {
	s := "yzyzyzyzyzyzyzy"
	k := 2
	fmt.Println(evenweekly.CanConstruct(s, k))
}

func checkOverlap() {
	radius := 1415
	x_center := 807
	y_center := -784
	x1 := -733
	y1 := 623
	x2 := -533
	y2 := 1005
	fmt.Println(evenweekly.CheckOverlap(radius, x_center, y_center, x1, y1, x2, y2))
}

func maxSatisfaction() {
	input := []int{-2,5,-1,0,3,-3}
	fmt.Println(evenweekly.MaxSatisfaction(input))
}

func minSubsequence() {
	input := []int{6}
	fmt.Println(weekly.MinSubsequence(input))
}

func NumSteps() {
	s := "10"
	fmt.Println(weekly.NumSteps(s))
}

func stoneGame() {
	input := []int{5,3,4,5}
	fmt.Println(dp.StoneGame(input))
}

func lfu() {
	cache := LFU.Constructor(2)

	cache.Put(1,1)
	cache.Put(2,2)
	fmt.Println(cache.Get(1))       // 返回 1
	cache.Put(3, 3)    // 去除 key 2
	fmt.Println(cache.Get(2))    // 返回 -1 (未找到key 2)
	fmt.Println(cache.Get(3))    // 返回 3
	cache.Put(4, 4)    // 去除 key 1
	fmt.Println(cache.Get(1))       // 返回 -1 (未找到 key 1)
	fmt.Println(cache.Get(3))       // 返回 3
	fmt.Println(cache.Get(4))       // 返回 4

	//cache := LFU.Constructor(0)
	//cache.Put(0,0)
	//fmt.Println(cache.Get(0))
}

func rotate() {
	input := [][]int{{5, 1, 9,11},{2, 4, 8,10},{13, 3, 6, 7},{15,14,12,16}}
	april.RotateOptimize(input)
	fmt.Println(input)
}

func movingCount() {
	m, n, k := 38,15,9
	fmt.Println(april.MovingCount(m, n, k))
}