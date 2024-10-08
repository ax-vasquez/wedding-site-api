//go:build integration
// +build integration

package models

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_EntreeModel_Integration(t *testing.T) {
	assert := assert.New(t)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	t.Run("Can find a single entree", func(t *testing.T) {
		id, _ := uuid.Parse("f8cd5ea3-bb29-42fc-9984-a6c37d8b99c3")
		entree, err := FindEntreeById(ctx, id)
		assert.Nil(err)
		assert.Equal("Caprese pasta", entree.OptionName)
	})
	t.Run("Can find all possible entrees", func(t *testing.T) {
		entrees, _ := FindEntrees(ctx)
		assert.Equal(5, len(entrees))
	})
	t.Run("Can find entrees for user", func(t *testing.T) {
		id, _ := uuid.Parse(FirstUserIdStr)
		entrees, err := FindEntreesForUser(ctx, id)
		assert.Nil(err)
		assert.Equal(1, len(entrees))
		assert.Equal("Caprese pasta", entrees[0].OptionName)
	})
	t.Run("Can create an entree", func(t *testing.T) {
		entrees := []Entree{{
			OptionName: "Cap'n Crunch",
		}}
		err := CreateEntrees(ctx, &entrees)
		assert.Nil(err)
		assert.Equal("Cap'n Crunch", entrees[0].OptionName)
		// Embedded test so we can easily-target the new record and delete it as part of the next test
		t.Run("Can delete an entree", func(t *testing.T) {
			id := entrees[0].ID
			result, err := DeleteEntree(ctx, id)
			assert.Nil(err)
			assert.Equal(1, int(*result))
		})
	})
}
