# PermBAC 🛡️

> **DEPRECATED:** Please use [VDL](https://github.com/varavelio/vdl) with a policy plugin.

<br/>
<img src="./mascot.png"  width="180" />
<br/>

Easy and flexible permission-based access control for Go applications.

PermBAC simplifies permission management in your Go projects, enabling detailed definition and efficient checking of permissions based on a structure that is configurable and adaptable to your needs.

## Installation

To install the PermBAC binary, use the following command:

```bash
go install github.com/eduardolat/permbac/cmd/permbac@latest
```

You can also download precompiled binaries from the [GitHub releases](https://github.com/eduardolat/permbac/releases).

## Getting Started

PermBAC operates through a clear and straightforward workflow, facilitating integration with your project:

1. **Initialization**: Generate a base configuration file with the command:

   ```bash
   permbac --init
   ```

2. **Permission Definition**: Write your permissions in the permission file. For example, in `permbac_perms.json`:

   ```json
   [
     {
       "name": "Example",
       "desc": "PermBAC permission example"
     },
     {
       "name": "...",
       "desc": "..."
     }
   ]
   ```

   - **name**: Permission name. It must be unique and follow the regex `^[A-Z][a-zA-Z0-9-_]*$`.
   - **desc**: Permission description. It is optional and used for documentation purposes.

   You can see the detailed schema definition and validation rules [here.](internal/schema/perms.json)

3. **Permission Generation**: Use `permbac --generate` to create your permissions. You can specify a custom configuration file with:

   ```bash
   permbac --generate --config ./permbac.json
   ```

You can also execute `permbac --help` to see the available commands and options.

```bash
Usage of permbac:
  -config string
        Path to the configuration file (default "./permbac.json")
  -generate
        Runs the PermBAC code generator using the configuration file
  -init
        Initialize a new PermBAC configuration file
```

### Configuration File

The configuration file allows customization of the permission generation process. An example `permbac.json` could be:

```json
{
  "perms": ["./permbac_perms.json"],
  "outdir": "./permbac",
  "package": "permbac"
}
```

1. **perms**: Array of paths to the permission files to include in the generation process. You can use glob patterns (`./perms/*.json`) to include multiple files.

2. **outdir**: Output directory for the generated files.

3. **package**: Package name for the generated files.

You can see the detailed schema definition and validation rules [here.](internal/schema/config.json)

## Contributions

Contributions are welcome! Feel free to open an issue or pull request on [our repository](https://github.com/eduardolat/permbac).

## License

PermBAC is licensed under the [MIT License](https://opensource.org/licenses/MIT), allowing its broad use in both commercial and private projects.
