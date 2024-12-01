package request

// Заявка
type Request struct {
	ReceiptTime float64
	StartTime   float64
	EndTime     float64
}

func NewRequest(ReceiptTime float64) *Request {
	return &Request{
		ReceiptTime: ReceiptTime,
	}
}

