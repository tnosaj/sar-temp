package queue

import (
	"encoding/json"
	"os"
)

type FileQueue struct {
	path string
}

func NewFileQueue(path string) *FileQueue {
	return &FileQueue{path: path}
}

func (f *FileQueue) Add(r Reading) error {
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	return enc.Encode(r)
}

func (f *FileQueue) PopAll() ([]Reading, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var readings []Reading
	dec := json.NewDecoder(file)
	for dec.More() {
		var r Reading
		if err := dec.Decode(&r); err == nil {
			readings = append(readings, r)
		}
	}
	os.Remove(f.path)
	return readings, nil
}
