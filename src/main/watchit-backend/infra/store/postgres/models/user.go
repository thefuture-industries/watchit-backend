package models

import "time"

type UserCore struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserUUID  string    `json:"user_uuid"`
	Email     *string   `json:"email",omitempty`
	Verified  bool      `json:"verified"`
}

type UserProfile struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserUUID  string    `json:"user_uuid"`
	Avatar    []byte    `json:"avatar"`
	Username  string    `json:"username"`
	Bio       *string   `json:"bio",omitempty`
}

type UserLimit struct {
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	UserUUID              string    `json:"user_uuid"`
	LimitId               uint      `json:"limit_id"`
	MaxQueryLengthUsage   uint      `json:"max_query_length_usage"`
	DailySearchLimitUsage uint      `json:"daily_search_limit_usage"`
	SearchPriorityUsage   uint      `json:"search_priority_usage"`
}
