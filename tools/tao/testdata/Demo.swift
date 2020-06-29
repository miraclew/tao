import Foundation


class DemoRpcService {
  let app = "Core"
  static let shared = DemoRpcService()
  private init() {}

  func create(req: NewThing, completion: @escaping (NewThingResult?, Error?) -> ()) {
    APIClient.shared.rpc(app: app, path: "/v1/demoservice/create", req: req, completion: completion)
  }
  
}

class DemoSocketService {
  let app = "Core"
  static let shared = DemoSocketService()
  private init() {}

  func sendClientMessage(req: ClientMessage) {
    SocketClient.shared.send(data: req)
  }
  
  func recvServerMessage(handler: @escaping (msg: ServerMessage) -> ()) {
    SocketClient.shared.subscribe("ServerMessage", handler: handler)
  }

}



struct ClientMessage: Codable {
}

struct ServerMessage: Codable {
}

struct NewThing: Codable {
  var mobile: String
  var code: String
}

struct NewThingResult: Codable {
  var code: String
  var version: String
}

