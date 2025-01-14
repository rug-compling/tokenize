package tokenize

import (
	"github.com/rug-compling/tokenize/br"
	"github.com/rug-compling/tokenize/nobr"
)

func Dutch(text string, withLineBreaks bool) (string, error) {
	if withLineBreaks {
		return br.Dutch(text)
	}
	return nobr.Dutch(text)
}
