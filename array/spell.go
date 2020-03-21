package array

import (
	"sync"
)

var charMap map[int32]int

var chs []byte

func CountCharacters(words []string, chars string) int {
	//if len(chars) == 0 || len(words) == 0 {
	//	return 0
	//}
	//
	//charMap = make(map[int32]int)
	//for _, c := range chars {
	//	if _, ok := charMap[c]; ok {
	//		charMap[c] += 1
	//	} else {
	//		charMap[c] = 1
	//	}
	//}
	//
	//var wg sync.WaitGroup
	//res := 0
	//ch := make(chan int, len(words))
	//for _, word := range words {
	//	wg.Add(1)
	//	go isContain(word, chars, &wg, ch)
	//}
	//
	//wg.Wait()
	//close(ch)
	//for v := range ch {
	//	res += v
	//}
	//
	//return res

	if len(words) == 0 {
		return 0
	}

	// 构建字符哈希表
	chs = make([]byte, 'z'-'A'+1)
	for _, ch := range chars {
		chs[ch-'A']++
	}

	// 初始化长度之和
	res := 0
	var wg sync.WaitGroup
	ch := make(chan int, len(words))

	// 遍历字符串数组
	for _, str := range words {
		wg.Add(1)
		go isContained(str, &wg, ch)
	}

	wg.Wait()
	close(ch)
	for v := range ch {
		res += v
	}

	return res
}

func isContained(str string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	ws := make([]byte, 'z'-'A'+1)
	for _, w := range str {
		ws[w-'A']++
	}
	// 一一比较两个哈希表的元素
	for i := 0; i < len(ws); i++ {
		if len(ws) > len(chs) {
			return
		}
		if ws[i] > chs[i] {
			return
		}
	}
	ch<-len(str)
	return
}

func isContain(word string, chars string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	wordMap := make(map[int32]int)
	for _, c := range word {
		if _, ok := wordMap[c]; ok {
			wordMap[c] += 1
		} else {
			wordMap[c] = 1
		}
	}

	for k, v := range wordMap {
		if _, ok := charMap[k]; ok {
			if v > charMap[k] {
				return
			}
		} else {
			return
		}
	}

	//for i:=0;i<len(word);i+=1 {
	//	pos := strings.Index(chars, string(word[i]))
	//	if pos < 0 {
	//		return
	//	} else {
	//		chars = strings.Replace(chars, string(word[i]), "", 1)
	//	}
	//}
	ch<-len(word)
	return
}
