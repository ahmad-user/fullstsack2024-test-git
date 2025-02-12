package user

import (
	"time"

	"gorm.io/gorm"
)

type MyClient struct {
	ID           int            `gorm:"primaryKey;autoIncrement"`
	Name         string         `gorm:"type:char(250);not null"`
	Slug         string         `gorm:"type:char(100);not null"`
	IsProject    string         `gorm:"type:varchar(30);not null;default:'0';check:is_project IN ('0', '1')"`
	SelfCapture  string         `gorm:"type:char(1);not null;default:'1'"`
	ClientPrefix string         `gorm:"type:char(4);not null"`
	ClientLogo   string         `gorm:"type:char(255);not null;default:'no-image.jpg'"`
	PhoneNumber  string         `gorm:"type:char(50)"`
	City         string         `gorm:"type:char(50)"`
	CreatedAt    time.Time      `gorm:"type:timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp;index"`
}

type MyClientModel struct {
	Connection *gorm.DB
}

func (m *MyClientModel) AddClient(newClient MyClient) error {
	err := m.Connection.Create(&newClient).Error
	return err
}

func (m *MyClientModel) GetAllClients() ([]MyClient, error) {
	var clients []MyClient
	err := m.Connection.Find(&clients).Error
	return clients, err
}

func (m *MyClientModel) GetClientByID(id int) (MyClient, error) {
	var client MyClient
	err := m.Connection.Where("id = ?", id).First(&client).Error
	return client, err
}

func (m *MyClientModel) UpdateClient(id int, updatedData MyClient) error {
	err := m.Connection.Model(&MyClient{}).Where("id = ?", id).Updates(updatedData).Error
	return err
}

func (m *MyClientModel) DeleteClient(id int) error {
	err := m.Connection.Where("id = ?", id).Delete(&MyClient{}).Error
	return err
}
