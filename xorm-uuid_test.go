package xorm_uuid

import (
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
	"xorm.io/xorm"
)

var (
	exampleXormUuid  = XormUUID(uuid.MustParse("5d1e75b1-c608-47ed-bb41-a72b2d6976fc"))
	exampleXormUuid2 = XormUUID(uuid.MustParse("3ca856ae-7d41-4567-a8f4-6b0ac8f02e9c"))
	exampleXormUuid3 = XormUUID(uuid.MustParse("6b32ccf5-9f02-4fe5-b319-f76b13371d6e"))
)

type testTable struct {
	Uuid  XormUUID `xorm:"pk binary(16)"`
	Other XormUUID `xorm:"binary(16)"`
	First int
}

func TestXormUuid(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", ":memory:")
	if err != nil {
		return
	}
	// Sync database
	t.Log("// Sync database")
	err = engine.Sync(&testTable{})
	assert.NoError(t, err)

	// Test insert
	t.Log("// Test insert")
	_, err = engine.Insert(&testTable{
		Uuid:  exampleXormUuid,
		Other: exampleXormUuid2,
		First: 1,
	})
	assert.NoError(t, err)

	// Test get
	t.Log("// Test get")
	a := testTable{Uuid: exampleXormUuid}
	a2, err := engine.Get(&a)
	assert.NoError(t, err)
	assert.True(t, a2)
	assert.Equal(t, exampleXormUuid2, a.Other)
	assert.Equal(t, 1, a.First)

	// Test update
	t.Log("// Test update")
	_, err = engine.Update(&testTable{
		Other: exampleXormUuid3,
		First: 2,
	}, testTable{Uuid: exampleXormUuid})
	assert.NoError(t, err)

	// Check update
	t.Log("// Check update")
	b := testTable{Uuid: exampleXormUuid}
	b2, err := engine.Get(&b)
	assert.NoError(t, err)
	assert.True(t, b2)
	assert.Equal(t, exampleXormUuid3, b.Other)
	assert.Equal(t, 2, b.First)
}
