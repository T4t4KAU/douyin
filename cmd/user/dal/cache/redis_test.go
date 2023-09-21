package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := User{}
	f.AddUser(ctx, 1)
}

func TestCheckUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := User{}
	ok := f.CheckUser(ctx, 1)
	fmt.Println(ok)
}

func TestDelUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := User{}
	f.DelUser(ctx, 1)

	ok := f.CheckUser(ctx, 1)
	if ok {
		fmt.Println("failed to delete")
	}
	fmt.Println("success")
}

func TestExistUser(t *testing.T) {
	Init()

	ctx := context.Background()
	f := User{}
	ok := f.CheckUser(ctx, 1)
	fmt.Println(ok)
}
