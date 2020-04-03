# Message Pattern

Communication message patterns of systems:

 1. Request/Response: synchronized api call, usually used in http api call.
 2. Message: bi-direction realtime data, usually used in websocket/socket connection.
 3. Command: async data send to a specified service, backend internal only
 4. Event: an event happens in a service, send to a topic, other services subscribe to the topic if interested.

| Message Pattern  | Protocol            | Client / Server    | Internal | External | Sync |
|------------------|---------------------|--------------------|----------|----------|------|
| Request/Response | Http                | App,H5 To Service  | Y        | Y        | Y    |
| Message          | Websocket OR others | App,H5 To Service  | N        | Y        | N    |
| Command          | Message Queue       | Service To Service | Y        | N        | N    |
| Event            | Pubsub              | Service To Service | Y        | N        | N    |
