package storage

import (
	"context"
	"testing"
)

// static image path
var image = "/Users/macbookpro/go/src/github.com/baxromumarov/cloud-storage/image_2024-12-09_09-14-27.png"

func TestCreate(t *testing.T) {
	// TestCreate creates a new file in the storage
	var file File
	t.Run("Create", func(t *testing.T) {
		// Create a new file in the storage
		err := file.Create(context.Background(), &File{})
		if err != nil {
			t.Errorf("Create() error = %v", err)
			return
		}
	})
}
