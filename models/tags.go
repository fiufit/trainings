package models

import "github.com/fiufit/trainings/contracts"

var validTags = map[string]struct{}{
	"strength":    {},
	"speed":       {},
	"endurance":   {},
	"lose weight": {},
	"gain weight": {},
	"sports":      {},
}

type Tag struct {
	Name string `gorm:"primaryKey;not null;index;unique"`
}

func ValidateTags(tagStrings ...string) ([]Tag, error) {
	tags := make([]Tag, len(tagStrings))
	for i, tag := range tagStrings {
		if _, exists := validTags[tag]; !exists {
			return []Tag{}, contracts.ErrInvalidTag
		}
		tags[i] = Tag{Name: tag}
	}
	return tags, nil
}
