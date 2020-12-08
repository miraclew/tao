Tao is a protocol first micro-service approach and toolset.

This project is in very early development stage, things may changes frequently.

## The Tao approach
The Tao approach includes

 * A command tool ```tao```
 * A simple ioc container for service locate
 * Several support packages ```github.com/miraclew/tao/pkg```

This project is inspired by a [Youtube presentation](https://www.youtube.com/watch?v=j6ow-UemzBc)
 
### The ```tao``` command
```tao``` command is the core tool to generate code/documents/sql etc.

Execute ```go get -u github.com/miraclew/tao/cmd/tao``` to install ```tao``` command

sub-commands

 * ```tao proto``` generate a skeleton .proto [Google ProtoBuffer](https://developers.google.com/protocol-buffers) syntax file
 * ```tao api``` generate api/event/client go code
 * ```tao svc``` generate service implementation go code, include service, handler etc.
 * ```tao sql``` generate SQL schema of model message in .proto file
 * ```tao doc``` generate OpenAPI v3 files
 * ```tao dart``` generate Flutter/Dart client code
 * ```tao kotlin``` generate Android/Kotlin client sdk.
 * ```tao swift``` generate iOS/Swift client sdk.

### How ```tao``` generators works

![image.png](doc/images/tao-generators.png)
 
### locator: the ioc container
We introduce a locator package as the IOC container.
 * Services register to the locator 
 * Service obtain it's dependency services from locator 
 
## Steps to build a micro-service
 * Write a service proto file
 * Generate api/event/client code
 * Implement service

### Examples

 
## TODO
 * Support JAVA project
 
