package idgenerator

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

var counter = 'A'

func GenerateID(pattern string) string {
	mu.Lock()
	defer mu.Unlock()
	uniqueID := fmt.Sprintf("%s-%s-%c", pattern, time.Now().Format("20241220"), counter)
	counter++
	return uniqueID
}
