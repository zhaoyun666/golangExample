package string

import (
    "testing"
    "strconv"
    "fmt"
)

func TestVBI(t *testing.T) {

    vbi := "0200000018000000000000000014000010000000"
    for i := 0; i < 20; i++ {
        vbi_byte, _ := strconv.ParseInt(vbi[i*2:i*2+2], 16, 64)
        //fmt.Println(vbi_byte)
        if vbi_byte == 0 {
            continue
        }
        // To binary
        vbi_binary := fmt.Sprintf("%08b", vbi_byte)
        fmt.Println(vbi_binary)
        // To find the VBI serial number
        for index, b := range vbi_binary {
            if b == 49 {
                // serial number
                serial := fmt.Sprintf("%d", i*8+index+1)
                fmt.Println(serial)
                // Append brandids
                //brandis = append(brandis, v.Vbis[serial].Id)

            }
        }
    }
}
