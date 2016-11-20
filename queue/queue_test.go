package queue

import "testing"

func TestEnque(t *testing.T) {
	queue := New()

	queue.Enque(10)
	var actual_result, err = queue.Peek()
	var expected_result = 10

	if err != nil || actual_result != expected_result {
		t.Errorf("Got Front Value as %v, But expected Value as %v\n", actual_result, expected_result)
	}

	queue.Enque(20)
	queue.Enque(30)

	var actual_result_v1, err_v1 = queue.Peek()
	if err_v1 != nil || actual_result_v1 != expected_result {
		t.Errorf("Got Front Value as %v, But expected Value as %v\n", actual_result_v1, expected_result)
	}
}

func TestPeek(t *testing.T) {
	queue := New()
	var _, err = queue.Peek()
	if err == nil {
		t.Errorf("Queue should be Empty, But its value is %v\n", queue.QueueSize())
	}

}

func TestDeque(t *testing.T) {
	queue := New()
	queue.Enque(10)
	queue.Deque()
	var _, err = queue.Peek()
	if err == nil {
		t.Errorf("Queue should be Empty, But its value is %v\n", queue.QueueSize())
	}

	queue.Enque(10)
	queue.Enque(20)
	var actual_result, err_v1 = queue.Peek()
	var expected_result = 10
	if err_v1 != nil || actual_result != expected_result {
		t.Errorf("Got Front Value as %v, But expected Value as %v\n", actual_result, expected_result)
	}
}
