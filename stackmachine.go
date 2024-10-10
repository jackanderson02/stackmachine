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

func (stack *Stack) Push(value int) {
	stack.StackNumbers = append(stack.StackNumbers, value)
	stack.lastElementIndex += 1
}

func (stack *Stack) Duplicate(value int){
	// Non empty stack
	if stack.lastElementIndex >= 0{
		stack.Push(stack.getLastElement())
	}
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

func StackMachine(commands string)(int, error) {
return 0, errors.New("")
}

func main() {
// main is unused - run using 
// go test ./...
}