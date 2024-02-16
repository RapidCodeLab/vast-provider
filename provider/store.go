package provider

import (
	"encoding/json"
	"os"
	"sync"
)

const (
	itemStorePath = "items.json"
	statStorePath = "stats.json"
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
		ss.data[itemID] = itemEvents{
			ID: itemID,
		}
	}

	e := ss.data[itemID]

	switch valueName {
	case EventTypeNotify:
		e.Notify++
	case EventTypeImpression:
		e.Impression++
	case EventTypeClick:
		e.Click++
	case EventTypeStart:
		e.Start++
	case EventTypeFirstQuartile:
		e.FirstQuartile++
	case EventTypeMidpoint:
		e.Midpoint++
	case EventTypeThirdQuartile:
		e.ThirdQuartile++
	case EventTypeComplete:
		e.Complete++
	default:
	}

	ss.data[itemID] = e
}

func readFromFile(filePath string, state interface{}) error {
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(state)
}

func writToFile(filePath string, state interface{}) error {
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}
