package stset

import "fmt"

//PairMap - Keep hashed pair.
// PM can be strict or non-strict
// strict PM must have all values unique (ex: [A:B]; [C:A] - is invalid as 'A' is not unique)
// non-strict PM must have only key and values be unique among them selves
//
//(ex non-stict: [A:B]; [C:A] - is valid 'A' is both unique Key and unique Value)
type PairMap struct {
	keyValMap map[string]string
	valKeyMap map[string]string
	strict    bool
}

func NewStrict() *PairMap {
	ss := PairMap{}
	ss.keyValMap = make(map[string]string)
	ss.valKeyMap = make(map[string]string)
	ss.strict = true
	return &ss
}

func NewNonStrict() *PairMap {
	ss := PairMap{}
	ss.keyValMap = make(map[string]string)
	ss.valKeyMap = make(map[string]string)
	ss.strict = false
	return &ss
}

func (ss *PairMap) Add(key, val string) error {
	ss.keyValMap[key] = val
	ss.valKeyMap[val] = key
	switch ss.strict {
	case true:
		if err := checkStrict(ss); err != nil {
			delete(ss.keyValMap, key)
			delete(ss.valKeyMap, val)
			return fmt.Errorf("integrity check failed: %v", err)
		}
	case false:
		if err := checkNonStrict(ss); err != nil {
			delete(ss.keyValMap, key)
			delete(ss.valKeyMap, val)
			return fmt.Errorf("integrity check failed: %v", err)
		}
	}
	return nil
}

func (ss *PairMap) GetValue(key string) (string, bool) {
	if val, ok := ss.keyValMap[key]; ok {
		return val, ok
	}
	if val, ok := ss.valKeyMap[key]; ok {
		return val, ok
	}
	return "", false
}

func (ss *PairMap) Delete(key string) {
	delete(ss.keyValMap, key)
	delete(ss.valKeyMap, key)
}

func checkStrict(ss *PairMap) error {
	numMap := make(map[string]int)
	for _, v := range ss.keyValMap {
		numMap[v]++
		if numMap[v] > 1 {
			return fmt.Errorf("value %v met %v times", v, numMap[v])
		}
	}
	for _, v := range ss.valKeyMap {
		numMap[v]++
		if numMap[v] > 1 {
			return fmt.Errorf("value %v met %v times", v, numMap[v])
		}
	}
	return nil
}

func checkNonStrict(ss *PairMap) error {
	numMap := make(map[string]int)
	for _, v := range ss.keyValMap {
		numMap[v]++
		if numMap[v] > 1 {
			return fmt.Errorf("key %v met %v times", v, numMap[v])
		}
	}
	numMap = make(map[string]int)
	for _, v := range ss.valKeyMap {
		numMap[v]++
		if numMap[v] > 1 {
			return fmt.Errorf("value %v met %v times", v, numMap[v])
		}
	}
	return nil
}
