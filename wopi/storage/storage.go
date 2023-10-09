package storage

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	ErrFileNotFound error = errors.New("file with id `%d` could not be found")
)

type Storage struct {
	Root  string
	Files map[string]File
}

type File struct {
	ID        string
	Name      string
	Path      string
	Extension string
}

func New(path string) Storage {
	return Storage{
		Root:  path,
		Files: make(map[string]File),
	}
}

func (s Storage) Open(id string, read bool) (File, []byte, error) {
	file, ok := s.Files[id]
	if !ok {
		return File{}, []byte{}, fmt.Errorf(ErrFileNotFound.Error(), id)
	}

	if read {
		raw, err := s.read(file.Path)
		if err != nil {
			return File{}, []byte{}, err
		}

		return file, raw, nil
	}

	return file, []byte{}, nil
}

func (s Storage) read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}

	raw, err := io.ReadAll(f)
	if err != nil {
		return []byte{}, err
	}

	return raw, err
}

func (s *Storage) Update(id string, content []byte) error {
	file, ok := s.Files[id]
	if !ok {
		return fmt.Errorf(ErrFileNotFound.Error(), id)
	}

	f, err := os.Create(file.Path)
	if err != nil {
		return err
	}

	_, err = f.Write(content)

	return err
}

func (s *Storage) Create(content []byte) (File, error) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	extension := "docx"
	id := strconv.Itoa(r.Int())

	for s.Contains(string(id)) {
		id = strconv.Itoa(r.Int())
	}

	file := File{
		ID:        string(id),
		Extension: "docx",
		Name:      fmt.Sprintf("%s.%s", id, extension),
	}
	file.Path = fmt.Sprintf("%s/%s", s.Root, file.Name)

	s.Files[string(id)] = file

	f, err := os.Create(fmt.Sprintf("%s/%s", s.Root, file.Name))
	if err != nil {
		return File{}, err
	}

	_, err = f.Write(content)
	if err != nil {
		return File{}, err
	}

	return file, nil
}

func (s *Storage) Delete(id string) error {
	file, ok := s.Files[string(id)]
	if !ok {
		return fmt.Errorf(ErrFileNotFound.Error(), id)
	}
	if err := os.Remove(file.Path); err != nil {
		return err
	}

	delete(s.Files, id)
	return nil
}

func (s Storage) Contains(id string) bool {
	_, ok := s.Files[id]

	return ok
}
