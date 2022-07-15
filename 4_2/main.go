package main

import (
	"fmt"
	"sort"
	"strings"
)

// Частотный анализ
// Написать функцию, которая получает на вход текст и возвращает
// 10 самых часто встречающихся слов без учета словоформ

func main() {
	fmt.Println(GetTopWords(1, "a b b a"))           // a
	fmt.Println(GetTopWords(2, "a b c"))             // a b
	fmt.Println(GetTopWords(2, ""))                  //
	fmt.Println(GetTopWords(2, "a b c a b a c"))     // a b
	fmt.Println(GetTopWords(1, "b b a c a b a c"))   // b
	fmt.Println(GetTopWords(3, "a c b b c a"))       // a c b
	fmt.Println(GetTopWords(2, "a b c d a b c d c")) // a c
	fmt.Println(GetTopWords(2, "a b c c b a"))       // a b
	fmt.Println(GetTopWords(5, "a"))                 // a

	lorem := "lorem ipsum is simply dummy text of the printing and typesetting industry lorem ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book it has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged it was popularised in the 1960s with the release of letraset sheets containing lorem ipsum passages, and more recently with desktop publishing software like aldus pagemaker including versions of lorem ipsum"
	fmt.Println(GetTopWords(10, lorem))
}

func GetTopWords(count int, text string) []string {
	if text == "" {
		return []string{}
	}

	topWords := []string{}

	wordsCount := map[string]int{}
	for _, word := range strings.Split(text, " ") {
		if _, hasWord := wordsCount[word]; hasWord {
			wordsCount[word]++
		} else {
			wordsCount[word] = 0
		}
	}

	countsWord := map[int][]string{}
	counts := []int{}
	for word, count := range wordsCount {
		countsWord[count] = append(countsWord[count], word)
	}

	for count := range countsWord {
		counts = append(counts, count)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	leftCount := count
	for i := 0; i < len(counts) && leftCount > 0; i++ {
		count := counts[i]
		words := countsWord[count]
		wordCount := len(words)

		if wordCount <= leftCount {
			topWords = append(topWords, words[:wordCount]...)
			leftCount -= wordCount
		} else {
			topWords = append(topWords, words[:leftCount]...)
			leftCount = 0
		}
	}

	return topWords
}
