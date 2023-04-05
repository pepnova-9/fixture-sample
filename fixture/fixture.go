package fixture

import (
	"fmt"
	"github.com/pepnova-9/fixture-sample/entity"
	"testing"
)

type Fixture struct {
	DataModels        []any
	MappingDataModels []any

	Notes []*entity.Note
	Tags  []*entity.Tag
	Likes []*entity.Like

	NoteTagMappings []*entity.NoteTagMapping
}

func (f *Fixture) Setup() {
	for _, dataModel := range f.DataModels {
		fmt.Printf("insert %#v \n", dataModel)
	}
	for _, dataModel := range f.MappingDataModels {
		fmt.Printf("insert %#v \n", dataModel)
	}
}

func Build(t *testing.T, modelConnectors ...*ModelConnector) *Fixture {
	fixture := &Fixture{}
	for _, modelConnector := range modelConnectors {
		modelConnector.addToFixtureAndConnect(t, fixture)
	}
	return fixture
}
