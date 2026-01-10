# fischerchess - Fischer Random Chess engine

A small Go library that provides Fischer Random (Chess960) starting positions and helpers to create valid chess boards from those positions. It builds on the [github.com/elaxer/standardchess](https://github.com/elaxer/standardchess) and
[github.com/elaxer/chess](https://github.com/elaxer/chess) packages.

## Features

- Provides the full list of 960 valid starting piece placements in `Variants`.
- Convenience function `NewBoard()` that returns a random Chess board arranged according to one of the 960 valid starting positions.
- Test coverage validating position symmetry, bishop coloring, and castling availability.

## Quick start

Install:

```bash
go get github.com/elaxer/fischerchess
```

Basic usage:

```go
package main

import (
	"fmt"
	"github.com/elaxer/fischerchess"
	"github.com/elaxer/standardchess/encoding/fen"
)

func main() {
	board := fischerchess.NewBoard()
	fmt.Println(fen.Encode(board)) // example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
}
```

Decode a known variant (from `Variants`):

```go
import "github.com/elaxer/standardchess/encoding/fen"

board, err := fen.Decode(fischerchess.Variants[0])
if err != nil {
	panic(err)
}
```

## API

- `func NewBoard() chess.Board` — returns a new board with a randomly selected Fischer Random starting layout.
- `var Variants []string` — array containing 960 valid starting piece placement FENs (piece placement only, without side-to-move, castling, etc.).

## Testing

Run tests:

```bash
go test ./...
```

The test suite checks that each starting variant decodes correctly, pieces are mirrored between White and Black ranks, bishops are on opposite-colored squares, and castling rights are available.

## Contributing

Bug reports and contributions are welcome. Please open issues or pull requests against this repository. Keep changes small and add tests for new behavior.

## License

The GNU General Public License
