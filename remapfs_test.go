package remapfs_test

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/l4go/remapfs"
)

//go:embed test/root
var raw_rootFS embed.FS
var rootFS = remapfs.MustSub(raw_rootFS, "test/root")

//go:embed test/foo
var raw_fooFS embed.FS
var fooFS = remapfs.MustSub(raw_fooFS, "test/foo")

//go:embed test/foobar
var raw_foobarFS embed.FS
var foobarFS = remapfs.MustSub(raw_foobarFS, "test/foobar")

func ExampleMustNew() {
	mfs := remapfs.MustNew(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})

	buf, err := fs.ReadFile(mfs, "foo.txt")
	if err != nil {
		return
	}

	fmt.Print(string(buf))
	// Output:
	// foo
}

func ExampleNew() {
	mfs, err := remapfs.New(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})
	if err != nil {
		return
	}

	buf, err := fs.ReadFile(mfs, "foo/bar.txt")
	if err != nil {
		return
	}

	fmt.Print(string(buf))
	// Output:
	// bar
}

func ExampleRemapFS_Open() {
	var mfs fs.FS
	var err error

	mfs, err = remapfs.New(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})

	_, err = mfs.Open("foo.txt")
	fmt.Println(err == nil)
	_, err = mfs.Open("bar.txt")
	fmt.Println(err == nil)
	_, err = mfs.Open("foo/bar.txt")
	fmt.Println(err == nil)
	_, err = mfs.Open("foo/baz.txt")
	fmt.Println(err == nil)
	_, err = mfs.Open("foo/bar/baz.txt")
	fmt.Println(err == nil)
	// Output:
	// true
	// false
	// true
	// false
	// true
}

func ExampleRemapFS_Stat() {
	var mfs fs.FS
	var err error

	mfs, err = remapfs.New(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})

	_, err = fs.Stat(mfs, "foo.txt")
	fmt.Println(err == nil)
	_, err = fs.Stat(mfs, "bar.txt")
	fmt.Println(err == nil)
	_, err = fs.Stat(mfs, "baz.txt")
	fmt.Println(err == nil)
	_, err = fs.Stat(mfs, "foo/bar.txt")
	fmt.Println(err == nil)
	_, err = fs.Stat(mfs, "foo/baz.txt")
	fmt.Println(err == nil)
	_, err = fs.Stat(mfs, "foo/bar/baz.txt")
	fmt.Println(err == nil)

	// Output:
	// true
	// false
	// false
	// true
	// false
	// true
}

func ExampleRemapFS_ReadFile() {
	var mfs fs.FS
	var err error

	mfs, err = remapfs.New(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})
	if err != nil {
		return
	}

	var buf []byte
	buf, err = fs.ReadFile(mfs, "foo.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Print(string(buf))
	}

	buf, err = fs.ReadFile(mfs, "bar.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Print(string(buf))
	}

	buf, err = fs.ReadFile(mfs, "foo/bar.txt")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Print(string(buf))
	}

	// Output:
	// foo
	// error
	// bar
}

func ExampleRemapFS_ReadDir() {
	var mfs fs.FS
	var err error

	mfs, err = remapfs.New(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})
	if err != nil {
		return
	}

	if ents, er := fs.ReadDir(mfs, "."); er == nil {
		for _, el := range ents {
			fmt.Println(el.Name())
		}
	}
	if ents, er := fs.ReadDir(mfs, "foo/bar"); er == nil {
		for _, el := range ents {
			fmt.Println(el.Name())
		}
	}

	// Unordered output:
	// foo.txt
	// hoge.txt
	// baz.txt
}

func ExampleGlob() {
	var mfs fs.FS
	var err error

	mfs, err = remapfs.New(remapfs.FSMap{
		".":       rootFS,
		"foo":     fooFS,
		"foo/bar": foobarFS,
	})
	if err != nil {
		return
	}

	if lst, er := fs.Glob(mfs, "*.txt"); er == nil {
		for _, name := range lst {
			fmt.Println(name)
		}
	}
	if lst, er := fs.Glob(mfs, "foo/*.txt"); er == nil {
		for _, name := range lst {
			fmt.Println(name)
		}
	}
	if lst, er := fs.Glob(mfs, "foo/bar/*.txt"); er == nil {
		for _, name := range lst {
			fmt.Println(name)
		}
	}

	// Unordered output:
	// foo.txt
	// hoge.txt
	// foo/bar.txt
	// foo/bar/baz.txt
}
