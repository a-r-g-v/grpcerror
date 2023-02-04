grpcerror
===

grpcerror は gRPC駆動アプリケーションのためのエラーハンドリングユーティリティです。
* 依存gRPCサービスから返却されたエラー翻訳を行うための Translate関数と gRPC Interceptorを提供します
* `status.New` や `errdetails` の設定のためのヘルパ関数を提供します

# エラー翻訳
gRPCアプリケーション向けに、別のgRPCサービスが生成したエラーの翻訳ユーティリティを提供します。これにより、あなたのアプリケーションでは依存gRPCサービスのエラーを起因とする誤ったエラーハンドリングを防止することができます。

## 問題意識

あるgRPC APIハンドラが、別のgRPCサービスを呼び出した結果エラーを得たとき、API実行の継続可否を参考にした上で、エラーハンドリングを行う必要があります。
* API実行を継続できるエラーであれば、エラーを対処して実行を継続する。
* API実行を継続できないエラーであれば、エラーを適切なエラーに翻訳した後に、クライアントに返却する。

このうち、後者であればエラー翻訳を行う必要があります。


エラー翻訳の方法は状況と文脈に依存してケースバイケースとなります。 APIが返却するべきエラーは、以下のような考え方によって決めることが多いと考えます。

* それがコンテキストやタイムアウトに起因するエラーであれば、codes.Canceled / codes.DeadlineExceeded に翻訳する
* 呼び出し側クライアントに起因するエラーであれば codes.InvalidArgument, NotFound, FailedPrecondition などに翻訳する
* サーバー実装に起因するエラーや、依存gRPCサービスの実装に起因するエラーであれば codes.Internal / codes.Unknown に翻訳する

これを踏まえて、依存gRPCサービスが返却したエラーのハンドリング方法は以下のようになることが多いと考えます。

| 依存gRPCサービスが生成したエラーコード | APIが返却するべきエラーコード    | 理由
|:----------------------|:--------------------| :--- | 
| Canceled              | Canceled            | クライアントがキャンセルした場合はAPIはキャンセルコードを返すべき（クライアントに責任がある）
| DeadlineExceeded      | DeadlineExceeded    | クライアントが設定したタイマーにタイムアウトした場合はタイムアウトコードを返すべき（クライアントに責任がある）サーバーが設定したタイマーである場合はこの限りではない
| Unknown               | Unknown or Internal | 依存gRPCサービス実装に起因するエラーであり、サーバーがそのサービスに依存している以上、サーバーの責任である
| InvalidArgument       | Internal | サーバーが依存gRPCのリクエスト契約に違反しているため、サーバーの責任である
| それ以外                  | Internal | 基本的には同上。

もちろん、状況と文脈次第で、例外ケースも多々あります。例えば、NotFound エラーを例にして、クライアントが外部gRPCサービスが保有しているリソースのIDを指定し、APIサーバーにリクエストを行う場合を考えます。
このような場合、APIサーバーが外部gRPCサービスにリクエストした結果 NotFound エラーとなるのであれば、APIサーバーはクライアントに NotFound エラーを返却するべきでしょう。

これらを踏まえると、特別なケースではエラー翻訳を都度行うべきであるが、デフォルトのエラー翻訳ルールを適用できる場合はそれを利用するべきであると考えます。
このユーティリティは、このポリシをアプリケーションに簡単に実装できるヘルパーを提供しています。

## 利用方法

デフォルトのエラー翻訳は、gRPC Server Interceptorとして設定します。

```go
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.TranslateUnaryServerInterceptor(interceptor.DefaultTranslator()),
		),
	)
```

`interceptor.DefaultTranslator()` は 上記表に規定されたルールに従って翻訳を行います。必要に応じて、`interceptor.MapTranslator` を利用することで挙動をカスタマイズできます。

都度の外部gRPCサービス起因のエラー対処は、以下のようになります。

```go
 resp, err := mayOccurGRPCError()
 if err != nil {
    return nil, err // デフォルトのエラー翻訳ポリシーを適用したい場合は、何もせず return する
 }

 resp, err := mayOccurGRPCError()
 if err != nil {
    return nil, fmt.Errorf("xxx.yyy  failed: %w", err) // Wrapしても、デフォルトのエラー翻訳ポリシーが採用される
 }

 resp, err := mayOccurGRPCError()
 if err != nil {
    if codes.Code(err) == codes.NotFound {
      return nil, grpcerror.Translate(err, grpcerror.NotFound("not found")) // NotFound に明示的にエラー翻訳することで、デフォルトのエラー翻訳ポリシの適用を回避する
    }
	return nil, err
 }
```


# ヘルパ関数
gRPCエラーを生成するためのヘルパ関数を提供します。 これにより、あなたのアプリケーションに存在するgRPCエラーを返却する場合のボイラープレート（or あなたのプロジェクトに存在するヘルパー）を削減することができます。

## 単純な例

```go
 resp, err := mayOccurGRPCError()
 if err != nil {
    return nil, status.Error(codes.AlreadyExists, "already exists")
 }
```

ヘルパ関数を用いることで、上記のようなコードを、以下のように置き換えできます。

```go
 resp, err := mayOccurGRPCError()
 if err != nil {
    return nil, grpcerror.AlreadyExistsError("already exists")
 }
```

## エラー詳細の例

```go
resp, err := mayOccurGRPCError()
if err != nil {
    serr := status.New(codes.InvalidArgument, "invalid argument")
    detail, err := serr.WithDetails(&errdetails.BadRequest{
        FieldViolations: []*errdetails.BadRequest_FieldViolation{
            {
                Field:       "name",
                Description: "len(name) must be [1, 32]",
            },
        },
    })
    if err != nil {
        return nil, status.Error(serr.Code(), fmt.Sprintf("%s: with defailt failed: %v", serr.Err(), err))
    }
    return nil, detail.Err()
}
```

このようなコードを以下のように書き換えることができます。

```go
resp, err := mayOccurGRPCError()
if err != nil {
    return nil, grpcerror.InvalidArgumentError("invalid argument", grpcerror.BadRequest(
        &errdetails.BadRequest{
            FieldViolations: []*errdetails.BadRequest_FieldViolation{
                {
                    Field:       "name",
                    Description: "len(name) must be [1, 32]",
                },
            },
        }))
}
```