package format

import (
	"fmt"
	"unisat-telegram-bot/types"
)

func FormatRunesDetail(detail types.RunesData) string {
	return fmt.Sprintf(
		"Runes: %s (%s)\nHolder: %d\nSupply: %s\nBurned: %s",
		detail.SpacedRune,
		detail.Symbol,
		int(detail.Holders),
		detail.Supply,
		detail.Burned,
	)
}
