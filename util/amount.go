// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util

import (
	"github.com/danannet/danad/domain/consensus/utils/constants"
	"github.com/pkg/errors"
	"math"
	"strconv"
)

// AmountUnit describes a method of converting an Amount to something
// other than the base unit of a dana. The value of the AmountUnit
// is the exponent component of the decadic multiple to convert from
// an amount in dana to an amount counted in units.
type AmountUnit int

// These constants define various units used when describing a dana
// monetary amount.
const (
	AmountMegaDANA  AmountUnit = 6
	AmountKiloDANA  AmountUnit = 3
	AmountDANA      AmountUnit = 0
	AmountMilliDANA AmountUnit = -3
	AmountMicroDANA AmountUnit = -6
	AmountSompi    AmountUnit = -8
)

// String returns the unit as a string. For recognized units, the SI
// prefix is used, or "Sompi" for the base unit. For all unrecognized
// units, "1eN DANA" is returned, where N is the AmountUnit.
func (u AmountUnit) String() string {
	switch u {
	case AmountMegaDANA:
		return "MDANA"
	case AmountKiloDANA:
		return "kDANA"
	case AmountDANA:
		return "DANA"
	case AmountMilliDANA:
		return "mDANA"
	case AmountMicroDANA:
		return "μDANA"
	case AmountSompi:
		return "Sompi"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + " DANA"
	}
}

// Amount represents the base dana monetary unit (colloquially referred
// to as a `Sompi'). A single Amount is equal to 1e-8 of a dana.
type Amount uint64

// round converts a floating point number, which may or may not be representable
// as an integer, to the Amount integer type by rounding to the nearest integer.
// This is performed by adding or subtracting 0.5 depending on the sign, and
// relying on integer truncation to round the value to the nearest Amount.
func round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}
	return Amount(f + 0.5)
}

// NewAmount creates an Amount from a floating point value representing
// some value in dana. NewAmount errors if f is NaN or +-Infinity, but
// does not check that the amount is within the total amount of dana
// producible as f may not refer to an amount at a single moment in time.
//
// NewAmount is for specifically for converting DANA to Sompi.
// For creating a new Amount with an int64 value which denotes a quantity of Sompi,
// do a simple type conversion from type int64 to Amount.
// TODO: Refactor NewAmount. When amounts are more than 1e9 DANA, the precision
// can be higher than one sompi (1e9 and 1e9+1e-8 will result as the same number)
func NewAmount(f float64) (Amount, error) {
	// The amount is only considered invalid if it cannot be represented
	// as an integer type. This may happen if f is NaN or +-Infinity.
	switch {
	case math.IsNaN(f):
		fallthrough
	case math.IsInf(f, 1):
		fallthrough
	case math.IsInf(f, -1):
		return 0, errors.New("invalid dana amount")
	}

	return round(f * constants.SompiPerDana), nil
}

// ToUnit converts a monetary amount counted in dana base units to a
// floating point value representing an amount of dana.
func (a Amount) ToUnit(u AmountUnit) float64 {
	return float64(a) / math.Pow10(int(u+8))
}

// ToDANA is the equivalent of calling ToUnit with AmountDANA.
func (a Amount) ToDANA() float64 {
	return a.ToUnit(AmountDANA)
}

// Format formats a monetary amount counted in dana base units as a
// string for a given unit. The conversion will succeed for any unit,
// however, known units will be formated with an appended label describing
// the units with SI notation, or "Sompi" for the base unit.
func (a Amount) Format(u AmountUnit) string {
	units := " " + u.String()
	return strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64) + units
}

// String is the equivalent of calling Format with AmountDANA.
func (a Amount) String() string {
	return a.Format(AmountDANA)
}

// MulF64 multiplies an Amount by a floating point value. While this is not
// an operation that must typically be done by a full node or wallet, it is
// useful for services that build on top of dana (for example, calculating
// a fee by multiplying by a percentage).
func (a Amount) MulF64(f float64) Amount {
	return round(float64(a) * f)
}
