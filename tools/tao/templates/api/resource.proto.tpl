syntax = "proto3";

import "tao.proto";

package {{.Resource}}.v1;

option (app) = "TODO";
option (resource) = "{{.Resource}}";

// Resource
message {{.Resource}} {
    option (model) = true;
    int64 Id = 1;
    string Title = 2;
    string Content = 3;
    int64 UserId = 4;

    Time CreatedAt = 10;

    Key PK_Id = 100;
    Key K_UserId = 101;
}

// Service (APIs in the service)
service Service {
    rpc Create(CreateRequest) returns (CreateResponse);
    rpc Delete(DeleteRequest) returns (DeleteResponse);
    rpc Update(UpdateRequest) returns (UpdateResponse);
    rpc Get(GetRequest) returns (GetResponse);
    rpc Query(QueryRequest) returns (QueryResponse);
}

message CreateRequest {
    {{.Resource}} {{.Resource|lower}} = 1;
}

message CreateResponse {
    int64 Id = 1;
}

message DeleteRequest {
    int64 Id = 1;
}

message DeleteResponse {
    string Result = 1;
}

message UpdateRequest {
    int64 Id = 1;
    map<string, Any> Values = 2;
}

message UpdateResponse {
    string Result = 1;
}

message GetRequest {
    int64 Id = 1;
    map<string, Any> Filter = 2;
}

message GetResponse {
    {{.Resource}} Result = 1;
}

message QueryRequest {
    repeated int64 Ids = 1;
    map<string, Any> Filter = 2;
    int64 Offset = 3;
    int32 Limit = 4;
    string Sort = 5;
}

message QueryResponse {
    repeated {{.Resource}} Results = 1;
}

// Event (subscribe events from this service)
service Event {
    option (internal) = true;
    rpc HandleCreated(CreatedEvent) returns (Empty);
    rpc HandleDeleted(DeletedEvent) returns (Empty);
    rpc HandleUpdated(UpdatedEvent) returns (Empty);
}

message CreatedEvent {
    {{.Resource}} Data = 1;
}

message DeletedEvent {
    {{.Resource}} Data = 1;
}

message UpdatedEvent {
    int64 Id = 1;
    map<string, Any> Values = 3;
}
