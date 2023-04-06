module github.com/baohan10/reedsolomon

go 1.17

require github.com/klauspost/cpuid/v2 v2.1.1

require golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e // indirect

retract (
	v1.11.2 // https://github.com/baohan10/reedsolomon/pull/229
	[v1.11.3, v1.11.5] // https://github.com/baohan10/reedsolomon/pull/238
	v1.11.6 // https://github.com/baohan10/reedsolomon/issues/240
)
