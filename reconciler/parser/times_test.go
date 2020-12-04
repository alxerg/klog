package parser

import (
	"github.com/stretchr/testify/assert"
	"main/entry"
	"testing"
)

func TestParseEntryWithTimes(t *testing.T) {
	yaml := `
date: 1985-03-14
hours:
- time: 2:00
- time: 5:00
`
	e, _ := Parse(yaml)
	assert.Equal(t, e.Times, []entry.Minutes{entry.Minutes(2 * 60), entry.Minutes(5 * 60)})
}

func TestParseEntryWithMalformedTimesFails(t *testing.T) {
	yaml := `
date: 1985-03-14
hours:
- time: asdf
`
	e, err := Parse(yaml)
	assert.Equal(t, e, entry.Entry{})
	assert.Error(t, err)
}

func TestParseEntryWithInvalidTimesFails(t *testing.T) {
	yaml := `
date: 1985-03-14
hours:
- time: 23:87
`
	e, err := Parse(yaml)
	assert.Equal(t, e, entry.Entry{})
	assert.Error(t, err)
}