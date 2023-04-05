package fixture

import (
	"testing"
)

func TestFixture(t *testing.T) {
	cleaningTag := Tag("掃除")
	f := Build(t,
		Note("掃除機の掛け方").Connect(
			cleaningTag,
			Like("田中"),
			Like("佐藤"),
		),
		Note("風呂掃除の仕方").Connect(
			cleaningTag,
		),
		Note("生姜焼きの作り方").Connect(
			Tag("料理"),
		),
	)

	f.Setup()
}
