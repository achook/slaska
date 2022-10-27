package main

import (
	"fmt"
	"time"
)

type MoneyData struct {
	LastUpdate int64
	Money      string
}

type FormattedDelta struct {
	Delta int
	Ago   string
}

func (m MoneyData) MinutesSinceLastUpdate() FormattedDelta {
	current := time.Now().Unix()
	delta := int((current - m.LastUpdate) / 60)

	if delta == 1 {
		return FormattedDelta{delta, "minutę"}
	}

	if delta%10 <= 1 || delta%10 >= 5 || (delta%100 >= 10 && delta%100 <= 19) {
		return FormattedDelta{delta, "minut"}
	}

	return FormattedDelta{delta, "minuty"}
	
}

func (f FormattedDelta) String() string {
	return fmt.Sprintf("%d %s", f.Delta, f.Ago)
}
