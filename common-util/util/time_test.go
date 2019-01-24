package util

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/utils"
	"testing"
	"time"
)

//Used to make sure I had calculated time correctly.  Leaving because it's handy for testing timex
func TestGetDurationDelta(t *testing.T) {
	dfltDuration := time.Second * 5
	now := time.Now()
	future := utils.ToMilliSeconds(now.Add(time.Second * 2))
	delta := future - utils.ToMilliSeconds(time.Now()) //do milliseconds since that's what you need

	delta2 := time.Millisecond*time.Duration(delta) + dfltDuration

	durationDelta := now.Add(delta2)
	fmt.Printf("Delta: %v :: Delta2: %v :: timeDelta: %v\n", delta, delta2, durationDelta)
}

func TestGetEpochMinute(t *testing.T) {
	dispatchEpoch := time.Date(2018, time.December, 4, 0, 0, 0, 0, time.UTC)
	minutesSinceEpoch := time.Now().Sub(dispatchEpoch).Minutes()

	fmt.Printf("Epoch: %v\nMinutes Since Epoch: %v", dispatchEpoch.UnixNano(), int64(minutesSinceEpoch))
}

func TestFormat(tst *testing.T) {

	t := time.Now()

	fmt.Printf("%d-%02d-%02d-%02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
