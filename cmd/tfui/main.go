package main

import (
	"tfui/internal"
	"tfui/internal/tf"
)

func main() {
	tf.GetStateList()
	internal.RenderList()
}
