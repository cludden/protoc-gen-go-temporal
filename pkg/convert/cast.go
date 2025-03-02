package convert

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	constraints.Signed
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	constraints.Unsigned
}

// Float is a constraint that permits float32 or float64.
type Float interface {
	~float32 | ~float64
}

// Number is a union of the Signed, Unsigned, and Float constraints.
type Number interface {
	Signed | Unsigned | Float
}

// SafeCast attempts to convert a value of type T to type U. If the value
// cannot be represented by U without overflow or underflow, it returns
// an error.
func SafeCast[T Number, U Number](val T) (U, error) {
	// 1) Convert 'val' to a float64 or a *64 integer for range checks.
	switch any(val).(type) {
	case float32, float64:
		return castFromFloat[T, U](float64(val))
	default:
		// Assume integer
		return castFromInt[T, U](val)
	}
}

// castFromFloat handles the case where T is a float.
func castFromFloat[T Number, U Number](f float64) (U, error) {
	min, max := rangeOf[U]()

	if f < min || f > max {
		return zero[U](), fmt.Errorf("value %.4g out of range [%g, %g] for target type",
			f, min, max)
	}

	// If U is an integer type, we still want to check that the integer part
	// fits exactly (e.g. no fractional part if a pure integer is expected).
	switch any(*new(U)).(type) {
	case float32, float64:
		// Floating types: safe to just cast.
		return U(f), nil
	default:
		// Integer types: ensure there's no fractional part
		truncated := math.Trunc(f)
		if truncated != f {
			return zero[U](), fmt.Errorf("float value %.4g has a fractional component, cannot cast to integer", f)
		}

		// Now check again with truncated bounds
		intVal := int64(truncated)
		iMin, iMax := intRangeOf[U]()
		if intVal < iMin || intVal > iMax {
			return zero[U](), fmt.Errorf("value %d out of range [%d, %d] for target integer type",
				intVal, iMin, iMax)
		}

		return U(intVal), nil
	}
}

// castFromInt handles the case where T is an integer type.
func castFromInt[T Number, U Number](val T) (U, error) {
	// We convert 'val' to int64 or uint64 for range checks. Because T may be
	// signed or unsigned, we handle that separately.

	switch any(*new(T)).(type) {
	case int, int8, int16, int32, int64:
		sVal := int64(val)
		return castSignedTo[T, U](sVal)
	default:
		uVal := uint64(val)
		return castUnsignedTo[T, U](uVal)
	}
}

// castSignedTo casts from a signed int64 source to the target U with checks.
func castSignedTo[T Number, U Number](sVal int64) (U, error) {
	// Get the integer range for U (if it's an integer), or float range if it's float.
	switch any(*new(U)).(type) {
	case float32, float64:
		// Floats: check range roughly, then cast.
		f := float64(sVal)
		min, max := rangeOf[U]()
		if f < min || f > max {
			return zero[U](), fmt.Errorf("value %d out of range [%g, %g] for float type", sVal, min, max)
		}
		return U(f), nil
	default:
		// Integer types
		iMin, iMax := intRangeOf[U]()
		if sVal < iMin || sVal > iMax {
			return zero[U](), fmt.Errorf("value %d out of range [%d, %d] for target integer type", sVal, iMin, iMax)
		}
		return U(sVal), nil
	}
}

// castUnsignedTo casts from an unsigned uint64 source to the target U with checks.
func castUnsignedTo[T Number, U Number](uVal uint64) (U, error) {
	switch any(*new(U)).(type) {
	case float32, float64:
		// Floats: check range roughly, then cast.
		f := float64(uVal)
		min, max := rangeOf[U]()
		if f < min || f > max {
			return zero[U](), fmt.Errorf("value %d out of range [%g, %g] for float type", uVal, min, max)
		}
		return U(f), nil
	default:
		// Integer types
		iMin, iMax := intRangeOf[U]()

		// If the target is signed, iMin could be negative. For an unsigned value,
		// the minimum relevant bound effectively becomes 0.
		var minCheck int64
		if iMin > 0 {
			minCheck = iMin
		}

		if uVal < uint64(minCheck) || uVal > uint64(iMax) {
			return zero[U](), fmt.Errorf("value %d out of range [%d, %d] for target integer type", uVal, iMin, iMax)
		}
		return U(uVal), nil
	}
}

// rangeOf returns the [min, max] range in float64 form for a given numeric type U.
func rangeOf[U Number]() (float64, float64) {
	switch any(*new(U)).(type) {
	case int8:
		return float64(math.MinInt8), float64(math.MaxInt8)
	case int16:
		return float64(math.MinInt16), float64(math.MaxInt16)
	case int32:
		return float64(math.MinInt32), float64(math.MaxInt32)
	case int:
		// Typically 32-bit on some systems, 64-bit on others. Adjust as needed.
		// For simplicity, assume 64-bit here:
		return float64(math.MinInt64), float64(math.MaxInt64)
	case int64:
		return float64(math.MinInt64), float64(math.MaxInt64)

	case uint8:
		return 0, float64(math.MaxUint8)
	case uint16:
		return 0, float64(math.MaxUint16)
	case uint32:
		return 0, float64(math.MaxUint32)
	case uint:
		// Typically 32-bit or 64-bit; assume 64:
		return 0, float64(^uint64(0)) // math.MaxUint64 as float64
	case uint64:
		// math.MaxUint64 is 1.844674407e+19, which will be approximated in float64
		return 0, float64(^uint64(0))

	case float32:
		return -math.MaxFloat32, math.MaxFloat32
	case float64:
		return -math.MaxFloat64, math.MaxFloat64
	}

	// Fallback (shouldn't happen in normal usage)
	return -math.MaxFloat64, math.MaxFloat64
}

// intRangeOf is similar to rangeOf but returns a signed integer range. Useful
// for integer-target checks without going through float64 first.
func intRangeOf[U Number]() (int64, int64) {
	switch any(*new(U)).(type) {
	case int8:
		return math.MinInt8, math.MaxInt8
	case int16:
		return math.MinInt16, math.MaxInt16
	case int32:
		return math.MinInt32, math.MaxInt32
	case int:
		// For a 64-bit assumption:
		return math.MinInt64, math.MaxInt64
	case int64:
		return math.MinInt64, math.MaxInt64

	case uint8:
		return 0, math.MaxUint8
	case uint16:
		return 0, math.MaxUint16
	case uint32:
		return 0, math.MaxUint32
	case uint:
		return 0, int64(^uint64(0) >> 1) // if 64-bit, ^uint64(0) is all bits set
	case uint64:
		// for intRangeOf an unsigned, we can only represent it in int64 up to 2^63-1
		return 0, math.MaxInt64

	default:
		// Should not happen
		return math.MinInt64, math.MaxInt64
	}
}

// zero is a small helper that returns the zero value of a type.
func zero[T any]() T {
	var z T
	return z
}
