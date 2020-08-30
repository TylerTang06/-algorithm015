package week01

type MyCircularDeque struct {
	q     []int
	rear  int
	front int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		q:     make([]int, k+1),
		rear:  1,
		front: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.front] = value
	this.front = (this.front - 1 + len(this.q)) % len(this.q)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.rear] = value
	this.rear = (this.rear + 1) % len(this.q)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % len(this.q)
	this.q[this.front] = 0
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + len(this.q)) % len(this.q)
	this.q[this.rear] = 0
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[(this.front+1)%len(this.q)]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[(this.rear-1+len(this.q))%len(this.q)]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if (this.front+1)%len(this.q) == this.rear {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.rear == this.front {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
