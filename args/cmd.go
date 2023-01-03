package args

type Handler func(kv KV, cmd Cmd) (string, bool, error)

type Chain []Handler

func (c Chain) Add(h Handler) Chain {
	return append(c, h)
}

func (c Chain) Handle(kv KV, cmd Cmd) (string, bool, error) {
	var message string
	for _, h := range c {
		msg, brk, err := h(kv, cmd)
		if err != nil {
			return "", brk, err
		}
		if brk {
			return message, brk, nil
		}
		message = message + "\n" + msg
	}
	return message, false, nil
}

func Command(h ...Handler) (string, error) {
	msg, _, err := Chain(h).Handle(Parse())
	return msg, err
}
