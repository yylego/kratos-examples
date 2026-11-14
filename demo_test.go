// Package examples_test compares base (upstream cache) and fork (workspace) code
// Contains test functions to show, compare, and generate difference reports
//
// examples_test 比较基准（上游缓存）和 fork（工作区）代码的差异
// 包含显示、比较和生成差异报告的测试函数
package examples_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/yylego/kratos-compare/comparekratos"
	examples "github.com/yylego/kratos-examples"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath"
)

func TestMain(m *testing.M) {
	demo1Base, demo1Fork := examples.GetDemo1BasePath(), examples.GetDemo1ForkPath()
	fmt.Println(demo1Base, "vs", demo1Fork)
	if demo1Base == demo1Fork {
		os.Exit(0)
	}
	demo2Base, demo2Fork := examples.GetDemo2BasePath(), examples.GetDemo2ForkPath()
	fmt.Println(demo2Base, "vs", demo2Fork)
	if demo2Base == demo2Fork {
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestShowDemo1Changes(t *testing.T) {
	comparekratos.ComparePath(examples.GetDemo1BasePath(), examples.GetDemo1ForkPath())
}

func TestShowDemo2Changes(t *testing.T) {
	comparekratos.ComparePath(examples.GetDemo2BasePath(), examples.GetDemo2ForkPath())
}

func TestCompareDemo1Modules(t *testing.T) {
	comparekratos.ComparePath(
		filepath.Join(examples.GetDemo1BasePath(), "go.mod"),
		filepath.Join(examples.GetDemo1ForkPath(), "go.mod"),
	)
}

func TestCompareDemo2Modules(t *testing.T) {
	comparekratos.ComparePath(
		filepath.Join(examples.GetDemo2BasePath(), "go.mod"),
		filepath.Join(examples.GetDemo2ForkPath(), "go.mod"),
	)
}

func TestShowDemo1ReadableChanges(t *testing.T) {
	comparekratos.ShowReadableChanges(examples.GetDemo1BasePath(), examples.GetDemo1ForkPath())
}

func TestShowDemo2ReadableChanges(t *testing.T) {
	comparekratos.ShowReadableChanges(examples.GetDemo2BasePath(), examples.GetDemo2ForkPath())
}

//go:generate go test -run ^TestGenerateDemo1Changes$
func TestGenerateDemo1Changes(t *testing.T) {
	outputRoot := osmustexist.ROOT(runpath.PARENT.Join("changes"))
	outputPath := filepath.Join(outputRoot, "demo1.md")
	comparekratos.GenerateChangesFile(examples.GetDemo1BasePath(), examples.GetDemo1ForkPath(), outputPath)
}

//go:generate go test -run ^TestGenerateDemo2Changes$
func TestGenerateDemo2Changes(t *testing.T) {
	outputRoot := osmustexist.ROOT(runpath.PARENT.Join("changes"))
	outputPath := filepath.Join(outputRoot, "demo2.md")
	comparekratos.GenerateChangesFile(examples.GetDemo2BasePath(), examples.GetDemo2ForkPath(), outputPath)
}

//go:generate go test -run ^TestGenerateAsideChanges$
func TestGenerateAsideChanges(t *testing.T) {
	root := runpath.PARENT.Path()
	excludeNames := []string{
		"changes",
		filepath.Base(examples.GetDemo1ForkPath()),
		filepath.Base(examples.GetDemo2ForkPath()),
	}
	outputPath := filepath.Join(root, "changes", "aside.md")
	comparekratos.GenerateTreeChanges(root, excludeNames, outputPath)
}
