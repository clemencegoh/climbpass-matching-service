# List of TODO:
- Create "Events"
    - Events CRUD
    - Add ability for Users to join event (Just string userID?)
- Create "Chatrooms" for events
    - Websocket implementation
    - gRPC implementation
    - QUIC implementation (HTTP3)



# Completed
- Climber CRUD
    - Link to Login user
- Gyms CRUD
- Auth with jwtTokens


# Generation
- Generation tool (Custom): https://github.com/clemencegoh/gogen
- Generate with `gogen new -t /home/clemence/workspace/projects/gogen/templates -w climbpass-matching-service -c <component>`
- **Remember to add `AddHandlers` method in router.go**