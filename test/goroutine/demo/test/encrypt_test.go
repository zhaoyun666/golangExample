package test

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	TestString := "Simple"

	md5Inst := md5.New()
	md5Inst.Write([]byte(TestString))
	result := md5Inst.Sum([]byte("xx"))
	fmt.Printf("%x\n\n", result)
}

func TestSha1(t *testing.T) {
	str := "Ni Shi Shui"
	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(str))
	result := sha1Instance.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)

}

func TestEncryptFile(t *testing.T) {
	filepath := "./test.txt"
	fmt.Println("md5: ", readFile(filepath, 1))
	fmt.Println("sha1: ", readFile(filepath, 2))
}

func TestCrc32(t *testing.T) {
	c := "LW0011E12000020T"
	//crc32Instance := crc32.MakeTable(crc32.)
	fmt.Println(crc32.ChecksumIEEE([]byte(c)))
	ct := crc32.MakeTable(crc32.IEEE)
	fmt.Println(crc32.Checksum([]byte(c), ct))
	fmt.Println(fmt.Sprintf("%08x", crc32.Checksum([]byte(c), ct)))
	fmt.Println(1429601838>>24&0xff, 1429601838>>16&0xff, 1429601838>>8&0xff, 1429601838&0xff)
	bg := make([]byte, 2)
	binary.BigEndian.PutUint16(bg, 5120)
	fmt.Println(bg)
	fmt.Println(5120>>8&0xff, 5120&0xff)
	fmt.Println(20 << 8)
	b := make([]byte, 0)
	b = append(b, 1429601838>>24&0xff, 1429601838>>16&0xff, 1429601838>>8&0xff, 1429601838&0xff)
	fmt.Println(binary.BigEndian.Uint32(b))

}

func readFile(filepath string, tp int) string {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		log.Fatalf("err:%s", err.Error())
	}
	switch tp {
	case 1:
		md5Instance := md5.New()
		io.Copy(md5Instance, f)
		return fmt.Sprintf("%x", md5Instance.Sum([]byte("")))
	case 2:
		sha1Instance := sha1.New()
		io.Copy(sha1Instance, f)
		return fmt.Sprintf("%x", sha1Instance.Sum([]byte("")))
	default:
		return "error format"
	}
}
