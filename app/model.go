package models

type BiEvent struct {
	artifact  uint   `json:"artifact"`
	source    string `json:"source"`
	target    string `json:"target"`
	version   uint   `json:"version"`
	timestamp uint   `json:"timestamp"`
	data      string `json:"data"`
}
