package wallet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountState(t *testing.T) {

	state, err := getAccountState("n1dZZnqKGEkb1LHYsZRei1CH6DunTio1j1q")
	assert.Equal(t, nil, err)
	fmt.Println(state)
}
