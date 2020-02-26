# TODO

## Ideas

 * proto file decouples with model in db
 * more general proto define
 * dynamic data access
 * data access without predefined model struct?

## New features
 * More generic proto file and api generation
 * No special single model message
 * Support more then one model messages
 * No service implementation for none-standard CRUD
 * Support none CRUD method
 * Support multi-word resource name
 * Handler support none POST method
 * Not depend on Echo
 * Remove handler generation ?
 
 
## v2

 * api generate interface definition, model messages, go client
 * svc generate service scaffold, endpoints
 * dao generate database, cache code
 * dart generate dart client
 * doc generate open api v3 

## Issues
 * Naming collision in other languages like dart
 
## Messages
 * Normal protocol message
 * Model message (which generate db code)
 
## Tao pkg

 
## TDA Tao data access

 * Interface: redis like instead of sql oriented interface
 * Remove the limitation of sqlx's needs of db tag  
