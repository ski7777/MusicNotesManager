package migration20201001001m

type Permission struct {
	ID                 int `gorm:"primaryKey"`
	Name, FriendlyName string
}
