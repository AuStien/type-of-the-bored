package terminal

import "golang.org/x/term"

func GetCols() (int, error) {
	w, _, err := term.GetSize(0)
	if err != nil {
		return -1, err
	}

	return w, nil
}
