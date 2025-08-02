package models

import (
	"errors"

	"github.com/danilobml/posts-api/cmd/api/initializers"
	"gorm.io/gorm"
)

var ErrPostNotFound = errors.New("post not found")

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p *Post) All() ([]*Post, error) {
	var posts []*Post

	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (p *Post) FindOne() (*Post, error) {
	result := initializers.DB.First(&p, p.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrPostNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (p *Post) Create() error {
	result := initializers.DB.Create(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *Post) Update() error {
	var post Post
	if err := initializers.DB.First(&post, p.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPostNotFound
		}
		return err
	}

	result := initializers.DB.Save(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *Post) Delete() error {
	result := initializers.DB.Delete(&p)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrPostNotFound
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}
