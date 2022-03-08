package repository

type (
	IFetchRepository interface {
	}

	FetchRepository struct {
	}
)

func NewFetchRepository() *FetchRepository {
	return &FetchRepository{}
}
