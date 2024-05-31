package genv_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	genv "github.com/abecodes/genv"
)

const (
	boolKey     = "BOOL"
	complexKey  = "COMPLEX"
	durationKey = "DURATION"
	floatKey    = "FLOAT"
	intKey      = "INT"
	invalidKey  = "INVALID"
	stringKey   = "STRING"
	timeKey     = "TIME"
	uintKey     = "UINT"

	boolValue     = "true"
	complexValue  = "10+10i"
	durationValue = "10s"
	floatValue    = "1.0"
	intValue      = "1"
	stringValue   = "string"
	timeValue     = "Thu, 30 May 2024 20:06:14 GMT"

	unBoolKey   = "UNBOOL"
	unNumberKey = "UNNUMBER"

	unBoolValue   = "not a bool"
	unNumberValue = "not an int"
)

var timeLocation, _ = time.LoadLocation("GMT")

var testEnvs = map[string]string{
	boolKey:     boolValue,
	complexKey:  complexValue,
	durationKey: durationValue,
	floatKey:    floatValue,
	intKey:      intValue,
	stringKey:   stringValue,
	timeKey:     timeValue,
	uintKey:     intValue,

	unBoolKey:   unBoolValue,
	unNumberKey: unNumberValue,
}

func complexComparer[T complex64 | complex128](got any, want any) bool {
	var t T

	switch any(t).(type) {
	case complex128:
		x, _ := got.(complex128)
		y, _ := want.(complex128)

		return real(x) == real(y) && imag(x) == imag(y)
	case complex64:
		x, _ := got.(complex64)
		y, _ := want.(complex64)

		return real(x) == real(y) && imag(x) == imag(y)
	}

	return false
}

