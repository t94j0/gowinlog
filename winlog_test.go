// +build windows

package winlog

import (
	"fmt"
	. "testing"
)

func TestWinlogWatcherConfiguresRendering(t *T) {
	watcher, err := NewWinLogWatcher()
	assertEqual(err, nil, t)

	watcher.SetRenderMessage(false)
	watcher.SetRenderLevel(false)
	watcher.SetRenderTask(false)
	watcher.SetRenderProvider(false)
	watcher.SetRenderOpcode(false)
	watcher.SetRenderChannel(false)
	watcher.SetRenderId(false)

	assertEqual(watcher.renderMessage, false, t)
	assertEqual(watcher.renderLevel, false, t)
	assertEqual(watcher.renderTask, false, t)
	assertEqual(watcher.renderProvider, false, t)
	assertEqual(watcher.renderOpcode, false, t)
	assertEqual(watcher.renderChannel, false, t)
	assertEqual(watcher.renderId, false, t)

	watcher.SetRenderMessage(true)
	watcher.SetRenderLevel(true)
	watcher.SetRenderTask(true)
	watcher.SetRenderProvider(true)
	watcher.SetRenderOpcode(true)
	watcher.SetRenderChannel(true)
	watcher.SetRenderId(true)

	assertEqual(watcher.renderMessage, true, t)
	assertEqual(watcher.renderLevel, true, t)
	assertEqual(watcher.renderTask, true, t)
	assertEqual(watcher.renderProvider, true, t)
	assertEqual(watcher.renderOpcode, true, t)
	assertEqual(watcher.renderChannel, true, t)
	assertEqual(watcher.renderId, true, t)

	watcher.Shutdown()
}

func ExampleNewWinLogWatcher() {
	watcher, _ := winlog.NewWinLogWatcher()
	// Recieve any future messages
	watcher.SubscribeFromNow("Application")
	for {
		select {
		case evt := <-watcher.Event():
			// Print the event struct
			fmt.Printf("Event: %v\n", evt)
			// Print the updated bookmark for that channel
			fmt.Printf("Bookmark XML: %v\n", evt.Bookmark)
		case err := <-watcher.Error():
			fmt.Printf("Error: %v\n\n", err)
		}
	}
}
