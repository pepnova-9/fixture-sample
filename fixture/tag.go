package fixture

import (
	"github.com/pepnova-9/fixture-sample/entity"
	"testing"
)

func Tag(id string) *ModelConnector {
	tag := &entity.Tag{
		ID: id,
	}
	return &ModelConnector{
		Model: tag,
		addToFixture: func(t *testing.T, f *Fixture) {
			f.DataModels = append(f.DataModels, tag)
			f.Tags = append(f.Tags, tag)
		},
	}
}
