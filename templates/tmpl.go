// Code generated by "esc -include=.*\.tmpl -o=tmpl.go -pkg=templates ./"; DO NOT EDIT.

package templates

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/err/assertion.tmpl": {
		name:    "assertion.tmpl",
		local:   "err/assertion.tmpl",
		size:    1528,
		modtime: 1680667219,
		compressed: `
H4sIAAAAAAAC/7RTTWvjMBC951fMmhySksr3Qm8NSy+b0oTtWU1GXoEsZUdjShH674tkryOH0I9l64ux
R/Pe05s3IRxQaYtQSe+RWDtbwXWMMwCAEK6hvmocvx7xBhrNv7pnsXdtvX9BbUzdOEbPvtaWkaw0desO
aLzY3UmWV/UERysQj8gdWb8mcjSU0qMVMIsXaXlNBN9uwWoDYSynfsb2aCSPMjNEBaJAOVGhPcQ4m13u
/OGK3uL07NyJ4dj/9qIHF+vfnTSZYsErCCHZs7HmdWpRKX8vjcmqQ0DjMUYkCiGrXxX+iR50WRrvCBbW
MYht95xVLkE8kLZ8b48d+1FbelagWhbbYyqrRVUqaNF72WAWUa0m3uoMlCvLcg4jdP+3rmG3udsMJoDX
BwRUCvfs+3lRvv6bcxkH+BUpnSu4uR1jNcxqYEyDQqLT/UjaBkHs0PMj+s6wL3z8m/l7/0Sascz7eTAx
ZaGCuYoxcTCLEJ6k5aQj2fwgSbbpQ2yZtG0Wy8lUfNP3LifcfUbOOLOkuSqTtrH4U5ru/OxnNZZRMNpi
f/QdnW9o/Sz/d8f9bnyMMGfr8p/T1+UU9kpiLBc57+87WxajGtZ1eE1Qk9hU+QgOrOAf9nJK/icAAP//
s0lIJvgFAAA=
`,
	},

	"/err/call.tmpl": {
		name:    "call.tmpl",
		local:   "err/call.tmpl",
		size:    241,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/0SOQWrDQAxFryKMFy0YHaDQA3hTSlvatRjLrsCeFo2SEITuHsY4mdWHP2/el/vEs2SG
LtG6dhHuF7FfwA9OLGfW2sgM+c8Ax/JpekoWYYbunKf6eicBI1qLb7RxxJO7Ul4Yehmg5xVeXgHfSWlj
Yy2HvZeIAR5/296PitUbzJB0KU2/K+riTuPX9Z9xLN+kQpOkCMTG7vF85C0AAP//ZQi8iPEAAAA=
`,
	},

	"/err/function.tmpl": {
		name:    "function.tmpl",
		local:   "err/function.tmpl",
		size:    1748,
		modtime: 1680673385,
		compressed: `
H4sIAAAAAAAC/3xVTU/jMBA9J79iVCHUInDuRRx2BStxgVWplgNarUwyCdE6TmVPFlWW//vKH0mTpoUD
safz8ea9ycSYAstaIizKTuZUt3JhbWrMDWRXVUv7Ha6hqumje2d522T5J9ZCZFVLqElntSRUkousaQsU
mm3vOfGrDG5ijosS1nfArE1Tlx6MYVvU9MQbtHZJcOXS1LJi2xWYNHEhnzV9ANtgjvU/VNamiTfXJbBH
/UKqy8kbB+uPGkWhgy1xgKH0FtDe2eWN3orLCo8CEmP83YH08PY7jD+5EJRFvA01e9PofHR0qFybP7ni
DRIqX8xD46qaABvBmkf4gt40QzeqOK3vdXGkG+NxOKoLaxu+e9Okaln9DtWNQaHR2rfD3SWIoFzoUrYU
41fWSt4ghAzRNaI4oVdPApfFQbQj3qNG4TFQ6yHFxvuUc2XOqPAF90niiXf/TsSMBNig7gTpvs4rl/QV
90PJDVKnpH5Qqo0cfHJJD0oBOtORYo7kLIPt8/3zGr4VBTjVIOcaNfOClq0CY5wIrQL20r0HWYMkj/qJ
/8VitfK4RiI7jXpd/0SVroEojEPQN/Qa0plRA32NSVeORCFQWBuSEN1Oxe9F7h0nUFyEe94a4zoHH0Vs
08nlWdhELF4DeLc15nviJMLhvFxNQQ7v8LlxPbNgZkPoafQDtt+hd+bK2stYLY4I+8VFh9aaPsWZvZMY
w8IeXAMRCy8CG22j60OCwxZKTqym2SXWO7FM+jZfVU1D95Mls76Dy/c9oWbfu7JEZb4sCKM/YwibneCE
sKhaWsBF6Qf0YM25EMGcno3kWqPyn6HoOB9RsBbsCiIK/61JbGrTtJf9fwAAAP//6AOmMdQGAAA=
`,
	},

	"/err/got.tmpl": {
		name:    "got.tmpl",
		local:   "err/got.tmpl",
		size:    206,
		modtime: 1680673359,
		compressed: `
H4sIAAAAAAAC/3SOQQrDQAhFr/IJs0xzgEKXpfveIDBOEIIDzmQl3r0YhpYWuvLr84lmmQoLYdpqn3Bx
h5mushESz0i043rD8qR27L25m3FBYvcZZiQ5Jo/asUQYPReskt9ShH6otLtq1S+Tyy8k1UFx4qp/78Rj
n92zvgIAAP//I/eJPM4AAAA=
`,
	},

	"/err/header.tmpl": {
		name:    "header.tmpl",
		local:   "err/header.tmpl",
		size:    142,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/0TMMQ7CMAyF4d2nsDrBQC7BxIK4gkUebYXiViGb9e6OlAi6/bL1voiM1+rQaYFl1ImU
iGo+Q9N1KwXePmRE6g941gspuz3fNkMj0mMkKbKWfatNT4dw65cB3K2AHJO2/DhSzv/6BgAA///GzMM9
jgAAAA==
`,
	},

	"/err/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "err/inline.tmpl",
		size:    49,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NYCAgAA
