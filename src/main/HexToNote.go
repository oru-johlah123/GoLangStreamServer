var intervalls = map[string]byte {
	"0":0x00,
	"1":0x40,
	"2":0x80,
	"3":0xC0,
}

var okt = map[string]byte {
	"0":0x00,
	"1":0x10,
	"2":0x20,
	"3":0x30,
}

var note = map[string]byte {
	"c":0x00,
	"c#":0x01,
	"d":0x02,
	"d#":0x03,
	"e":0x04,
	"f":0x05,
	"f#":0x06,
	"g":0x07,
	"g#":0x08,
	"a":0x09,
	"a#":0x0A,
	"b":0x0B,
}

type HexToNote struct {
	
}