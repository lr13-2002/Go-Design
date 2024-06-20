package prototype

import (
	"encoding/json"
	"time"
)

type KeyWord struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

func (k *KeyWord) Clone() *KeyWord {
	var newKeyWod KeyWord
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyWod)
	return &newKeyWod
}

type KeyWords map[string]*KeyWord

func (words KeyWords) Clone(updateWords []*KeyWord) KeyWords {
	newKeyWords := KeyWords{}
	for k, v := range words {
		newKeyWords[k] = v
	}

	for _, word := range updateWords {
		newKeyWords[word.word] = word.Clone()
	}

	return newKeyWords
}
