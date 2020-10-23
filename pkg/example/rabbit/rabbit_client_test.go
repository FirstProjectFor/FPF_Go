package rabbit

import "testing"

func TestPushMessage(t *testing.T) {
	pushMessage(10000000)
}

func TestReadMessage(t *testing.T) {
	readMessage()
}


func TestReadLog(t *testing.T) {
	readLog()
}
