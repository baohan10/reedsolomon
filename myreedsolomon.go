package reedsolomon

type ReedSolomon struct {
	r   Encoder
	ecK int
	ecN int
	ecP int
}

func (r *ReedSolomon) Encode(all [][]byte) {
	r.r.Encode(all)
}

func (r *ReedSolomon) Reconstruct(all [][]byte) error {
	return r.r.ReconstructData(all)
}

func GetReedSolomonIns(k int, n int) *ReedSolomon {
	enc := ReedSolomon{
		ecN: n,
		ecK: k,
		ecP: n - k,
	}
	enc.r, _ = New(k, n-k)
	return &enc
}
