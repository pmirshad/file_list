// Copyright 2013 Irshad Pananilath. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.

// Package file_list provides a function to list the files and/or
// directories under a root directory

package file_list

import (
  "os"
  "path/filepath"
  "sort"
  "time"
)

type FileInfo struct {
  AbsPath string      // absolute path for the file
  Name    string      // base name of the file
  Size    int64       // length in bytes for regular files
  Mode    os.FileMode // file mode bits
  ModTime time.Time   // modification time
  IsDir   bool        // abbreviation for Mode().IsDir()
}

// byModTime implements sort.Interface.
type byModTime []FileInfo

func (f byModTime) Len() int           { return len(f) }
func (f byModTime) Less(i, j int) bool { return f[i].ModTime.Before(f[j].ModTime) }
func (f byModTime) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// byName implements sort.Interface.
type byName []FileInfo

func (f byName) Len() int           { return len(f) }
func (f byName) Less(i, j int) bool { return f[i].Name < f[j].Name }
func (f byName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// bySize implements sort.Interface.
type bySize []FileInfo

func (f bySize) Len() int           { return len(f) }
func (f bySize) Less(i, j int) bool { return f[i].Size < f[j].Size }
func (f bySize) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// AllFiles enumerates a list of files (and directories) under the directory
// specified as argument. Directories are included if include_dirs is set
// to true. The function returns a slice of type FileInfo
func AllFiles(root string, include_dirs bool) (result []FileInfo) {
  files := []FileInfo{}

  absRoot, err := filepath.Abs(root)

  if err != nil {
    return
  }

  walkFunc := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if include_dirs || !info.IsDir() {
      files = append(files, FileInfo{
        path,
        info.Name(),
        info.Size(),
        info.Mode(),
        info.ModTime(),
        info.IsDir()})
    }
    return nil
  }

  if err := filepath.Walk(absRoot, walkFunc); err != nil {
    return
  }

  return files
}

// Sort the slice of FileInfos by Name reverse sorting if the reverse parameter
// is set
func SortByName(files *[]FileInfo, reverse bool) {
  if reverse {
    sort.Sort(sort.Reverse(byName(*files)))
  } else {
    sort.Sort(byName(*files))
  }
}

// Sort the slice of FileInfos by ModTime reverse sorting if the reverse parameter
// is set
func SortByModTime(files *[]FileInfo, reverse bool) {
  if reverse {
    sort.Sort(sort.Reverse(byModTime(*files)))
  } else {
    sort.Sort(byModTime(*files))
  }
}

// Sort the slice of FileInfos by Size reverse sorting if the reverse parameter
// is set
func SortBySize(files *[]FileInfo, reverse bool) {
  if reverse {
    sort.Sort(sort.Reverse(bySize(*files)))
  } else {
    sort.Sort(bySize(*files))
  }
}
