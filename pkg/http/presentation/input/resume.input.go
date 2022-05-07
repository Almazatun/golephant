package input

type ResumeInput[T any] struct {
	Data     T
	IsUpdate bool
}
