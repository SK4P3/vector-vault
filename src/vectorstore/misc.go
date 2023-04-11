package vectorstore

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	fmt.Printf("%s took %.12fs \n", name, time.Since(start).Seconds())
}
