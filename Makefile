# ========================================
# Start Kratos Development Adventure
# 开始你的 Kratos 开发之旅吧
# ========================================

# Format code in projects via command line
# 使用命令行整理项目中的代码
fmt:
	@echo "开始整理所有演示项目的代码..."
	cd demo1kratos && clang-format-batch --extensions=proto
	cd demo2kratos && clang-format-batch --extensions=proto
	@echo "✅ 所有项目的代码整理完成！"

# Build demo projects, includes proto generation, config processing, code generation, etc.
# 构建所有演示项目，包括 proto 生成、配置文件处理、代码生成等
all:
	@echo "开始构建所有演示项目..."
	cd demo1kratos && make all
	cd demo2kratos && make all
	@echo "✅ 所有项目构建完成！"

# ========================================
# Magic Command make orz - Auto Sync Proto Code with Service
# 魔法命令 make orz - 自动同步 Proto 代码与服务
# ========================================
# This is the most amazing feature! When you change proto files, run this command:
# Add interface → Auto add function implementation in service (adds stub, you implement the logic)
# Delete interface → Auto convert service method to unexported (avoids compile bugs)
# Sort functions → Sort service implementations based on proto definition sequence
# Usage:
#   1. Add CreateArticle interface in proto file
#   2. Run make orz
#   3. Service auto generates CreateArticle method stub, just add business logic!
#
# 这是最强大的功能！当你修改 proto 文件后，运行这个命令：
# 新增接口 → 自动在服务层添加对应的函数实现（新增个空函数，需要您实现函数内部逻辑）
# 删除接口 → 自动将服务代码中对应的方法改为非导出的（避免编译错误）
# 函数排序 → 根据你 proto 里定义的函数顺序重新排列服务里的实现代码
# 使用场景举例:
#   1. 在 proto 文件中新增了 CreateArticle 接口
#   2. 运行 make orz
#   3. 服务层自动生成 CreateArticle 方法框架，你只需要填充业务逻辑！
orz:
	@echo "开始执行 Proto-Service 自动同步..."
	cd demo1kratos && make all && orzkratos-srv-proto -auto
	cd demo2kratos && make all && orzkratos-srv-proto -auto
	@echo "✅ 同步完成！请检查生成的代码并完善业务逻辑"

# Regenerate the changes/ diffs of each fork; demo1kratos and demo2kratos each come with a separate make generate
# 重新生成各 fork 的 changes/ 差异文件；demo1kratos 与 demo2kratos 各自有自己的 make generate
generate:
	@go generate ./...

# ========================================
# TEMPLATE BEGIN: TEST AND COVERAGE CONFIG
# 模板开始: 测试和覆盖率配置
# ========================================
# Test and Coverage (GitHub Actions)
# 测试和覆盖率（GitHub Actions 自动执行）
# ========================================

# Coverage output DIR
# 覆盖率输出目录
COVERAGE_DIR ?= .coverage.out

# Reference: https://github.com/yylego/gormrepo/blob/main/Makefile
test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

# Run tests with custom flags
# 使用自定义参数运行测试
test-with-flags:
	@go test $(TEST_FLAGS) ./...

# ========================================
# TEMPLATE END: TEST AND COVERAGE CONFIG
# 模板结束: 测试和覆盖率配置
# ========================================
