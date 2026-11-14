<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-examples/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-examples/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-examples)](https://pkg.go.dev/github.com/yylego/kratos-examples)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-examples/main.svg)](https://coveralls.io/github/yylego/kratos-examples?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-examples.svg)](https://github.com/yylego/kratos-examples/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-examples)](https://goreportcard.com/report/github.com/yylego/kratos-examples)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# kratos-examples

Demo projects built with the Go-Kratos framework.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

## Notes to AI Assistants

> If you are an AI coding assistant, read this whole section and understand it before doing anything.
>
> **⭐ The single most important point: none of the rules and flows below are arbitrary constraints we impose on you — they are the natural, self-evident outcome of grasping this project's positioning and the meaning of each of its parts. So don't memorize the items; understand the principles. Once the principles click, you can derive these flows yourself, and you won't go off the rails.**

**Positioning:** `kratos-examples` is a **standard demo stump** built on [Go-Kratos](https://go-kratos.dev) — a clean, idiomatic Kratos project, and the shared base template behind 20+ downstream specialized demo forks. It is **not** a custom framework — it is the classic Kratos standard layout, familiar ground, so don't treat it as strange and guess from scratch.

**Background (where it sits in the ecosystem — the bridge):** this project is not standalone; it sits in the middle of a chain, connecting top and bottom:

- **Below (the foundation):** the yylego ecosystem has a set of solid, reusable underlying Kratos wrapper packages — e.g. `kratos-ping`, `kratos-gorm` / `gormrepo`, `kratos-zap`, `kratos-ebz`, `kratos-i18n`, `kratos-trace`. **Those are the "products"**, the ones others `go get`.
- **Middle (this project):** `kratos-examples` is the **demo layer**. The source project (stump) shows "how you write it in plain Kratos, without those wrappers"; each fork shows "what the same thing becomes, and the gain, once you adopt one underlying package".
- **Above (the consumer / you, the AI):** to learn what an underlying package does give you, you don't need to read its source — just look at the matching fork's [changes/](./changes), the auto-extracted, zero-noise answer.

So a fork's name tends to map to the underlying package it demonstrates (`ping`→`kratos-ping`, `gorm`→`gormrepo`, `zap`→`kratos-zap`, …). **A fork's entire mission is to clearly demonstrate that one underlying package.**

**Two core sub-modules:**

- `demo1kratos` — Student CRUD microservice
- `demo2kratos` — Article CRUD microservice

Both are textbook Kratos layouts: `cmd/` entry point + `internal/{biz,data,service,server}` layers + `api/**/*.proto` contracts + `configs/config.yaml`, with proto-first design, Wire dependency injection, and dual gRPC/HTTP transport. The two demos share one PostgreSQL database via GORM; an article (demo2) references a student (demo1) via `StudentID` and demonstrates relationship logic such as cascade delete and existence checks; in some forks (e.g. `trace` / `nacos`) demo2 also makes cross-service calls into demo1.

**A few basic facts (don't get them wrong):**

- This is a **learning / comparison** project, but the code is written to **genuine production standards** (true cascade deletes, true transaction locks, precise error codes) — that's what makes the comparison trustworthy and instructive, so **don't cut corners just because "it is a mere demo".**
- It is a **Go multi-module repo**: one root module + two sub-modules, `demo1kratos` and `demo2kratos`.
- **Each fork is its own independent repository** (an independent clone), not a git submodule.

**How to build / run (the commands sit in the Makefile — no need to fumble):**

- Enter a sub-project `demo1kratos` / `demo2kratos`, run `make all` (generates api+config+wire) then `make build`
- Run a service: the standard Kratos way — `go run ./cmd/demo1kratos -conf ./configs/config.yaml` (requires a running PostgreSQL)

**⚠️ Always use the built-in Makefile targets — do NOT improvise your own `go build` commands.** A bare `go build ./...` without `-o` litters stray binaries across the working tree; `make build` collects each artifact under the sub-project's `bin/` directory (the sub-project's `.gitignore` excludes it), keeping the tree clean.

**Before doing anything, work out which one you sit in right now — a "source project" / a "fork project"; here we just make the split plain:**

- **A source project** is the pure stump; its `changes/` always shows `✅ NO CHANGES`. It **cannot have any feature difference** — don't go hunting down "something extra" here, and don't force one in.
- **A fork project** grows from the source, each demonstrating just one underlying feature. **That "changed thing" shows up in a fork alone**, recorded in [changes/](./changes). Read changes/ to see what it adds — and **apart from that one feature, do not rewrite the parts shared with the source.**

**⛔ The single most important rule: apart from the one feature this fork demonstrates, everything else MUST be a perfect match with the source (the stump) — nothing added, nothing removed, nothing changed.** The entire value of this collection is _comparison_: the stump shows "how you write it by default", and a fork shows "the same thing improved, plus one extra capability, using some underlying feature". That gain shows up next to the stump, so the shared parts must line up down to the byte. What you must **not touch**:

- **Interfaces / proto**: rpc method count, message fields, error enums — a perfect match with the source
- **biz logic**: method count, model fields, error semantics (which failure returns which error code) must match
- **Relationship & transaction logic**: cascade delete, foreign-key / existence checks, transaction locks — do not simplify any of them away
- **Each detail**: sort direction, pagination guards, default values, edge-case handling — each of them follows the source

**The most common mistake — changing shared code because "I think this should improve".** Classic crashes: the source creates two tables and you make it one; the source has a pagination guard and you delete it; the source sorts DESC and you switch it to ASC; the source returns a set of precise error codes and you collapse them into a single generic `ServerError`. These **demonstrate no feature and pollute the comparison — each of them is wrong.**

**Before changing any shared code, ask yourself: does this change demonstrate the one feature this fork is meant to show?**

- Yes → change it; that is the whole point of this fork.
- No → then it belongs to **the source itself**: change it in the source (so each fork benefits), else leave it as is. **Do not sneak it into a single fork.**

**Two kinds of fork — each fork is one of these two:**

- **Pure-equivalent**: the same logic, written more elegantly with some wrapper — logic is 100% equivalent to the source; just the _style_ differs.
- **Pure-extension**: a genuine new capability added on top of the source, while the source's original base logic is **kept as-is**.

**⛔ Translate the source, do not rewrite it — this matters most in "pure-equivalent" forks:** when you rewrite biz with some wrapper, treat it as a **translation** of the source — you may switch language (English identifiers → Chinese) and switch style (plain gorm → gormrepo and other ecosystem idioms), but the **logic must replicate the source 100%**: cascade delete, foreign-key / existence checks, precise error codes — each has to come through intact, **not simplified on your own initiative**. The order stands: first replicate the source's logic, then let the wrapper-based version stand next to the source, to be compared.

**If you are asked to change / review a fork — remember "it builds ≠ it's correct", and mechanically align against the source, item by item:**

- **proto must diff 0 against the source** (rpc method count, message fields, error enums — none may change).
- **grep a few signals**: are errors all collapsed into one generic code (= translation not complete)? are the cascade / check references still there? is the service method count right? are any model fields missing?
- **Each demo must be green**: `go build ./...` + `go vet ./...` + `gofmt -l .` (empty output) + (in the `test` fork) `go test ./...`.
- **Report "done" once you see the green with your own eyes** — don't claim completion on a hunch.

**`changes/` is generated by tests — do not hand-edit those `.md` files.** To update them, change the code and re-run `go test ./...` to regenerate.

> **Committing `changes/` (multi-module, two-round release):** `changes/` is built by diffing the **published** sub-module versions from the module cache (e.g. `demo1kratos@v0.0.5`) against your local workspace code. While you edit the source before tagging new sub-module versions, your local code sits ahead of the published version, so `changes/` shows diffs — and **that diff captures the record of what you changed, so you commit it as part of round 1** (the sub-module release). Then in round 2 (the root release) you bump the root `go.mod` to the new sub-module versions and re-run `go test ./...`; now base == local, so `changes/` goes back to `✅ NO CHANGES`, which you commit too. So `changes/` travels with each round: diffs at round 1, `✅ NO CHANGES` by round 2 — commit it both times, do not withhold it.

**The full release / upgrade / sync flow is in the "Upgrade & Sync Approach" section below** — remember the takeaway: **releasing and tagging happens on the source project alone; a fork just gets pushed, stays untagged, and does not drag in tags when syncing from upstream.**

## Introduction

**kratos-examples** is a reference implementation demonstrating best practices when building microservices with the [Go-Kratos](https://go-kratos.dev) framework. It serves as:

- 🎯 **Foundation Project** - The upstream template of 20+ specialized demo projects in the kratos-orz ecosystem
- 🛠️ **Toolchain Integration Example** - Showcasing kratos-orz development tools in action
- 📚 **Learning Resource** - Complete microservice structure following Kratos conventions
- ⚡ **Fast Development** - Auto-sync proto and code through magic commands like make orz

## About Go-Kratos

[Go-Kratos](https://go-kratos.dev) is a concise and efficient microservice framework that provides:

- Clean architecture with distinct separation of concerns
- Comprehensive middleware and plugin ecosystem
- Built-in gRPC and HTTP transport
- Excellent documentation and active ecosystem

**kratos-examples builds upon this solid foundation**, adding enhanced tooling and automation to streamline the development workflow.

## Example Projects

The following projects are forked from kratos-examples, each demonstrating different features and integrations such as authentication, configuration, scheduling, databases, logging, tracing, frontend integration, and more:

|    demo     |                      repo                      |
| :---------: | :--------------------------------------------: |
|     ast     |     https://github.com/kratos-examples/ast     |
| custom-auth | https://github.com/kratos-examples/custom-auth |
| static-auth | https://github.com/kratos-examples/static-auth |
|   config    |   https://github.com/kratos-examples/config    |
|    cron     |    https://github.com/kratos-examples/cron     |
|     ebz     |     https://github.com/kratos-examples/ebz     |
|    cobra    |    https://github.com/kratos-examples/cobra    |
|    gorm     |    https://github.com/kratos-examples/gorm     |
|  gormzhcn   |  https://github.com/kratos-examples/gormzhcn   |
|    cors     |    https://github.com/kratos-examples/cors     |
|    i18n     |    https://github.com/kratos-examples/i18n     |
|    nacos    |    https://github.com/kratos-examples/nacos    |
| rate-limit  | https://github.com/kratos-examples/rate-limit  |
|   swaggo    |   https://github.com/kratos-examples/swaggo    |
|    trace    |    https://github.com/kratos-examples/trace    |
|    test     |    https://github.com/kratos-examples/test     |
|    vue3     |    https://github.com/kratos-examples/vue3     |
|  vue3zhcn   |  https://github.com/kratos-examples/vue3zhcn   |
|    wire     |    https://github.com/kratos-examples/wire     |
|     zap     |     https://github.com/kratos-examples/zap     |
|    zapzh    |    https://github.com/kratos-examples/zapzh    |
|   migrate   |   https://github.com/kratos-examples/migrate   |
|    ping     |    https://github.com/kratos-examples/ping     |
| supervisord | https://github.com/kratos-examples/supervisord |

## Core Features

### 🛠️ Recommended Toolchain

The kratos-orz ecosystem ships focused utilities. The first entries target the Kratos and proto workflow; the rest are shared Go helpers. Set up each one as you need — each command stands on its own, most IDEs show a run button on the block so a click does the setup, and each repo link opens the complete docs.

#### `orzkratos-add-proto`

Scaffolds a fresh proto file in the api path, so a new endpoint starts from a prepared skeleton.

Repo: https://github.com/yylego/kratos-orz

```bash
go install github.com/yylego/kratos-orz/cmd/orzkratos-add-proto@latest
```

#### `orzkratos-srv-proto`

Syncs the service code onto the proto contract: adds method stubs, hides removed methods, and reorders to match.

Repo: https://github.com/yylego/kratos-orz

```bash
go install github.com/yylego/kratos-orz/cmd/orzkratos-srv-proto@latest
```

#### `protoc-gen-orzkratos-errors`

Generates Go errors helpers from proto enums, with status codes and nested wrapping.

Repo: https://github.com/yylego/kratos-errgen

```bash
go install github.com/yylego/kratos-errgen/cmd/protoc-gen-orzkratos-errors@latest
```

#### `wirekratos`

Runs Wire DI in Kratos workspace mode, with framework-aware markers.

Repo: https://github.com/yylego/kratos-wire

```bash
go install github.com/yylego/kratos-wire/cmd/wirekratos@latest
```

#### `clang-format-batch`

Batch-formats proto and cpp sources in one shot.

Repo: https://github.com/yylego/clang-format

```bash
go install github.com/yylego/clang-format/cmd/clang-format-batch@latest
```

#### `depbump`

Upgrades module dependencies in a single pass, across the whole workspace if asked.

Repo: https://github.com/yylego/depbump

```bash
go install github.com/yylego/depbump/cmd/depbump@latest
```

#### `go-lint`

Wraps golangci-lint with auto-format, runs format plus static checks across modules.

Repo: https://github.com/yylego/go-lint

```bash
go install github.com/yylego/go-lint/cmd/go-lint@latest
```

#### `go-commit`

Automates git commits with Go formatting baked in.

Repo: https://github.com/yylego/go-commit

```bash
go install github.com/yylego/go-commit/cmd/go-commit@latest
```

#### `tago`

Manages git tags with semantic auto-increment, and handles the path prefix sub-module tags too.

Repo: https://github.com/yylego/tago

```bash
go install github.com/yylego/tago/cmd/tago@latest
```

### ⚡ Magic Command: `make orz`

The core feature - auto synchronization between proto files and service code:

```bash
make orz
```

**What it does:**

- ✅ New methods in proto → Auto generates service method stubs
- ✅ Deleted methods → Converts to unexported functions (preserves logic)
- ✅ Reordered methods → Auto rearranges service code to match proto sequence

**Workflow example:**

1. Add `CreateArticle` method to the proto file
2. Run `make orz`
3. Service generates `CreateArticle` method stub
4. Implement the business logic

## Upgrade & Sync Approach

This project is a multi-module repo: a root module plus two sub-modules (`demo1kratos`, `demo2kratos`). The root references the sub-modules, so a release runs in two rounds — tag the sub-modules first (with the path prefix, e.g. `demo1kratos/v0.0.X`), then bump the root onto the new sub-module versions and tag the root (`v0.0.X`).

**Grasp one basic principle first and each step below then follows:** the two sub-modules `demo1kratos` / `demo2kratos` are **real, referenced Go modules** — the root requires them by version, and `changes/` uses their **published** versions as the comparison baseline. So they must be formally released and tagged; a fork, by contrast, is not `go get`'d, so tagging it is meaningless. And because `changes/` is by definition "the published sub-module version vs. your local code", **each time you edit the source project you must complete the loop "release new sub-module versions → point the root at them → regenerate `changes/`" before the diff returns to zero (`✅ NO CHANGES`) — the release process is, in essence, driving that diff back to zero.**

**It follows that:** just the source project (this upstream) needs releasing and tagging; downstream forks (fork projects) are pure demos, not referenced, so they hold no tags at all — just push once editing wraps up.

### Source-project release flow (ordered, step by step)

1. **Upgrade all dependencies and toolchain, but lock the language version.** Bring each dependency and tool to the latest, leaving just the `go` directive (language version) in `go.mod` untouched, and use it to pin the whole flow (e.g. `GOTOOLCHAIN=go<version-in-go.mod>`). **Why: you're upgrading the _ecosystem_ (deps / tools), so the language baseline must stay put; otherwise "upgrading dependencies" quietly bumps the Go version too and turns into an uncontrolled, sweeping change.**
2. **Regenerate intermediate artifacts.** Once the toolchain is upgraded, regenerate the generated code with the tools (`wire`; `buf generate` when proto changed), so the generated output is fresh with zero drift.
3. **Verify end to end.** Run all unit tests + all static checks (`go build ./...`, `go vet ./...`, `gofmt -l .`); and **compile the services, start them up, and run a local API test** — just once everything passes do you enter the two release rounds below.
4. **Round 1 — release the sub-modules.** Commit this round's source changes together with `changes/` (which now carries the diff — precisely the record of "what changed this round") **→ push** → then tag **new versions** of the two sub-modules `demo1kratos` / `demo2kratos` **→ push the tags**.
5. **Round 2 — release the root.** Bump the root `go.mod` onto the new sub-module versions → re-run `go test ./...` at the root to regenerate `changes/` (now base == local, so it returns to `✅ NO CHANGES`) → **in this round, polish just the root project's own content (root README / comments / docs); do not touch any sub-project (`demo1kratos` / `demo2kratos`) code** → **commit → push** (including this `✅ NO CHANGES` version of changes/) → finally tag the **root module** (`v0.0.X`) **→ push the tag**.

> **Each round is the full sequence "commit → push → tag → push tags": a tag must sit on a commit that has already been pushed to the remote — do not tag straight off a commit that is not yet pushed.**

> **⛔ Round 2 must not touch sub-project code again — the reason:** once round 1 tags the sub-modules, that code is _frozen_ — the tag (e.g. `demo1kratos@v0.0.6`) is the very baseline each fork's `changes/` diffs against. If you edit sub-project code in round 2 without cutting a new tag, the tag stops matching the code; then each fork comparing against that tag gets spurious diffs that aren't its feature at all, and its `changes/` is ruined. If a sub-project does need more changes, it has to go through round 1 again (a fresh tag) — it must not be slipped into round 2.**

### Fork sync flow

**Principle:** a fork is a **living twin** of "the stump + one feature" — when the stump moves forward (e.g. a sweeping Kratos upgrade), that feature has to be re-applied onto the **new stump**. So syncing is not a plain pull; it's "swap in the new pot, then replant the same seedling". Downstream forks each focus on a single feature (trace, gorm, zap, i18n, ...) and sync from this upstream now and then to stay aligned with the newest Kratos version. Key points:

- **Do not sync tags**: don't drag in the upstream's own tags (a fork carries no tags anyway).
- **Finish the whole merge before pushing.** Resolve conflicts → re-apply this fork's customizations line by line → regenerate code → get tests and checks green → regenerate `changes/`, and just commit/push once everything has landed; **do not push while the merge is unfinished (`MERGE_HEAD` still dangling), and don't drift off mid-merge.**
- **When the upstream stump changes, the fork's customization must adapt.** Once merged, compare what the stump changed (e.g. a sweeping framework upgrade, API changes) and **re-apply this fork's small set of customizations onto the new code, adapting to the new API**; the blueprint is in `changes/`.
- A fork **does not merge back upstream** — each stands as a standalone example.

The utilities above map onto each task: `depbump` handles dependencies, `tago` handles tags, `go-lint` handles format and checks. The exact sequence is best driven one step at a time, not scripted, so the flow adapts to each situation.

### Code Changes

The [changes/](./changes) section contains markdown files documenting code differences:

- [changes/demo1.md](./changes/demo1.md) - Demo1 changes compared to source
- [changes/demo2.md](./changes/demo2.md) - Demo2 changes compared to source
- [changes/aside.md](./changes/aside.md) - Aside modules and sibling projects

Tests auto-generate these files:

```bash
go test -v -run TestGenerateDemo1Changes # Generate demo1.md
go test -v -run TestGenerateDemo2Changes # Generate demo2.md
go test -v -run TestGenerateAsideChanges # Generate aside.md
```

**In source project:** Files show `✅ NO CHANGES`

**In fork projects:** Files show code differences with syntax highlighting, making it simple to track customizations on GitHub.

## Project Structure

### Built-in Demos

Provides two demos to showcase the usage of various features:

- [demo1kratos](./demo1kratos) - Student CRUD microservice (simple Kratos example)
- [demo2kratos](./demo2kratos) - Article CRUD microservice (advanced features and integrations)

Both demos follow standard Kratos project structure with proto-first design, Wire DI, and gRPC/HTTP endpoints.

### Input Validation

The project demonstrates a two-stage validation pattern:

- **Service Stage** - Returns `ErrorBadParam` (HTTP 400) on invalid input, giving clients actionable feedback
- **Biz Stage** - Uses `must` assertions (panic on failure) as a safeguard. Since the service stage has done the validation, the biz stage is just a redundant check, so it uses simple assertions instead of returning verbose errors. This also ensures the biz module is self-protected — even if someone invokes it from a new context and forgets to validate inputs, the assertions catch it. Instead of scattering defensive checks throughout the code, assertions at the entrance make the module refuse to execute on invalid inputs, keeping the downstream logic clean and confident

### Data Access

Both demos use a shared PostgreSQL database via GORM.

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
