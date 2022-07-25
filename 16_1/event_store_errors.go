package main

type AddEventError struct{}

func (*AddEventError) Error() string {
	return "add event error"
}

type RemoveEventError struct{}

func (*RemoveEventError) Error() string {
	return "remove event error"
}

type UpdateEventError struct{}

func (*UpdateEventError) Error() string {
	return "update event error"
}

type AddHandlerError struct{}

func (*AddHandlerError) Error() string {
	return "add handler error"
}

type RemoveHandlerError struct{}

func (*RemoveHandlerError) Error() string {
	return "remove handler error"
}
