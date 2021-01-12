package mc

import "net/rpc"

type Client struct {
	selector *Selector
}

func NewClient(addrs ...string) *Client {
	if len(addrs) == 0 {
		addrs = append(addrs, "localhost:8000")
	}
	return &Client{selector: &Selector{addrs: addrs}}
}

func (c *Client) Get(key string) *Result {
	addr, err := c.selector.Select(key)
	if err != nil {
		return &Result{Error: Error(err.Error())}
	}
	return c.get(addr, key)
}

func (c *Client) get(addr, key string) *Result {
	cl, err := rpc.Dial("tcp", addr)
	if err != nil {
		return &Result{Error: Error(err.Error())}
	}
	res := new(Result)
	if err := cl.Call("mc.Get", &Arg{Key: key}, res); err != nil {
		res.Error = Error(err.Error())
		return res
	}
	return res
}

func (c *Client) Add(key string, value []byte) *Result {
	addr, err := c.selector.Select(key)
	if err != nil {
		return &Result{Error: Error(err.Error())}
	}
	return c.add(addr, key, value)
}

func (c *Client) add(addr, key string, value []byte) *Result {
	cl, err := rpc.Dial("tcp", addr)
	if err != nil {
		return &Result{Error: Error(err.Error())}
	}
	res := new(Result)
	if err := cl.Call("mc.Add", &Arg{Key: key, Value: value}, res); err != nil {
		res.Error = Error(err.Error())
		return res
	}
	return res
}

func (c *Client) Remove(key string) *Result {
	addr, err := c.selector.Select(key)
	if err != nil {
		return &Result{Error: Error(err.Error())}
	}
	return c.remove(addr, key)
}

func (c *Client) remove(addr, key string) *Result {
	cl, err := rpc.Dial("tcp", addr)
	if err != nil {
		return &Result{Error: Error(err.Error())}
	}
	res := new(Result)
	if err := cl.Call("mc.Remove", &Arg{Key: key}, res); err != nil {
		res.Error = Error(err.Error())
		return res
	}
	return res
}

var DefaultClient = NewClient("127.0.0.1:8000")

func Get(key string) *Result {
	return DefaultClient.Get(key)
}

func Add(key string, value []byte) *Result {
	return DefaultClient.Add(key, value)
}

func Remove(key string) *Result {
	return DefaultClient.Remove(key)
}
