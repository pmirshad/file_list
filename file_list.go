// Copyright 2013 Irshad Pananilath. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.

// Package file_list provides a function to list the files and/or
// directories under a root directory

package file_list

import (
  "os"
  "path/filepath"
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
