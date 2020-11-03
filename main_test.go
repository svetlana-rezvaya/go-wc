package main

import "testing"

func Test_countWords(test *testing.T) {
	wordCount := countWords("test test")
	if wordCount != 2 {
		test.Fail()
	}
}

func Test_countWords_withDuplicatedSpaces(test *testing.T) {
	wordCount := countWords("test  test")
	if wordCount != 2 {
		test.Fail()
	}
}

func Test_countWords_withPrependedSpaces(test *testing.T) {
	wordCount := countWords("  test test")
	if wordCount != 2 {
		test.Fail()
	}
}

func Test_countWords_withTrailingSpaces(test *testing.T) {
	wordCount := countWords("test test  ")
	if wordCount != 2 {
		test.Fail()
	}
}
