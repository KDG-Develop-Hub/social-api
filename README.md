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
| internal/app        | アプリケーションを実行させるための関数 |
| internal/controller | ハンドラー関数             |
| internal/domain     | ドメインモデル             |
| internal/repository | データベースの接続関数         |
| internal/service    | ビジネスロジック            |
| tmp                 | テンポラリディレクトリ         |
