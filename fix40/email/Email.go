//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a Email wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//EmailType is a required field for Email.
func (m Message) EmailType() (*field.EmailTypeField, errors.MessageRejectError) {
	f := &field.EmailTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEmailType reads a EmailType from Email.
func (m Message) GetEmailType(f *field.EmailTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTime is a non-required field for Email.
func (m Message) OrigTime() (*field.OrigTimeField, errors.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from Email.
func (m Message) GetOrigTime(f *field.OrigTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RelatdSym is a non-required field for Email.
func (m Message) RelatdSym() (*field.RelatdSymField, errors.MessageRejectError) {
	f := &field.RelatdSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRelatdSym reads a RelatdSym from Email.
func (m Message) GetRelatdSym(f *field.RelatdSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for Email.
func (m Message) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from Email.
func (m Message) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for Email.
func (m Message) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from Email.
func (m Message) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for Email.
func (m Message) LinesOfText() (*field.LinesOfTextField, errors.MessageRejectError) {
	f := &field.LinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from Email.
func (m Message) GetLinesOfText(f *field.LinesOfTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a required field for Email.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Email.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Email.
func (m Message) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Email.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Email.
func (m Message) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Email.
func (m Message) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Email messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Email.
func Builder(
	emailtype *field.EmailTypeField,
	linesoftext *field.LinesOfTextField,
	text *field.TextField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("C"))
	builder.Body().Set(emailtype)
	builder.Body().Set(linesoftext)
	builder.Body().Set(text)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "C", r
}