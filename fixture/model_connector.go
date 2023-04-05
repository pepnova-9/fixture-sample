package fixture

import "testing"

// ModelConnector はFixtureに入れるModelとそれに依存するModelを紐づけるための仕組みを提供します。
type ModelConnector struct {
	Model any

	// addToFixtureはFixtureに各Modelを追加する関数の定義です。
	// この関数は必ず実装されている必要があります。
	addToFixture func(t *testing.T, f *Fixture)

	// connectは各Modelとそれに依存しているModelを紐付ける関数の定義です。
	// Connect関数を呼び出す場合は関数は必ず実装されている必要があります。
	connect func(t *testing.T, f *Fixture, connectingModel any)

	// isAlreadyAddedToFixtureは既にfixtureに追加されているかを表すフラグで
	// 同一のModelConnectorを複数回fixtureに追加することを防ぎます。
	isAlreadyAddedToFixture bool

	// dependenciesはこのModelに依存しているModelが入ります。
	dependencies []*ModelConnector
}

// Connect は依存するModel(ModelConnector)の紐付けを行うためにdependenciesに依存するModelConnectorを追加します。
func (mc *ModelConnector) Connect(connectors ...*ModelConnector) *ModelConnector {
	mc.dependencies = append(mc.dependencies, connectors...)
	return mc // メソッドチェーンで記述できるようにする
}

func (mc *ModelConnector) addToFixtureAndConnect(t *testing.T, fixture *Fixture) {
	if mc.isAlreadyAddedToFixture {
		t.Logf("この%Tは既にfixtureに追加されているのでSKIPします: %v", mc.Model, mc.Model)
		return
	}

	if mc.addToFixture == nil {
		t.Fatalf("%T: addToFixture fieldは必ずセットされている必要があります。", mc.Model)
	}
	mc.addToFixture(t, fixture)
	mc.isAlreadyAddedToFixture = true

	for _, modelConnector := range mc.dependencies {
		if mc.connect == nil {
			t.Fatalf("%Tにconnectがセットされていないため, %Tをconnectできません.", mc.Model, modelConnector.Model)
		}

		mc.connect(t, fixture, modelConnector.Model)

		modelConnector.addToFixtureAndConnect(t, fixture)
	}
}
