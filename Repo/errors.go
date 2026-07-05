package repo

import "errors"

var ErrSearchTaskById = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task with id exists")
