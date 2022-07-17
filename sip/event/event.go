package event

import (
	"fmt"

	"github.com/KalbiProject/kalbi/interfaces"
	"github.com/KalbiProject/kalbi/log"
	"github.com/KalbiProject/kalbi/sip/message"
)

//SipEvent object that gets passed to the SipListener
type SipEvent struct {
	sipmsg *message.SipMsg
	tx     interfaces.Transaction
	lp     interfaces.ListeningPoint
}

//GetSipMessage returns message that created this event
func (se *SipEvent) GetSipMessage() *message.SipMsg {
	return se.sipmsg
}

//SetSipMessage sets message that created this event
func (se *SipEvent) SetSipMessage(msg *message.SipMsg) {
	se.sipmsg = msg
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Method - %v", string(se.sipmsg.Req.Method)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.UriType - %v", string(se.sipmsg.Req.UriType)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.StatusCode - %v", string(se.sipmsg.Req.StatusCode)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.StatusDesc - %v", string(se.sipmsg.Req.StatusDesc)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.User - %v", string(se.sipmsg.Req.User)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Host - %v", string(se.sipmsg.Req.Host)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Port - %v", string(se.sipmsg.Req.Port)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.UserType - %v", string(se.sipmsg.Req.UserType)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Src - %v", string(se.sipmsg.Req.Src)))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.UriType - %v", se.sipmsg.From.UriType))
	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Name - %v", string(se.sipmsg.From.Name)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.User - %v", string(se.sipmsg.From.User)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Host - %v", string(se.sipmsg.From.Host)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Port - %v", string(se.sipmsg.From.Port)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Tag - %v", string(se.sipmsg.From.Tag)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.UserType - %v", string(se.sipmsg.From.UserType)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Src - %v", string(se.sipmsg.From.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Host - %v", string(se.sipmsg.To.Host)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Name - %v", string(se.sipmsg.To.Name)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Port - %v", string(se.sipmsg.To.Port)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Src - %v", string(se.sipmsg.To.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Tag - %v", string(se.sipmsg.To.Tag)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.UriType - %v", string(se.sipmsg.To.UriType)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.User - %v", string(se.sipmsg.To.User)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.To.UserType - %v", string(se.sipmsg.To.UserType)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Expires - %v", string(se.sipmsg.Contact.Expires)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Host - %v", string(se.sipmsg.Contact.Host)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Name - %v", string(se.sipmsg.Contact.Name)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Port - %v", string(se.sipmsg.Contact.Port)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Qval - %v", string(se.sipmsg.Contact.Qval)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Src - %v", string(se.sipmsg.Contact.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Tran - %v", string(se.sipmsg.Contact.Tran)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.UriType - %v", string(se.sipmsg.Contact.UriType)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.User - %v", string(se.sipmsg.Contact.User)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Via - %v", se.sipmsg.Via))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Cseq.ID - %v", string(se.sipmsg.Cseq.ID)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Cseq.Method - %v", string(se.sipmsg.Cseq.Method)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Cseq.Src - %v", string(se.sipmsg.Cseq.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Algorithm - %v", string(se.sipmsg.Auth.Algorithm)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.CNonce - %v", string(se.sipmsg.Auth.CNonce)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Nc - %v", string(se.sipmsg.Auth.Nc)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Nonce - %v", string(se.sipmsg.Auth.Nonce)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.QoP - %v", string(se.sipmsg.Auth.QoP)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Realm - %v", string(se.sipmsg.Auth.Realm)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Response - %v", string(se.sipmsg.Auth.Response)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Src - %v", string(se.sipmsg.Auth.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.URI - %v", string(se.sipmsg.Auth.URI)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Username - %v", string(se.sipmsg.Auth.Username)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Ua.Src - %v", string(se.sipmsg.Ua.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Ua.Value - %v", string(se.sipmsg.Ua.Value)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Exp.Src - %v", string(se.sipmsg.Exp.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Exp.Value - %v", string(se.sipmsg.Exp.Value)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.MaxFwd.Src - %v", string(se.sipmsg.MaxFwd.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.MaxFwd.Value - %v", string(se.sipmsg.MaxFwd.Value)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.CallID.Src - %v", string(se.sipmsg.CallID.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.CallID.Value - %v", string(se.sipmsg.CallID.Value)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.ContType.Src - %v", string(se.sipmsg.ContType.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.ContType.Value - %v", string(se.sipmsg.ContType.Value)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.ContLen.Src - %v", string(se.sipmsg.ContLen.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.ContLen.Value - %v", string(se.sipmsg.ContLen.Value)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Src - %v", string(se.sipmsg.Src)))

	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Body - %v", se.sipmsg.Body))

	log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Attrib - %v", se.sipmsg.Sdp.Attrib))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.ConnData - %v", se.sipmsg.Sdp.ConnData))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.MediaDesc - %v", se.sipmsg.Sdp.MediaDesc))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.AddrType - %v", string(se.sipmsg.Sdp.Origin.AddrType)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.NetType - %v", string(se.sipmsg.Sdp.Origin.NetType)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.SessionID - %v", string(se.sipmsg.Sdp.Origin.SessionID)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.SessionVersion - %v", string(se.sipmsg.Sdp.Origin.SessionVersion)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.Src - %v", string(se.sipmsg.Sdp.Origin.Src)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.UniAddr - %v", string(se.sipmsg.Sdp.Origin.UniAddr)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.Username - %v", string(se.sipmsg.Sdp.Origin.Username)))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Time - %v", se.sipmsg.Sdp.Time))
        log.Log.Info(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Version - %v", se.sipmsg.Sdp.Version))

	log.Log.Info(fmt.Sprintf("SetSipMessage() msg - %v", msg))

}

//GetTransaction returns transaction related to the SIP message that created this event
func (se *SipEvent) GetTransaction() interfaces.Transaction {
	return se.tx
}

//SetTransaction sets transaction related to the SIP message that created this event
func (se *SipEvent) SetTransaction(tx interfaces.Transaction) {
	se.tx = tx
}

//SetListeningPoint gives ability to set interfaces.ListeningPoint
func (se *SipEvent) SetListeningPoint(lp interfaces.ListeningPoint) {
	se.lp = lp
}

//GetListeningPoint returns interfaces.ListeningPoint
func (se *SipEvent) GetListeningPoint() interfaces.ListeningPoint {
	return se.lp
}
