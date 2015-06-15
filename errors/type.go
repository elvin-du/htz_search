package errors

import(
    "encoding/json"
    "log"
)

type Error struct{
    Code int `json:"code"`
    Msg string `json:"msg"`
}

func NewErorr(code int, msg string)*Error{
    return &Error{code,msg}
}

func (this *Error)Error()string{
    bin,err := json.MarshalIndent(this,"", "  ")
    if nil != err{
        log.Println(err)
        return ""
    }

    return string(bin)
}

var(
    E_ARTICLE_NOT_FOUND = NewErorr(404,"article not found")
)
