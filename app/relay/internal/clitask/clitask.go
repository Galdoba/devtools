package clitask

type Task struct {
	Sender string `json:"From,omitempty"` //от какой утилиты
	Agent  string `json:"Agent"`          //кто должен выполнять
	//Schedule cronex.Schdule //когда можно делать
}
