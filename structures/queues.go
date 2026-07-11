package structures

import (
	"errors"
	"fmt"
	"log"
)

type Queue[T comparable] struct {
	items []T
}

func (q *Queue[T]) Push(item T) {
	q.items = append([]T{item}, q.items...)
}

func (q *Queue[T]) Pop() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("empty queue")
	}

	lastItem := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return lastItem, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("empty queue")
	}
	lastIndex := len(q.items) - 1
	return q.items[lastIndex], nil
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Search(item T) (int, bool) {
	for index, value := range q.items {
		if value == item {
			return index, true
		}
	}
	return -1, false
}

func Matchmake(queue *Queue[string], user [2]string) string {
	username := user[0]
	userAction := user[1]
	if userAction == "join" {
		queue.Push(username)
	} else if userAction == "leave" {
		index, found := queue.Search(username)
		if found {
			queue.items = append(queue.items[:index], queue.items[index+1:]...)
		}
	}
	if queue.Size() >= 4 {
		user1, err := queue.Pop()
		if err != nil {
			log.Fatal(err)
		}
		user2, err := queue.Pop()
		if err != nil {
			log.Fatal(err)
		}

		return fmt.Sprintf("%s has matched with %s!", user1, user2)
	}
	return "No match found!"
}
