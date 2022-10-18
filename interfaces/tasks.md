# Interfaces tasks

1. You have already seen the standard library contains Reader and Writer interfaces. Can you think of any others you might expect to find. Does Go's standard library contain them?

2. Try and write a LoggingWriter similar to the reader we saw in the slides

3. There are a couple of free weather APIs:
- https://open-meteo.com/en
- https://openweathermap.org/api
Write an interface which allows users to see if it is raining at a particular location, what the temperature is at a particular location and the wind speed (each one a different method). Implement it for one of the two APIs above using net/http and implement it for random weather. E.g. use a random number to pick the temperature.

Hint: use encoding/json to deocde the JSON from the APIs
