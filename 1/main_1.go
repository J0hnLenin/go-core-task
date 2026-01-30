package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func initData() []any {
	values := make([]any, 7)
	values[0] = 10
	values[1] = 012
	values[2] = 0x3A
	values[3] = 2.7
	values[4] = "Golang"
	values[5] = false
	values[6] = -1 + 3i

	return values
}

func toStringAndConcatinate(data []any) string {
	stringData := make([]string, len(data))
	for i, value := range(data) {
		stringData[i] = fmt.Sprint(value)
	}
	
	return strings.Join(stringData, " ")
}

func toRune(s string) []rune {
	return []rune(s)
}

func hash(r []rune) []byte {
	hasher := sha256.New()
    
	half := len(r)/2

	hasher.Write([]byte(string(r[:half])))
	hasher.Write([]byte("go-2024"))
	hasher.Write([]byte(string(r[half:])))
    return hasher.Sum(nil)
}

func main() {
	
	data := initData()
	fmt.Println("Слайс переменных: ")
	for _, value := range(data) {
		fmt.Println(fmt.Sprintf("%T", value), value)
	}
	fmt.Println()
	
	strData := toStringAndConcatinate(data)
	fmt.Println("Склеенная строка:")
	fmt.Println(strData)
	fmt.Println()

	runeData := toRune(strData)
	fmt.Println("Слайс рун:")
	fmt.Println(runeData)
	fmt.Println()

	fmt.Println("Хэш:")
	fmt.Println(hash(runeData))
}
