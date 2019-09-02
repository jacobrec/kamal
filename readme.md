# Kamal
[A tool for navigating ports](https://en.wikipedia.org/wiki/Kamal_(navigation))

Allows for very simple load balancing, and redirection of subdomains and ports.

```
Example Usage:

Add a new rule: kamal --add test.example.com localhost:8080
	This will forward all requests trying to reach test.example.com
	that hit this server to localhost on port 8080

Remove a rule: kamal --rm test.example.com localhost:8080
	This will remove the rule, such that requests to test.example.com
	will no longer go to localhost:8080

Run a program: kamal --run test.example.com 8080 ./my-binary my-binary-arg1 my-arg2
	This will run the program my-binary, passing through the specifed
	args and for the duriation of the running program will proxy
	all traffic for test.example.com to localhost:8080

Basic load balancing:
	 kamal --add test.example.com localhost:8000
	 kamal --add test.example.com localhost:8001
	Now traffic will be split between these 2 servers

```
