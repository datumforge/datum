package interceptors

import (
	"context"
	"time"

	"entgo.io/ent"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"go.uber.org/zap"
)

func QueryLogger(l *zap.SugaredLogger) ent.InterceptFunc {
	return func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query generated.Query) (ent.Value, error) {
			q, err := intercept.NewQuery(query)
			if err != nil {
				return nil, err
			}

			start := time.Now()
			defer func() {
				l.Infow("query duration", "duration", time.Since(start), "schema", q.Type())
			}()

			return next.Query(ctx, query)
		})
	}
}
