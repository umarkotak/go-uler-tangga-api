package model

const (
	STATUS_HEALTHY = "healthy"

	// waiting -> playing -> rolling_number -> moving -> end_turn
	STATE_WAITING        = "waiting"
	STATE_PLAYING        = "playing"
	STATE_ROLLING_NUMBER = "rolling_number"
	STATE_MOVING         = "moving"
	STATE_END_TURN       = "end_turn"
)

type (
	Player struct {
		Identity        Identity       `json:"identity"`
		Avatar          Avatar         `json:"avatar"`
		AvatarPosition  AvatarPosition `json:"avatar_position"` //
		IndexPosition   int64          `json:"index_position"`  // minField - maxField
		MapPosition     []Position     `json:"map_position"`    // index position mapped by map direction
		IsOnline        bool           `json:"is_online"`
		Winning         bool           `json:"winning"`
		WinningPosition int64          `json:"winning_position"`
		Status          string         `json:"status"`
		CurrentState    string         `json:"current_state"`
		NextState       string         `json:"next_state"`
		MoveAvailable   int64          `json:"move_available"`
		Items           []Item         `json:"items"`
	}

	Avatar struct {
		BaseColor string `json:"base_color"`
		Icon      string `json:"icon"`
	}

	AvatarPosition struct {
		MarginTop  string `json:"margin_top"`
		MarginLeft string `json:"margin_left"`
	}

	Position struct {
		Index  int64 `json:"index"`
		IsHere bool  `json:"is_here"`
	}

	Item struct {
		RandomID   string `json:"random_id"`
		EffectType string `json:"effect_type"`
	}
)

func (p *Player) Init(mapConfig MapConfig, playerConfigs []PlayerConfig) {
	playerConfig := playerConfigs[p.Identity.RoomPlayerIndex-1]
	p.Avatar = Avatar{
		BaseColor: playerConfig.HexColor,
		Icon:      playerConfig.FaIcon,
	}
	p.AvatarPosition = AvatarPosition{
		MarginTop:  playerConfig.MarginTop,
		MarginLeft: playerConfig.MarginLeft,
	}
	p.CalculateCurrentPosition(mapConfig, 0)
}

func (p *Player) CalculateCurrentPosition(mapConfig MapConfig, movingCount int64) int64 {
	if p.IndexPosition == int64(len(mapConfig.Direction)-1) {
		return 0
	}
	p.MapPosition = []Position{}
	for i := 0; i < int(mapConfig.Size); i++ {
		position := Position{
			Index:  int64(i),
			IsHere: false,
		}
		p.MapPosition = append(p.MapPosition, position)
	}
	if (p.IndexPosition + movingCount) >= int64(len(mapConfig.Direction)-1) {
		remainingStep := int64(len(mapConfig.Direction)) - p.IndexPosition - 1
		movingCount -= remainingStep
		movingCount = remainingStep - movingCount
	}
	p.IndexPosition += movingCount

	p.MapPosition[mapConfig.Direction[p.IndexPosition]] = Position{
		Index:  mapConfig.Direction[p.IndexPosition],
		IsHere: true,
	}

	return movingCount
}
