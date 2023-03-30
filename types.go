package main

import (
	"fmt"
	"time"
)

type MoneyData struct {
	LastUpdate int64
	Money      string
	NextUpdate int64
}

type FormattedDelta struct {
	Delta   int
	Minutes string
}

func FromInt(delta int) FormattedDelta {
	if delta == 1 {
		return FormattedDelta{delta, "minutę"}
	}

	if delta%10 <= 1 || delta%10 >= 5 || (delta%100 >= 10 && delta%100 <= 19) {
		return FormattedDelta{delta, "minut"}
	}

	return FormattedDelta{delta, "minuty"}
}

func (m MoneyData) MinutesSinceLastUpdate() FormattedDelta {
	current := time.Now().Unix()
	delta := int((current - m.LastUpdate) / 60)
	return FromInt(delta)
}

func (m MoneyData) MinutesUntilNextUpdate() FormattedDelta {
	current := time.Now().Unix()
	delta := int((m.NextUpdate - current) / 60)
	return FromInt(delta)
}

func (f FormattedDelta) String() string {
	return fmt.Sprintf("%d %s", f.Delta, f.Minutes)
}
