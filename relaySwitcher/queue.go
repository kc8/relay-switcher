package main

import (
	"bytes"
	"fmt"
)

type Queue struct { 
    nodes []interface{}
    head int 
    tail int
    length int // available slots
    size int // total size of queue
}

func New() *Queue {
    result := &Queue{}
    result.Create()
    return result
}

func (q* Queue) Create() *Queue {
    initalSize := 1
    q.nodes = make([]interface{}, initalSize)
    q.head = 0
    q.tail = 0
    q.size = initalSize
    q.length = 0
    return q
}

func (q* Queue) Deque() interface{} {
    if q.IsEmpty() == true {
        return nil
    }
    result := q.nodes[q.head]
    q.head = (q.head + 1) % q.size
    q.length -= 1 
    return result
}

func (q* Queue) IsFull() bool {
    if q.length == q.size {
        return true 
    }
    return false
}

func (q* Queue) IsEmpty() bool {
    if q.length == 0 {
        return true 
    }
    return false
}

func (q* Queue) Enqueue(node interface{}) {
    q.resize()
    q.nodes[q.tail] = node;
    q.tail = (q.tail + 1) % q.size
    q.length += 1
}

// Return error or the new size of the queue
func (q* Queue) resize() (int, error) {
    if (q.IsFull() == true) {
        newSize := q.size * 2
        tempNodes := make([]interface{}, newSize)
        //TODO we could try built in copy w/ array splicing
        for j := 0; j < q.size; j++ {
            // var pos int = q.head + j % q.size
            // fmt.Printf("pos %d, j %d, head %d, size %d, newsize %d\n" ,pos, j, q.head, q.size, newSize)
            if (q.head == 0) {
                tempNodes[j] = q.nodes[q.head + j % q.size];
            } else {
                tempNodes[j] = q.nodes[q.head % q.size];
            }
        }
        q.size = newSize
        q.head = 0
        q.tail = q.length
        q.nodes = tempNodes;
        return newSize , nil
    }
    return 0, nil
}

func (q *Queue) ToString() string {
    // TODO this does not actually display the order of the queue
    var result  bytes.Buffer
    result.WriteString("CONTENTS OF QUEUE ARE NOT IN QUEUE ORDER")
    for i, obj := range q.nodes {
        result.WriteString(fmt.Sprintf("pos %d | obj %v \n", i, obj))
    }
    return result.String()
}
