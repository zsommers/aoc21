package util

import (
	"bufio"
	"os"
	"strconv"
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
