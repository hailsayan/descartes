package main

func StartDecipher(senderChan chan string, decipherer func(encrypted string) string) chan string {
	receiverChan := make(chan string, 100)

	go func() {
		for {
			select {
			case msg := <-senderChan:
				receiverChan <- decipherer(msg)
			}
		}
	}()

	return receiverChan
}
