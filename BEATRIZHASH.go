package BEATRIZHASH

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

// BeatrizHash computes a hash using the BeatrizHash algorithm.
func BeatrizHash(lastHash string, expectedHash string, difficulty int) string {
	hasher := sha1.New()
	for i := 0; i <= difficulty*100; i++ {
		hasher.Write([]byte(lastHash + strconv.Itoa(i)))
		hash := hex.EncodeToString(hasher.Sum(nil)[:])
		if hash == expectedHash {
			return hash
		}
		hasher.Reset()
	}
	return ""
}
