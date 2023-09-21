package cache

import (
	"context"
	"douyin/cmd/relation/dal/cache"
	"fmt"
	"testing"
)

func TestAddFavorite(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	f.AddFavorite(ctx, 1, 2)
}

func TestCheckFavorite(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	ok := f.CheckFavorite(ctx, 2)
	fmt.Println(ok)
}

func TestDelFavorite(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	f.DelFavorite(ctx, 1, 2)

	ok := f.CheckFavorite(ctx, 2)
	if ok {
		fmt.Println("failed to delete")
	}
	fmt.Println("success")
}

func TestExistFavorite(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	ok := f.ExistFavorite(ctx, 1, 2)
	fmt.Println(ok)
}

func TestCountFavorite(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	cnt, err := f.CountFavorite(ctx, 2)
	if err != nil {
		return
	}
	fmt.Println(cnt)
}

func TestAddFavorited(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	f.AddFavorited(ctx, 1, 2)
}

func TestCheckFavorited(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	ok := f.CheckFavorited(ctx, 1)
	fmt.Println(ok)
}

func TestDelFavorited(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	f.DelFavorited(ctx, 1, 2)

	ok := f.CheckFavorite(ctx, 1)
	if ok {
		fmt.Println("failed to delete")
	}
	fmt.Println("success")
}

func TestExistFavorited(t *testing.T) {
	Init()

	ctx := context.Background()
	f := Favorites{}
	ok := f.ExistFavorited(ctx, 1, 2)
	fmt.Println(ok)
}

func TestRedis(t *testing.T) {
	Init()
	cache.Init()

	ctx := context.Background()
	keys, err := rdbFavorites.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		println(key)
	}
}
