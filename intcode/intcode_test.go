package intcode_test

import (
	"sync"
	"testing"

	"github.com/jecepeda/advent-code-2019/intcode"
	"github.com/stretchr/testify/require"
)

func TestGetInputsRight(t *testing.T) {
	instructions, err := intcode.ReadFile("example_1.txt")
	require.NoError(t, err)
	require.NotNil(t, instructions)
	intProgram := intcode.NewIntCodeProgram()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		intProgram.Exec(instructions)
	}()
	t.Run("base program works", func(tt *testing.T) {
		intProgram.Input <- 1
		intProgram.Input <- 1
		out, ok := <-intProgram.Output
		require.Equal(tt, true, ok)
		require.NotZero(tt, out)
		require.Equal(tt, true, intProgram.Finished)
		require.Equal(tt, true, intProgram.Started)
	})
	t.Run("the program can be reset", func(tt *testing.T) {
		require.Equal(tt, true, intProgram.Finished)
		require.Equal(tt, true, intProgram.Started)
		intProgram.Reset()
		require.Equal(tt, false, intProgram.Finished)
		require.Equal(tt, false, intProgram.Started)
	})
}
