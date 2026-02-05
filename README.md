## Requirements for testing in go
- It needs to be in a file with a name like xxx_test.go

- The test function must start with the word Test

- The test function takes one argument only t *testing.T

## TDD cycle 
1. Write a test

2. Make the compiler pass

3. Run the test, see that it fails and check the error message is meaningful

4. Write enough code to make the test pass

5. Refactor
