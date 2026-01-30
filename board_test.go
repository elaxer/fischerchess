package fischerchess

import (
	"slices"
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/encoding/fen"
	"github.com/elaxer/standardchess/encoding/pgn"
	"github.com/stretchr/testify/assert"
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

func TestMakeMoveWithUndo(t *testing.T) {
	tests := []struct {
		name       string
		initFENStr string
		pgnStr     string
		wantFENStr string
	}{
		{
			"https://www.chess.com/game/live/143486290046",
			"qnbrkrnb/pppppppp/8/8/8/8/PPPPPPPP/QNBRKRNB w KQkq - 0 1",
			`1. d4 b6 2. b3 Nc6 3. g4 d5 4. c4 Be6 5. cxd5 Bxd5 6. Bxd5 Rxd5 7. Nf3 Nb4 8.
Qc3 Nxa2 9. Qxc7 Rd7 10. Qc4 g6 11. e3 f5 12. gxf5 Rxf5 13. Ne5 Rd6 14. Qxg8+
Rf8 15. Qc4 Qe4 16. Ba3 Bxe5 17. Bxd6 Bxd6 18. Nc3 Bb4 19. Qb5+ Kf7 20. Qc4+ Qe6
21. Qxe6+ Kxe6 22. Rd3 Bxc3+ 23. Kd1 Rc8 24. d5+ Kd6 25. f4 Bf6 26. f5 Rc1+ 27.
Ke2 Nc3+ 28. Kd2 Na2 29. Rxc1 Nxc1 30. Kxc1 b5 31. fxg6 hxg6 32. e4 Be5 33. h3
g5 34. Kd2 Bf4+ 35. Ke2 a5 36. Kf3 Ke5 37. Rc3 a4 38. bxa4 bxa4 39. Ra3 g4+ 40.
hxg4 Kd6 41. Rxa4 Kd7 42. Kxf4 Kd6 43. g5 e6 44. dxe6 Kxe6 45. Ra6+ Kf7 46. Kf5
Ke7 47. g6 Kf8 48. e5 Kg7 49. e6 Kh8 50. Ra8+ Kg7 51. Ra7+ Kh8 52. e7 Kg7 53.
e8=Q+ Kh6 54. Qh8# 1-0`,
			"7Q/R7/6Pk/5K2/8/8/8/8 b - - 2 54",
		},
		{
			"https://www.chess.com/game/live/143178021453",
			"bbqrnkrn/pppppppp/8/8/8/8/PPPPPPPP/BBQRNKRN w KQkq - 0 1",
			`1. Ng3 b6 2. O-O Bxg2 3. Kxg2 Qb7+ 4. Kg1 Nd6 5. Nf3 Ne4 6. Nxe4 Qxe4 7. d3 Qg4+
8. Kh1 c6 9. Qe3 Bf4 10. Qe4 g5 11. e3 Bd6 12. Ne5 Qxe4+ 13. dxe4 Bxe5 14. Rg1
O-O-O 15. Rg2 g4 16. Rdg1 Rg6 17. Rxg4 Rxg4 18. Rxg4 Ng6 19. c4 Rg8 20. Bd3 h5
21. Rg5 h4 22. Rg1 Rg7 23. f4 Bf6 24. e5 Bxe5 25. fxe5 Kc7 26. b4 d6 27. exd6+
Kxd6 28. Bxg7 Ke6 29. Bxg6 fxg6 30. Rxg6+ Kf5 31. Rxc6 Ke4 32. c5 Kxe3 33. cxb6
axb6 34. Rxb6 Kd3 35. Re6 Kc4 36. a3 Kb5 37. Bb2 Ka4 38. Ra6+ Kb3 39. b5 Kxb2
40. a4 e5 41. a5 e4 42. Re6 Kb3 43. b6 Kb4 44. a6 Ka5 45. b7 e3 46. a7 e2 47.
b8=Q h3 48. a8=Q# 1-0`,
			"QQ6/8/4R3/k7/8/7p/4p2P/7K b - - 0 48",
		},
		{
			"https://www.chess.com/game/live/132674345717",
			"rnkqrnbb/pppppppp/8/8/8/8/PPPPPPPP/RNKQRNBB w KQkq - 0 1",
			`1. Nc3 Nc6 2. d4 d5 3. Qd2 Ne6 4. O-O-O g5 5. Kb1 Nexd4 6. Rc1 f5 7. e3 Ne6 8.
Ng3 f4 9. exf4 Nxf4 10. Nf5 Ne5 11. g3 d4 12. gxf4 gxf4 13. Ne4 Qd5 14. Ned6+
cxd6 15. Bxd5 Bxd5 16. f3 Nxf3 17. Qxf4 Nxe1 18. Rxe1 Bf6 19. Qg4 O-O-O 20.
Nxd6+ Kc7 21. Nxe8+ Rxe8 22. Bxd4 Bxd4 23. Qxd4 e6 24. Qe5+ Kc6 25. b3 Kd7 26.
c4 Be4+ 27. Rxe4 1-0`,
			"4r3/pp1k3p/4p3/4Q3/2P1R3/1P6/P6P/1K6 b - - 0 27",
		},
		{
			"https://www.chess.com/analysis/collection/double-brilliant-chess960-X2V49bAE/4VnusEgV7U/analysis",
			"rnnkbqrb/pppppppp/8/8/8/8/PPPPPPPP/RNNKBQRB w KQkq - 0 1",
			`1. d4 g6 2. e3 Nc6 3. a3 e6 4. Nb3 N8e7 5. Nc3 O-O-O 6. O-O-O a6 7. g3 f5 8. Nc5
Qf6 9. Nxa6 Nd5 10. Nxd5 exd5 11. Bxd5 Bf7 12. Bxf7 Qxf7 13. Nc5 d6 14. Nxb7
Rde8 15. Qa6 Kd7 16. Na5 Nxa5 17. Qxa5 Ra8 18. Qb5+ Ke7 19. Bb4 Rgb8 20. Qc6 Qa2
21. Qxc7+ Ke8 22. Bxd6 Qxb2+ 23. Kd2 1-0`,
			"rr2k2b/2Q4p/3B2p1/5p2/3P4/P3P1P1/1qPK1P1P/3R2R1 b - - 1 23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, err := fen.Decode(tt.initFENStr)
			require.NoError(t, err)

			pgn, err := pgn.FromString(tt.pgnStr)
			require.NoError(t, err)
			for _, move := range pgn.Moves() {
				result, err := board.MakeMove(move)
				require.NotNil(t, result)
				require.NoError(t, err)
			}

			require.Equal(t, tt.wantFENStr, fen.Encode(board).String())

			for i := range board.MoveHistory() {
				result, err := board.UndoLastMove()
				require.NoErrorf(t, err, "No %d", i+1)
				require.NotNil(t, result, "No %d", i+1)
			}

			afterFEN := fen.Encode(board)
			assert.Equal(t, tt.initFENStr, afterFEN.String())
		})
	}
}

func assertCastlings(t *testing.T, board chess.Board, variant string) {
	boardFen := fen.Encode(board)
	require.Equal(t, variant+" w KQkq - 0 1", boardFen.String())
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
