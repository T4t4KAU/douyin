package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := REQ_FRQ{}
	f.AddFrq(ctx, "1")
}

func TestCheckUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := REQ_FRQ{}
	ok := f.CheckFrq(ctx, "1")
	fmt.Println(ok)
}

func TestIncUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := REQ_FRQ{}
	num, err := f.IncrFrq(ctx, "1")
	if err != nil {
		fmt.Println("failed to incr")
	}
	fmt.Println(num)
}

func TestDelUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := REQ_FRQ{}
	f.DelFrq(ctx, "1")

	ok := f.CheckFrq(ctx, "1")
	if ok {
		fmt.Println("failed to delete")
	}
	fmt.Println("success")
}
