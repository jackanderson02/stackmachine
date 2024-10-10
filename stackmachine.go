package main


import(
	"errors"
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
		return errors.New("This value is not allowed on the stack.")
	}
	stack.StackNumbers = append(stack.StackNumbers, value)
	stack.lastElementIndex += 1

	return nil
}

func (stack *Stack) Duplicate(){
	// Non empty stack
	if stack.lastElementIndex >= 0{
		stack.Push(stack.getLastElement())
	}
}

func isNumberAllowedOnStack(number int) bool{
	return number >= 50000 || number <= 0
}

func (stack *Stack) PopTwoMostRecentNumbers(err *error) (int, int){
	firstNumber, errOne:= stack.Pop()
	secondNumber, errTwo := stack.Pop()

	if (errOne != nil || errTwo != nil){
		returnedError := (errors.New("Failed to pop two numbers from the stack."))
		err = &returnedError
		return 0,0 
	}else{
		err = nil
		return firstNumber, secondNumber
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

	if (stack.lastElementIndex >= 0){
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

func StackMachine(commands string)(int, error) {
	return 0, errors.New("")
}

func main() {
// main is unused - run using 
// go test ./...
}