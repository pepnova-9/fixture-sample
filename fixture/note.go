package fixture

import (
	"github.com/pepnova-9/fixture-sample/entity"
	"testing"
)

func Note(id string) *ModelConnector {
	note := &entity.Note{
		ID: id,
	}
	return &ModelConnector{
		Model: note,
		addToFixture: func(t *testing.T, f *Fixture) {
			f.DataModels = append(f.DataModels, note)
			f.Notes = append(f.Notes, note)
		},
		connect: func(t *testing.T, f *Fixture, dependentModel any) {
			switch dependentModel := dependentModel.(type) {
			case *entity.Like:
				dependentModel.NoteID = note.ID
			case *entity.Tag:
				mapping := &entity.NoteTagMapping{
					NoteID: note.ID,
					TagID:  dependentModel.ID,
				}
				f.NoteTagMappings = append(f.NoteTagMappings, mapping)
				f.MappingDataModels = append(f.MappingDataModels, mapping)
			default:
				t.Fatalf("%Tのconnectに%Tが定義されていないためconnectできません。", note, dependentModel)
			}
		},
	}
}
