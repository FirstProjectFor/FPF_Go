package mock

//go:generate mockgen -destination ./person_mock.go -package mock -source person.go

type Person interface {
	SetName(name string) string
}
