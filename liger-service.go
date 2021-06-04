package goobserver

import (
	"errors"
)

type ligerService struct {
	store map[string][]*funcType
}

func newLigerService() *ligerService {
	return &ligerService{
		store: map[string][]*funcType{},
	}
}

func (l *ligerService) addSubscriber(topic string, callback *funcType) error {
	if topic == "" {
		return errors.New(TopicInvalid)
	}

	l.store[topic] = append(l.store[topic], callback)

	return nil
}

func (l *ligerService) removeSubscriber(topic string, funcName string) error {
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

func (l *ligerService) getSubscriberNames(topic string) ([]string, error) {
	if _, ok := l.store[topic]; !ok {
		return nil, errors.New(TopicNotFound)
	}

	var funcNames []string

	for _, t := range l.store[topic] {
		funcNames = append(funcNames, t.FuncName)
	}

	return funcNames, nil
}

func (l *ligerService) getTopics() []string {
	var topicNames []string

	for k, _ := range l.store {
		topicNames = append(topicNames, k)
	}

	return topicNames
}

func (l *ligerService) deleteTopic(topic string) error {
	if _, ok := l.store[topic]; !ok {
		return errors.New(TopicNotFound)
	}

	delete(l.store, topic)

	return nil
}

func (l *ligerService) deleteAllTopics() {
	l.store = map[string][]*funcType{}
}

func (l *ligerService) publish(topic string, payload []byte) error {
	if _, ok := l.store[topic]; !ok {
		return errors.New(TopicNotFound)
	}

	for _, t := range l.store[topic] {
		t.FuncVal(payload)
	}

	return nil
}
