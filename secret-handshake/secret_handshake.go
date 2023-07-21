package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	var handshake []string
	for i, action := range actions {
		if code&(1<<i) != 0 {
			handshake = append(handshake, action)
		}
	}
	if code&(1<<4) != 0 {
		for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}

	}
	return handshake
}
