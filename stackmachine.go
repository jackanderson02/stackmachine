package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Stack struct{
	StackNumbers []int
	lastElementIndex int
}

func NewStack() *Stack{
	var emptyStackNumbers []int = []int{}
	return &Stack{
		emptyStackNumbers,
		-1,
	}
}


func NewStackFromNumbers(stackNumbers []int) *Stack{
	return &Stack{
		stackNumbers,
		len(stackNumbers) - 1,
	}
}

func (stack *Stack) getLastElement() int {
	return stack.StackNumbers[stack.lastElementIndex]
}

func (stack *Stack) Push(value int) error {
	if isNumberAllowedOnStack(value){
		stack.StackNumbers = append(stack.StackNumbers, value)
		stack.lastElementIndex += 1
		return nil
	}else{
		return errors.New("This value is not allowed on the stack.")
	}

}

func (stack *Stack) Duplicate(){
	// Non empty stack
	if stack.lastElementIndex >= 0{
		stack.Push(stack.getLastElement())
	}
}

func isNumberAllowedOnStack(number int) bool{
	return number <= 50000 &&  number >=0
}

func (stack *Stack) hasTwoNumbers() bool{
	return len(stack.StackNumbers) >= 2
}

func (stack *Stack) PopTwoMostRecentNumbers(err *error) (int, int){

	if stack.hasTwoNumbers(){
		firstNumber, _:= stack.Pop()
		secondNumber, _ := stack.Pop()
		err = nil
		return firstNumber, secondNumber
	}else{
		returnedError := (errors.New("Failed to pop two numbers from the stack."))
		err = &returnedError
		return 0,0
	}
	

}
func (stack *Stack) Plus() error{
	var err error;
	firstNumber, secondNumber := stack.PopTwoMostRecentNumbers(&err)
	if(err == nil){
		sum := firstNumber + secondNumber
		err = stack.Push(sum)
	}
	return err

}

func (stack *Stack) Pop() (int, error){

	if !stack.isEmpty(){
		lastElement := stack.getLastElement()
		// stack.StackNumbers = slices.Delete(stack.StackNumbers, lastElementIndex, lastElementIndex)
		stack.StackNumbers = stack.StackNumbers[:stack.lastElementIndex]
		stack.lastElementIndex -= 1
		return lastElement, nil
	}else{
		return -1, errors.New("Cannot remove from empty stack.")
	}

}

func (stack *Stack) Minus() error{
	var err error;
	firstNumber, secondNumber := stack.PopTwoMostRecentNumbers(&err)
	if(err == nil){
		err = stack.Push(firstNumber - secondNumber)
		return err
	}
	return err
}

func (stack *Stack) Multiply() error{
	var err error;
	firstNumber, secondNumber := stack.PopTwoMostRecentNumbers(&err)
	if(err == nil){
		err = stack.Push(firstNumber * secondNumber)
		return err
	}
	return err
}

func (stack *Stack) Clear() {

	stack.StackNumbers = []int{}
	stack.lastElementIndex = -1

}

func (stack *Stack) isEmpty() bool{
	return stack.lastElementIndex == -1
}
func (stack *Stack) Sum() error{
	if (stack.isEmpty()){
		return errors.New("Cannot sum over an empty stack.")
	}
	var err error
	for{
		if (err != nil){
			break
		}
		err = stack.Plus()
	}

	return nil
}

func StackMachine(commands string)(int, error) {
	individualCommands := strings.Split(commands, " ")
	stack := NewStack()
	var err error;
	for _, cmd := range individualCommands{
		fmt.Println(stack.StackNumbers)
		fmt.Println(cmd)
		switch cmd{
			case "POP":
				_, err = stack.Pop()
				if err != nil{
					return -1, err
				}
			case "DUP":
				stack.Duplicate()
			case "+":
				err = stack.Plus()
				if err != nil{
					return -1, err
				}
			case "-":
				err = stack.Minus()
				if err != nil{
					return -1, err
				}
			case "*":
				err = stack.Multiply()
				if err != nil{
					return -1, err
				}
			case "CLEAR":
				stack.Clear()
			case "SUM":
				err = stack.Sum()
				if err != nil{
					return -1, err
				}
			default:
				num, _:= strconv.Atoi(cmd)
				stack.Push(num)
		}
	}

	if stack.isEmpty(){
		return 0, errors.New("Empty stack, nothing to return.")
	}else{
		return stack.getLastElement(), nil
	}
}

func main() {
// main is unused - run using 
// go test ./...
}