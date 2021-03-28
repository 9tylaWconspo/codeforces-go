// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`"(name)is(age)yearsold"`, `[["name","bob"],["age","two"]]`, 
			`"bobistwoyearsold"`,
		},
		{
			`"hi(name)"`, `[["a","b"]]`, 
			`"hi?"`,
		},
		{
			`"(a)(a)(a)aaa"`, `[["a","yes"]]`, 
			`"yesyesyesaaa"`,
		},
		{
			`"(a)(b)"`, `[["a","b"],["b","a"]]`, 
			`"ba"`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, evaluate, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-234/problems/evaluate-the-bracket-pairs-of-a-string/
