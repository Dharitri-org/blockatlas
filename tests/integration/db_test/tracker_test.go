//go:build integration
// +build integration

package db_test

import (
	"testing"

	"github.com/Dharitri-org/blockatlas/tests/integration/setup"
	"github.com/stretchr/testify/assert"
)

func TestDb_SetBlock(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assert.Nil(t, database.SetLastParsedBlockNumber("ethereum", 0))

	block, err := database.GetLastParsedBlockNumber("ethereum")
	assert.Nil(t, err)
	assert.Equal(t, block.Height, int64(0))

	assert.Nil(t, database.SetLastParsedBlockNumber("ethereum", 110))

	newBlock, err := database.GetLastParsedBlockNumber("ethereum")
	assert.Nil(t, err)
	assert.Equal(t, newBlock.Height, int64(110))
}
