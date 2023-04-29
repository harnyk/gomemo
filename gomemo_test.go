package gomemo_test

import (
	"testing"

	"github.com/harnyk/gomemo"
	"github.com/ysmood/got"
)

func TestMemoize(t *testing.T) {
	t.Run("should memoize a function, which accepts a simple struct", func(t *testing.T) {
		type payload struct {
			a int
			b int
		}

		callCount := 0

		sum := func(input payload) (int, error) {
			callCount++
			return input.a + input.b, nil
		}

		memoizedSum := gomemo.Memoize(sum)

		// First call
		result, err := memoizedSum(payload{a: 1, b: 2})
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 1)

		// Second call
		result, err = memoizedSum(payload{a: 1, b: 2})
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 1)
	})

	t.Run("should memoize a function, which accepts a pointer to struct", func(t *testing.T) {
		type payload struct {
			a int
			b int
		}

		callCount := 0

		sum := func(input *payload) (int, error) {
			callCount++
			return input.a + input.b, nil
		}

		memoizedSum := gomemo.Memoize(sum)

		p0 := &payload{a: 1, b: 2}
		p1 := &payload{a: 1, b: 2}
		// First call
		result, err := memoizedSum(p0)
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 1)

		// Second call. Should NOT be memoized, because p0 and p1 are different pointers
		result, err = memoizedSum(p1)
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 2)

		// Third call with the same pointer as the first call. Should be memoized
		result, err = memoizedSum(p0)
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 2)
	})
}

func TestMemoizeWithHasher(t *testing.T) {
	t.Run("should use memoized value if hash matches", func(t *testing.T) {

		type hashablePayload struct {
			a byte
			b byte
		}

		callCount := 0

		sum := func(input hashablePayload) (int, error) {
			callCount++
			return int(input.a) + int(input.b), nil
		}

		hash := func(input hashablePayload) int16 {
			return int16(input.a)<<8 | int16(input.b)
		}

		memoizedSum := gomemo.MemoizeWithHasher(sum, hash)

		// First call
		result, err := memoizedSum(hashablePayload{a: 1, b: 2})
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 1)

		// Second call
		result, err = memoizedSum(hashablePayload{a: 1, b: 2})
		got.T(t).Nil(err)
		got.T(t).Equal(result, 3)
		got.T(t).Equal(callCount, 1)
	})
}
