package mc

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

var (
	ErrNotFound = errors.New("cache not found")
)

type Error string

type Result struct {
	Value []byte
	Error Error
}

type Arg struct {
	Key   string
	Value []byte
}

func init() {
	gob.Register(Arg{})
	gob.Register(Result{})
}

type cache struct {
	c Cache
}

func (c *cache) Get(arg *Arg, res *Result) error {
	v, ok := c.c.Get(arg.Key)
	if !ok {
		res.Error = Error("nil")
		return nil
	}
	res.Value = v.([]byte)
	return nil
}

func (c *cache) Add(arg *Arg, res *Result) error {
	c.c.Add(arg.Key, arg.Value)
	res.Value = arg.Value
	return nil
}

func (c *cache) Remove(arg *Arg, res *Result) error {
	c.c.Remove(arg.Key)
	return nil
}

type RPCServer struct {
	opt *Option
}

func NewRPCServer(opts ...OptionFunc) *RPCServer {
	s := &RPCServer{opt: &Option{}}
	for _, o := range opts {
		o(s.opt)
	}
	return s
}

func (s *RPCServer) Run(addr string) {
	c := cache{c: s.opt.cache}
	rpc.RegisterName("mc", &c) // nolint: errcheck
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net listen: %s", err)
	}
	log.Printf("serving on %s\n", addr)
	for {
		conn, err := l.Accept()
		fmt.Printf("%s -> %s\n", conn.RemoteAddr().String(), conn.RemoteAddr().Network())
		if err != nil {
			log.Printf("accept: %s, skipped", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
