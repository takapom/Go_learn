package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todos []string

func main() {
	//入力から読み取りのための準備
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo アプリ ---")
		fmt.Println("1. 追加")
		fmt.Println("2. 一覧表示")
		fmt.Println("3. 終了")
		fmt.Print("番号を入力してください: ")

		//ユーザーがエンターキーを押すまで待機
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Todoを入力: ")
			scanner.Scan()
			todo := scanner.Text()
			todos = append(todos, todo)
			fmt.Println("追加しました！")
		case "2":
			fmt.Println("\n--- Todo一覧 ---")
			for i, t := range todos {
				fmt.Printf("%d. %s\n", i+1, t)
			}
		case "3":
			fmt.Println("終了します。")
			return
		default:
			fmt.Println("無効な入力です。")
		}
	}
}
