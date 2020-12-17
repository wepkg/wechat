package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWXML(t *testing.T) {
	m := WXML{
		"appid":     "wx2421b1c4370ec43b",
		"partnerid": "10000100",
		"prepayid":  "WX1217752501201407033233368018",
		"package":   "Sign=WXPay",
		"noncestr":  "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"timestamp": "1514363815",
	}

	x, err := FormatMap2XML(m)

	assert.Nil(t, err)

	r, err := ParseXML2Map([]byte(x))

	assert.Nil(t, err)
	assert.Equal(t, m, r)
}

func TestUint32Bytes(t *testing.T) {
	i := uint32(250)
	b := EncodeUint32ToBytes(i)

	assert.Equal(t, i, DecodeBytesToUint32(b))
}
