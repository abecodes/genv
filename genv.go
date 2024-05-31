package genv

import (
	"os"
	"strconv"
	"time"
)

type envType interface {
	string |
		bool |
		uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		complex64 | complex128 |
		time.Time | time.Duration
}

var (
	base   = 10
	bit8   = 8
	bit16  = 16
	bit32  = 32
	bit64  = 64
	bit128 = 128
)

// Get will try to retrieve a value for a given key from the env and return it as type T.
// If this is not possible, the default value for the given type T will be returned.
//
// If a time.Time value is requested, it needs to be available in the env in the RFC1123 format, e.g:
// "Thu, 30 May 2024 21:01:37 GMT"
func Get[T envType](key string) T {
	val := os.Getenv(key)

	return getVal[T](val)
}

// GetWithDefault will try to retrieve a value for a given key from the env and return it as type T.
// If this is not possible, the passed in default value will be returned.
//
// If a time.Time value is requested, it needs to be available in the env in the RFC1123 format, e.g:
// "Thu, 30 May 2024 21:01:37 GMT"
func GetWithDefault[T envType](key string, defaultVal T) T {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	return getVal[T](val)
}

func getVal[T envType](val string) T {
	var v T

	if val == "" {
		return v
	}

	switch any(v).(type) {
	case string:
		v = any(val).(T)
	case bool:
		x, err := strconv.ParseBool(val)
		if err != nil {
			return v
		}

		return any(bool(x)).(T)
	case uint:
		x, err := strconv.ParseUint(val, base, 0)
		if err != nil {
			return v
		}

		return any(uint(x)).(T)
	case uint8:
		x, err := strconv.ParseUint(val, base, bit8)
		if err != nil {
			return v
		}

		return any(uint8(x)).(T)
	case uint16:
		x, err := strconv.ParseUint(val, base, bit16)
		if err != nil {
			return v
		}

		return any(uint16(x)).(T)
	case uint32:
		x, err := strconv.ParseUint(val, base, bit32)
		if err != nil {
			return v
		}

		return any(uint32(x)).(T)
	case uint64:
		x, err := strconv.ParseUint(val, base, bit64)
		if err != nil {
			return v
		}

		return any(x).(T)
	case int:
		x, err := strconv.ParseInt(val, base, 0)
		if err != nil {
			return v
		}

		return any(int(x)).(T)
	case int8:
		x, err := strconv.ParseInt(val, base, bit8)
		if err != nil {
			return v
		}

		return any(int8(x)).(T)
	case int16:
		x, err := strconv.ParseInt(val, base, bit16)
		if err != nil {
			return v
		}

		return any(int16(x)).(T)
	case int32:
		x, err := strconv.ParseInt(val, base, bit32)
		if err != nil {
			return v
		}

		return any(int32(x)).(T)
	case int64:
		x, err := strconv.ParseInt(val, base, bit64)
		if err != nil {
			return v
		}

		return any(x).(T)
	case float32:
		x, err := strconv.ParseFloat(val, bit32)
		if err != nil {
			return v
		}

		return any(float32(x)).(T)
	case float64:
		x, err := strconv.ParseFloat(val, bit64)
		if err != nil {
			return v
		}

		return any(x).(T)
	case complex64:
		x, err := strconv.ParseComplex(val, bit64)
		if err != nil {
			return v
		}

		return any(complex64(x)).(T)
	case complex128:
		x, err := strconv.ParseComplex(val, bit128)
		if err != nil {
			return v
		}

		return any(x).(T)
	case time.Duration:
		x, err := time.ParseDuration(val)
		if err != nil {
			return v
		}

		return any(x).(T)
	case time.Time:
		x, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return v
		}

		return any(x).(T)
	}

	return v
}
