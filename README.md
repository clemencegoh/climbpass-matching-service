# climbpass-matching-service

Experimental matching service using business logic of climbpass as an example

Goal is to see how well this can be implemented in Go.

Mainly just for learning Golang's web-related capabilities


# Functions
1. Climbing/Climber Profiles
    - CRUD
    - Restricted to personal profiles
2. Gym Profiles
    - CRUD
    - Serves only as data, should not need to support
    CRUD from client side
3. Authentication 
    - Auth sign in/sign up with JWT tokens
4. Matching Offer
    - Service to put up an "Event"/"Jio"
    - Should support CRUD for users to be able to create/join
    - should support input of gym location and users inside
5. Chatrooms
    - Chatting with users
    - Auto create chatrooms for Events


# Development
## Adding dependencies
Use go modules to add deps
- (if not already) `go mod init`
- `go build` should add new deps 



# Disclaimer
- For learning purposes only 
- Not meant to be used in a real business use case since there are other factors to consider
