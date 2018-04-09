package dateformat

import (
	"testing"
	"fmt"
	"time"
)

const (
	GOTIMELAYOUT     = "20060102150405"
	OUTPUTTIMELAYOUT = "2006-01-02 15:04:05"
	REDISCACHENAME   = "redis"
)

func TestTimeFormat(t *testing.T) {
	fmt.Println(time.Now().Format(GOTIMELAYOUT)[0:8])
}
