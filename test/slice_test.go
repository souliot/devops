package test

import (
	"reflect"
	"sort"
	"testing"
)

func TestSliceEqual(t *testing.T) {
	a1 := []string{"192.168.0.2", "192.168.0.3", "192.168.0.4"}
	a2 := []string{"192.168.0.3", "192.168.0.4", "192.168.0.2"}

	addrs_old := sort.StringSlice(a1)
	addrs := sort.StringSlice(a2)
	sort.Sort(addrs_old)
	sort.Sort(addrs)
	t.Log(reflect.DeepEqual(addrs_old, addrs))
}
