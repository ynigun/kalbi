//Basic SIP Server example 
package main

import ("github.com/KalbiProject/Kalbi"
        "github.com/KalbiProject/Kalbi/interfaces"
	    "github.com/KalbiProject/Kalbi/sip/method"
	    )


type Server struct {
	stack       *kalbi.SipStack
}


func (s *Server) HandleRegister(tx interfaces.Transaction) {


}

//HandleRequests 
func (s *Server) HandleRequests(event interfaces.SipEventObject){
	tx := event.GetTransaction()
	switch string(tx.GetLastMessage().Req.Method) {
	case method.CANCEL:
		
	case method.INVITE:
	
	case method.REGISTER:
		go s.HandleRegister(tx)
	case method.BYE:
		
	case method.ACK:
		
	default:
	}

}

//HandleResponses
func (s *Server) HandleResponses(event interfaces.SipEventObject){

}

func (s *Server) Start(host string, port int) {
	s.stack = kalbi.NewSipStack("Basic Client Example")
	s.stack.CreateListenPoint("udp", host, port)
	s.stack.SetSipListener(s)
	go s.stack.Start()
}

func main() {
	s := new(Server)
	s.Start("0.0.0.0", 5060)
	select{}//blocking action
}