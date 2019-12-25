package main

import (
	"flag"
	"fmt"
	"github.com/peng19940915/go_proxy/modules/http_proxy/g"
	"github.com/peng19940915/go_proxy/modules/http_proxy/proxy"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())

}
func init() {
	prepare()
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	handleVersion(*version)
	handleHelp(*help)
	handleConfig(*cfg)
	g.InitLog()
}

func main() {
	http := flag.String("http", ":8080", "proxy listen addr")
	auth := flag.String("auth", "", "basic credentials(username:password)")
	genAuth := flag.Bool("genAuth", false, "generate credentials for auth")
	flag.Parse()
	server := proxy.NewServer(*http, *auth, *genAuth)
	server.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()
	select {}
}

func handleVersion(displayVersion bool) {
	if displayVersion {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
}

func handleHelp(displayHelp bool) {
	if displayHelp {
		flag.Usage()
		os.Exit(0)
	}
}

func handleConfig(configFile string) {
	err := g.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
func main() {
	http := flag.String("http", ":8080", "proxy listen addr")
	auth := flag.String("auth", "", "basic credentials(username:password)")
	genAuth := flag.Bool("genAuth", false, "generate credentials for auth")
	flag.Parse()
	server := gsproxy.NewServer(*http, *auth, *genAuth)
	server.Start()
}
*/
