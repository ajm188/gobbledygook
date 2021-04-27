package gobbledygook

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

// Word returns a random word from the dictionary.
//
// It panics if called before InitWords().
func Word() string {
	return getWords(1, "Word")[0]
}

// Words returns n random words from the dictionary.
//
// It panics if called before InitWords().
func Words(n int) []string {
	return getWords(n, "Words")
}

func getWords(n int, funcName string) []string {
	if len(words) == 0 {
		panic("must call InitWords before calling " + funcName)
	}

	ws := make([]string, n)
	for i := range ws {
		ws[i] = words[rand.Intn(len(words))]
	}

	return ws
}

var words []string

// InitWords reads in the system dictionary from /usr/share/dict/words. It must
// called before calling either Word or Words.
func InitWords() error {
	if len(words) != 0 {
		return nil
	}

	b, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		return err
	}

	ws := strings.Split(string(b), "\n")
	if len(ws[len(ws)-1]) == 0 {
		ws = ws[:len(ws)-1]
	}

	words = ws
	return nil
}
