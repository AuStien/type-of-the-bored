package ansi

type CODE struct {
	Name  string
	Value string
}

var (
	RESET = CODE{
		Name:  "RESET",
		Value: "\033[0m",
	}

	CLEAR_LINE = CODE{
		Name:  "CLEAR",
		Value: "\033[2K\r",
	}

	MOVE_CURSOR_UP = CODE{
		Name:  "MOVE_CURSOR_UP",
		Value: "\033[1A",
	}

	UNDERLINE = CODE{
		Name:  "UNDERLINE",
		Value: "\033[4m",
	}

	RESET_UNDERLINE = CODE{
		Name:  "RESET_UNDERLINE",
		Value: "\033[24m",
	}

	GREEN = CODE{
		Name:  "GREEN",
		Value: "\033[32m",
	}

	RED = CODE{
		Name:  "RED",
		Value: "\033[41m",
	}

	HIDE_CURSOR = CODE{
		Name:  "HIDE_CURSOR",
		Value: "\033[?25l",
	}

	SHOW_CURSOR = CODE{
		Name:  "SHOW_CURSOR",
		Value: "\033[?25h",
	}

	NONE = CODE{
		Name:  "NONE",
		Value: "",
	}
)

func (c CODE) String() string {
	return c.Value
}
