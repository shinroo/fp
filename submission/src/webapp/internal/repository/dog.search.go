package repository

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/dog.search.sql
var dogSearchQuery string

func (r *Dog) Search(ctx context.Context, searchKeyword string, filters []DogSearchFilter) ([]*models.Dog, error) {
	query := dogSearchQuery

	args := []interface{}{searchKeyword, searchKeyword}

	if len(filters) > 0 {
		filterConditions := make([]string, len(filters))
		for i, filter := range filters {
			filterConditions[i] = filter.ToWhereCondition()
			args = append(args, filter.GetValue())
		}
		query += " AND " + strings.Join(filterConditions, " AND ")
	}

	var dogs []*models.Dog

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	query = r.db.Rebind(query)

	fmt.Println("query:", query)

	err = r.db.SelectContext(ctx, &dogs, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search for dogs: failed to select context: %w", err)
	}

	return dogs, nil
}
