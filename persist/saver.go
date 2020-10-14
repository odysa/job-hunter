package persist

import (
	"job-hunter/engine"
	"log"
)

func Saver(index string) (chan engine.Item, error) {
	in := make(chan engine.Item)
	go func() {
		itemCount := 1
		for {
			item := <-in
			log.Printf("Got Item #%d:%v", itemCount, item)
			itemCount++
		}
	}()
	return in, nil
}