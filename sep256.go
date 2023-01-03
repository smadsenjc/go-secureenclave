//go:build darwin && cgo
// +build darwin,cgo

package secureenclave

// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/lib -lSEP256
// #include "SEP256.h"
import "C"

func CreateP256PrivateKey() []byte {
	privateKeyHandle := make([]C.char, 300)
	length := C.int(300)
	result := C.GenerateSEP256Key(&privateKeyHandle[0], &length)
	if !result {
		panic("")
	}

	privKey := make([]byte, length)
	for i := 0; i < (int)(length); i++ {
		privKey[i] = (byte)(privateKeyHandle[i])
	}

	// marshal x, y into x9.62 format
	// x962pubKey := make([]byte, 0)
	// x962pubKey = append(x962pubKey[:], 0x04)
	// x962pubKey = append(x962pubKey[:], pubKey.ECCParameters.Point.XRaw...)
	// x962pubKey = append(x962pubKey[:], pubKey.ECCParameters.Point.YRaw...)

	return privKey
}
