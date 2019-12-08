package intcode_test

import (
	"fmt"
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
		require.Equal(tt, true, <-intProgram.Finish)
		require.Equal(tt, true, intProgram.Started)
	})
	t.Run("the program can be reset", func(tt *testing.T) {
		require.Equal(tt, true, intProgram.Finished)
		require.Equal(tt, true, intProgram.Started)
		intProgram.Reset()
		require.Equal(tt, false, intProgram.Finished)
		require.Equal(tt, false, intProgram.Started)
	})
	t.Run("the program can be run twice", func(tt *testing.T) {
		intcodes := []int{
			3, 26, 1001, 26, -4,
			26, 3, 27, 1002, 27,
			2, 27, 1, 27, 26,
			27, 4, 27, 1001, 28, -1,
			28, 1005, 28, 6, 99, 0, 0, 5,
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			intProgram.Exec(intcodes)
		}()
		intProgram.Input <- 9
		intProgram.Input <- 0
		out, ok := <-intProgram.Output
		require.Equal(tt, true, ok)
		require.NotZero(tt, out)
		require.Equal(tt, false, intProgram.Finished)
		require.Equal(tt, true, intProgram.Started)
		intProgram.Input <- 9
		intProgram.Input <- 0
		out, ok = <-intProgram.Output
		require.Equal(tt, true, ok)
		require.NotZero(tt, out)
		require.Equal(tt, false, intProgram.Finished)
		require.Equal(tt, true, intProgram.Started)
		select {
		case _ = <-intProgram.Finish:
			require.Fail(tt, "the program has finished")
		default:
			fmt.Println("this is ok")
		}
	})
}
