package pattern

import (
	"fmt"
	. "github.com/andeyfedoseev/browser-demux/browser"
	"github.com/minio/minio/pkg/wildcard"
)

type Pattern struct {
	Pattern string
	Browser *Browser
}

func (p *Pattern) Matches(url string) bool {
	pattern := p.Pattern
	if len(pattern) == 0 {
		return false
	}
	if pattern[0] != '*' {
		pattern = fmt.Sprintf("*%s", pattern)
	}
	if pattern[len(pattern)-1] != '*' {
		pattern = fmt.Sprintf("%s*", pattern)
	}
	return wildcard.MatchSimple(pattern, url)
}
