package main

import "strings"

func numUniqueEmails(emails []string) int {
	m1 := make(map[string]int, 0)
	for i := 0; i < len(emails); i++ {
		email := strings.Split(emails[i], "@")
		local := strings.Replace(email[0], ".", "", -1)
		local = strings.Split(local, "+")[0]
		e := local + "@" + email[1]
		m1[e]++
	}
	count := 0
	for _, _ = range m1 {
		count++
	}
	return count
}
