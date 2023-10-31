package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
)

// https://www.includehelp.com/cryptography/digital-signature-algorithm-dsa.aspx
func main() {
	// 1. Key Generation
	pub, priv := generateKey()

	fmt.Println("public key:", pub)
	fmt.Println("private key:", priv)

	// 2. Signature Generation
	// For a message M:
	M := "hello world, this is my message"

	// a. Compute its hash, H(M).
	hashedPayload := sha1.Sum([]byte(M))

	s := new(big.Int)
	r := new(big.Int)
	k := new(big.Int)
	var err error
	for {
		// b. Pick a random number k from [1, q-1].
		k, err = rand.Int(rand.Reader, new(big.Int).Sub(pub.q, big.NewInt(1)))
		if err != nil {
			panic(err)
		}

		fmt.Println("in r k:", k)

		// c. Compute r = (g^k mod p) mod q. If r = 0, choose another k.
		r = new(big.Int).Mod(new(big.Int).Exp(pub.g, k, pub.p), pub.q)
		if r.Cmp(big.NewInt(0)) != 1 {
			continue
		}

		break
	}

	for {
		fmt.Println("in s k:", k)
		// d. Compute k^-1 mod q, the modular inverse of k with respect to q.
		kInverse := new(big.Int).ModInverse(k, pub.q)

		// e. Compute s = k^-1 * (H(M) + x*r) mod q. If s = 0, choose another k and repeat from step c.
		s = new(big.Int).Mod(new(big.Int).Mul(kInverse, new(big.Int).Add(new(big.Int).SetBytes(hashedPayload[:]), new(big.Int).Mul(priv.x, r))), pub.q)
		if s.Cmp(big.NewInt(0)) != 1 {
			k, err = rand.Int(rand.Reader, new(big.Int).Sub(pub.q, big.NewInt(1)))
			if err != nil {
				panic(err)
			}

			continue
		}

		break
	}

	// The signature for the message M is the pair (r, s). M is appended in order to verify the signature later on.
	sig := &signature{r, s, M}

	// 3. Signature Verification

	// a. Compute its hash, H(M).
	hashedPayload = sha1.Sum([]byte(sig.m))

	// b. Check if 0 < r < q and 0 < s < q, otherwise reject the signature.
	if sig.r.Cmp(big.NewInt(0)) != 1 || sig.r.Cmp(pub.q) != -1 || sig.s.Cmp(big.NewInt(0)) != 1 || sig.s.Cmp(pub.q) != -1 {
		panic("invalid signature")
	}

	// c. Compute w = s^-1 mod q.
	w := new(big.Int).ModInverse(sig.s, pub.q)

	// d. Compute u1 = H(M) * w mod q.
	u1 := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).SetBytes(hashedPayload[:]), w), pub.q)

	// e. Compute u2 = r * w mod q.
	u2 := new(big.Int).Mod(new(big.Int).Mul(sig.r, w), pub.q)

	// f. Compute v = ((g^u1 * y^u2) mod p) mod q.
	// If v = r, then the signature is valid; otherwise, it's invalid.
	v := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(pub.g, u1, pub.p), new(big.Int).Exp(pub.y, u2, pub.p)), pub.p)
	v.Mod(v, pub.q)
	if v.Cmp(sig.r) != 0 {
		panic("invalid signature")
	}

	// The signature is valid.
	fmt.Println("congratulations, the signature is valid")
}

type signature struct {
	r *big.Int
	s *big.Int
	m string
}

type publicKey struct {
	p *big.Int
	q *big.Int
	g *big.Int
	y *big.Int
}

type privateKey struct {
	x *big.Int
}

func generateKey() (*publicKey, *privateKey) {
	// Firstly, choose a prime number q, which is called the prime divisor in this.
	// q is typically 160 bits long.
	q, err := rand.Prime(rand.Reader, 160)
	if err != nil {
		panic(err)
	}

	// Then, choose another primer number p, such that p-1 mod q = 0. p is called the prime modulus in this.
	p := new(big.Int)
	for {
		// Select a random 864-bit number L. (Note: 1024 (bit length of p) - 160 (bit length of q) = 864).
		L, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 863)) // Random 864-bit number
		if err != nil {
			panic(err)
		}

		// Compute pCandidate = L * q + 1.
		pCandidate := new(big.Int).Add(new(big.Int).Mul(L, q), big.NewInt(1))
		if pCandidate.ProbablyPrime(20) { // Adjust the confidence level as necessary
			p = pCandidate
			break
		}
	}

	// Then, choose an integer g, such that 1 < g < p, g**q mod p = 1 and g = h**((p–1)/q) mod p. q is also called g's multiplicative order modulo p in this algorithm.
	g := new(big.Int)
	for {
		// choose h in order to calculate g later on
		// h is called the auxiliary value in this.
		// h should be a number between 1 and p-1, such that g = h^((p-1)/q) mod p > 1
		h, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
		if err != nil {
			panic(err)
		}

		fmt.Println("find g", g)

		g = new(big.Int).Exp(h, new(big.Int).Div(new(big.Int).Sub(p, big.NewInt(1)), q), p)
		if g.Cmp(big.NewInt(1)) == 1 {
			break
		}
	}

	// Then, choose an integer, such that 0 < x < q - 1 for this.
	x, err := rand.Int(rand.Reader, new(big.Int).Sub(q, big.NewInt(1)))
	if err != nil {
		panic(err)
	}

	// Now, compute y as g**x mod p.
	y := new(big.Int).Exp(g, x, p)

	return &publicKey{p, q, g, y}, &privateKey{x}
}
