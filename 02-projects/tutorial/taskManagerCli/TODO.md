設定ファイルが破損している（無効なJSON/YAML）：現在は静かに無視されるため、ユーザに警告した方がよい。
同時実行（複数プロセスで同一ファイルを操作）：競合でデータ損失が起きる。fcntl/lockfile等の排他ロックが必要。
ファイル権限エラー（読み書き不可）、ディスク満杯：書き込み失敗をハンドルしてユーザへ通知。
空の説明（add の引数が無い） — cobra が最低1引数を要求しているため制御済み。

## Taskman: 簡潔な設計ノート

以下は `main.go`（taskman CLI）の実装に関する整理・改善案です。

### 1. 概要

- 目的: ローカルJSONにタスクを保存し、`add`/`list`/`complete`/`delete` を提供するCLI。主に `cobra`（コマンド実装）と `viper`（設定管理）を使用。

### 2. 詳細フロー（短縮）

- 起動時にフラグ/設定を読み、`PersistentPreRun` で `loadTasks` を実行して状態をロード。
- サブコマンド（`add`/`list`/`complete`/`delete`）が `taskManager` を操作。
- 実行後に `PersistentPostRun` で `saveTasks` を呼び、JSONへ永続化する。

### 3. 想定される問題点（現状の留意点）

- 設定読み込み: `viper.AddConfigPath("$HOME")` は展開されないため `os.UserHomeDir()` を使うべき。
- エラーハンドリング: `ReadFile`/`Unmarshal`/`WriteFile` の戻り値を無視している箇所があり、失敗時の情報がユーザに届かない。
- 同時アクセス: ファイルロックがなく、複数プロセス同時操作で競合・データ破損が起きる可能性がある。
- ファイル/ディスクのエラー（権限、容量）への耐性が弱い。
- 文字列切り詰め: `truncate` がバイト単位で切る場合、多バイト文字で表示が壊れる可能性がある。

### 4. Code Quality & Refactoring（推奨）

- viper の設定パスをホームディレクトリで正しく指定し、CLIフラグと viper を `BindPFlag` で同期する。
- `loadTasks`/`saveTasks` をエラーを返す `PreRunE`/`PostRunE`（あるいは `PersistentPreRunE`/`PersistentPostRunE`）に置き換え、失敗したら明確にエラーメッセージを返す。
- 書き込み前にディレクトリ作成（`os.MkdirAll`）・一時ファイル→原子置換パターンを利用し、破損リスクを下げる。
- ファイルロック（`flock` など）で同時実行の競合を防止する。

#### PreRunE / PostRunE へ移行（追加）

`PersistentPreRun` / `PersistentPostRun` は実行時にエラーを返せないため、エラーを起点とした安全な停止や適切なユーザ通知が行いづらいです。`*RunE` 系を使うと `error` を返せるため、より堅牢な実装になります。

簡単な移行例:

```go
// 署名は error を返す
func loadTasksE(cmd *cobra.Command, args []string) error {
	// ...ファイル読み込み・Unmarshal...
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}
	return nil
}

func saveTasksE(cmd *cobra.Command, args []string) error {
	// ...Marshal・WriteFile...
	if err := os.WriteFile(taskManager.FilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}
	return nil
}

// rootCmd の定義で:
// PersistentPreRun: loadTasks, -> PersistentPreRunE: loadTasksE
// PersistentPostRun: saveTasks, -> PersistentPostRunE: saveTasksE
```

この変更により、読み込み／保存の失敗を `rootCmd.Execute()` の呼び出し元で受け取り、適切な終了コードとエラーメッセージで終了できます。

### 5. エッジケース（要チェック項目）

- 設定ファイルが破損している場合は警告を出す、もしくはバックアップから復元する戦略。
- 同時実行時の排他制御（ロック）と、書き込み失敗時のロールバックや一時ファイル戦略。
- 多バイト文字列の切り詰め処理（`truncate` をランニングルネンスで扱う等）。
- `NextID` の衝突や桁あふれ（大規模運用では UUID に置き換える）。

---

（このファイルは `main.go` の実装メモ／改善TODOです。必要なら上記の変更を `main.go` に反映するパッチを作成します。）
