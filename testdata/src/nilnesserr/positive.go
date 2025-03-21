package a

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
)

func Do() error {
	if rand.Float64() > 0.5 {
		return fmt.Errorf("do err")
	}
	return nil
}

func Do2() error {
	if rand.Float64() > 0.5 {
		return fmt.Errorf("do err")
	}
	return nil
}

func Do3() (int, error) {
	if rand.Float64() > 0.5 {
		return 1, fmt.Errorf("do err")
	}
	return 0, nil
}

func Empty() int {
	var a int
	a += 1
	return a
}

func Call() error {
	err1 := Do()
	if err1 != nil {
		return err1
	}
	err2 := Do2()
	if err2 != nil {
		a := 1
		a = a + 2
		fmt.Println(a)
		if a > 10 {
			fmt.Println(a)
			if a > 11 {
				return err1 // want `return a nil value error after check error`
			}
		}
	}
	return nil
}

func Call2() error {
	err := Do()
	if err != nil {
		return err
	}
	return err
}

func Call3() error {
	err := Do()
	if err == nil {
		return err
	}
	return err
}

func Call4() (error, error) {
	err := Do()
	if err != nil {
		return nil, err
	}
	err2 := Do2()
	if err2 != nil {
		return err, err2 // want `return a nil value error after check error`
	}
	return nil, nil
}

func Call12() (err error) {
	err = Do()
	if err != nil {
		return err
	}
	err2 := Do2()
	if err2 != nil {
		return // want `return a nil value error after check error`
	}
	return
}

func Call15() error {
	err := Do()
	if err != nil {
		return err
	} else if err2 := Do2(); err2 == nil {
		return err2
	} else {
		return err // want `return a nil value error after check error`
	}
}

type localError struct{}

func (e localError) Error() string {
	return "localErr"
}

func Call18() error {
	err := Do()
	if err != nil {
		return err
	}

	_ = fmt.Errorf("call Do2 got err %w", err) // want `call fmt.Errorf with a nil value error after check error`
	_ = errors.Is(err, io.EOF)                 // want `call errors.Is with a nil value error after check error`
	_ = errors.Join(io.EOF, err)               // want `call errors.Join with a nil value error after check error`
	_ = errors.As(err, new(localError))        // want `call errors.As with a nil value error after check error`
	_ = errors.Unwrap(err)                     // want `call errors.Unwrap with a nil value error after check error`

	_ = fmt.Sprintf("call Do2 got err %+v", err)

	return nil
}
