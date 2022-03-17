package foo

import "github.com/pkg/errors"

func Foobar() error {
	return errors.New("foobar")

}
