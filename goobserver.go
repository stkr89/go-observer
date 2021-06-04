package goobserver

type GoObserver struct {
	*goObserverService
}

func NewGoObserver() *GoObserver {
	return &GoObserver{
		newGoObserverService(),
	}
}

//Publish publish payload to a topic
func (l *GoObserver) Publish(topic string, payload []byte) error {
	return l.goObserverService.publish(topic, payload)
}

//Subscribe subscribe to a new or existing topic
func (l *GoObserver) Subscribe(topic string, callback func([]byte)) error {
	return l.goObserverService.addSubscriber(topic, &funcType{
		FuncName: getFunctionName(callback),
		FuncVal:  callback,
	})
}

//Unsubscribe unsubscribe from an existing topic
func (l *GoObserver) Unsubscribe(topic string, callback func([]byte)) error {
	return l.goObserverService.removeSubscriber(topic, getFunctionName(callback))
}

//UnsubscribeAll remove all subscribers of topic
func (l *GoObserver) UnsubscribeAll(topic string) error {
	return l.goObserverService.removeSubscriber(topic, "")
}

//GetSubscribers get subscribers of topic
func (l *GoObserver) GetSubscribers(topic string) ([]string, error) {
	return l.goObserverService.getSubscriberNames(topic)
}

//GetTopics get all topics
func (l *GoObserver) GetTopics() []string {
	return l.goObserverService.getTopics()
}

//DeleteTopic delete topic
func (l *GoObserver) DeleteTopic(topic string) error {
	return l.goObserverService.deleteTopic(topic)
}

//DeleteAllTopics delete all topics along with their subscribers
func (l *GoObserver) DeleteAllTopics() {
	l.store = map[string][]*funcType{}
}
