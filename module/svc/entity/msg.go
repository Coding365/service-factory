package entity

const STATUS_SUCC = 0
const STATUS_FAIL = 1

type Msg struct {
	Status int
	Memo   string
	Result interface{}
}

func (this *Msg) SetMsgStatus(status int) {
	this.Status = status
}

func (this *Msg) SetMsgMemo(memo string) {
	this.Memo = memo
}

func (this *Msg) SetMsgResult(result interface{}) {
	this.Result = result
}

func GetMsg(memo string, status int) *Msg {
	msg := new(Msg)
	msg.Memo = memo
	msg.Status = status
	return msg
}

func GetMsgSucc() *Msg {
	msg := new(Msg)
	msg.Memo = "接口调用成功"
	msg.Status = STATUS_SUCC
	return msg
}

func GetMsgFail() *Msg {
	msg := new(Msg)
	msg.Memo = "接口调用失败"
	msg.Status = STATUS_FAIL
	return msg
}

func GetMsgSuccMemo(memo string) *Msg {
	msg := new(Msg)
	msg.Memo = memo
	msg.Status = STATUS_SUCC
	return msg
}

func GetMsgFailMemo(memo string) *Msg {
	msg := new(Msg)
	msg.Memo = memo
	msg.Status = STATUS_FAIL
	return msg
}

func GetMsgAll(memo string, status int, result interface{}) *Msg {
	msg := new(Msg)
	msg.Memo = memo
	msg.Status = status
	msg.Result = result
	return msg
}
