package healthservice

import "context"

type Health interface {
	Check(context.Context) error
}
