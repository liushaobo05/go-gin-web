package util

// base64 encode
import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"math"

	jsoniter "github.com/json-iterator/go"
)

// 定义JSON操作
var (
	json              = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal       = json.Marshal
	JSONUnmarshal     = json.Unmarshal
	JSONMarshalIndent = json.MarshalIndent
	JSONNewDecoder    = json.NewDecoder
	JSONNewEncoder    = json.NewEncoder
)

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64 decode
func Decode(str string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(str)
	return string(s), e
}

// 字符串转[]byte
func EncodeString(s string) []byte {
	return []byte(s)
}

// []byte转字符串
func DecodeToString(b []byte) string {
	return string(b)
}

// bool 转[]byte
func EncodeBool(b bool) []byte {
	if b == true {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// 自动识别int类型长度，转换为[]byte
func EncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return EncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return EncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return EncodeInt32(int32(i))
	} else {
		return EncodeInt64(int64(i))
	}
}

// 自动识别uint类型长度，转换为[]byte
func EncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return EncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return EncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return EncodeUint32(uint32(i))
	} else {
		return EncodeUint64(uint64(i))
	}
}

func EncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func EncodeUint8(i uint8) []byte {
	return []byte{byte(i)}
}

func EncodeInt16(i int16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(i))
	return bytes
}

func EncodeUint16(i uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, i)
	return bytes
}

func EncodeInt32(i int32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(i))
	return bytes
}

func EncodeUint32(i uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, i)
	return bytes
}

func EncodeInt64(i int64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(i))
	return bytes
}

func EncodeUint64(i uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, i)
	return bytes
}

func EncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func EncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func DecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(DecodeToUint8(b))
	} else if len(b) < 3 {
		return int(DecodeToUint16(b))
	} else if len(b) < 5 {
		return int(DecodeToUint32(b))
	} else {
		return int(DecodeToUint64(b))
	}
}

// 将二进制解析为uint类型，根据[]byte的长度进行自动转换
func DecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(DecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(DecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(DecodeToUint32(b))
	} else {
		return uint(DecodeToUint64(b))
	}
}

// 将二进制解析为bool类型，识别标准是判断二进制中数值是否都为0，或者为空
func DecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Compare(b, make([]byte, len(b))) == 0 {
		return false
	}
	return true
}

func DecodeToInt8(b []byte) int8 {
	return int8(b[0])
}

func DecodeToUint8(b []byte) uint8 {
	return uint8(b[0])
}

func DecodeToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(fillUpSize(b, 2)))
}

func DecodeToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(fillUpSize(b, 2))
}

func DecodeToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(fillUpSize(b, 4)))
}

func DecodeToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(fillUpSize(b, 4))
}

func DecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(fillUpSize(b, 8)))
}

func DecodeToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(fillUpSize(b, 8))
}

func DecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(fillUpSize(b, 4)))
}

func DecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(fillUpSize(b, 8)))
}

// 当b位数不够时，进行高位补0
func fillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b
	}
	c := make([]byte, 0)
	c = append(c, b...)
	for i := 0; i < l-len(b); i++ {
		c = append(c, 0x00)
	}
	return c
}

// JSONMarshalToString JSON编码为字符串
func JSONMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}
