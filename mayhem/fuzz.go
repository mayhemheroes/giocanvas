package fuzz

import "github.com/ajstarks/giocanvas"

func mayhemit(bytes []byte) int {
    content := string(bytes)
    
    giocanvas.ColorLookup(content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}