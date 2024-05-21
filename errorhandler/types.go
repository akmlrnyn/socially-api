package errorhandler

type ErrrorNotFound struct{
	Message string `json:"message"`
}

type BadRequestError struct{
	Message string `json:"message"`
}

type InternalServerError struct{
	Message string `json:"message"`
}

type Unauthorized struct{
	Message string `json:"message"`
}

func (e *ErrrorNotFound) Error() string{
	return e.Message
}

func (e *BadRequestError) Error() string{
	return e.Message
}

func (e *InternalServerError) Error() string{
	return e.Message
}

func (e *Unauthorized) Error() string{
	return e.Message
}