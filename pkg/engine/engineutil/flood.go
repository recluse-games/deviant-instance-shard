package engineutil

import deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"

//FloodFill Flood fills a grid of tiles from one location to another.
func FloodFill(startx int32, starty int32, x int32, y int32, filledID string, blockedID string, limit int32, tiles []*[]*deviant.Tile) {
	if (*tiles[x])[y].Id != blockedID && (*tiles[x])[y].Id != filledID {
		var apCostX int32
		var apCostY int32

		if startx > x {
			apCostX = startx - x
		} else if startx < x {
			apCostX = x - startx
		} else {
			apCostX = 0
		}

		if starty > y {
			apCostY = starty - y
		} else if starty < y {
			apCostY = y - starty
		} else {
			apCostY = 0
		}

		newTile := &deviant.Tile{}
		newTile.X = int32(x)
		newTile.Y = int32(y)
		newTile.Id = filledID
		(*tiles[x])[y] = newTile

		if limit-apCostX-apCostY >= 0 {
			if x+1 < int32(len(tiles)) {
				FloodFill(startx, starty, x+1, y, filledID, blockedID, limit, tiles)
			}

			if y+1 < int32(len(*tiles[x])) {
				FloodFill(startx, starty, x, y+1, filledID, blockedID, limit, tiles)
			}

			if x-1 >= 0 {
				FloodFill(startx, starty, x-1, y, filledID, blockedID, limit, tiles)
			}

			if y-1 >= 0 {
				FloodFill(startx, starty, x, y-1, filledID, blockedID, limit, tiles)
			}
		}
	}
}
