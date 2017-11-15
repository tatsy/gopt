package core

const (
	maxUint = ^uint(0)
)

// Random is a random number generator with XOR shift.
type Random struct {
	seed [4]uint
}

// NewRandom creates a new generator with the seed.
func NewRandom(seed int64) *Random {
	r := new(Random)
	if seed == 0 {
		r.seed = [4]uint{123456789, 362436069, 521288629, 88675123}
	} else {
		s := uint(seed)
		for i := 1; i <= 4; i++ {
			r.seed[i-1] = s
			s = 1812433253*(s^(s>>30)) + uint(i)
		}
	}
	return r
}

// Uint returns a random integer of uint type
func (r *Random) Uint() uint {
	t := r.seed[0] ^ (r.seed[0] << 11)
	r.seed[0] = r.seed[1]
	r.seed[1] = r.seed[2]
	r.seed[2] = r.seed[3]
	r.seed[3] = (r.seed[3] ^ (r.seed[3] >> 19)) ^ (t ^ (t >> 8))
	return r.seed[3]
}

// Int returns a random integer of int type
func (r *Random) Int() int {
	return int(r.Uint())
}

// Float32 returns a random floating point number of float32 type
func (r *Random) Float32() float32 {
	return float32(r.Uint()) / float32(maxUint)
}

// Float64 returns a random floating point number of float64 type
func (r *Random) Float64() float64 {
	return float64(r.Uint()) / float64(maxUint)
}
