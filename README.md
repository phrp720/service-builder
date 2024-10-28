# service-builder

service-builder is a simple , customizable   Go library that  makes the creation of Linux and Windows services easier.

## Installation

```bash
go get github.com/phrp720/service-builder
```
## Features
- Builds and run systemd services
- Builds Windows services (coming soon)

## Usage

### systemd

#### Example : Creating a systemd service
Running the code below will create a systemd service file named `dummy.service` in the `/etc/systemd/system/` directory with the content defined in the `service` object. After creating the file, it enables and starts the service automatically.

```go
builder := systemd.NewServiceBuilder()
service := builder.
    // Unit
    Description("Dummy Service").
    Before("network.target").
    After("network.target").
    BindsTo("dummy.target").
    Conflicts("shutdown.target").
    Documentation("https://example.com").
    OnFailure("reboot.target").
    PartOf("multi-user.target").
    Requires("network.target").
    Wants("network-online.target").
    // Service
    ExecStart("/usr/bin/dummy").
    ExecStartPre("/usr/bin/dummy-pre").
    ExecStartPost("/usr/bin/dummy-post").
    ExecStop("/usr/bin/dummy-stop").
    // Install
    RequiredBy("multi-user.target").
    WantedBy("multi-user.target").
    Build()

err := systemd.GenerateAndStart(service, "dummy.service",true)
if err != nil {
    fmt.Print(err)
    return
}
```


### `Generate` functions explained
| Function           | Description                                                                                                        | Parameters                                      |
|--------------------|--------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| `Generate`         | Writes the .service file content to the specified path.                                                            | `s ServiceConfig`, `path string`                |
| `GenerateDefault`  | Writes the .service file content to `/etc/systemd/system/`.                                                        | `s ServiceConfig`, `file string`                |
| `GenerateAndStart` | Writes the .service file content to `/etc/systemd/system/` and starts the service. Optionally enables the service. | `s ServiceConfig`, `file string`, `enable bool` |
