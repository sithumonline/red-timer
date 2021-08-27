package post

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/sithumonline/red-timer/database"
)

type Post struct {
	gorm.Model `swaggerignore:"true"`
	Body       string `json:"body,omitempty"`
}

func (obj *Post) Save() error {
	if err := database.Database().Create(&obj).Error; err != nil {
		log.Errorf("failed to create post: %+v: %v", obj, err)
		return err
	}

	return nil
}

func (obj *Post) GetList() ([]Post, error) {
	list := make([]Post, 0)

	if err := database.Database().Find(&list).Error; err != nil {
		log.Errorf("failed to find post: %+v: %v", obj, err)
		return nil, err
	}

	return list, nil
}

func (obj *Post) Get(id string) (*Post, error) {
	t := &Post{}

	if err := database.Database().Where("ID = ?", id).First(t).Error; err != nil {
		log.Errorf("failed to find post: %+v: %v", obj, err)
		return t, err
	}

	return t, nil
}

func (obj *Post) Delete(id string) error {
	if err := database.Database().Model(&obj).Where("ID = ?", id).Delete(obj).Error; err != nil {
		log.Errorf("failed to delete post: %+v: %v", obj, err)
		return err
	}

	return nil
}

func (obj *Post) Update(id string) error {
	if err := database.Database().Model(&obj).Where("ID = ?", id).Updates(obj).Error; err != nil {
		log.Errorf("failed to update post: %+v: %v", obj, err)
		return err
	}

	return nil
}

func (obj *Post) Migrate() error {
	if err := database.Database().Migrator().AutoMigrate(obj); err != nil {
		log.Errorf("failed to migrate post: %+v: %v", obj, err)
		return err
	}

	return nil
}
