package operations

import "errors"

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ゼロによる除算はできません")
	}
	return a / b, nil
}
