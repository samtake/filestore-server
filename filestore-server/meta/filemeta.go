package meta

//文件信息结构
type FileMeta struct {
	FileSha1 string //唯一标志
	FileName string
	FileSize int64
	Location string
	UploadAt string //上传信息
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//通过fileSha1获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
