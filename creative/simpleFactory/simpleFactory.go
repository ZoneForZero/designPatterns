package simpleFactory

type IDownload interface {
	Download()
}

type SmbDownloader struct{}

func (s *SmbDownloader) Download() {
	println("smb download")
}

type NfsDownloader struct{}

func (n *NfsDownloader) Download() {
	println("nfs download")
}

func NewDownloader(t string) IDownload {
	switch t {
	case "smb":
		return &SmbDownloader{}
	case "nfc":
		return &NfsDownloader{}
	}
	return nil
}

func init() {
	//测试：根据协议类型，创建不同类型的下载器
	smbDownloader := NewDownloader("smb")
	smbDownloader.Download()

	nfsDownloader := NewDownloader("nfc")
	nfsDownloader.Download()
}
