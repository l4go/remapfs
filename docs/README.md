# l4go/remapfs ライブラリ

使われるfs.FSをパスで振り分ける、仮想的なfs.FSを提供します。　

httpモジュールでは、以下のようなコードで、パスでhttpハンドラを呼び分けることができます。
``` go
http.Handle("/"、rootHandler)
http.Handle("/foo"、fooHandler)
http.Handle("/foo/bar"、barHandler)
http.Handle("/hoge"、hogeHandler)
```

この処理のようなパスでの処理の振り分けを、fs.FS上でも実現する為のモジュールです。

以下のようなコードで、使われるfs.FSをパスで振り分ける、仮想的なfs.FSを作成できます。

``` go
mfs, err := remapfs.New(remapfs.FSMap{
    ".":       rootFS,
    "foo":     fooFS,
    "foo/bar": foobarFS,
    "hoge":    hogeFS,
})
```

エラーが起きないと分かってる場合は、以下のように書くこともできます。 
``` go
var virtualFS = remapfs.MustNew(remapfs.FSMap{
    ".":       rootFS,
    "foo":     fooFS,
    "foo/bar": foobarFS,
    "hoge":    hogeFS,
})
```

## 詳細仕様

* [remapfs.RemapFS](RemapFS.md)
  * mapによる指定で、複数のfs.FSを1つのfs.FSにする機能を提供します。

