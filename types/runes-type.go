package types

type RunesDetailResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data RunesData `json:"data"`
}

type RunesData struct {
	RuneID       string `json:"runeid"`
	Rune         string `json:"rune"`
	SpacedRune   string `json:"spacedRune"`
	Number       int    `json:"number"`
	Height       int    `json:"height"`
	TxIdx        int    `json:"txidx"`
	Timestamp    int64  `json:"timestamp"`
	Divisibility int    `json:"divisibility"`
	Symbol       string `json:"symbol"`
	Etching      string `json:"etching"`
	Premine      string `json:"premine"`
	Terms        Terms  `json:"terms"`
	Mints        string `json:"mints"`
	Burned       string `json:"burned"`
	Holders      int    `json:"holders"`
	Transactions int    `json:"transactions"`
	Supply       string `json:"supply"`
	Start        *int   `json:"start"`
	End          *int   `json:"end"`
	Mintable     bool   `json:"mintable"`
	Remaining    string `json:"remaining"`
}

type Terms struct {
	Amount      string `json:"amount"`
	Cap         string `json:"cap"`
	HeightStart int    `json:"heightStart"`
	HeightEnd   int    `json:"heightEnd"`
	OffsetStart *int   `json:"offsetStart"`
	OffsetEnd   *int   `json:"offsetEnd"`
}
