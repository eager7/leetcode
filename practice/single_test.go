package practice

import "testing"
import "golang.org/x/sync/singleflight"

func TestSingle(t *testing.T) {
	s1 := singleflight.Group{}
	_, _, _ = s1.Do("test", func() (i interface{}, err error) {
		return nil, nil
	})
}
