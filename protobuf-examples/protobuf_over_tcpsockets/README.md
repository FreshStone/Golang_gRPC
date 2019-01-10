Implementation details

- communication done over tcp sockets
- data is send in protobufs encoded format
- 2 different msgs types can be sent in the current implementation [individual, Enterprise]
- length of each msg and msg type is sent along with the msg to assist their unmarshaling
- support multiple msgs from multiple concurrent clients 
