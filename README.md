# go-graphql-server-template

```
github.com/99designs/gqlgen/graphql
github.com/99designs/gqlgen/graphql/handler
github.com/99designs/gqlgen@v0.17.49
gorm.io/driver/sqlite
```

1. 建 domain
2. 建 repo
3. 建 graphql 檔設定 input 和 model
4. 將 query / mutation 加到 schema.graphql
5. 執行
   ```powershell=
   go get github.com/99designs/gqlgen@v0.17.49; go run github.com/99designs/gqlgen
   ```
   建立 `/internal/autogen`
6. 前往 `/internal/resolvers` ，將 autogen 產生的 method 補齊業務邏輯
7. 前往 http://localhost:8080/ 執行 graphql
