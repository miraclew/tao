# Tao generator design

## Core components
Tao generator include 3 components:

 * Parser parse the proto3 files into proto models
 * Mapper maps proto models into schema models
 * Generator apply schema models to template and output target files
 * Engine drive the whole generation process

## Engine

How does engine works:
 1. Project detecting, get all proto files, go.mod
 2. Parse proto files
 3. Load template files
 4. Load mapper to map proto to template data
 5. Execute template to render output file/files

## Commands 
 * Project commands
    1. locator 
    2. doc 
    3. integration tests 
    4. app sdk
    
 * Resource commands
    1. proto, 
    2. api(include api/event/client), 
    3. svc 
    4. unit tests

## Targets
  
  Resource:
  (proto def, mapper) -> data,  (templates, data) -- output files

### Schema mapping

| Command | Schema | Proto | Schema/Mapper | Data(model) | Template | Output files
|---------|---------|------|------|----------|---------|---------|
| doc     | OpenAPIv3 | N | 1 | N | 1 | N |
| sql     | SQL | N | 1 | N | 1 | N |
| api     | N/A | 1 | 3 | 1 | 3 | 3 |
| svc     | N/A | 1 | 3 | 1 | 3 | 3 |
| locator | N/A | N | 1 | 1 | 1 | 1 |

