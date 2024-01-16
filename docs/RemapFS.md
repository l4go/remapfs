
## type FSMap map[string]fs.FS

fs.FS形式のパスと、fs.FSをマッピングするmapです。

## type RemapFS

複数のfs.FSを1つのfs.FSにする機能を提供します。

fs.FSの以下の機能に対応しています。

- [fs.FS](https://pkg.go.dev/io/fs#FS)
- [fs.StatFS](https://pkg.go.dev/io/fs#StatFS)
- [fs.ReadDirFS](https://pkg.go.dev/io/fs#ReadDirFS)
- [fs.ReadFileFS](https://pkg.go.dev/io/fs#ReadFileFS)
- [fs.Glob()](https://pkg.go.dev/io/fs#Glob)

### `func New(fs_tbl FSMap) (*RemapFS, error)`

`FSMap`型のmapをつかって、`RemapFS`を生成します。  
指定されたfs.FSのパスが間違っているとerrorになります。

### `RemapFS`の作成サンプル

以下のようなコードで`RemapFS`が生成できます。

```go
mfs, err := remapfs.New(remapfs.FSMap{
    ".":       rootFS,
    "foo":     fooFS,
    "foo/bar": foobarFS,
})
if err != nil {
    return
}
```

### `func MustNew(fsmap FSMap) *RemapFS`

`FSMap`型のmapをつかって、`RemapFS`を生成します。  
指定されたfs.FSのパスが間違っているとpanicします。

グローバル変数の初期化などに使います。

以下のようなコードで`RemapFS`が生成できます。

```go
var virtualFS = remapfs.MustNew(remapfs.FSMap{
    ".":       rootFS,
    "foo":     fooFS,
    "foo/bar": foobarFS,
})
```

### `func MustSub(fsys fs.FS, sub string) fs.FS`

[fs.FSのSub()](https://pkg.go.dev/io/fs#SubFS)のラッパー関数です。

`Sub(dir string) (FS, error)`がエラーを返すと、panicします。  

グローバル変数の初期化などに使います。

