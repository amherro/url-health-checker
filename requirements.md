## 🟡 Level 2 — Dependency Injection, Mocking, Concurrency, & Select

*You know: DI, mocking, goroutines, channels, select*

**Project: URL Health Checker**
A CLI tool that takes a list of URLs and checks if they're up (HTTP 200), reporting their status and response time.
- Use **goroutines + channels** to check URLs concurrently (this is where Go clicks)
- Mock the HTTP client in tests so you're not hitting real servers
- Output a simple table: `✓ google.com (45ms)` / `✗ broken.site (timeout)`

This is the moment Go's concurrency model will feel genuinely exciting. Checking 20 URLs simultaneously vs. one by one is a visceral demo of why people love Go.


## My Notes
- Take in a slice of strings (URLs) in terminal
- In a goroutine & channel, check if each URL returns HTTP 200
- Output table in terminal