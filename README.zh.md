<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/kratos-examples/release.yml?branch=main&label=BUILD)](https://github.com/yylego/kratos-examples/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/kratos-examples)](https://pkg.go.dev/github.com/yylego/kratos-examples)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/kratos-examples/main.svg)](https://coveralls.io/github/yylego/kratos-examples?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/kratos-examples.svg)](https://github.com/yylego/kratos-examples/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/kratos-examples)](https://goreportcard.com/report/github.com/yylego/kratos-examples)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# kratos-examples

基于 Go-Kratos 框架构建的演示项目

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

## 给 AI 助手的说明

> 如果你是 AI 编程助手，动手前先把这一段读完、读懂。
>
> **⭐ 最关键的一句：下面所有的规矩和流程，都不是我们硬性强加给你的教条——它们全都是你一旦真正理解了这个项目的定位、以及各个部分的含义之后，自然而然、顺理成章推出来的结果。所以别去硬背条目，要去理解原理；原理一通，这些流程你自己都能推导出来，也就不会跑偏瞎搞了。**

**项目定位：** `kratos-examples` 是一个基于 [Go-Kratos](https://go-kratos.dev) 的**标准演示树桩**——一个规范、地道的 Kratos 工程，也是下游 20+ 个专用演示 fork 的共同基础模板。它**不是**自定义框架，就是经典 Kratos 的标准布局，你早就认识它，别当成陌生项目从头猜。

**背景（它在生态里的位置——上通下达）：** 这项目不是孤立的，它夹在一条链的中间，起的是"承上启下"的作用：

- **下（底层）**：yylego 生态里有一批真正可复用的底层 Kratos 封装包——比如 `kratos-ping`、`kratos-gorm` / `gormrepo`、`kratos-zap`、`kratos-ebz`、`kratos-i18n`、`kratos-trace` 等。**这些才是"产品"**，是给别人 `go get` 引用的。
- **中（本项目）**：`kratos-examples` 是**演示层**。源项目（树桩）演示"不用这些封装、裸 Kratos 怎么写"；每个 fork 演示"用上某一个底层包之后，同一件事变成什么样、好在哪"。
- **上（使用者 / 你这个 AI）**：想知道某个底层包到底带来了什么，不用通读它的源码——直接看对应 fork 的 [changes/](./changes)，那就是被机械提取出来的、零噪声的答案。

所以每个 fork 的名字通常就对应它演示的那个底层包（`ping`→`kratos-ping`、`gorm`→`gormrepo`、`zap`→`kratos-zap`……）。**一个 fork 的全部使命，就是把"它对应的那一个底层包"演示清楚。**

**两个核心子模块：**

- `demo1kratos` —— 学生（Student）CRUD 微服务
- `demo2kratos` —— 文章（Article）CRUD 微服务

两者都是教科书式 Kratos 结构：`cmd/` 入口 + `internal/{biz,data,service,server}` 分层 + `api/**/*.proto` 契约 + `configs/config.yaml`，proto-first、Wire 依赖注入、gRPC/HTTP 双协议。两个 demo 通过 GORM 共用同一个 PostgreSQL 库；文章（demo2）通过 `StudentID` 关联学生（demo1），并演示级联删除、存在性校验等关系逻辑；有的 fork（如 `trace` / `nacos`）里 demo2 还会跨服务调用 demo1。

**几条基本常识（别搞错）：**

- 这是**学习 / 对比**项目，但代码是按**真实生产质量**写的（真级联删除、真事务锁、真精确错误码）——这样"对比"才可信、才有教学价值，**别因为"反正是 demo"就偷工减料。**
- 它是 **Go 多模块仓库**：一个根模块 + `demo1kratos` + `demo2kratos` 两个子模块。
- **各个 fork 是各自独立的仓库**（各自独立 clone），不是 git submodule。

**怎么构建 / 运行（命令都在 Makefile 里，别晕头转向）：**

- 进入子项目 `demo1kratos` / `demo2kratos`，先 `make all`（生成 api+config+wire）再 `make build`
- 运行服务：标准 Kratos 跑法 —— `go run ./cmd/demo1kratos -conf ./configs/config.yaml`（需要一个可用的 PostgreSQL）

**⚠️ 一律用 Makefile 里现成的 target，别凭自己的想法乱敲 `go build`。** 直接 `go build ./...` 不带 `-o`，会把一堆二进制文件散落在工作区里到处都是、脏得像随地大小便；`make build` 会在子项目里统一把产物放进 `bin/` 目录（子项目的 `.gitignore` 已忽略它），干净利落。

**动手前先分清你现在在「源项目」还是「分叉项目」——你自己能判断出来，这里只把两者的区别说清：**

- **源项目**是纯树桩，`changes/` 永远显示 `✅ NO CHANGES`。它**不可能有任何特性差异**，别在这里瞎找"多出来的东西"，也别硬往里加。
- **分叉项目**从源头长出来、每个只演示一个底层特性。**那个"变化的东西"只有分叉项目才有**，记录在 [changes/](./changes) 里。看 changes/ 就知道它加了什么——**除了那个特性，别擅自改写与源头共享的部分。**

**⛔ 最核心的铁律：除了本 fork 要演示的那一个特性，其余一切必须跟源头（树桩）一模一样——不加、不删、不改。** 整个合集的价值全在"对比"：树桩演示"正常怎么写"，fork 演示"用某个底层特性把同一件事写得更好、或多加一个能力"；而"更好"只有跟树桩并排才看得出来，所以共享的部分必须逐字对齐。具体哪些**绝不许动**：

- **接口 / proto**：rpc 方法数、消息字段、错误枚举，全照源头
- **biz 逻辑**：方法数、模型字段、错误语义（哪个失败返回哪个错误码）都要一致
- **关系与事务逻辑**：级联删除、外键/存在性校验、事务锁，一个都不能简化掉
- **各种细节**：排序方向、分页守卫、默认值、边界处理，统统照源头

**最容易犯的错——凭"我觉得这里该更好"去动共享代码。** 典型翻车：源头建两张表你改成一张、源头有分页守卫你删了、源头排序 DESC 你改成 ASC、源头一组精确错误码你塌成一个笼统的 `ServerError`。这些**既没演示特性、又污染了对比，一律是错的**。

**想改共享代码前，先自问：这个改动演示了本 fork 要演示的那个特性吗？**

- 是 → 改，这正是这个 fork 存在的意义。
- 不是 → 那它属于**源头本身**：要么去改源头、让所有 fork 都受益，要么根本不该改。**绝不能在单个 fork 里偷偷分叉出去。**

**两类 fork，每个 fork 必属其一、且只属一类：**

- **纯等效**：同样的逻辑，换某个封装写得更优雅——逻辑跟源头 100% 等效，只有"写法"不同。
- **纯扩展**：在源头上正经加一个新能力，但源头原有的基础逻辑**原样保留**。

**⛔ 忠实翻译（尤其"纯等效"类 fork）：** 用某个封装重写 biz 时，把它当成对源头的**翻译**——可以换语言（英文标识符→中文）、换写法（裸 gorm→gormrepo 等生态用法），但**逻辑必须 100% 忠实复刻源头**：级联删除、外键/存在性校验、精确错误码，一样都要"译"过来，**绝不许自作主张简化**。顺序永远是：先忠实复刻源头逻辑，再让高阶特性能跟源头并排对比。

**如果你被叫去改 / 复查一个 fork——记住"build 过 ≠ 对"，要机械对齐源头逐项核：**

- **proto 与源头 diff 应为 0**（rpc 方法数、消息字段、错误枚举都不许动）。
- **grep 几个信号**：错误是不是全塌成一个通用码（=翻译没翻全）、级联 / 校验的引用还在不在、service 方法数够不够、模型字段有没有少。
- **每个 demo 跑绿**：`go build ./...` + `go vet ./...` + `gofmt -l .`（输出为空）+（若是 `test` fork）`go test ./...`。
- **改完亲眼看到绿灯再说"好了"**，别凭感觉就宣称完成。

**`changes/` 是测试自动生成的，别用手去改那些 `.md`。** 要更新就改代码、再跑 `go test ./...` 重生成。

> **提交 `changes/`（多模块、两轮发布）：** `changes/` 是拿模块缓存里**已发布**的子项目版本（如 `demo1kratos@v0.0.5`）跟本地工作区代码比对生成的。你在源项目里改了代码、还没给子项目打新标签之前，本地代码领先于已发布版本，`changes/` 会冒出差异——而**这份差异正是"我们改了啥"的记录，所以轮1（发子模块）时就把它一起提交进去**。等到轮2（发根项目）把根 `go.mod` 的依赖升到子项目新版本、在根目录重跑 `go test ./...`，此时基准=新版=本地，`changes/` 回到 `✅ NO CHANGES`，这份也提交。所以 `changes/` 随每一轮走：轮1 后带差异、轮2 后回 `✅ NO CHANGES`——**两次都提交，别扣着不提交。**

**完整的发版 / 升级 / 同步流程见下面《升级与同步思路》一节**——记住结论：**发版打标签只有源项目需要；fork 只是简单推送、永不打标签，同步上游时也不要拉标签。**

## 项目简介

**kratos-examples** 是使用 [Go-Kratos](https://go-kratos.dev) 框架构建微服务的最佳实践参考实现。它的作用是：

- 🎯 **基础项目** - kratos-orz 生态系统中 20+ 专用演示项目的基础模板
- 🛠️ **工具链集成示例** - 展示 kratos-orz 开发工具的实践应用
- 📚 **学习资源** - 完整的微服务架构，遵循 Kratos 规范
- ⚡ **快速开发** - 通过 make orz 等魔法命令实现 proto 和代码自动同步

## 关于 Go-Kratos

[Go-Kratos](https://go-kratos.dev) 是一个简洁高效的微服务框架，它提供：

- 清晰的架构和关注点分离
- 完善的中间件和插件生态
- 内置 gRPC 和 HTTP 协议支持
- 优秀的文档和活跃的生态

**kratos-examples 在这个坚实的基础上**，添加了增强的工具链和自动化能力，简化开发工作流程。

## 示例项目

以下项目均基于 kratos-examples 进行 fork，分别展示不同的功能特性和集成方案，涵盖认证、配置、定时任务、数据库、日志、链路追踪、前端集成等多个方面：

|    演示     |                      仓库                      |
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

## 核心功能

### 🛠️ 推荐工具链

kratos-orz 生态提供一组专注的小工具。前几个面向 Kratos 与 proto 工作流，其余是通用的 Go 小工具。按需取用即可——每条命令各自独立（不编号，方便日后增减），多数 IDE 会在代码块上显示运行按钮，点一下即可安装，每个仓库链接里有完整文档。

#### `orzkratos-add-proto`

在 api 目录下生成一个新的 proto 文件，新接口从准备好的骨架起步。

仓库：https://github.com/yylego/kratos-orz

```bash
go install github.com/yylego/kratos-orz/cmd/orzkratos-add-proto@latest
```

#### `orzkratos-srv-proto`

让 service 代码与 proto 契约保持同步：补全方法框架、隐藏已删方法、按 proto 顺序重排。

仓库：https://github.com/yylego/kratos-orz

```bash
go install github.com/yylego/kratos-orz/cmd/orzkratos-srv-proto@latest
```

#### `protoc-gen-orzkratos-errors`

从 proto 枚举生成 Go 错误码辅助函数，带状态码和错误嵌套。

仓库：https://github.com/yylego/kratos-errgen

```bash
go install github.com/yylego/kratos-errgen/cmd/protoc-gen-orzkratos-errors@latest
```

#### `wirekratos`

在 Kratos 工作区模式下运行 Wire 依赖注入，带框架感知标识。

仓库：https://github.com/yylego/kratos-wire

```bash
go install github.com/yylego/kratos-wire/cmd/wirekratos@latest
```

#### `clang-format-batch`

一次性批量格式化 proto 和 cpp 源码。

仓库：https://github.com/yylego/clang-format

```bash
go install github.com/yylego/clang-format/cmd/clang-format-batch@latest
```

#### `depbump`

一次性升级模块依赖，需要时可覆盖整个工作区。

仓库：https://github.com/yylego/depbump

```bash
go install github.com/yylego/depbump/cmd/depbump@latest
```

#### `go-lint`

封装 golangci-lint 并自动格式化，跨模块运行格式化与静态检查。

仓库：https://github.com/yylego/go-lint

```bash
go install github.com/yylego/go-lint/cmd/go-lint@latest
```

#### `go-commit`

自动化 git 提交，内置 Go 代码格式化。

仓库：https://github.com/yylego/go-commit

```bash
go install github.com/yylego/go-commit/cmd/go-commit@latest
```

#### `tago`

管理 git 标签并语义化自增，也处理带路径前缀的子模块标签。

仓库：https://github.com/yylego/tago

```bash
go install github.com/yylego/tago/cmd/tago@latest
```

### ⚡ 魔法命令：`make orz`

最强大的功能 - proto 文件与服务代码之间的自动同步：

```bash
make orz
```

**它的作用：**

- ✅ proto 中新增方法 → 自动生成服务方法框架
- ✅ 删除的方法 → 转换成非导出函数（保留逻辑）
- ✅ 重排序方法 → 自动调整服务代码顺序以匹配 proto

**工作流示例：**

1. 在 proto 文件中添加 `CreateArticle` 方法
2. 运行 `make orz`
3. 服务自动生成 `CreateArticle` 方法框架
4. 实现业务逻辑

## 升级与同步思路

本项目是多模块仓库：一个根模块加两个子模块（`demo1kratos`、`demo2kratos`）。根模块引用子模块，所以发布分两轮——先给子模块打标签（带路径前缀，如 `demo1kratos/v0.0.X`），再把根模块升级到子模块的新版本并给根模块打标签（`v0.0.X`）。

**先理解一条基本原理，后面每一步就都顺理成章了：** 两个子项目 `demo1kratos` / `demo2kratos` 是**真会被引用的 Go 模块**——根模块按版本号引用它们，`changes/` 也拿它们**已发布**的版本当比对基准。所以它们必须正式发版、打标签；而 fork 没人 `go get`，打标签毫无意义。又因为 `changes/` 本就是"已发布的子项目版本 vs 本地代码"的差异，**源项目每改一轮代码，都得走完"给子项目发新版 → 根模块引用新版 → 重生成 `changes/`"这一圈，差异才会重新归零（`✅ NO CHANGES`）——发版流程的本质，就是让这个差异重新归零。**

**由此推出：只有源项目（本上游）才需要发版本、打标签；下游 fork（分叉项目）是纯演示、不会被引用，一律不打任何标签，改完简单推送即可。**

### 源项目发版流程（有序，逐步来）

1. **全面升级依赖与工具链，但锁定语言版本。** 把所有依赖和工具包升到最新，唯独 `go.mod` 里的 `go` 指令（语言版本）保持不动，并用它锁住整个流程（如 `GOTOOLCHAIN=go<go.mod 里的版本>`）。**为什么：升的是"生态"（依赖 / 工具），语言基线得稳住；否则"升依赖"顺手把 Go 版本也跳了，就成了不可控的大变动。**
2. **重新生成中间产物。** 工具链升级后，用工具重新生成那些"生成出来的代码"（`wire`；proto 动过才 `buf generate`），确保生成物是新鲜的、零漂移。
3. **全面验证。** 跑全部单元测试 + 全部静态检查（`go build ./...`、`go vet ./...`、`gofmt -l .`）；并且**真正把服务编译、启动起来，做一遍本地接口测试**——全部通过，才进入下面的两轮发布。
4. **轮 1——发子模块。** 把这一轮的源码改动，连同 `changes/`（此时它带着差异、正是"这轮改了啥"的记录）一起**提交 → 推送** → 再给 `demo1kratos` / `demo2kratos` 两个子项目打**新标签 → 推送标签**。
5. **轮 2——发根模块。** 把根 `go.mod` 的依赖升到子项目新版本 → 在根目录重跑 `go test ./...` 重新生成 `changes/`（此时基准=新版=本地，回到 `✅ NO CHANGES`）→ **这一轮只润色根项目自己的内容（根 README / 注释 / 文档），绝不再碰任何子项目（`demo1kratos` / `demo2kratos`）的代码** → **再提交 → 推送**（含这份 `✅ NO CHANGES` 的 changes/）→ 最后给**根模块**打**标签 → 推送标签**（`v0.0.X`）。

> **每一轮都是完整的"提交 → 推送 → 打标签 → 推送标签"：标签必须落在已经推送到远程的提交上，绝不能提交完不推送就直接打标签。**

> **⛔ 轮 2 为什么绝不能再动子项目代码：** 子项目在轮 1 打完标签后就"冻结"了——那个标签（如 `demo1kratos@v0.0.6`）正是**所有 fork 的 `changes/` 用来做比对的基准**。你若在轮 2 又改了子项目代码却没重新打标签，标签就跟实际源码对不上了；将来每个 fork 拿这个标签做对比时，都会冒出"根本不属于本 fork 特性"的**假差异**，那份 `changes/` 就废了。真要再改子项目，就得给它重新走一遍轮 1（重新打新标签），绝不能塞进轮 2。

### 分叉项目同步流程

**原理：** 一个 fork 就是"树桩 + 一个特性"的**活镜像**——树桩往前走（如 Kratos 大版本升级），那个特性就得在**新树桩**上重新贴一遍。所以同步不是简单拉取，而是"换了新盆、把原来那棵苗重新栽进去"。下游 fork 各自专注一个功能（trace、gorm、zap、i18n 等），会时而不时从本上游同步、以跟上最新 Kratos 版本。要点：

- **同步时不要同步标签**：上游自带的标签别拉进来（fork 本就不打标签）。
- **合并流程一定要完整走完再推送。** 解冲突 → 逐行补回本 fork 的定制 → 重新生成代码 → 跑绿测试与检查 → 重生成 `changes/`，全部落定后才提交推送；**绝不允许合并没走完（`MERGE_HEAD` 还挂着）就去推送或干别的**。
- **上游树桩变了，fork 的定制要跟着适配。** 合并后对照看树桩改了什么（如框架大版本升级、API 变更），把本 fork 那一小撮定制**逐行复刻到新版代码上、适配新 API**；蓝图就在 `changes/` 里。
- Fork **永不合并回上游**——每个都作为独立示例长期存在。

上面的工具正好对应各个环节：`depbump` 管依赖、`tago` 管标签、`go-lint` 管格式化与检查。具体顺序最好一步一步来、不固化成脚本，这样流程能随具体情况灵活调整。

### 代码变更

[changes/](./changes) 包含记录代码差异的 markdown 文件：

- [changes/demo1.md](./changes/demo1.md) - Demo1 相对源项目的变更
- [changes/demo2.md](./changes/demo2.md) - Demo2 相对源项目的变更
- [changes/aside.md](./changes/aside.md) - 附属模块和兄弟项目

这些文件通过测试自动生成：

```bash
go test -v -run TestGenerateDemo1Changes # 生成 demo1.md
go test -v -run TestGenerateDemo2Changes # 生成 demo2.md
go test -v -run TestGenerateAsideChanges # 生成 aside.md
```

**在源项目中：** 文件显示 `✅ NO CHANGES`

**在 fork 项目中：** 文件显示代码差异并带语法高亮，简单直观地在 GitHub 上查看定制化的改动。

## 项目结构

### 内置演示

提供两个内置演示项目，展示各种功能的使用：

- [demo1kratos](./demo1kratos) - Student CRUD 微服务（Kratos 简单示例）
- [demo2kratos](./demo2kratos) - Article CRUD 微服务（高级功能和集成）

两个演示都遵循标准的 Kratos 项目结构，采用 proto-first 设计、Wire 依赖注入、gRPC/HTTP 双协议端点。

### 参数校验

项目演示了双层校验模式：

- **Service 层** - 返回 `ErrorBadParam`（HTTP 400）提示客户端输入无效，给予可操作的反馈
- **Biz 层** - 使用 `must` 断言（失败时 panic）作为安全兜底。由于 Service 层已经做过校验，Biz 层只是冗余检查，因此采用简单的断言而非返回详细错误。这也确保了 Biz 模块自身的安全性——即使其他调用方在新的场景中复用该模块时忘记校验参数，断言也能及时捕获问题。相比在代码中到处散布防御性检查，在入口处用断言直接拒绝非法参数，让后续的业务逻辑保持简洁和确定性

### 数据层

两个演示通过 GORM 使用同一个 PostgreSQL 数据库。

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->
