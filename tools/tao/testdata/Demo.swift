import Foundation




protocol DemoSocketDelegate {
  func recvServerMessage(data: ServerMessage)
}

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

  var delegate: DemoSocketDelegate?

  func sendClientMessage(req: ClientMessage) {
    SocketClient.shared.send(data: req)
  }

}

struct ClientMessage: Codable {
  var id: Int
  var userId: Int
  var type: Int
  var subType: String?
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


