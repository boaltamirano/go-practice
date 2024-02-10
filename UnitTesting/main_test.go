package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	c := require.New(t)

	result := AddTwoNumber(20, 2)
	expect := 22

	c.Equal(expect, result)

}
