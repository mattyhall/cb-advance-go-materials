# Mocking tasks
1. Write a function in your weather program that calls each of the methods on your interface and returns them as a struct. Let's call it CoallesceWeatherInfo

2. Use GoMock to create a mock implementation of your interface

3. Use this mock to write a couple of unit tests of your function:
- Test that the values you return from each method get put in the struct correctly
- Test that if the middle function you call returns an error then you get the error from CoallesceWeatherInfo
