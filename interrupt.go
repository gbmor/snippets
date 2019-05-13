
// Watch for SIGINT aka ^C
// Close the log file then exit
func watchForInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for sigint := range c {
			log.Printf("\n\nCaught %v. Cleaning up ...\n", sigint)

			if !confObj.stdoutLogging {
				// signal to close the log file
				closelog <- true
				time.Sleep(20 * time.Millisecond)
			}

			close(closelog)
			os.Exit(0)
		}
	}()
}
