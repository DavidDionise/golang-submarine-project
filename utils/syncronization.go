package utils

import "sync"

var (
	Mutex     = sync.Mutex{}
	WaitGroup = sync.WaitGroup{}
)
