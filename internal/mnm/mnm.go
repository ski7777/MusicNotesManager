package mnm

import (
	"github.com/ski7777/MusicNotesManager/internal/filestore"
	"gorm.io/gorm"
	"sync"
)

type MusicNotesManager struct {
	db     *gorm.DB
	dblock sync.Mutex
	fs     *filestore.FileStore
}

func (m *MusicNotesManager) Run() error {
	return nil
}

func (m *MusicNotesManager) lockDB(_ *gorm.DB) {
	m.dblock.Lock()
}

func (m *MusicNotesManager) unlockDB(_ *gorm.DB) {
	m.dblock.Unlock()
}

func NewMusicNotesManager(db *gorm.DB, dbts bool, fs *filestore.FileStore) *MusicNotesManager {
	m := new(MusicNotesManager)
	m.db = db
	m.fs = fs
	if !dbts {
		m.db.Callback().Create().Before("gorm:create").Register("mnm:gorm:lock", m.lockDB)
		m.db.Callback().Create().After("gorm:create").Register("mnm:gorm:unlock", m.unlockDB)

		m.db.Callback().Query().Before("gorm:query").Register("mnm:gorm:lock", m.lockDB)
		m.db.Callback().Query().After("gorm:query").Register("mnm:gorm:unlock", m.unlockDB)

		m.db.Callback().Delete().Before("gorm:delete").Register("mnm:gorm:lock", m.lockDB)
		m.db.Callback().Delete().After("gorm:delete").Register("mnm:gorm:unlock", m.unlockDB)

		m.db.Callback().Update().Before("gorm:update").Register("mnm:gorm:lock", m.lockDB)
		m.db.Callback().Update().After("gorm:update").Register("mnm:gorm:unlock", m.unlockDB)

		m.db.Callback().Row().Before("gorm:row").Register("mnm:gorm:lock", m.lockDB)
		m.db.Callback().Row().After("gorm:row").Register("mnm:gorm:unlock", m.unlockDB)

		m.db.Callback().Raw().Before("gorm:raw").Register("mnm:gorm:lock", m.lockDB)
		m.db.Callback().Raw().After("gorm:raw").Register("mnm:gorm:unlock", m.unlockDB)
	}
	return m
}
