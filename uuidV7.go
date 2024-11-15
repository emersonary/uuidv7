// This is a UUID v7 generator and parser.
// For public purposes, it has millisecond granularity.
// Three bytes, assumed to be random, are used to store nanosecond information.
// Therefore, for private purposes, it has nanosecond granularity.
package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func putUint48(b []byte, millis, nanos uint64) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(millis >> 40)
	b[1] = byte(millis >> 32)
	b[2] = byte(millis >> 24)
	b[3] = byte(millis >> 16)
	b[4] = byte(millis >> 8)
	b[5] = byte(millis)

	b[6] = byte(nanos >> 16)
	b[7] = byte(nanos >> 8)
	byt := byte(nanos)
	b[9] = byte(rand.Intn(255))
	b[8] = (b[8] & 0x0F) | (byt >> 4)
	b[9] = (b[9] & 0x0F) | (byt << 4)

	for i := 10; i < 16; i++ {
		b[i] = byte(rand.Intn(255))
	}

}

func uint48(b []byte) uint64 {
	return uint64(b[5]) |
		uint64(b[4])<<8 |
		uint64(b[3])<<16 |
		uint64(b[2])<<24 |
		uint64(b[1])<<32 |
		uint64(b[0])<<40
}

func uint24(b []byte) uint64 {
	return uint64(b[2]) |
		uint64(b[1])<<8 |
		uint64(b[0])<<16
}

func NewUUIDv7() uuid.UUID {
	return UUIDv7FromTimeStamp(time.Now())
}

func UUIDv7FromTimeStamp(timeStamp time.Time) uuid.UUID {
	unixTimeMilli := timeStamp.UnixMilli()
	unixTimeNano := timeStamp.UnixNano()
	nanosAfterMili := unixTimeNano - unixTimeMilli*1000000
	var uuid uuid.UUID
	putUint48(uuid[:], uint64(unixTimeMilli), uint64(nanosAfterMili))
	uuid[6] = (uuid[6] & 0x0F) | 0x70
	uuid[8] = (uuid[8] & 0x3F) | 0x80
	return uuid
}

func TimeStampFromUUIDv7(uuid uuid.UUID) time.Time {
	timestampMilli := uint48(uuid[0:6])
	byt := uuid[6:10]
	byt[0] = (byt[0] & 0x0F)
	byt[2] = (byt[2]&0x0F)<<4 + byt[3]>>4
	timeStampNano := uint24(byt[:3])
	return time.UnixMilli(int64(timestampMilli)).Add(time.Nanosecond * time.Duration(timeStampNano))
}
