// Package fischerchess implements Fischer Random Chess (Chess960) specific
// functionality on top of the standard chess library.
package fischerchess

import (
	"math/rand/v2"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/encoding/fen"
)

// NewBoardFilled creates a new Fischer Random Chess board with pieces arranged
// according to one of the 960 valid starting positions, selected randomly using
// the provided rand.Rand source.
func NewBoardFilled() chess.Board {
	//nolint:gosec
	fenStr := Variants[rand.IntN(len(Variants))]

	board, err := fen.Decode(fenStr)
	if err != nil {
		panic(err)
	}

	return board
}
