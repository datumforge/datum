package customtypes

type JSONObject map[string]interface{}

// type Pair struct {
//	K, V []byte
//}
//
//// Value implements the driver Valuer interface
// func (p Pair) Value() (driver.Value, error) {
//	var b bytes.Buffer
//	if err := gob.NewEncoder(&b).Encode(p); err != nil {
//		return nil, err
//	}
//
//	return b.Bytes(), nil
//}
//
//// Scan implements the Scanner interface.
//func (p *Pair) Scan(value interface{}) (err error) {
//	switch v := value.(type) {
//	case nil:
//	case []byte:
//		err = gob.NewDecoder(bytes.NewBuffer(v)).Decode(p)
//	default:
//		err = fmt.Errorf("unexpected type %T", v)
//	}
//
//	return
//}
