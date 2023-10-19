package mysql

import "time"

type UserDto struct {
	ID          string     `json:"id"`
	DisplayName string     `json:"displayName"`
	Gender      string     `json:"gender"`
	Birthdate   *time.Time `json:"birthdate"`
	Addr        string     `json:"addr"`
	Avatar      string     `json:"avatar"`
	Type        int32      `json:"type"`
	IsDisabled  bool       `json:"isDisabled"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (u *User) Dto() UserDto {
	return UserDto{
		ID:          u.ID,
		DisplayName: u.DisplayName,
		Gender:      u.Gender,
		Birthdate:   u.Birthdate,
		Addr:        u.Addr,
		Avatar:      u.Avatar,
		Type:        u.Type,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func UserDtoFunc(u User) UserDto {
	return u.Dto()
}
