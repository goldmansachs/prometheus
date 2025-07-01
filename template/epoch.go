package template

import (
       "fmt"
       "time"
)

func now() string {
       ep := time.Now().UTC().Unix() * 1000 // Grafana expects timestamp in msec
       return fmt.Sprintf("%d", ep)
}

func now6h() string {
       ep := time.Now().UTC().Unix() * 1000
       ep -= 6 * int64(time.Hour/1000/1000)
       return fmt.Sprintf("%d", ep)
}
