package wildcardserver

import (
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	dns.Server
}

func (s *Server) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

	if r.MsgHdr.Opcode != dns.OpcodeQuery {
		return
	}

	replayMsg := &dns.Msg{}
	replayMsg.SetReply(r)
	for _, question := range r.Question {
		log.Infof("accept request for <%s>", question.Name)
		ip, err := findIPv4InDomain(question.Name)
		if err != nil {
			log.Infof("<%s> parse failed", question.Name)
			continue
		}

		dnsRR := &dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    300,
			},
			A: ip,
		}
		dnsTXT := &dns.TXT{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeTXT,
				Class:  dns.ClassINET,
				Ttl:    300,
			},
			Txt: []string{question.Name + " wildcard dns provide by github.com/codysk/wildcard-ip"},
		}

		switch question.Qtype {
		case dns.TypeTXT:
			log.Infof("TXT: <%s> -> %s", question.Name, dnsTXT.String())
			replayMsg.Answer = append(replayMsg.Answer, dnsTXT)
		case dns.TypeA, dns.TypeAAAA:
			log.Infof("Resolved: <%s> -> %s", question.Name, ip.String())
			replayMsg.Answer = append(replayMsg.Answer, dnsRR)
			replayMsg.Extra = append(replayMsg.Extra, dnsTXT)
		default:
		}
	}
	_ = w.WriteMsg(replayMsg)

}

func NewServer(protocol string, address string) *Server {
	server := &Server{
		dns.Server{
			Addr:       address,
			Net:        protocol,
			TsigSecret: nil,
		},
	}
	server.Handler = server
	return server
}
