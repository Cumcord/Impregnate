package middle

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type BrowserLocation struct {
	Location string
	Dir      bool
	Drive    string
}

const BrowserVFSPathDefault = "computer://"

func BrowserVFSLocationReal(vfsPath string) bool {
	return vfsPath != BrowserVFSPathDefault
}

func BrowserVFSList(vfsPath string) []BrowserLocation {
	vfsEntries := []BrowserLocation{}

	if vfsPath == BrowserVFSPathDefault {
		if runtime.GOOS == "windows" {
			for i := 0; i < 26; i++ {
				drive := string(rune('A'+i)) + ":\\"
				_, err := ioutil.ReadDir(drive)
				if err == nil {
					vfsEntries = append(vfsEntries, BrowserLocation{
						Location: drive,
						Dir:      true,
						Drive:    drive,
					})
				}
			}
			return vfsEntries
		}
		home := os.Getenv("HOME")
		return []BrowserLocation{
			{
				Drive:    "Home",
				Location: home,
				Dir:      true,
			},
			{
				Drive:    "Root",
				Location: "/",
				Dir:      true,
			},
			{
				Drive:    "PWD",
				Location: ".",
				Dir:      true,
			},
		}
	}
	fileInfos, err := ioutil.ReadDir(vfsPath)
	if err == nil {
		for _, fi := range fileInfos {
			vfsEntries = append(vfsEntries, BrowserLocation{
				Location: filepath.Join(vfsPath, fi.Name()),
				Dir:      fi.IsDir(),
			})
		}
	}

	return vfsEntries
}
