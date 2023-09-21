package cache

import (
	"context"
)

type (
	REQ_FRQ struct{}
)

func (f REQ_FRQ) AddFrq(ctx context.Context, k string) {
	add(rdbLIM, ctx, k, 1)
}

func (f REQ_FRQ) CheckFrq(ctx context.Context, k string) bool {
	return check(rdbLIM, ctx, k)
}

func (f REQ_FRQ) IncrFrq(ctx context.Context, k string) (int64, error) {
	return incr(rdbLIM, ctx, k)
}

func (f REQ_FRQ) DelFrq(ctx context.Context, k string) {
	del(rdbLIM, ctx, k, 1)
}
