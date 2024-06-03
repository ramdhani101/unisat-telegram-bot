package format

import "fmt"

func FormatBRC20Detail(detail map[string]interface{}) string {
	return fmt.Sprintf(
		"Ticker: %s\nSelf Mint: %v\nHolders Count: %d\nHistory Count: %d\nMax: %s\nConfirmed Minted: %s\nCreator: %s\nTransaction ID: %s",
		detail["ticker"],
		detail["selfMint"],
		int(detail["holdersCount"].(float64)),
		int(detail["historyCount"].(float64)),
		detail["max"],
		detail["confirmedMinted"],
		detail["creator"],
		detail["txid"],
	)
}
