Package file_list
=================

Generate a list of files / directories in a directory in Go

Import the package as

    import "github.com/pmirshad/file_list"

Invoke the function as

    files := file_list.AllFiles("/home/pmirshad", true)

AllFiles returns a slice of type file_list.FileInfo

Sort Functions
--------------

The first argument to the sort functions is a reference to the slice to be sorted. The second argument is a `bool` which decides whether the list is sorted in reverse order.

1. By **modification time**

        file_list.SortByModTime(&files, false)

2. By **name**

        file_list.SortByName(&files, false)

3. By **size**

        file_list.SortBySize(&files, false)

type FileInfo
-------------

    type FileInfo struct {
      AbsPath string      // absolute path for the file
      Name    string      // base name of the file
      Size    int64       // length in bytes for regular files
      Mode    os.FileMode // file mode bits
      ModTime time.Time   // modification time
      IsDir   bool        // abbreviation for Mode().IsDir()
    }
