/**
https://leetcode-cn.com/problems/valid-parentheses/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: tru
 */
package main

import (
	"errors"
	"fmt"
)

func main(){
	var str string = "()"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	l := len(s)
	stack := make(Stack,0,l)
	for i:=0;i<l;i++ {
		t := s[i:i+1]
		if t == "(" || t == "[" || t == "{" {
			stack.Push(t)
		} else {
			t2,_ := stack.Pop()
			if !( ( t2 == "(" && t == ")" ) || ( t2 == "[" && t == "]") || ( t2 == "{" && t == "}" ) ) {
				return false
			}
		}
	}
	if len(stack) !=0 {
		return false
	}

	return true
}



type Stack []string

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool  {
	return len(stack) == 0
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value string)  {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error)  {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack) - 1], nil
}

func (stack *Stack) Pop() (string, error)  {
	theStack := *stack
	if len(theStack) == 0 {
		return "", errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}