package remapfs

import (
	"io/fs"
	"strings"
)

type RemapFS struct {
	tbl map[string]fs.FS
}

func MustSub(fsys fs.FS, sub string) fs.FS {
	subfs, err := fs.Sub(fsys, sub)
	if err != nil {
		panic("remapfs: " + err.Error())
	}

	return subfs
}

type FSMap map[string]fs.FS

func MustNew(fsmap FSMap) *RemapFS {
	fsys, err := New(fsmap)
	if err != nil {
		panic("remapfs: " + err.Error())
	}

	return fsys
}

func New(fs_tbl FSMap) (*RemapFS, error) {
	for k, v := range fs_tbl {
		if v == nil {
			return nil, &fs.PathError{Op: "new", Path: k, Err: fs.ErrInvalid}
		}
		if !fs.ValidPath(k) {
			return nil, &fs.PathError{Op: "new", Path: k, Err: fs.ErrInvalid}
		}
	}

	return &RemapFS{tbl: fs_tbl}, nil
}

func (rmfs *RemapFS) find(name string) (fs.FS, string) {
	mfs := fs.FS(nil)
	rt := name
	sb := "."

	for {
		var ok bool
		mfs, ok = rmfs.tbl[rt]
		if ok {
			break
		}

		if rt == "." {
			return nil, name
		}

		si := strings.LastIndex(rt, "/")
		switch {
		case si < 0:
			rt = "."
			sb = name
		case si == 0 || si >= len(name)-1:
			panic("invalid fs.FS name")
		default:
			rt = name[:si]
			sb = name[si+1:]
		}
	}

	return mfs, sb
}

func (rmfs *RemapFS) Open(name string) (fs.File, error) {
	subfs, sub := rmfs.find(name)
	if subfs == nil {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
	}
	return subfs.Open(sub)
}

func (rmfs *RemapFS) Stat(name string) (fs.FileInfo, error) {
	subfs, sub := rmfs.find(name)
	if subfs == nil {
		return nil, &fs.PathError{Op: "stat", Path: name, Err: fs.ErrNotExist}
	}
	return fs.Stat(subfs, sub)
}

func (rmfs *RemapFS) ReadFile(name string) ([]byte, error) {
	subfs, sub := rmfs.find(name)
	if subfs == nil {
		return nil, &fs.PathError{Op: "readfile", Path: name, Err: fs.ErrNotExist}
	}
	return fs.ReadFile(subfs, sub)
}

func (rmfs *RemapFS) ReadDir(name string) ([]fs.DirEntry, error) {
	subfs, sub := rmfs.find(name)
	if subfs == nil {
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: fs.ErrNotExist}
	}
	return fs.ReadDir(subfs, sub)
}
