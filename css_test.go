package css

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleStatement(t *testing.T) {
	var buf bytes.Buffer
	setOutput(&buf)

	Set("body", Style{
		"width": 10,
	})

	expectedOutput := "body { width: 10; }\n"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestMultipleStatements(t *testing.T) {
	var buf bytes.Buffer
	setOutput(&buf)

	Set("body", Style{
		"width":  10,
		"height": 20,
	})

	Set(".padding", Style{
		"left":  10,
		"right": 20,
	})

	expectedOutput := `body { width: 10; height: 20; }
.padding { left: 10; right: 20; }
`
	assert.Equal(t, expectedOutput, buf.String())
}

func TestInheritance(t *testing.T) {
	var buf bytes.Buffer
	setOutput(&buf)

	Set("body", Style{
		"height": 20,

		".title": Style{
			"color":      "black",
			".sub-title": Style{"font-size": 10},
		},
	})

	expectedOutput := `body { height: 20; }
body .title { color: "black"; }
body .title .sub-title { font-size: 10; }
`
	assert.Equal(t, expectedOutput, buf.String())
}
