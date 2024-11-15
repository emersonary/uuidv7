package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var now = time.Now().Truncate(time.Second)

func calcedTimeFromInt(i int64) time.Time {

	return now.Add(time.Minute * time.Duration(i)).
		Add(time.Second * time.Duration(i)).
		Add(time.Millisecond * time.Duration(i)).
		Add(time.Microsecond * time.Duration(i)).
		Add(time.Nanosecond * time.Duration(i))

}

func TestParse(t *testing.T) {

	loops := 10000000

	for i := range loops {

		calcedTime := calcedTimeFromInt(int64(i))
		uuid := UUIDv7FromTimeStamp(calcedTime)
		extractedTime := TimeStampFromUUIDv7(uuid)
		require.Equal(t, calcedTime.UTC().Hour(), extractedTime.UTC().Hour())
		require.Equal(t, calcedTime.Minute(), extractedTime.Minute())
		require.Equal(t, calcedTime.Second(), extractedTime.Second())

	}

}

type TimeAndUUID struct {
	timeStamp time.Time
	uuid      string
}

func TestOrder(t *testing.T) {

	loops := 10000000

	listTimeAndUUID := make([]TimeAndUUID, loops)

	for i := range loops {

		calcedTime := calcedTimeFromInt(int64(i))
		uuid := UUIDv7FromTimeStamp(calcedTime)
		listTimeAndUUID[i].timeStamp = calcedTime
		listTimeAndUUID[i].uuid = uuid.String()

		if i > 0 {
			require.True(t, listTimeAndUUID[i-1].timeStamp.Before(listTimeAndUUID[i].timeStamp))
			require.True(t, listTimeAndUUID[i-1].uuid < listTimeAndUUID[i].uuid)
		}
	}

}
