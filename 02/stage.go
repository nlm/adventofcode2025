package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/nlm/adventofcode2025/internal/stage"
	"github.com/nlm/adventofcode2025/internal/utils"
)

func IsInvalid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stage.Println(s[0:len(s)/2], s[len(s)/2:])
	return s[0:len(s)/2] == s[len(s)/2:]
}

func Stage1(input io.Reader) (any, error) {
	invalids := 0
	data := string(utils.Must(io.ReadAll(input)))
	for _, entry := range strings.Split(data, ",") {
		bounds := strings.SplitN(entry, "-", 2)
		start, end := utils.MustAtoi(bounds[0]), utils.MustAtoi(bounds[1])
		if start > end {
			return nil, fmt.Errorf("invalid start end")
		}
		stage.Println(start, "->", end)
		for i := start; i <= end; i++ {
			stage.Println(i)
			if IsInvalid(fmt.Sprint(i)) {
				invalids += i
			}
		}
		stage.Println()
	}
	return invalids, nil
}

func IsInvalid2(s string) bool {
	stage.Println("inval2:", s)
nextlen:
	for l := 1; l <= len(s)/2; l++ {
		if len(s)%l != 0 {
			stage.Println("skip len", l)
			continue
		}
		stage.Println("len", l)
		lastVal := ""
		match := false
		for i := 0; i < len(s); i += l {
			val := fmt.Sprint(s[i : i+l])
			stage.Println("->", val)
			if lastVal == "" {
				lastVal = val
			} else if lastVal == val {
				match = true
			} else {
				continue nextlen
			}
		}
		if match {
			return true
		}
	}
	return false
}

func Stage2(input io.Reader) (any, error) {
	invalids := 0
	data := string(utils.Must(io.ReadAll(input)))
	for _, entry := range strings.Split(data, ",") {
		bounds := strings.SplitN(entry, "-", 2)
		start, end := utils.MustAtoi(bounds[0]), utils.MustAtoi(bounds[1])
		if start > end {
			return nil, fmt.Errorf("invalid start end")
		}
		stage.Println(start, "->", end)
		for i := start; i <= end; i++ {
			if IsInvalid2(fmt.Sprint(i)) {
				stage.Println("invalid", i)
				invalids += i
			}
			stage.Println()
		}
		stage.Println()
	}
	return invalids, nil
}
