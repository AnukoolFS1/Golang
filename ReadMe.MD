# 🐹 Go Learning Notes

Personal notes from working through the Go programming language — modules, packages, and tooling.

---

## 📦 Module Management

### Replace a Module Path (Local Development)

When working with a local module that isn't published yet, use `replace` to point Go to the local path:

```bash
go mod edit -replace example.com/greetings=../greetings
```

This tells Go: *"whenever you see `example.com/greetings`, use the local `../greetings` directory instead."*

---

### Sync Dependencies with `go mod tidy`

After editing `go.mod`, run tidy to resolve and download any missing modules:

```bash
go mod tidy
```

**Example output:**

```
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

> The long version string (`v0.0.0-00010101000000-...`) is a pseudo-version Go assigns to local/replaced modules that don't have a proper version tag yet.

---

## 🛠️ Installing & Running Go Programs

### Find the Install Target Path

To see where `go install` will place your compiled binary:

```bash
go list -f '{{.Target}}'
```

### Add Go Binaries to PATH

So you can run installed Go programs from anywhere in your terminal:

```bash
export PATH=$PATH:/path/to/your/install/directory
```

> 💡 To make this permanent, add the line to your `~/.bashrc`, `~/.zshrc`, or equivalent shell config file.

---

## 🔗 References

- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)