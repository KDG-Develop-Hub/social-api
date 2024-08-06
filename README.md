# kdg socialのバックエンド

## セットアップ

```shell
# パッケージと必要なコマンドをインストール
make install

# Makefileを実行
make run
```

## ディレクトリ構成

| ディレクトリ名             | 説明                  |
|---------------------|---------------------|
| cmd                 | アプリケーションのエントリーポイント  |
| configs             | 設定ファイル              |
| internal            | アプリケーションのコアロジック     |
| pkg                 | プロジェクトに依存しないパッケージ   |
| internal/app        | アプリケーションを実行させるための実装 |
| internal/controllers | ハンドラー関数             |
| internal/enitities     | エンティティモデル             |
| internal/repositories | データベースの接続         |
| internal/services    | ビジネスロジック            |
| tmp                 | テンポラリディレクトリ         |
