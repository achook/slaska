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

	if delta%10 > 1 && delta%10 < 5 {
		return FormattedDelta{delta, "minuty"}
	}

	return FormattedDelta{delta, "minut"}
}

func (f FormattedDelta) String() string {
	return fmt.Sprintf("%d %s", f.Delta, f.Ago)
}
