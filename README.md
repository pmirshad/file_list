Package file_list
=================

Generate a list of files / directories in a directory in Go

Import the package as

    import "github.com/pmirshad/file_list"

Invoke the function as

    files := file_list.AllFiles("/home/pmirshad", true)

AllFiles returns a slice of type file_list.FileInfo

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
