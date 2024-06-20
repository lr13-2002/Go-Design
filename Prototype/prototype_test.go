package prototype

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestKeyWordsClone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := KeyWords{
		"testA": &KeyWord{
			word:     "testA",
			visit:    1,
			UpdateAt: &updateAt,
		},
		"testB": &KeyWord{
			word:     "testA",
			visit:    2,
			UpdateAt: &updateAt,
		},
		"testC": &KeyWord{
			word:     "testA",
			visit:    3,
			UpdateAt: &updateAt,
		},
	}
	now := time.Now()

	updatedWords := []*KeyWord{
		{
			word:     "testB",
			visit:    10,
			UpdateAt: &now,
		},
	}

	got := words.Clone(updatedWords)
	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
