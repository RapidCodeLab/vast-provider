package provider

import (
	"encoding/json"
	"os"
	"sync"
)

const (
	itemStorePath = ""
	statStorePath = ""
)

type (
	itemStore struct {
		data map[string]item
		mu   sync.Mutex
	}

	statStore struct {
		data map[string]itemEvents
		mu   sync.Mutex
	}
)

func NewItemStore() *itemStore {
	return &itemStore{
		data: make(map[string]item),
	}
}

func NewStatStore() *statStore {
	return &statStore{
		data: make(map[string]itemEvents),
	}
}

func (ss *statStore) IncrementValue(
	valueName,
	itemID string,
) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if ss.data[itemID] == (itemEvents{}) {
		ss.data[itemID] = itemEvents{}
	}

	e := ss.data[itemID]

	switch valueName {
	case EventTypeNotify:
		e.Notify = 0
	case EventTypeImpression:
	case EventTypeClick:
	case EventTypeStart:
	case EventTypeFirstQuartile:
	case EventTypeMidpoint:
	case EventTypeThirdQuartile:
	case EventTypeComplete:
	default:
	}

	ss.data[itemID] = e
}

func readFromFile(filePath string, state interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(state)
}

func writToFile(filePath string, state interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}
