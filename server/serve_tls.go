package server

import "errors"

type TLSMode int

const (
	TLSModeNone TLSMode = iota + 1
	TLSModeAuto
	TLSModeStatic
)

func (s *server) HTTPTLSMode() TLSMode {
	if s.config.HTTPTLSCertPath != "" && s.config.HTTPTLSCertKey != "" {
		return TLSModeStatic
	}

	if len(s.config.HTTPTLSAutoDomains) > 0 {
		return TLSModeAuto
	}

	return TLSModeNone
}

func (s *server) HTTPTLSEnabled() bool {
	return s.HTTPTLSMode() != TLSModeNone
}

func (s *server) serveHTTPWithStaticTLS(err chan<- error) {
	go func() {
		err <- errors.New("not implemented yet")
		//err <- s.http.StartTLS(net.JoinHostPort(s.Addr(), strconv.Itoa(s.config.HTTPTLSPort)), s.config.HTTPTLSCertPath, s.config.HTTPTLSCertKey)
	}()
}

func (s *server) serveHTTPWithAutoTLS(err chan<- error) {
	go func() {
		err <- errors.New("not implemented yet")
	}()
}
