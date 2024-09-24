package components

import "errors"

var ErrInvalidEventType = errors.New("invalid event type")
var ErrNotLiveEvent = errors.New("mutation event is not a live mutation event")
