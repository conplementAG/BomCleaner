# Introduction 

This tool will remove component entries from an already created CycloneDX-Bom-xml file of a dotnet framework application.
It will analyse the given *.deps.json file of a dotnet project to identify all framework specific libraries.
These libraries will be removed from the bom.xml because these named dependency are not used by the application 
but their dotnet framework version installed on the executing system.

For example a System.Buffers/4.5.1 dependency will be removed because the application will use the System.Buffers.dll
from the installed runtime.

This also means updating the nuget package providing the "System.Buffers/4.5.1" dll will not fix any vulnerability 
because it is not used. Only updating the runtime will fix an issue in the System.Buffers.dll.


# Getting Started
- Install go
- Information to deps.json : https://github.com/dotnet/sdk/blob/main/documentation/specs/runtime-configuration-file.md

# Build

Build:
``` 
cd cmd/bomcleaner
go build .
```   

Build Linux version (if on windows): 
``` 
$env:GOOS = "linux"   
cd cmd/bomcleaner
go build .
 ```

# Test / Example

After building:

``` 
cd example
cmd/dotnetcleaner.exe bom.xml ToDoApi.deps.json
 ```
--> in the created cleanbom.xml should no runtime library like "System.Buffers" should be listed

# How Deps.json is evaluated

The "targets" path of the deps.json file will be analysed.
As described here https://github.com/dotnet/sdk/blob/main/documentation/specs/runtime-configuration-file.md :
> Each property under targets describes a "target", which is a collection of libraries required by the application 
> when run or compiled in a certain framework and platform context. A target must specify a Framework name, and may 
> specify a Runtime Identifier. Targets without Runtime Identifiers represent the dependencies and assets which are 
> platform agnostic. ...

This means every library that is listed under "targets" but has no runtime entry will be seen as runtime specific. 
This means not the listed dependency is used but the dotnet framework version installed on the executing system.
This is reflected by the deployment folder that will not contain any runtime dlls.  

For example a System.Buffers/4.5.1 dependency will be removed because the application will use the System.Buffers.dll
from the installed runtime. This means not the System.Buffers nuget package is the dependency but the runtime version of dotnet.

