package darwin

import (
	crand "crypto/rand"
	rand "math/rand"

	"encoding/binary"
	"log"
)

var src cryptoSource
var RandSource = rand.New(src)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// Random number in range [0, 1)
func Random() float64 {
	// seed := int64(time.Now().Nanosecond())
	return RandSource.Float64()
}

func RandIntn(n int) int {
	// seed := int64(time.Now().Nanosecond())
	intn := RandSource.Intn(n)
	// fmt.Printf("%d ", intn)
	return intn
}

// Flip and coin with prob and report if is head
func Flip(prob float64) bool {
	// seed := int64(time.Now().Nanosecond())
	chance := RandSource.Float64()
	return chance <= prob
}