func TestGenv(t *testing.T) {
	for k, v := range testEnvs {
		_ = os.Setenv(k, v)
	}

	tests := []struct {
		name     string
		got      any
		want     any
		comparer func(got any, want any) bool
	}{
		{
			name: "test string value in env",
			got:  genv.Get[string](stringKey),
			want: stringValue,
		},
		{
			name: "test string value not in env",
			got:  genv.Get[string](invalidKey),
			want: "",
		},
		{
			name: "test string value in env with default",
			got:  genv.GetWithDefault[string](stringKey, "default"),
			want: stringValue,
		},
		{
			name: "test string value not in env with default",
			got:  genv.GetWithDefault[string](invalidKey, "default"),
			want: "default",
		},
		{
			name: "test bool value in env",
			got:  genv.Get[bool](boolKey),
			want: true,
		},
		{
			name: "test bool value not in env",
			got:  genv.Get[bool](invalidKey),
			want: false,
		},
		{
			name: "test bool value in env with default",
			got:  genv.GetWithDefault[bool](boolKey, false),
			want: true,
		},
		{
			name: "test bool value not in env with default",
			got:  genv.GetWithDefault[bool](invalidKey, false),
			want: false,
		},
		{
			name: "test bool value with unparseable value",
			got:  genv.Get[bool](unBoolKey),
			want: false,
		},
		{
			name: "test uint value in env",
			got:  genv.Get[uint](uintKey),
			want: uint(1),
		},
		{
			name: "test uint value not in env",
			got:  genv.Get[uint](invalidKey),
			want: uint(0),
		},
		{
			name: "test uint value in env with default",
			got:  genv.GetWithDefault[uint](uintKey, 0),
			want: uint(1),
		},
		{
			name: "test uint value not in env with default",
			got:  genv.GetWithDefault[uint](invalidKey, 1),
			want: uint(1),
		},
		{
			name: "test uint value with unparseable value",
			got:  genv.Get[uint](unNumberKey),
			want: uint(0),
		},
		{
			name: "test uint8 value in env",
			got:  genv.Get[uint8](uintKey),
			want: uint8(1),
		},
		{
			name: "test uint8 value not in env",
			got:  genv.Get[uint8](invalidKey),
			want: uint8(0),
		},
		{
			name: "test uint8 value in env with default",
			got:  genv.GetWithDefault[uint8](uintKey, 0),
			want: uint8(1),
		},
		{
			name: "test uint8 value not in env with default",
			got:  genv.GetWithDefault[uint8](invalidKey, 1),
			want: uint8(1),
		},
		{
			name: "test uint8 value with unparseable value",
			got:  genv.Get[uint8](unNumberKey),
			want: uint8(0),
		},
		{
			name: "test uint16 value in env",
			got:  genv.Get[uint16](uintKey),
			want: uint16(1),
		},
		{
			name: "test uint16 value not in env",
			got:  genv.Get[uint16](invalidKey),
			want: uint16(0),
		},
		{
			name: "test uint16 value in env with default",
			got:  genv.GetWithDefault[uint16](uintKey, 0),
			want: uint16(1),
		},
		{
			name: "test uint16 value not in env with default",
			got:  genv.GetWithDefault[uint16](invalidKey, 1),
			want: uint16(1),
		},
		{
			name: "test uint16 value with unparseable value",
			got:  genv.Get[uint16](unNumberKey),
			want: uint16(0),
		},
		{
			name: "test uint32 value in env",
			got:  genv.Get[uint32](uintKey),
			want: uint32(1),
		},
		{
			name: "test uint32 value not in env",
			got:  genv.Get[uint32](invalidKey),
			want: uint32(0),
		},
		{
			name: "test uint32 value in env with default",
			got:  genv.GetWithDefault[uint32](uintKey, 0),
			want: uint32(1),
		},
		{
			name: "test uint32 value not in env with default",
			got:  genv.GetWithDefault[uint32](invalidKey, 1),
			want: uint32(1),
		},
		{
			name: "test uint32 value with unparseable value",
			got:  genv.Get[uint32](unNumberKey),
			want: uint32(0),
		},
		{
			name: "test uint64 value in env",
			got:  genv.Get[uint64](uintKey),
			want: uint64(1),
		},
		{
			name: "test uint64 value not in env",
			got:  genv.Get[uint64](invalidKey),
			want: uint64(0),
		},
		{
			name: "test uint64 value in env with default",
			got:  genv.GetWithDefault[uint64](uintKey, 0),
			want: uint64(1),
		},
		{
			name: "test uint64 value not in env with default",
			got:  genv.GetWithDefault[uint64](invalidKey, 1),
			want: uint64(1),
		},
		{
			name: "test uint64 value with unparseable value",
			got:  genv.Get[uint64](unNumberKey),
			want: uint64(0),
		},
		{
			name: "test int value in env",
			got:  genv.Get[int](intKey),
			want: int(1),
		},
		{
			name: "test int value not in env",
			got:  genv.Get[int](invalidKey),
			want: int(0),
		},
		{
			name: "test int value in env with default",
			got:  genv.GetWithDefault[int](intKey, 0),
			want: int(1),
		},
		{
			name: "test int value not in env with default",
			got:  genv.GetWithDefault[int](invalidKey, 1),
			want: int(1),
		},
		{
			name: "test int value with unparseable value",
			got:  genv.Get[int](unNumberKey),
			want: int(0),
		},
		{
			name: "test int8 value in env",
			got:  genv.Get[int8](intKey),
			want: int8(1),
		},
		{
			name: "test int8 value not in env",
			got:  genv.Get[int8](invalidKey),
			want: int8(0),
		},
		{
			name: "test int8 value in env with default",
			got:  genv.GetWithDefault[int8](intKey, 0),
			want: int8(1),
		},
		{
			name: "test int8 value not in env with default",
			got:  genv.GetWithDefault[int8](invalidKey, 1),
			want: int8(1),
		},
		{
			name: "test int8 value with unparseable value",
			got:  genv.Get[int8](unNumberKey),
			want: int8(0),
		},
		{
			name: "test int16 value in env",
			got:  genv.Get[int16](intKey),
			want: int16(1),
		},
		{
			name: "test int16 value not in env",
			got:  genv.Get[int16](invalidKey),
			want: int16(0),
		},
		{
			name: "test int16 value in env with default",
			got:  genv.GetWithDefault[int16](intKey, 0),
			want: int16(1),
		},
		{
			name: "test int16 value not in env with default",
			got:  genv.GetWithDefault[int16](invalidKey, 1),
			want: int16(1),
		},
		{
			name: "test int16 value with unparseable value",
			got:  genv.Get[int16](unNumberKey),
			want: int16(0),
		},
		{
			name: "test int32 value in env",
			got:  genv.Get[int32](intKey),
			want: int32(1),
		},
		{
			name: "test int32 value not in env",
			got:  genv.Get[int32](invalidKey),
			want: int32(0),
		},
		{
			name: "test int32 value in env with default",
			got:  genv.GetWithDefault[int32](intKey, 0),
			want: int32(1),
		},
		{
			name: "test int32 value not in env with default",
			got:  genv.GetWithDefault[int32](invalidKey, 1),
			want: int32(1),
		},
		{
			name: "test int32 value with unparseable value",
			got:  genv.Get[int32](unNumberKey),
			want: int32(0),
		},
		{
			name: "test int64 value in env",
			got:  genv.Get[int64](intKey),
			want: int64(1),
		},
		{
			name: "test int64 value not in env",
			got:  genv.Get[int64](invalidKey),
			want: int64(0),
		},
		{
			name: "test int64 value in env with default",
			got:  genv.GetWithDefault[int64](intKey, 0),
			want: int64(1),
		},
		{
			name: "test int64 value not in env with default",
			got:  genv.GetWithDefault[int64](invalidKey, 1),
			want: int64(1),
		},
		{
			name: "test int64 value with unparseable value",
			got:  genv.Get[int64](unNumberKey),
			want: int64(0),
		},
		{
			name: "test float32 value in env",
			got:  genv.Get[float32](floatKey),
			want: float32(1),
		},
		{
			name: "test float32 value not in env",
			got:  genv.Get[float32](invalidKey),
			want: float32(0),
		},
		{
			name: "test float32 value in env with default",
			got:  genv.GetWithDefault[float32](floatKey, 0),
			want: float32(1),
		},
		{
			name: "test float32 value not in env with default",
			got:  genv.GetWithDefault[float32](invalidKey, 1),
			want: float32(1),
		},
		{
			name: "test float32 value with unparseable value",
			got:  genv.Get[float32](unNumberKey),
			want: float32(0),
		},
		{
			name: "test float64 value in env",
			got:  genv.Get[float64](floatKey),
			want: float64(1),
		},
		{
			name: "test float64 value not in env",
			got:  genv.Get[float64](invalidKey),
			want: float64(0),
		},
		{
			name: "test float64 value in env with default",
			got:  genv.GetWithDefault[float64](floatKey, 0),
			want: float64(1),
		},
		{
			name: "test float64 value not in env with default",
			got:  genv.GetWithDefault[float64](invalidKey, 1),
			want: float64(1),
		},
		{
			name: "test float64 value with unparseable value",
			got:  genv.Get[float64](unNumberKey),
			want: float64(0),
		},
		{
			name:     "test complex64 value in env",
			got:      genv.Get[complex64](complexKey),
			want:     complex64(complex(10, 10)),
			comparer: complexComparer[complex64],
		},
		{
			name:     "test complex64 value not in env",
			got:      genv.Get[complex64](invalidKey),
			want:     complex(0, 0),
			comparer: complexComparer[complex64],
		},
		{
			name:     "test complex64 value in env with default",
			got:      genv.GetWithDefault[complex64](complexKey, complex(0, 0)),
			want:     complex64(complex(10, 10)),
			comparer: complexComparer[complex64],
		},
		{
			name:     "test complex64 value not in env with default",
			got:      genv.GetWithDefault[complex64](invalidKey, complex(1, 1)),
			want:     complex64(complex(1, 1)),
			comparer: complexComparer[complex64],
		},
		{
			name:     "test complex64 value with unparseable value",
			got:      genv.Get[complex64](unNumberKey),
			want:     complex(0, 0),
			comparer: complexComparer[complex64],
		},
		{
			name:     "test complex128 value in env",
			got:      genv.Get[complex128](complexKey),
			want:     complex128(complex(10, 10)),
			comparer: complexComparer[complex128],
		},
		{
			name:     "test complex128 value not in env",
			got:      genv.Get[complex128](invalidKey),
			want:     complex(0, 0),
			comparer: complexComparer[complex128],
		},
		{
			name:     "test complex128 value in env with default",
			got:      genv.GetWithDefault[complex128](complexKey, complex(0, 0)),
			want:     complex128(complex(10, 10)),
			comparer: complexComparer[complex128],
		},
		{
			name:     "test complex128 value not in env with default",
			got:      genv.GetWithDefault[complex128](invalidKey, complex(1, 1)),
			want:     complex128(complex(1, 1)),
			comparer: complexComparer[complex128],
		},
		{
			name:     "test complex128 value with unparseable value",
			got:      genv.Get[complex128](unNumberKey),
			want:     complex(0, 0),
			comparer: complexComparer[complex128],
		},
		{
			name: "test time value in env",
			got:  genv.Get[time.Time](timeKey),
			want: time.Date(2024, 5, 30, 20, 6, 14, 0, timeLocation),
			comparer: func(got, want any) bool {
				x, _ := got.(time.Time)
				y, _ := want.(time.Time)

				return x.Equal(y)
			},
		},
		{
			name: "test time value not in env",
			got:  genv.Get[time.Time](invalidKey),
			want: time.Time{},
		},
		{
			name: "test time value in env with default",
			got:  genv.GetWithDefault[time.Time](timeKey, time.Time{}),
			want: time.Date(2024, 5, 30, 20, 6, 14, 0, timeLocation),
			comparer: func(got, want any) bool {
				x, _ := got.(time.Time)
				y, _ := want.(time.Time)

				return x.Equal(y)
			},
		},
		{
			name: "test time value not in env with default",
			got: genv.GetWithDefault[time.Time](
				invalidKey,
				time.Date(2000, 5, 5, 0, 0, 0, 0, time.UTC),
			),
			want: time.Date(2000, 5, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "test time value with unparseable value",
			got:  genv.Get[time.Time](unNumberKey),
			want: time.Time{},
		},
		{
			name: "test duration value in env",
			got:  genv.Get[time.Duration](durationKey),
			want: time.Duration(10 * time.Second),
		},
		{
			name: "test duration value not in env",
			got:  genv.Get[time.Duration](invalidKey),
			want: 0 * time.Second,
		},
		{
			name: "test duration value in env with default",
			got:  genv.GetWithDefault[time.Duration](durationKey, 0*time.Second),
			want: 10 * time.Second,
		},
		{
			name: "test duration value not in env with default",
			got:  genv.GetWithDefault[time.Duration](invalidKey, 100*time.Second),
			want: 100 * time.Second,
		},
		{
			name: "test duration value with unparseable value",
			got:  genv.Get[time.Duration](unNumberKey),
			want: 0 * time.Second,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.comparer != nil {
				if !tt.comparer(tt.got, tt.want) {
					t.Errorf("Get = %v, want %v", tt.got, tt.want)
				}

				return
			}

			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Get = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
