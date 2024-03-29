package mongo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Foo struct {
	BasicDocument `bson:",inline"`
	Action        string
}

func (f Foo) DocumentName() string { return "foo" }

func dummyConnect() (Client, error) {
	clientConfig := ClientConfig{
		Host:     "localhost",
		Port:     27017,
		Database: "test_db",
	}

	return NewClient(clientConfig)
}

// WARNING: Those tests should be run in order or at least TestConnect should
// be launched before any other test to initialize the client

var (
	testClient   Client
	connectError error
)

func TestConnect(t *testing.T) {
	testClient, connectError = dummyConnect()

	assert.Nil(t, connectError)
	assert.NotNil(t, testClient)
}

func TestInsertFoo(t *testing.T) {
	assert.NotNil(t, testClient)

	foo := Foo{
		Action: "Bar",
	}

	err := testClient.Persist(&foo)

	os.Setenv("TEST_UUID", foo.GetID())

	assert.Nil(t, err)
}

func TestInsertExistingFoo(t *testing.T) {
	assert.NotNil(t, testClient)

	existingUUID := os.Getenv("TEST_UUID")

	foo := Foo{
		BasicDocument: BasicDocument{
			ID: existingUUID,
		},
		Action: "Bar",
	}

	err := testClient.Persist(&foo)

	assert.NotNil(t, err)
}

func TestFindOneByID(t *testing.T) {
	assert.NotNil(t, testClient)

	foo := Foo{}

	existingUUID := os.Getenv("TEST_UUID")

	err := testClient.FindOneById(&foo, existingUUID)

	assert.Nil(t, err)
	assert.NotNil(t, foo)
	assert.NotNil(t, foo.CreatedAt)
	assert.Equal(t, existingUUID, foo.GetID())
	assert.Equal(t, "Bar", foo.Action)
}

func TestReplaceOrPersistReplace(t *testing.T) {
	assert.NotNil(t, testClient)

	existingUUID := os.Getenv("TEST_UUID")

	foo := Foo{
		Action: "Bar Replaced",
		BasicDocument: BasicDocument{
			ID: existingUUID,
		},
	}

	err := testClient.ReplaceOrPersist(&foo)

	assert.Nil(t, err)
}

func TestReplaceOrPersistPersist(t *testing.T) {
	assert.NotNil(t, testClient)

	foo := Foo{
		Action: "Bar Persisted",
	}

	err := testClient.ReplaceOrPersist(&foo)

	os.Setenv("TEST_UUID_2", foo.GetID())

	assert.Nil(t, err)
}

func TestFindOneByIDNewlyPersisted(t *testing.T) {
	assert.NotNil(t, testClient)

	foo := Foo{}

	existingUUID := os.Getenv("TEST_UUID_2")

	err := testClient.FindOneById(&foo, existingUUID)

	assert.Nil(t, err)
	assert.NotNil(t, foo)
	assert.Equal(t, existingUUID, foo.GetID())
}

func TestDelete(t *testing.T) {
	assert.NotNil(t, testClient)

	existingUUID := os.Getenv("TEST_UUID_2")

	foo := Foo{
		Action: "Bar Persisted",
		BasicDocument: BasicDocument{
			ID: existingUUID,
		},
	}

	err := testClient.Delete(&foo)

	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	assert.NotNil(t, testClient)

	foo := Foo{
		Action: "Updated testing the test of the testers",
	}

	existingUUID := os.Getenv("TEST_UUID")

	foo.Action = "Updated testing the test of the testers"

	err := testClient.Update(&foo, existingUUID, foo)

	assert.Nil(t, err)
}
