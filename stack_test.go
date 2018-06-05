package main

import (
	"testing"
	"bytes"
)

func TestStack_Push_Pop(t *testing.T) {
	s := NewStack()

	s.Push([]byte("1"))

	ret := s.Pop()
	if !bytes.Equal(ret, []byte("1")) {
		t.Errorf("Expected %v, but got: %v", []byte("1"), ret)
	}
}

func TestStack_Len(t *testing.T) {
	s := NewStack()

	s.Push([]byte("1"))
	if s.Len() != 1 {
		t.Errorf("Expected len of stack is equal 1, but got: %v", s.Len())
	}

	s.Push([]byte("2"))
	if s.Len() != 2 {
		t.Errorf("Expected len of stack is equal 2, but got: %v", s.Len())
	}

	s.Pop()
	if s.Len() != 1 {
		t.Errorf("Expected len of stack is equal 1, but got: %v", s.Len())
	}
}


func TestStack_Pop_LIFO(t *testing.T) {
	s := NewStack()

	s.Push([]byte("1"))
	s.Push([]byte("2"))
	s.Push([]byte("3"))

	got := s.Pop()
	if !bytes.Equal(got, []byte("3")) {
		t.Errorf("Expected %v, but got: %v", []byte("3"), got)
	}
}