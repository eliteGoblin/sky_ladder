package _535_encode_and_decode_tinyurl

import (
	"fmt"
	"strconv"
)

type Codec struct {
	dict      map[string]string
	tinyTable []string
}

func Constructor() Codec {
	res := Codec{}
	res.dict = make(map[string]string)
	res.tinyTable = make([]string, 0, 4096)
	return res
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if k, ok := this.dict[longUrl]; ok {
		return this.dict[k]
	}
	tiny := fmt.Sprintf("%x", len(this.dict))
	this.tinyTable = append(this.tinyTable, tiny)
	this.dict[longUrl] = tiny

	return tiny
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	index, _ := strconv.ParseInt(shortUrl, 0, 64)
	return this.tinyTable[index]
}
