package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithValidData(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, 0, client.Points)
}

func TestNewClientWithInvalidData(t *testing.T) {
	client, err := NewClient("", "j@j.com")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client name is required")

	client, err = NewClient("John Doe", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client email is required")
}

func TestAddPointsBatch(t *testing.T) {
	pointsTable := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, point := range pointsTable {
		client, _ := NewClient("John", "j@j.com")
		err := client.AddPoints(point)
		assert.Nil(t, err)

		if client.Points != point {
			t.Errorf("Points expected: %d, got: %d", point, client.Points)
		}
	}
}

func FuzzClient_AddPoints(f *testing.F) {
	seeding := []int{1, 2, 3, 4, 5, 6, 7}

	for _, seed := range seeding {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, points int) {
		client, _ := NewClient("John Doe", "j@j.com")
		err := client.AddPoints(points)
		if err != nil {
			return
		}
		if client.Points != points {
			t.Errorf("Points expected: %d, got: %d", points, client.Points)
		}
	})
}
