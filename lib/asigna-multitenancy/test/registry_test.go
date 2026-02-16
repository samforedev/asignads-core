package test

import (
	"testing"

	asignamultitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRegistry_Isolation(t *testing.T) {
	registry := asignamultitenancy.NewConnectionRegistry()

	dbA := &gorm.DB{Config: &gorm.Config{AllowGlobalUpdate: true}}
	dbB := &gorm.DB{Config: &gorm.Config{AllowGlobalUpdate: false}}

	registry.Set("tenant-a", "dsn-a", dbA)
	registry.Set("tenant-b", "dsn-b", dbB)

	recovereddbA, recovereddsnA, existsA := registry.Get("tenant-a")
	recovereddbB, _, existsB := registry.Get("tenant-b")

	assert.True(t, existsA)
	assert.True(t, existsB)

	assert.Equal(t, "dsn-a", recovereddsnA)
	assert.Equal(t, dbA, recovereddbA)

	assert.NotEqual(t, recovereddbA, recovereddbB)
}
