package transport

type Stream interface {
	SerializeString(*string) error
	SerializeInt(*int) error
}
