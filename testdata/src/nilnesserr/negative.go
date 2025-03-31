package a

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

func Call5() error {
	err := Do()
	if err != nil {
		return err
	}
	a, err := Do3()
	if err != nil {
		return err
	}
	_ = a
	if err := Do2(); err != nil {
		return err
	}
	return err
}

func Call6(ctx context.Context, in string) (int, error) {
	if err := Do(); err != nil {
		return 0, err
	}
	if !strings.Contains(in, "${{") || !strings.Contains(in, "}}") {
		return 1, nil
	}
	res, err := Do3()
	if err != nil {
		return 23, err
	}
	_ = res + 1
	if err := Do2(); err != nil {
		return 4, err
	}
	return 5, err
}

func Call7() error {
	var a any = int(1)
	switch a.(type) {
	case int:
		err := Do()
		if err != nil {
			return err
		}
		if err2 := Do2(); err2 != nil {
			return err2
		}
		return err
	case string:
	default:
		return nil
	}

	return nil
}

func Call8() error {
	err := Do()
	if err != nil {
		return err
	}

	err2 := Do()
	if err2 == nil {
		return err
	}
	return nil
}

func Call9() (err error) {
	if err = Do(); err != nil {
		return
	} else if err = Do2(); err != nil {
		return
	}

	_, _ = Do3()
	return
}

func Call10() (int, error) {
	res, err := Do3()
	if err == nil {
		num, err := Do3()
		num += res
		if err != nil {
			return 0, err
		}
		if num > 0 {
			err := Do()
			if err != nil {
				return num, err
			}
		}
		return num, err
	}
	return 0, err
}

func Call11() error {
	_, err := Do3()
	if err != nil {
		return err
	} else if err = Do2(); err != nil {
		return err
	}
	return err
}

func Call14() error {
	err := Do()
	if err == nil {
		return err
	}

	err2 := Do2()
	if errors.Is(err, err2) {
		return nil
	}
	return nil
}

func Call16() (int, error) {
	err := Do()
	if err != nil {
		return 0, err
	}
	res, err := Do3()
	if err != nil {
		return 0, err
	}
	return res, err
}

func Call17(ctx context.Context) error {
	err := Do2()
	if err == nil {
		return nil
	}

	num2, err := Do3()
	if err != nil {
		_ = num2
		return err
	}
	return err
}

func Call13() error {
	err := Do2()
	if err != nil {
		_, err := Do3()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func Do4() (error, error) {
	return nil, nil
}

func Call19() (remoteErr error, err error) {
	remoteErr, err = Do4()
	if remoteErr != nil {
		remoteErr = fmt.Errorf("run Call19 got %w", remoteErr)
		err = nil
		return
	}
	if err != nil {
		///remoteErr = nil
		err = fmt.Errorf("run Call19 got %w", err)
		return
	}
	return
}
