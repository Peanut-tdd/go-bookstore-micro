package source

import (
	"context"
	"fmt"
)

func (m *defaultBookModel) List(ctx context.Context, ids string) (resp []BookModel, err error) {

	sql := fmt.Sprintf("select * from %s where `id` in (?)", m.table)
	err = m.conn.QueryRowsCtx(ctx, &resp, sql, ids)

	if err != nil {
		return nil, err
	}
	return
}
