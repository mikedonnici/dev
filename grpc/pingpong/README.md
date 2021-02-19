# pingpong


A simple exercise in gRPC practice whereby a Go container will play a Python container in a game of ping pong.


The rules:

- Each running container starts it's own version of pingpong

- All gRPC requests are delayed by a random about of time, say less than 3 seconds
- Responses do not have to be delayed
- This will ensure some randomness when both sides are allowed to make a request
- Once running, a request for a game is issued and first to receive the request should respond
- Requestor servers first

- Server:
   - 1 in 10 chance of a missed serve / missed hit
   - 1 in 10 chance of an ace serve / winning hit
   - 8 in 10 chance of a normal serve / normal hit

- Receiver:
   - 1 in 10 chance of a miss hit
   - 1 in 10 chance of winning return
   - 8 in 10 chance of normal return


- Server changes after 5 points
- Game ends when one score is >= 21 and greater score exceeds by 2 or more points


- At any time a container is either server or receiver
   - receiving   



Messages:

- Request game
- Accept game request
- Serve / Hit
- Update score
- Declare victory
- Shake hands



