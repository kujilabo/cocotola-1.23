package firestore_test

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	r := regexp.MustCompile(`<([^>]*)>`)
	text := "Let's try <anyway>. <anyway>"
	x := r.FindAllStringSubmatchIndex(text, -1)
	log.Printf("%v\n", x)
	list := make([][]int, 0)
	list = append(list, []int{0, 0})
	index := 0
	assert.Equal(t, "", x)
	for _, pos := range x {
		log.Printf("%v, %v\n", pos[0], pos[1])
		if pos[0] == 0 {
			list[index] = []int{pos[3] + 1, 0}
		} else {
			list[index] = []int{list[index][0], pos[0]}
			log.Printf("<<%s>>\n", text[list[index][0]:list[index][1]])
			index++
			list = append(list, []int{pos[3] + 1, 0})
		}
	}
}
