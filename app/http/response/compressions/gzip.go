package compressions

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"fmt"
)

func GzipCompression(content string) (string, error) {
	var compressedBytes bytes.Buffer

	gzipWriter := gzip.NewWriter(&compressedBytes)

	_, err := gzipWriter.Write([]byte(content))
	if err != nil {
		return "", err
	}

	err = gzipWriter.Close()
	if err != nil {
		return "", err
	}

	return compressedBytes.String(), err
}

func PrintHexDump(data []byte) string {
	hexStr := hex.EncodeToString(data)
	var output string
	for i := 0; i < len(hexStr); i += 2 {
		output += fmt.Sprintf("%s", hexStr[i:i+2])
		if (i+2)%16 == 0 {
			output += fmt.Sprintln() // New line after 8 two-byte groups (16 characters)
		} else {
			output += fmt.Sprint(" ") // Space between two-byte groups
		}
	}
	output += "\n"

	return output
}
