package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagPrepareEvaluation(t *testing.T) {
	t.Run("happy code path", func(t *testing.T) {
		f := GenFixtureFlag()
		assert.NoError(t, f.PrepareEvaluation())
		assert.NotNil(t, f.FlagEvaluation.VariantsMap)
	})
}

func TestFlagPreload(t *testing.T) {
	t.Run("happy code path", func(t *testing.T) {
		f := GenFixtureFlag()
		db := PopulateTestDB(f)
		defer db.Close()

		err := f.Preload(db)
		assert.NoError(t, err)
	})
}
