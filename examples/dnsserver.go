// http://breenmachine.blogspot.com/2014/09/transfer-file-over-dns-in-windows-with.html
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/miekg/dns"
)

var chunks map[string]string

func chunk(data []byte) {
	str := base64.StdEncoding.EncodeToString(data)
	size := 240
	count := (len(str)/size) + 1

	for i := 0; i < count; i++ {
		iStr := strconv.Itoa(i)
		begin := i * size
		end := begin + size

		if end > len(str) {
			end = len(str)
		}

		chunks[iStr] = str[begin:end]
	}
}

type handler struct{}
func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Question[0].Qtype {
	case dns.TypeNS:
	    log.Printf("NS request for %s\n", r.Question[0].Name)

	    rr, err := dns.NewRR(fmt.Sprintf("uuid.domain.com NS ns.domain.com"))
	    if err != nil {
	        return
	    }

	    msg.Answer = append(msg.Answer, rr)

	case dns.TypeTXT:
		log.Printf("TXT request for %s\n", r.Question[0].Name)

		key := strings.Split(r.Question[0].Name, ".")[0]
		val, ok := chunks[key]; if !ok {
			val = ""
		}

	    rr, err := dns.NewRR(fmt.Sprintf("%s.domain.com TXT %s", key, val))
		if err == nil {
			msg.Answer = append(msg.Answer, rr)
	    }
	}

	w.WriteMsg(&msg)
}

func main() {
	chunks = make(map[string]string)

	data, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Print("Failed to read file.")
	} else {
		chunk(data)
	}

	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
