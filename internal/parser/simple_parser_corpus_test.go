package parser

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fuzzCorpusDir = "test/fuzz/corpus"
)

// TestFuzzCorpus runs the current fuzzing corpus.
func TestFuzzCorpus(t *testing.T) {
	assert := assert.New(t)

	corpusFiles, err := filepath.Glob(filepath.Join(fuzzCorpusDir, "*"))
	assert.NoError(err)

	for _, corpusFile := range corpusFiles {
		t.Run(filepath.Base(corpusFile), _TestCorpusFile(corpusFile))
	}
}

func _TestCorpusFile(file string) func(*testing.T) {
	return func(t *testing.T) {
		assert := assert.New(t)

		data, err := ioutil.ReadFile(file)
		assert.NoError(err)
		content := string(data)

		parser := New(content)
		for {
			stmt, _, ok := parser.Next()
			if !ok {
				break
			}
			assert.NotNil(stmt)
		}
	}
}
