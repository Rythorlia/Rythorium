package transaction

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
	Timestamp int64
	Signature []byte
}
