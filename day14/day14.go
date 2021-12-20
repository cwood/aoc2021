package day14

import (
	"log"
	"strings"
)

func insertInto(arr []string, index int, v string) []string {
	arr = append(arr[:index], arr[index:]...)
	arr[index] = v

	return arr

}

func PolymoreTemplateAfterN(initial string, rules map[string]string, n int) int {
	str := strings.Split(initial, "")

step:
	for i := 0; i <= n-1; i++ {

		currC := 0
		for {
			if currC+1 == len(str)-1 {
				continue step
			}

			r := str[currC] + str[currC+1]
			if newC, ok := rules[r]; ok {
				str = insertInto(str, currC+1, newC)
				log.Printf("%d inserting %s %s %v", currC, newC, r, str)
			}
			currC++
		}
	}

	log.Printf("%v", str)

	return 0
}
