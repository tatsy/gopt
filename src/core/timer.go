package core

import (
    "fmt"
    "time"
)

type Timer struct {
    startTime, endTime time.Time
}

func NewTimer() *Timer {
    return &Timer{}
}

func (t *Timer) Start() {
    t.startTime = time.Now()
}

func (t *Timer) Stop() {
    t.endTime = time.Now()
    duration := t.endTime.Sub(t.startTime)
    hours := int(duration.Hours())
    minutes := int(duration.Minutes()) % 60
    seconds := int(duration.Seconds()) % 60
    fmt.Printf("Time: ")
    if hours != 0 {
        fmt.Printf("%d hours ", hours)
    }
    if hours != 0 || minutes != 0 {
        fmt.Printf("%d min ", minutes)
    }
    if hours != 0 || minutes != 0 {
        fmt.Printf("%d sec\n", seconds)
    } else {
        fmt.Printf("%.2f sec\n", duration.Seconds())
    }
}