//+q60H/MQAAAA==
`,
	},

	"/err/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "err/inputs.tmpl",
		size:    177,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/0yNMaoDMQxE+38KsWz58QECOUCaEMgJFCwvLqwESVsJ3T1YpHAlzWN4416pdSbYOn9O
0y3CfW9wuUKZb2/Ab4PyPF9GarqyOw6qEWbFnbhGMA76h1/I3t7KQzrbLeUTCvJByVFwkJFoKlAOLe5J
5/TiWc/fNwAA//94+RPrsQAAAA==
`,
	},

	"/err/message.tmpl": {
		name:    "message.tmpl",
		local:   "err/message.tmpl",
		size:    201,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/zyN4WqDQBCE//sUiyi0oPsAhT5A/xRpS/9f4mgW9GLuTkNY9t2DB/HXDDPDN6o9BvGg
ckaMbkRJrVmhKgP5ayL+XU8JMUWz+sakCt+bqd4lXYh/cIZsCHvCf48F/O+mFWZ8DPnbzTB7y0Tugvj0
5Zd1B6oG50dQJQ1VmOjjk7hzwc1ICLmXgSoxa16/9XZws7wXqi1l+wwAAP//kC65UskAAAA=
`,
	},

	"/test/call.tmpl": {
		name:    "call.tmpl",
		local:   "test/call.tmpl",
		size:    0,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/function.tmpl": {
		name:    "function.tmpl",
		local:   "test/function.tmpl",
		size:    19,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/6qu1tdSSFRIzs/NTc0rUdDSr+UCBAAA//+6o7WcEwAAAA==
`,
	},

	"/test/header.tmpl": {
		name:    "header.tmpl",
		local:   "test/header.tmpl",
		size:    0,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "test/inline.tmpl",
		size:    0,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "test/inputs.tmpl",
		size:    0,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/message.tmpl": {
		name:    "message.tmpl",
		local:   "test/message.tmpl",
		size:    0,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/test/results.tmpl": {
		name:    "results.tmpl",
		local:   "test/results.tmpl",
		size:    0,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/testify/call.tmpl": {
		name:    "call.tmpl",
		local:   "testify/call.tmpl",
		size:    241,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/0SOQWrDQAxFryKMFy0YHaDQA3hTSlvatRjLrsCeFo2SEITuHsY4mdWHP2/el/vEs2SG
LtG6dhHuF7FfwA9OLGfW2sgM+c8Ax/JpekoWYYbunKf6eicBI1qLb7RxxJO7Ul4Yehmg5xVeXgHfSWlj
Yy2HvZeIAR5/296PitUbzJB0KU2/K+riTuPX9Z9xLN+kQpOkCMTG7vF85C0AAP//ZQi8iPEAAAA=
`,
	},

	"/testify/function.tmpl": {
		name:    "function.tmpl",
		local:   "testify/function.tmpl",
		size:    2929,
		modtime: 1680671749,
		compressed: `
H4sIAAAAAAAC/6xWTW/jNhA9y79iIAQLK/DSdy/20GLTIpfNIgm6h0VRMNZIK5SiXHLUwCD43wsO9WXZ
cjdFc4jFEWfmzbzhE53Lsag0QiqtRUNVo1N47/3KufewvS0bOh5wB2VF39sXsW/q7f4VK6W2ZUNoyW4r
TWi0VNu6yVFZ8fxJkrzdcowYUtz91Up1Z0xjijVtgEi8Sk13xmzAuaoA8aDV8RGpNdryNu+dI6wPShJC
updKpSCCEZVF79EY51DnHcqqgMbAWjcE4ql9YVgZiC+m0nSvDy1ZBgMAsIGiJvF0CK+KdTrNUqO1skRO
lAZg46uKg/CbjDOizjnksPJ+tRo7iaHe1Ptp+Vzov8D0vujq6subRq1tmYY3PxIHNvAfiruSvGj1nmej
6/lNAbuPwWf1vwxKCA/OiWe09FnW6P2a4Db4VboUzxm4VRLyvFb0HcQj7rH6G433qyTpJkDc2ycy7Z7Y
OFh/qVDlNtqSgBAKtoDlzSFut9tIXeLMIXGO16FQhnc8YPdqJD6shpy9afI8ewyoQplfpJE1EhpOxtCk
KU+ATWCde3BCNp2hm2Q8zc9EBOKcYxyh1bn3tTx8s2QqXf4es/cH7du4DgE6UME1jh/7Z95rWSPECP3w
JEt89U2QOh9Jm/W94yj+DK1lSF3hfchzZhZYuNL7JOHGh38XfCYEPKJtFdk+z1ep6Vrvh5Sn0hacO/0D
DKYZY6HJ2y08P3x62MFPeQ6BNdhLi1YwoUVjom6uGzNKQEfJvf0s/8Q8yxjXhOTAUc/rHx1LQYvjOER+
Y60xnJsU0Oc4qSo0USlU3scgRB9Oye9J7jeeQAke4feDcyyn7EXisdXrRdhEoltG8EE1znXiIsLheZ2d
ghzO8NK4LgjM2RByG3nAjgfkzdJ4/27Q1Gj/TaoWvXd9iAXdSZwTUQd34XMZD4KYqNFmDDCqUHJBms4W
Xb4LYtKX+dVUNFR/IjK7j/Du5Uhoxc9tUaBxP5KwG4M4nvMvfXZuf9DIXcpgQDZ+rUw8gincFDzg81tC
MC+hmJ9DmP1No00uQ4shl5ThUivhwt80X7wycK7uijTqS/heDxyIJ5bZMMaTz7sto2828hHV8nJehndT
XOr6gs9b8U7vF6rSGLcuYF4CeaWGt+L5taF4w3kbgOFoLY74XCLBe/AZTG+JiV/xNSm6/BMAAP//Sw1a
k3ELAAA=
`,
	},

	"/testify/header.tmpl": {
		name:    "header.tmpl",
		local:   "testify/header.tmpl",
		size:    142,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/0TMMQ7CMAyF4d2nsDrBQC7BxIK4gkUebYXiViGb9e6OlAi6/bL1voiM1+rQaYFl1ImU
iGo+Q9N1KwXePmRE6g941gspuz3fNkMj0mMkKbKWfatNT4dw65cB3K2AHJO2/DhSzv/6BgAA///GzMM9
jgAAAA==
`,
	},

	"/testify/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "testify/inline.tmpl",
		size:    49,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NYCAgAA
