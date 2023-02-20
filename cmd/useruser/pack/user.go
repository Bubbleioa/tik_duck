package pack

import (
	"time"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	allerrors "github.com/808-not-found/tik_duck/pkg/allerrors"
)

type Follow struct {
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
}

func (m *Follow) TableName() string {
	return "follow"
}

func User(m *db.User, fromID int64) (*useruser.User, error) {
	if m == nil {
		return nil, allerrors.ErrDBUserToRPCUserVali()
	}
	var IsFollowShip bool
	var reserr error
	f := Follow{}

	err := db.DB.Where("from_user_id = ? AND to_user_id = ?", fromID, m.ID).Find(&f).Error
	if err == nil {
		IsFollowShip = false
	} else {
		IsFollowShip = true
	}
	return &useruser.User{
		Id:            m.ID,
		Name:          m.Name,
		FollowCount:   &m.FollowCount,
		FollowerCount: &m.FollowerCount,
		IsFollow:      IsFollowShip,
	}, reserr
}

func Users(ms []*db.User, fromID int64) ([]*useruser.User, error) {
	users := make([]*useruser.User, 0)
	var reserr error
	for _, m := range ms {
		userinfo, err := User(m, fromID)
		if err != nil {
			reserr = err
		}
		users = append(users, userinfo)
	}
	return users, reserr
}
