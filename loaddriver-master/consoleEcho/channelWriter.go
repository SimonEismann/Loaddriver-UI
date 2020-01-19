package consoleEcho

// channelWriter implements io.Writer and is used to write to channels
type channelWriter struct {
	channel chan []byte
}

// Write writes to its underlying channel
func (c *channelWriter) Write(p []byte) (int, error) {
	c.channel <- p
	return len(p), nil
}

// NewChannelWriter creates a new instance of private struct channelWriter
func NewChannelWriter(channel chan []byte) *channelWriter {
	return &channelWriter{channel: channel}
}
