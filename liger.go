package goObserver

type Liger struct {
	*ligerService
}

func NewLiger() *Liger {
	return &Liger{
		newLigerService(),
	}
}

//Publish publish payload to a topic
func (l *Liger) Publish(topic string, payload []byte) error {
	return l.ligerService.publish(topic, payload)
}

//Subscribe subscribe to a new or existing topic
func (l *Liger) Subscribe(topic string, callback func([]byte)) error {
	return l.ligerService.addSubscriber(topic, &funcType{
		FuncName: getFunctionName(callback),
		FuncVal:  callback,
	})
}

//Unsubscribe unsubscribe from an existing topic
func (l *Liger) Unsubscribe(topic string, callback func([]byte)) error {
	return l.ligerService.removeSubscriber(topic, getFunctionName(callback))
}

//UnsubscribeAll remove all subscribers of topic
func (l *Liger) UnsubscribeAll(topic string) error {
	return l.ligerService.removeSubscriber(topic, "")
}

//GetSubscribers get subscribers of topic
func (l *Liger) GetSubscribers(topic string) ([]string, error) {
	return l.ligerService.getSubscriberNames(topic)
}

//GetTopics get all topics
func (l *Liger) GetTopics() []string {
	return l.ligerService.getTopics()
}

//DeleteTopic delete topic
func (l *Liger) DeleteTopic(topic string) error {
	return l.ligerService.deleteTopic(topic)
}

//DeleteAllTopics delete all topics along with their subscribers
func (l *Liger) DeleteAllTopics() {
	l.store = map[string][]*funcType{}
}
