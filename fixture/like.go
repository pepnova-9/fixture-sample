package fixture

import (
	"github.com/pepnova-9/fixture-sample/entity"
	"testing"
)

func Like(likeBy string) *ModelConnector {
	like := &entity.Like{
		LikeBy: likeBy,
	}
	return &ModelConnector{
		Model: like,
		addToFixture: func(t *testing.T, f *Fixture) {
			f.DataModels = append(f.DataModels, like)
			f.Likes = append(f.Likes, like)
		},
	}
}
