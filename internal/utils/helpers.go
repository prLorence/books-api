package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
)

func IsAuthorized(r *http.Request) bool {
	return false
}

func Int32ToBytes(i int32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func BytesToInt32(b []byte) int32 {
	var i int32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &i)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return i
}
