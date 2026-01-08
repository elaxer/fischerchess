package fischerchess

import (
	"slices"
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/encoding/fen"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	for _, variant := range Variants {
		t.Run(variant, func(t *testing.T) {
			t.Parallel()

			board, err := fen.Decode(variant)
			require.NoError(t, err)

			t.Run("check castlings are available", func(t *testing.T) {
				assertCastlings(t, board, variant)
			})
			t.Run("check mirrored pieces exist", func(t *testing.T) {
				assertMirroredPieces(t, board)
			})
			t.Run("check bishops are on opposite colors", func(t *testing.T) {
				assertBishops(t, board)
			})
			t.Run("check the king is between the rooks", func(t *testing.T) {
				assertKingBetweenRooks(t, board)
			})
		})
	}
}

func assertCastlings(t *testing.T, board chess.Board, variant string) {
	boardFen := fen.Encode(board)
	require.Equal(t, variant+" w KQkq - 0 1", boardFen)
}

func assertMirroredPieces(t *testing.T, board chess.Board) {
	from := chess.NewPosition(chess.FileNull, chess.Rank1)
	for pos, piece := range board.Squares().IterByDirection(from, chess.DirectionRight) {
		require.NotNilf(t, piece, "expected piece at %s, got nil", pos)

		mirroredPos := chess.NewPosition(pos.File, chess.Rank8)
		mirroredPiece, err := board.Squares().FindByPosition(mirroredPos)
		require.NoError(t, err)
		require.NotNilf(t, piece, "expected mirrored piece at %s, got nil", mirroredPos)
		require.Falsef(
			t,
			mirroredPiece.Notation() != piece.Notation() || mirroredPiece.Color() == piece.Color(),
			"expected mirrored piece at %s to be %s of opposite color, got %v",
			mirroredPos,
			piece.Notation(),
			mirroredPiece,
		)
	}
}

func assertBishops(t *testing.T, board chess.Board) {
	bishops := slices.Collect(
		board.Squares().GetPieces(standardchess.NotationBishop, chess.ColorWhite),
	)
	require.Equalf(t, 2, len(bishops), "expected 2 white bishops, got %d", len(bishops))

	blackBishopPos := board.Squares().GetByPiece(bishops[0])
	whiteBishopPos := board.Squares().GetByPiece(bishops[1])
	require.Truef(
		t,
		blackBishopPos.File%2 != whiteBishopPos.File%2,
		"white bishops are not on opposite colors: %s and %s",
		blackBishopPos,
		whiteBishopPos,
	)
}

func assertKingBetweenRooks(t *testing.T, board chess.Board) {
	_, kingPos := board.Squares().FindPiece(standardchess.NotationKing, chess.ColorWhite)
	require.False(t, kingPos.IsEmpty(), "king is not found")

	hasLeftRook := hasRook(board.Squares(), kingPos, chess.DirectionLeft, chess.ColorWhite)
	require.Truef(t, hasLeftRook, "no rook found to the left of the king at %s", kingPos)

	hasRightRook := hasRook(board.Squares(), kingPos, chess.DirectionRight, chess.ColorWhite)
	require.Truef(t, hasRightRook, "no rook found to the right of the king at %s", kingPos)
}

func hasRook(squares *chess.Squares, from, direction chess.Position, color chess.Color) bool {
	for _, piece := range squares.IterByDirection(from, direction) {
		if piece != nil && piece.Notation() == standardchess.NotationRook &&
			piece.Color() == color {
			return true
		}
	}

	return false
}
