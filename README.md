## TDD cycle 
1. Write a test

2. Make the compiler pass

3. Run the test, see that it fails and check the error message is meaningful

4. Write enough code to make the test pass

5. Refactor

## Requirements for testing in go
- It needs to be in a file with a name like xxx_test.go

- The test function must start with the word Test

- The test function takes one argument only t *testing.T

## Setting up Example tests
- The function must be also in a file that ends in _test.go
  
- The function name must follow this pattern "Example{NameOfFunction}" eg ExampleAdder

- Output of function to be tested is written under comments within the Example function a format like this : 
// Output: 6

where 6 is the expected output of the function being test

- Examples without output comments are useful for demonstrating code that cannot run as unit tests, such as that which accesses the network, while guaranteeing the example at least compiles.

## Runing tests
$ go run -v     # v for verbose
$ go test -cover