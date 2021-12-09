package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ReadLines(path string) []string {
	f, err := os.Open(path)
	CheckErr(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	ls := []string{}
	for s.Scan() {
		ls = append(ls, s.Text())
	}
	CheckErr(s.Err())

	return ls
}

func Atoi(s string) (i int) {
	var err error
	if i, err = strconv.Atoi(s); err != nil {
		panic(err.Error())
	}
	return
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ReadIntString(s string) []int {
	result := []int{}
	for _, n := range strings.Split(s, ",") {
		result = append(result, Atoi(n))
	}
	return result
}
