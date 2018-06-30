package mechanics

//define type of tiles with diffrent rules that apply
//eg : can walk
type TileRules int

const (
	TILE_WALL  TileRules =  0
	TILE_GRASS TileRules = 1
)
