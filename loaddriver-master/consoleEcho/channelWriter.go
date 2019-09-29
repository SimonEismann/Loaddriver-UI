package consoleEcho

type channelWriter struct {
	channel chan []byte
}

func (c *channelWriter) Write(p []byte) (int, error) {
	c.channel <- p
	return len(p), nil
}

func NewChannelWriter(channel chan []byte) *channelWriter {
	return &channelWriter{channel: channel}
}
