package robotname

import (
	"errors"
	"math/rand"
	"strings"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var names = map[string]bool{}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(names) == 676000 {
		return "", errors.New("Max number of names reached")
	}

	var nameBuilder strings.Builder
	for {
		nameBuilder.WriteRune(rand.Int31n(26) + 65)
		nameBuilder.WriteRune(rand.Int31n(26) + 65)
		nameBuilder.WriteRune(rand.Int31n(10) + 48)
		nameBuilder.WriteRune(rand.Int31n(10) + 48)
		nameBuilder.WriteRune(rand.Int31n(10) + 48)
		if alreadyUsed := names[nameBuilder.String()]; !alreadyUsed {
			names[nameBuilder.String()] = true
			break
		}
		nameBuilder.Reset()
	}
	r.name = nameBuilder.String()
	return nameBuilder.String(), nil
}

func (r *Robot) Reset() {
	r.name = ""
}
