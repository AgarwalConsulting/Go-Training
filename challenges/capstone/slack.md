# Implement a Slack-like chat server in Go

Your chat server manages a set of users. Unlike Slack, there is just one channel that everybody chats on. There are two APIs: (a) send a message to the channel and (b) get messages from the channel.

Use HTTP and JSON. You can test your server with curl; no need to write a fancy GUI client.

Assume that the set of users is fixed and available in a static JSON file at server startup; you don't need to implement signup.

Assume that there are fewer than a hundred users, and you only need to retain up to a 1000 messages, and that messages disappear after one hour. Don't use a database; keep the messages in the server's memory.

To recap:

1. Implement a chat server in Go. Implement basic authentication using a static JSON file that's read at start up.

2. Implement a `SendMessage(message string)` API. Use HTTP basic authentication. Refer: [example code](https://github.com/AgarwalConsulting/Go-Training/blob/master/examples/11-A-01-net/auth/basic/main.go)

3. Implement a `GetMessages() []Message` API. Use basic auth again. Return a JSON object containing the messages in a channel.
Figure out what fields would go into a message object (the text, user, date, etc.)

4. Run the server on your laptop, and test it using curl.

---

## Create a UI for the above application using React

---

## Stretch goals

Some ideas for what do next:

Product Features

1. Add a sign up API. Get rid of the static user file at start up, allow users to sign up.

2. Message retention. Add a database or persistent storage of some sort. Retain 10k messages for up to 30 days.

3. Group chat aka Channels -- implement slack-like channels where a subset of users can chat with each other.

4. Implement a search API for messages.

5. Implement the client side as a simple CLI.

### Scale

How fast is your server? Benchmark your server with lots of users and lots of messages. Use the `hey` package to generate load, or write your own code to do it. Do you find any limits? What happens after these limits are reached?

Profile your server to find out more about the bottlenecks in your code. Look for CPU, memory and (if applicable) mutex profiles.

Redo your benchmarks once you use a database to retain messages. Are the limits different now?

How might you scale up to more users, more messages, and more retention?

### Security

Spend some time trying to attack your own server. Can you attempt to brute force the authentication? Can you impersonate the server to an unsuspecting user? Can you make a denial of service attack against it? Anything else?

Implement mitigations for these problems.

### Operations

Package your server in a Docker container and deploy it to a cloud server. Purchase a domain name and point it at your server. Share the URL with your colleagues.

Make a code change and deploy an upgrade to your server. How might you do this with minimal downtime to your users?
