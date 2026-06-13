# go-web-textbook-examples

## 書籍『現場で使えるGo言語Webアプリ開発: Gin・認証・Cloud Runまで実装で学ぶ本番設計』 — サンプルコード
[書籍『現場で使えるGo言語Webアプリ開発: Gin・認証・Cloud Runまで実装で学ぶ本番設計](https://github.com/forest6511/go-web-textbook-examples)

## airによるホットリロード開発ループ
- [air](https://github.com/air-verse/air)
  - go install github.com/air-verse/air@latest
## その他ライブラリ
- [uuid](https://github.com/google/uuid)
  - go get github.com/google/uuid
- [CORS:gin-contrib/cors](https://github.com/gin-contrib/cors)
  - go get github.com/gin-contrib/cors
- [rate](https://pkg.go.dev/golang.org/x/time/rate)
  - go get golang.org/x/time/rate
- [gzip](https://github.com/gin-contrib/gzip)
  - go get github.com/gin-contrib/gzip
- [pgx](https://github.com/jackc/pgx)
  - go get github.com/jackc/pgx/v5
- [golang-migrate](https://github.com/golang-migrate/migrate)
  - go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1
  - (例)
    - migrate create -ext sql -dir internal/db/migrations -seq create_users
    - migrate create -ext sql -dir internal/db/migrations -seq create_tasks
    - migrate create -ext sql -dir internal/db/migrations -seq create_audits
    - migrate create -ext sql -dir internal/db/migrations -seq seed_dev_user
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
  - go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0
  - go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## Ch04-postgres-sqlc
- Postgresql(Docker)の起動
  - docker compose up -d --wait
- Postgresql(Docker)の接続
  - psql postgresql://app:app@localhost:5432/app
- migrate -path internal/db/migrations -database 'postgres://app:app@localhost:5432/app?sslmode=disable' up
  - 実行結果
    - 1/u create_users (58.79917ms)
    - 2/u create_tasks (126.458135ms)