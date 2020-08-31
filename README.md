我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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
