package argss

type Variant struct {
	Key     string
	Index   *int
	Weight  *float64
	Payload interface{}
}

func New(payload interface{}) *Variant {
	vr := Variant{}
	vr.Payload = payload
	// for _, meta := range metadata {

	// }
	return &vr
}

// type ctxInfo struct {
// 	key    *string
// 	index  *int
// 	weight *float64
// }

// func WithKey(k string) *ctxInfo {
// 	return ctxInfo{
// 		key: &k
// 	}
// }

// func WithIndex(i int) *ctxInfo {
// 	return ctxInfo{
// 		index: &i
// 	}
// }

// func WithWeight(w float64) *ctxInfo {
// 	return ctxInfo{
// 		weight: &w
// 	}
// }
