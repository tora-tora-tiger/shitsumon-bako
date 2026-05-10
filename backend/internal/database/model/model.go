package model

import (
	"crypto/rand"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// TODO: 今後置き換える
func generateUUID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant RFC 4122
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func (i *ImageFile) BeforeCreate(tx *gorm.DB) (err error) {
	i.Id = generateUUID()
	return
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	q.Id = generateUUID()
	return
}

type TimeStamp struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ImageFile struct {
	TimeStamp
	Id           string `gorm:"primaryKey;type:uuid"`
	OriginalName string
	MimeType     string
	Size         int64
	Url          string
	QuestionId   string `gorm:"type:uuid;index"`
}

type Question struct {
	TimeStamp
	Id                string       `gorm:"primaryKey;type:uuid"`
	Status            string       `gorm:"type:text;not null;default:'unread'" sql:"type:ENUM('unread', 'read', 'answered')"`
	RecipientId       string       `gorm:"type:uuid;index;not null"`
	SenderId          *string      `gorm:"type:uuid;index"`
	Content           string       `gorm:"type:text;not null"`
	AttachedImageList *[]ImageFile `gorm:"foreignKey:QuestionId"` //;constraint:OnDelete:CASCADE"`
	IsAnonymous       bool         `gorm:"not null;default:true"`
	Answer            *Answer      `gorm:"foreignKey:QuestionId"` //;constraint:OnDelete:CASCADE"`
}

type Answer struct {
	TimeStamp
	Id                string       `gorm:"primaryKey;type:uuid"`
	QuestionId        string       `gorm:"type:uuid;index"`
	Content           string       `gorm:"type:text;not null"`
	AttachedImageList *[]ImageFile `gorm:"foreignKey:QuestionId"` //;constraint:OnDelete:CASCADE"`
	IsPublic          bool         `gorm:"not null;default:false"`
}
