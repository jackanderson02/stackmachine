package main

import (
	"errors"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	stack := NewStackFromNumbers([]int{1, 2, 3})

	poppedValue, err := stack.Pop()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if poppedValue != 3 {
		t.Errorf("expected popped value to be 3, but got %d", poppedValue)
	}

	if len(stack.StackNumbers) != 2 {
		t.Errorf("expected stack size to be 2 after pop, but got %d", len(stack.StackNumbers))
	}

	emptyStack := NewStack()
	_, err = emptyStack.Pop()
	if err == nil {
		t.Error("expected an error when popping from an empty stack, but got none")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()

	stack.Push(5)

	if len(stack.StackNumbers) != 1 {
		t.Errorf("expected stack size to be 1 after pushing, but got %d", len(stack.StackNumbers))
	}

	if stack.StackNumbers[0] != 5 {
		t.Errorf("expected top element to be 5 after pushing, but got %d", stack.StackNumbers[0])
	}

	stack.Push(10)
	stack.Push(15)

	if len(stack.StackNumbers) != 3 {
		t.Errorf("expected stack size to be 3 after pushing multiple elements, but got %d", len(stack.StackNumbers))
	}

	if stack.StackNumbers[2] != 15 {
		t.Errorf("expected top element to be 15 after pushing, but got %d", stack.StackNumbers[2])
	}

	stackWithNumbers := NewStackFromNumbers([]int{1, 2, 3})
	stackWithNumbers.Push(4)

	if len(stackWithNumbers.StackNumbers) != 4 {
		t.Errorf("expected stack size to be 4 after pushing, but got %d", len(stackWithNumbers.StackNumbers))
	}

	if stackWithNumbers.StackNumbers[3] != 4 {
		t.Errorf("expected top element to be 4 after pushing, but got %d", stackWithNumbers.StackNumbers[3])
	}

	stackWithBigNumbers := NewStackFromNumbers([]int{20001, 30000})
	err := stackWithBigNumbers.Plus()
	if (err == nil){
		t.Errorf("expected an error when summing 20001 and 30000 but did not receive one.")
	}
}

func TestStack_Duplicate(t *testing.T) {
	stack := NewStackFromNumbers([]int{1, 2, 3})

	stack.Duplicate()

	if len(stack.StackNumbers) != 4 {
		t.Errorf("expected stack size to be 4 after DUP, but got %d", len(stack.StackNumbers))
	}

	if stack.StackNumbers[3] != 3 {
		t.Errorf("expected top element to be 3 after DUP, but got %d", stack.StackNumbers[3])
	}

	emptyStack := NewStack()

	if len(emptyStack.StackNumbers) != 0 {
		t.Errorf("expected stack size to remain 0 after DUP on empty stack, but got %d", len(emptyStack.StackNumbers))
	}
}

func TestStack_Plus(t *testing.T) {
	stack := NewStackFromNumbers([]int{10, 20})

	err := stack.Plus()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if len(stack.StackNumbers) != 1 {
		t.Errorf("expected stack size to be 1 after addition, but got %d", len(stack.StackNumbers))
	}

	if stack.StackNumbers[0] != 30 {
		t.Errorf("expected top element to be 30 after addition, but got %d", stack.StackNumbers[0])
	}

	stackWithOverflow := NewStackFromNumbers([]int{25000, 26000})

	err = stackWithOverflow.Plus()
	if err == nil {
		t.Fatal("expected an overflow error, but got none")
	}

	stackWithOneElement := NewStackFromNumbers([]int{10})

	err = stackWithOneElement.Plus()
	if err == nil {
		t.Fatal("expected an error due to insufficient elements, but got none")
	}

}

func TestStack_Minus(t *testing.T) {
	stack := NewStackFromNumbers([]int{20, 10})

	err := stack.Minus()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if len(stack.StackNumbers) != 1 {
		t.Errorf("expected stack size to be 1 after subtraction, but got %d", len(stack.StackNumbers))
	}

	if stack.StackNumbers[0] != 10 {
		t.Errorf("expected top element to be 10 after subtraction, but got %d", stack.StackNumbers[0])
	}

	stackWithNegativeResult := NewStackFromNumbers([]int{10, 20})

	err = stackWithNegativeResult.Minus()
	if err == nil {
		t.Fatal("expected a result below zero error, but got none")
	}

	if err.Error() != "Result is below zero." {
		t.Errorf("expected result below zero error message, but got: %v", err)
	}

	stackWithOneElement := NewStackFromNumbers([]int{10})

	err = stackWithOneElement.Minus()
	if err == nil {
		t.Fatal("expected an error due to insufficient elements, but got none")
	}

	if err.Error() != "Could not pop second number off stack." {
		t.Errorf("expected insufficient elements error message, but got: %v", err)
	}
}

/*
All these tests must pass for completion
*/

func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		expectedErr error
	}{
		{name: "empty error", commands: "", expected: 0, expectedErr: errors.New("")},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, expectedErr: errors.New("")},
		{name: "too few add", commands: "99 +", expected: 0, expectedErr: errors.New("")},
		{name: "too few minus", commands: "99 -", expected: 0, expectedErr: errors.New("")},
		{name: "too few multiply", commands: "99 *", expected: 0, expectedErr: errors.New("")},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "sum single value", commands: "99 SUM", expected: 99, expectedErr: nil},
		{name: "sum empty", commands: "SUM", expected: 0, expectedErr: errors.New("")},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, expectedErr: nil},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("")},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, expectedErr: nil},
		{name: "single integer", commands: "9876", expected: 9876, expectedErr: nil},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, expectedErr: errors.New("")},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil},
		{name: "minus", commands: "2 5 -", expected: 3, expectedErr: nil},
		{name: "underflow minus", commands: "5 2 -", expected: 0, expectedErr: errors.New("")},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, expectedErr: nil},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, expectedErr: nil},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, expectedErr: errors.New("")},
		{name: "overflow single value", commands: "50001", expected: 0, expectedErr: errors.New("")},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, expectedErr: nil},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("")},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, expectedErr: nil},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, expectedErr: nil},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: nil},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("")},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, expectedErr: nil},
	}

	for _, test := range tests {

		got, err := StackMachine(test.commands)

		if test.expectedErr != nil {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("%s (%s) got error %v, want %v", test.name, test.commands, err, test.expectedErr)
			}
		} else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
}
