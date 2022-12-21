package util

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"io"
	"os"
)

// EncodeGob encode any to byte
func EncodeGob(s interface{}) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}

// EncodeMD5 md5 encryption
func EncodeMD5(value interface{}) string {
	m := md5.New()
	m.Write(EncodeGob(value))
	return hex.EncodeToString(m.Sum(nil))
}

// EncodeMD5 md5 encryption
func EncodeMD5File(file *os.File) string {
	m := md5.New()
	io.Copy(m, file)
	return hex.EncodeToString(m.Sum(nil))
}
