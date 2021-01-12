package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mivinci/lru"
	"github.com/mivinci/mc"
)

var (
	maxCache = flag.Int("c", 32, "max cache size")
	addr     = flag.String("address", "127.0.0.1:8000", "addresses to serve on")
	server   = flag.Bool("server", false, "run as a server")
	client   = flag.Bool("client", false, "run as a client")
)

func main() {
	flag.Parse()

	addrs := parseAddr(*addr)

	if *server && *client {
		fmt.Println("mc can be run either a server or a client")
		os.Exit(1)
	}

	if (*server && !*client) || (!*server && !*client) {
		n := len(addrs)
		if n != 1 {
			fmt.Printf("only 1 address is needed when running as a server, but %d were given\n", n)
			os.Exit(1)
		}
		if *maxCache <= 0 {
			*maxCache = 32
		}
		s := mc.NewRPCServer(
			mc.WithCache(lru.New(*maxCache)),
		)
		s.Run(addrs[0])
		return
	}
	cl := mc.NewClient(addrs...)
	sc := bufio.NewScanner(os.Stdin)
	prompt := getPrompt(addrs)
	println("MC v1.0.0, press [ctrl c] to exit.")
	print(prompt)
	for sc.Scan() {
		c, k, v, ok := parse(sc.Text())
		if !ok {
			println("invalid command")
			print(prompt)
			continue
		}
		switch c {
		case "get":
			fmt.Println(mc.String(cl.Get(k)))
		case "add":
			fmt.Println(mc.String(cl.Add(k, []byte(v))))
		case "remove":
			fmt.Println(mc.String(cl.Remove(k)))
		}
		print(prompt)
	}
}

func getPrompt(addrs []string) string {
	if len(addrs) == 1 {
		addr := addrs[0]
		if addr[0] == ':' {
			addr = fmt.Sprintf("127.0.0.1%s", addr)
		}
		return fmt.Sprintf("%s> ", addr)
	}
	return "cluster> "
}

func parse(cmd string) (subcmd string, key string, value string, ok bool) {
	re := regexp.MustCompile(`(get|remove|add) \w+`)
	if ok = re.MatchString(cmd); !ok {
		return
	}
	ss := strings.Split(cmd, " ")
	switch ss[0] {
	case "get", "remove":
		return ss[0], ss[1], "", true
	case "add":
		if len(ss) != 3 {
			return
		}
		return ss[0], ss[1], ss[2], true
	}
	return
}

func parseAddr(addr string) []string {
	addrs := strings.Split(addr, ",")
	for i, v := range addrs {
		if v[0] == ':' {
			addrs[i] = fmt.Sprintf("127.0.0.1%s", v)
		}
	}
	return addrs
}
