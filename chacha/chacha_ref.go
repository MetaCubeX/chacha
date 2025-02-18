// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

//go:build !amd64 && !386

package chacha

import "encoding/binary"

func init() {
	useSSE2 = false
	useSSSE3 = false
	useAVX = false
	useAVX2 = false
}

func initialize(state *[64]byte, key []byte, nonce *[16]byte) {
	binary.LittleEndian.PutUint32(state[0:], sigma0)
	binary.LittleEndian.PutUint32(state[4:], sigma1)
	binary.LittleEndian.PutUint32(state[8:], sigma2)
	binary.LittleEndian.PutUint32(state[12:], sigma3)
	copy(state[16:], key[:])
	copy(state[48:], nonce[:])
}

func xorKeyStream(dst, src []byte, block, state *[64]byte, rounds int) int {
	return xorKeyStreamGeneric(dst, src, block, state, rounds)
}

func hChaCha20(out *[32]byte, nonce *[16]byte, key *[32]byte) {
	hChaCha20Generic(out, nonce, key)
}
