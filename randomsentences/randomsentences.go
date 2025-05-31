package randomsentences

import (
	"math/rand"
	"time"
)

var sentences = []string{
	"Go is fast and efficient.",
	"Always check errors.",
	"Use goroutines wisely.",
	"Keep your code Complex.",
	"Keep your code simple.",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random returns a random sentence from the list.
func Random() string {
	return sentences[rand.Intn(len(sentences))]
}
