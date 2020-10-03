package mnm

import "github.com/ski7777/MusicNotesManager/internal/database/models"

func (m *MusicNotesManager) createFile(data []byte) (int, error) {
	var file models.File
	if m.fs == nil {
		file.Data = data
	} else {
		m.fs.PreAddFile()
		defer m.fs.PostAddFile()
	}
	if err := m.db.Create(&file).Error; err != nil {
		return 0, err
	}
	if m.fs != nil {
		if err := m.fs.AddFile(file.ID, data); err != nil {
			return 0, err
		}
	}
	return file.ID, nil
}

func (m *MusicNotesManager) updateFile(id int, data []byte) error {
	var file models.File
	if err := m.db.First(&file, id).Error; err != nil {
		return err
	}
	if file.Database {
		file.Data = data
	}
	m.db.Save(file)
	if !file.Database {
		if m.fs == nil {
			return NoFileStoreError
		}
		if err := m.fs.UpdateFile(id, data); err != nil {
			return err
		}
	}
	return nil
}

func (m *MusicNotesManager) readFile(id int) ([]byte, error) {
	var file models.File
	if err := m.db.First(&file, id).Error; err != nil {
		return nil, err
	}
	if file.Database {
		return file.Data, nil
	} else {
		if m.fs == nil {
			return nil, NoFileStoreError
		}
		if data, err := m.fs.Read(id); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	}
}

func (m *MusicNotesManager) deleteFile(id int) error {
	var file models.File
	if err := m.db.First(&file, id).Error; err != nil {
		return err
	}
	if err := m.db.Delete(file).Error; err != nil {
		return err
	}
	if !file.Database {
		if m.fs == nil {
			return NoFileStoreError
		}
		if err := m.fs.Delete(id); err != nil {
			return err
		}
	}
	return nil
}
