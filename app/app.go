package main

type appResponse struct {
	Response interface{}
	Error    string
}

func responseMessage(msg string) map[string]string {
	return map[string]string{
		"msg": msg,
	}
}
