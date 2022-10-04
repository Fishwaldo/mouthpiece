package server

import (
	"github.com/Fishwaldo/mouthpiece/pkg"
)

// Server is the main server object
var server *mouthpiece.MouthPiece

func Set(s *mouthpiece.MouthPiece) {
	server = s
}

func Get() *mouthpiece.MouthPiece {
	return server
}