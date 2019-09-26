package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	S := []byte(s)
	helper(S, "", 1, res)
	return res
}

func helper(s []byte, temp string, count int, res *[]string) {
	if s == nil {
		return
	}

	if count == 4 {
		if isvalid(string(s)) {
			temp = temp + string(s)
			*res = append(*res, temp)
		} else {
			return
		}
	}

	for i := 1; i < 4 && i < len(s); i++ {
		if isvalid(string(s[:i])) {
			helper(s[:i], string(s[:i])+".", count+1, res)
		} else {
			return
		}
	}

	return
}

func isvalid(part string) bool {
	if len(part) != 1 && strings.HasPrefix(part, "0") {
		return false
	}

	if len(part) == 1 {
		return true
	} else if len(part) == 2 {
		dig, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
		if dig >= 10 && dig < 100 {
			return true
		}
	} else if len(part) == 3 {
		dig, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
		if dig >= 100 && dig <= 255 {
			return true
		}
	}

	return false
}
