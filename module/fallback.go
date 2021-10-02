package module

type FallbackStrategy interface {
	ClientIndex() int
	Failure(index int)
	Success(index int)
}
