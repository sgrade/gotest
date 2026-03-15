// 1622. Fancy Sequence
// https://leetcode.com/problems/fancy-sequence/

package fancysequence

// Based on Editorial's Approach 2: Use multiplicative inverse during the append operation.
const mod = 1000000007

type Fancy struct {
	vals []int
	a, b int // affine transform: val → a*val + b
}

func Constructor() Fancy {
	return Fancy{a: 1}
}

// modpow returns x^y % mod via binary exponentiation.
func modpow(x, y int) int {
	result := 1
	x %= mod
	for y > 0 {
		if y&1 != 0 {
			result = result * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return result
}

// modinv returns the modular inverse of x (mod is prime).
func modinv(x int) int {
	return modpow(x, mod-2)
}

// Append stores val after reversing the current affine transform.
func (f *Fancy) Append(val int) {
	f.vals = append(f.vals, (val-f.b+mod)%mod*modinv(f.a)%mod)
}

func (f *Fancy) AddAll(inc int) {
	f.b = (f.b + inc) % mod
}

func (f *Fancy) MultAll(m int) {
	f.a = f.a * m % mod
	f.b = f.b * m % mod
}

func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.vals) {
		return -1
	}
	return (f.a*f.vals[idx]%mod + f.b) % mod
}
