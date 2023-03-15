package learn

import (
	"context"
	"time"
)

func Test() {
	child, cancleFunc := context.WithTimeout(context.Background(), time.Second*10)
	cancleFunc()
	child.Done()

}
