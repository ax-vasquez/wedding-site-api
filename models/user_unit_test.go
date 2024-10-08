//go:build unit
// +build unit

package models

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/ax-vasquez/wedding-site-api/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_UserModel_Unit(t *testing.T) {
	os.Setenv("USE_MOCK_DB", "true")
	assert := assert.New(t)
	u := User{
		BaseModel: BaseModel{
			ID: uuid.New(),
		},
		Role:      "GUEST",
		IsGoing:   true,
		FirstName: "Booples",
		LastName:  "McFadden",
		Email:     "fake@email.place",
	}
	errMsg := "arbitrary database error"
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	t.Run("CreateUsers - database error returns error", func(t *testing.T) {
		_, mock, _ := Setup()
		mock.ExpectBegin()
		mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO "users" ("created_at","updated_at","deleted_at","role","is_going","first_name","last_name","email","password","token","refresh_token","hors_doeuvres_selection_id","entree_selection_id","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING "id"`)).WithArgs(
			test.AnyTime{},
			test.AnyTime{},
			nil,
			u.Role,
			u.IsGoing,
			u.FirstName,
			u.LastName,
			u.Email,
			u.Password,
			u.Token,
			u.RefreshToken,
			u.HorsDoeuvresSelectionId,
			u.EntreeSelectionId,
			u.ID,
		).WillReturnError(fmt.Errorf(errMsg))
		mock.ExpectRollback()
		mock.ExpectCommit()

		users := &[]User{u}
		err := CreateUsers(ctx, users)

		assert.NotNil(err)
		assert.Equal(errMsg, err.Error())

	})
	t.Run("FindUsers - database error returns error", func(t *testing.T) {
		someId := uuid.New()
		_, mock, _ := Setup()
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT "id","role","is_going","first_name","last_name","email","entree_selection_id","hors_doeuvres_selection_id" FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`)).WithArgs(
			someId,
		).WillReturnError(fmt.Errorf(errMsg))
		mock.ExpectRollback()
		mock.ExpectCommit()

		users, err := FindUsers(ctx, []uuid.UUID{someId})

		assert.Empty(users)
		assert.NotNil(err)
		assert.Equal(errMsg, err.Error())
	})
	t.Run("DeleteUser - database error returns error", func(t *testing.T) {
		someId := uuid.New()
		_, mock, _ := Setup()
		mock.ExpectBegin()
		mock.ExpectExec(
			regexp.QuoteMeta(`UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`)).WithArgs(
			test.AnyTime{},
			someId,
		).WillReturnError(fmt.Errorf(errMsg))
		mock.ExpectRollback()
		mock.ExpectCommit()

		count, err := DeleteUser(ctx, someId)

		assert.Zero(count)
		assert.NotNil(err)
		assert.Equal(errMsg, err.Error())
	})
	t.Run("UpdateUser - database error returns error", func(t *testing.T) {
		_, mock, _ := Setup()
		mock.ExpectBegin()
		mock.ExpectQuery(
			regexp.QuoteMeta(`UPDATE "users" SET "updated_at"=$1,"role"=$2,"is_going"=$3,"first_name"=$4,"last_name"=$5,"email"=$6 WHERE "users"."deleted_at" IS NULL AND "id" = $7 RETURNING "users"."role","users"."first_name","users"."last_name","users"."email","users"."hors_doeuvres_selection_id","users"."entree_selection_id"`)).WithArgs(
			test.AnyTime{},
			u.Role,
			u.IsGoing,
			u.FirstName,
			u.LastName,
			u.Email,
			u.ID,
		).WillReturnError(fmt.Errorf(errMsg))
		mock.ExpectRollback()
		mock.ExpectCommit()

		err := UpdateUser(ctx, &u)

		assert.NotNil(err)
		assert.Equal(errMsg, err.Error())
	})
	t.Run("SetIsGoing - database error returns error", func(t *testing.T) {
		_, mock, _ := Setup()
		mock.ExpectBegin()
		mock.ExpectExec(
			regexp.QuoteMeta(`UPDATE "users" SET "updated_at"=$1,"is_going"=$2 WHERE "users"."deleted_at" IS NULL AND "id" = $3`)).WithArgs(
			test.AnyTime{},
			u.IsGoing,
			u.ID,
		).WillReturnError(fmt.Errorf(errMsg))
		mock.ExpectRollback()
		mock.ExpectCommit()

		err := SetIsGoing(ctx, &u)

		assert.NotNil(err)
		assert.Equal(errMsg, err.Error())
	})
}
