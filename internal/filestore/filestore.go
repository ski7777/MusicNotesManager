package filestore

import (
	"github.com/spf13/afero"
	"os"
	pathlib "path"
	"strconv"
	"strings"
	"sync"
)

type FileStore struct {
	fs        afero.Fs
	locks     map[int]*sync.Mutex
	lockslock sync.Mutex
}

func (f *FileStore) PreAddFile() { f.lockslock.Lock() }

func (f *FileStore) AddFile(id int, data []byte) error {
	if file, err := f.fs.Create(strconv.Itoa(id)); err != nil {
		return err
	} else {
		if _, err := file.Write(data); err != nil {
			return err
		}
	}
	return nil
}

func (f *FileStore) PostAddFile() { f.lockslock.Unlock() }

func (f *FileStore) getLock(id int) *sync.Mutex {
	f.lockslock.Lock()
	defer f.lockslock.Unlock()
	if l, ok := f.locks[id]; ok {
		return l
	} else {
		f.locks[id] = &sync.Mutex{}
		return f.locks[id]
	}
}

func (f *FileStore) Read(id int) ([]byte, error) {
	l := f.getLock(id)
	l.Lock()
	defer l.Unlock()
	return afero.ReadFile(f.fs, strconv.Itoa(id))
}

func (f *FileStore) Delete(id int) error {
	l := f.getLock(id)
	l.Lock()
	defer l.Unlock()
	return f.fs.Remove(strconv.Itoa(id))
}

func (f *FileStore) UpdateFile(id int, data []byte) error {
	l := f.getLock(id)
	l.Lock()
	defer l.Unlock()
	if file, err := f.fs.Open(strconv.Itoa(id)); err != nil {
		return err
	} else {
		_, err = file.Write(data)
		return err
	}
}

func NewFileStore(path string) (*FileStore, error) {
	if path == "" {
		return nil, nil
	}
	if !strings.HasPrefix(path, "/") {
		if cwd, err := os.Getwd(); err != nil {
			return nil, err
		} else {
			path = pathlib.Join(cwd, path)
		}
	}
	f := new(FileStore)
	f.locks = make(map[int]*sync.Mutex)
	f.fs = afero.NewBasePathFs(afero.NewOsFs(), path)
	return f, nil
}
