package pack

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

func Comment(ctx context.Context, m *db.Comment, myID int64) (*userplat.Comment, error) {
	var res *userplat.Comment
	if m == nil {
		return res, nil
	}
	// TODO

	// return res, nil
	return &userplat.Comment{
		Id: m.ID,
	}, nil
}

func Comments(ctx context.Context, ms []*db.Comment, myID int64, vdID int64) ([]*userplat.Comment, error) {
	comments := make([]*userplat.Comment, 0)
	for _, m := range ms {
		n, err := Comment(ctx, m, myID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, n)
	}
	return comments, nil
}
