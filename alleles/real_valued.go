package alleles

import (
	"math/rand"
	"strconv"
)

// RealAllele is an allele based on a floating point value.
type RealAllele float64

// Copy returns a copy of this real allele
func (ra RealAllele) Copy() Allele {
	return RealAllele(0.0)
}

func (ra RealAllele) String() string {
	return strconv.FormatFloat(float64(ra), 'e', 4, 64)
}

// UniformBoundedRealFactory produces a new real valued
// allele between two bounds (min and max)
type UniformBoundedRealFactory struct {
	Min, Max RealAllele
}

// NormalRealFactory produces normally distributed alleles
// according to the relevant mean and standard deviation
type NormalRealFactory struct {
	Mean, Deviation RealAllele
}

// Random produces a new uniformly distributed allele.
func (u UniformBoundedRealFactory) Random() Allele {
	return RealAllele(rand.Float64())*(u.Max-u.Min) + u.Min
}

// Random produces a normally distributed allele.
func (n NormalRealFactory) Random() Allele {
	return RealAllele(rand.NormFloat64())*n.Deviation + n.Mean
}

// UinformBoundedRealMutator produces a mutator
// factory that will mutate values between min and max.
type UinformBoundedRealMutator struct {
	Min, Max RealAllele
}

// NormalRealMutator produces a mutator that uses a
// normal distribution for mutating values according to the given
// mean and distribution.
type NormalRealMutator struct {
	Mean, Deviation RealAllele
}

// Mutate produces a mutator that will mutate the allele
// between min and  max in the factory.
func (u UinformBoundedRealMutator) Mutate(a Allele) Allele {
	return a.(RealAllele)*(u.Max-u.Min)*RealAllele(rand.Float64()) + u.Min
}

// Mutate returns a new muator for the given normal distribution
func (n NormalRealMutator) Mutate(a Allele) Allele {
	delta := RealAllele(rand.NormFloat64())*n.Deviation + n.Mean
	return a.(RealAllele) + delta
}
