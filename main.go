package main

import (
	"learnGoWithTests/mock"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	mock.Countdown(os.Stdout, sleeper)
}
