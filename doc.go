/*
Package winlog parses Windows Logs using winevt.dll

Example:

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
*/
package winlog
