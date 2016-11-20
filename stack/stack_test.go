package stack

import "testing"

func TestPush(t *testing.T) {
	stack := New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	var expected_size int = 4
	var expected_topEle int = 40
	var actual_size int = stack.curr_size
	var actual_topEle, err = stack.Peek()

	if err != nil {
		t.Errorf("Invalid Output")
	}
	if actual_size != expected_size {
		t.Errorf("Got Stack size %v, But expected Stack size is %v\n", actual_size, expected_size)
	}
	if actual_topEle != expected_topEle {
		t.Errorf("Got top Element %v, But expected top element is %v\n", actual_topEle, expected_topEle)
	}
}

func TestPeek(t *testing.T) {
	stack := New()

	var expected_peek int = 0
	var actual_peek, err = stack.Peek()

	if expected_peek != actual_peek && err == nil {
		t.Errorf("Got Top Element as %v, But expected Peek Element is %v\n", actual_peek, expected_peek)
	}

	stack.Push(10)
	stack.Pop()
	actual_peek, err = stack.Peek()

	if expected_peek != actual_peek && err == nil {
		t.Errorf("Got Top Element as %v, But expected Peek Element is %v\n", actual_peek, expected_peek)
	}

	stack.Pop()
	stack.Pop()
	actual_peek, err = stack.Peek()

	if expected_peek != actual_peek && err == nil {
		t.Errorf("Got Top Element as %v, But expected Peek Element is %v\n", actual_peek, expected_peek)
	}

	stack.Push(10)
	stack.Push(20)

	var expected_peek_v1 int = 20
	var actual_peek_v1, err_v1 = stack.Peek()

	if expected_peek_v1 != actual_peek_v1 || err_v1 != nil {
		t.Errorf("Got Top Element as %v, But expected Peek Element is %v\n", actual_peek_v1, expected_peek_v1)
	}
}

func TestPop(t *testing.T) {
	stack := New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Pop()

	var expected_topEle int = 20
	var actual_topEle, err = stack.Peek()
	var expected_size int = 2
	var actual_size int = stack.curr_size

	if actual_size != expected_size || err != nil {
		t.Errorf("Got Stack size as %v, But the actual size is %v\n", actual_size, expected_size)
	}
	if actual_topEle != expected_topEle {
		t.Errorf("Got the actual Top Element %v, But the expected Top Element is %v\n", actual_topEle, expected_topEle)
	}
}
