package goobserver

import (
	"errors"
)

type goObserverService struct {
	store map[string][]*funcType
}

func newGoObserverService() *goObserverService {
	return &goObserverService{
		store: map[string][]*funcType{},
	}
}

func (l *goObserverService) addSubscriber(topic string, callback *funcType) error {
	if topic == "" {
		return errors.New(TopicInvalid)
	}

	l.store[topic] = append(l.store[topic], callback)

	return nil
}

func (l *goObserverService) removeSubscriber(topic string, funcName string) error {
	if _, ok := l.store[topic]; !ok {
		return errors.New(TopicNotFound)
	}

	if funcName == "" {
		l.store[topic] = []*funcType{}

		return nil
	}

	for i, t := range l.store[topic] {
		if t.FuncName == funcName {
			l.store[topic] = append(l.store[topic][:i], l.store[topic][i+1:]...)

			return nil
		}
	}

	return nil
}

func (l *goObserverService) getSubscriberNames(topic string) ([]string, error) {
	if _, ok := l.store[topic]; !ok {
		return nil, errors.New(TopicNotFound)
	}

	var funcNames []string

	for _, t := range l.store[topic] {
		funcNames = append(funcNames, t.FuncName)
	}

	return funcNames, nil
}

func (l *goObserverService) getTopics() []string {
	var topicNames []string

	for k, _ := range l.store {
		topicNames = append(topicNames, k)
	}

	return topicNames
}

func (l *goObserverService) deleteTopic(topic string) error {
	if _, ok := l.store[topic]; !ok {
		return errors.New(TopicNotFound)
	}

	delete(l.store, topic)

	return nil
}

func (l *goObserverService) deleteAllTopics() {
	l.store = map[string][]*funcType{}
}

func (l *goObserverService) publish(topic string, payload []byte) error {
	if _, ok := l.store[topic]; !ok {
		return errors.New(TopicNotFound)
	}

	for _, t := range l.store[topic] {
		t.FuncVal(payload)
	}

	return nil
}
