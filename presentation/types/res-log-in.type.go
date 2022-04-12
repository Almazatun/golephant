package types

import "time"

type ResLogIn[T any] struct {
	LogInEntityData   T
	Token             string
	ExperationTimeJWT time.Time
}
