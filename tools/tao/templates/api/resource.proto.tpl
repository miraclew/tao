{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/api.Proto*/ -}}
syntax = "proto3";

import "google/protobuf/descriptor.proto";
import "google/protobuf/empty.proto";

extend google.protobuf.FileOptions {
    string resource = 50002;
}

extend google.protobuf.FieldOptions {
    bool is_time = 50020;
}

package {{.Resource}}.v1;

option (resource) = "{{.Resource}}";

message Time {}
message Any {}
message Empty {}

// Resource
message {{.Resource}} {
    int64 Id = 1;
    int32 Type = 2;
    int32 ObjectType = 3;
    string ObjectId = 4;
    int64 UserId = 5;
    string Content = 6;
    int32 State = 7;
    Time CreatedAt = 8;
}

// Service
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
    string result = 1;
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

// Event
service Event {
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
