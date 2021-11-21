package simplefactory

import "fmt"

//API is interface
type API interface {
	Say(name string) string
}

//NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//hiAPI is one of API接口 implement
type hiAPI struct{}

//Say hi to name
func (hi *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API接口 implement
type helloAPI struct{}

//Say hello to name
func (hello *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
