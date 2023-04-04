package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 1. コマンドライン引数の解析
	contest := flag.String("contest", "", "contest name")
	numProblems := flag.Int("problems", 8, "number of problems")
	flag.Parse()

	if *contest == "" {
		fmt.Println("Error: contest name is required")
		os.Exit(1)
	}

	// 2. 問題名のリストを生成
	problems := generateProblemNames(*contest, *numProblems)

	// 3. ディレクトリ構造の作成
	baseDir := filepath.Join("contests", *contest)
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}

	for _, problem := range problems {
		problemDir := filepath.Join(baseDir, problem)
		if err := os.Mkdir(problemDir, os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		// 4. テンプレートファイルのコピー
		src := "template/main.go"
		dst := filepath.Join(problemDir, "main.go")
		if err := copyFile(src, dst); err != nil {
			fmt.Println("Error copying file:", err)
			os.Exit(1)
		}
	}
}

// generateProblemNames は、指定されたコンテスト名と問題数に基づいて問題名のリストを生成します。
func generateProblemNames(contest string, numProblems int) []string {
	problems := make([]string, 0, numProblems)
	for i := 0; i < numProblems; i++ {
		problem := contest + "_" + generateSuffix(i)
		problems = append(problems, problem)
	}
	return problems
}

// generateSuffix は、指定されたインデックスに対応するアルファベットの接尾辞を生成します。
func generateSuffix(index int) string {
	const alpha = "abcdefghijklmnopqrstuvwxyz"

	if index < 26 {
		return string(alpha[index])
	}

	prefix := index / 26
	suffix := index % 26

	return string(alpha[prefix-1]) + string(alpha[suffix])
}

// copyFile は src から dst へファイルをコピーします。
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	return nil
}
