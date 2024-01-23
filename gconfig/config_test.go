package gconfig

import "testing"

func TestConfig(t *testing.T) {
	cfg := newConfig("prog", "tag1", "tag2")
	cfg.fillTest()
}
