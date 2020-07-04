package tiles

import deviant "github.com/recluse-games/deviant-protobuf/genproto/go"

// CreateTiles Creates a deviant.Tiles from a set of tileIds
func CreateTiles(tileIds [][]string) *deviant.Tiles {
	tilesRowsLiteral := []*deviant.TilesRow{}

	for _, tileIDRow := range tileIds {
		tileRowLiteral := []*deviant.Tile{}

		for _, tileID := range tileIDRow {
			tileLiteral := &deviant.Tile{
				Id: tileID,
			}

			tileRowLiteral = append(tileRowLiteral, tileLiteral)
		}

		tileRow := &deviant.TilesRow{
			Tiles: tileRowLiteral,
		}

		tilesRowsLiteral = append(tilesRowsLiteral, tileRow)
	}

	tiles := &deviant.Tiles{
		Tiles: tilesRowsLiteral,
	}

	return tiles
}