//+q60H/MQAAAA==
`,
	},

	"/testify/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "testify/inputs.tmpl",
		size:    177,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/0yNMaoDMQxE+38KsWz58QECOUCaEMgJFCwvLqwESVsJ3T1YpHAlzWN4416pdSbYOn9O
0y3CfW9wuUKZb2/Ab4PyPF9GarqyOw6qEWbFnbhGMA76h1/I3t7KQzrbLeUTCvJByVFwkJFoKlAOLe5J
5/TiWc/fNwAA//94+RPrsQAAAA==
`,
	},

	"/testify/message.tmpl": {
		name:    "message.tmpl",
		local:   "testify/message.tmpl",
		size:    201,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/zyN4WqDQBCE//sUiyi0oPsAhT5A/xRpS/9f4mgW9GLuTkNY9t2DB/HXDDPDN6o9BvGg
ckaMbkRJrVmhKgP5ayL+XU8JMUWz+sakCt+bqd4lXYh/cIZsCHvCf48F/O+mFWZ8DPnbzTB7y0Tugvj0
5Zd1B6oG50dQJQ1VmOjjk7hzwc1ICLmXgSoxa16/9XZws7wXqi1l+wwAAP//kC65UskAAAA=
`,
	},

	"/testify/results.tmpl": {
		name:    "results.tmpl",
		local:   "testify/results.tmpl",
		size:    168,
		modtime: 1656969383,
		compressed: `
H4sIAAAAAAAC/1yNTQrCQAyFr/Iosyw9gOBS3HsDoRkJlAy8ma5C7i6pRcFVfr4vee6rVDXBROn7NvoU
AXc+7SUoOqPIhssVy+ODI9y1omjEDHexNTf3NrBkc85a82DstH4jG1MW8uQ4hMbv0385A3/uUd8BAAD/
/7BPz2GoAAAA
`,
	},

	"/": {
		name:  "/",
		local: `./`,
		isDir: true,
	},

	"/err": {
		name:  "err",
		local: `err`,
		isDir: true,
	},

	"/test": {
		name:  "test",
		local: `test`,
		isDir: true,
	},

	"/test_empty": {
		name:  "test_empty",
		local: `test_empty`,
		isDir: true,
	},

	"/testify": {
		name:  "testify",
		local: `testify`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"./": {},

	"err": {
		_escData["/err/assertion.tmpl"],
		_escData["/err/call.tmpl"],
		_escData["/err/function.tmpl"],
		_escData["/err/got.tmpl"],
		_escData["/err/header.tmpl"],
		_escData["/err/inline.tmpl"],
		_escData["/err/inputs.tmpl"],
		_escData["/err/message.tmpl"],
	},

	"test": {
		_escData["/test/call.tmpl"],
		_escData["/test/function.tmpl"],
		_escData["/test/header.tmpl"],
		_escData["/test/inline.tmpl"],
		_escData["/test/inputs.tmpl"],
		_escData["/test/message.tmpl"],
		_escData["/test/results.tmpl"],
	},

	"test_empty": {},

	"testify": {
		_escData["/testify/call.tmpl"],
		_escData["/testify/function.tmpl"],
		_escData["/testify/header.tmpl"],
		_escData["/testify/inline.tmpl"],
		_escData["/testify/inputs.tmpl"],
		_escData["/testify/message.tmpl"],
		_escData["/testify/results.tmpl"],
	},
}
