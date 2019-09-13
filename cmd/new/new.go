package new

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


const (
	githubSource   = "https://github.com/wanhello/iris-admin.git"
	//giteeSource    = "https://gitee.com/lyric/iris-admin.git"
	defaultPkgName = "github.com/wanhello/iris-admin"
)


// Config 配置参数
type Config struct {
	Dir       string
	PkgName   string
	UseMirror bool
}


// Command 创建项目命令
type Command struct {
	cfg *Config
}


func Exec(cfg Config) error {
	cmd := &Command{cfg: &cfg}
	return cmd.Exec()
}


// Exec 执行命令
func (a *Command) Exec() error  {
	dir, err := filepath.Abs(a.cfg.Dir)
	if err != nil {
		return err
	}
	log.Printf("generating directory：%s", dir)

	isClone := false
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = a.gitClone(dir)
			if err != nil {
				return err
			}
			isClone = true
		}
	}

	if pkgName := a.cfg.PkgName; pkgName != "" {
		err := a.changePkgName(dir, a.cfg.PkgName)
		if err != nil {
			return err
		}

		err = a.readAndReplaceFile(a.cfg.PkgName, fmt.Sprintf("%s/%s", dir, "go.mod"))
		if err != nil {
			return err
		}
	}

	if isClone {
		err = a.gitInit(dir)
		if err != nil {
			return err
		}
	}

	fmt.Printf("\nproject created：%s\n", dir)
	fmt.Println(TplProjectStructure)

	return nil
}

func (a *Command) execGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	return cmd.Run()
}

func (a *Command) gitClone(dir string) error {
	var args [] string
	args = append(args, "clone")
	args = append(args, "-q")
	args = append(args, "-b", "master")

	source := githubSource
	args = append(args, source)
	args = append(args, dir)

	log.Printf("Executing: git %s", strings.Join(args, " "))

	return a.execGit("", args...)
}

func (a *Command) gitInit(dir string) error {
	os.RemoveAll(filepath.Join(dir, ".git"))

	err := a.execGit(dir, "init")
	if err != nil {
		return err
	}

	err = a.execGit(dir, "add", "-A")
	if err != nil {
		return err
	}

	err = a.execGit(dir, "commit", "-m", "'Initial commit'")
	if err != nil {
		return err
	}

	return nil
}


func (a *Command) checkInDirs(dir, path string) bool {
	includeDirs := []string{"cmd", "internal", "pkg"}
	for _, d := range includeDirs {
		p := fmt.Sprintf("%s/%s", dir, d)
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

func (a *Command) changePkgName(dir, pkgName string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".go" || info.IsDir() || !a.checkInDirs(dir, path) {
			return nil
		}

		return a.readAndReplaceFile(pkgName, path)
	})
}

func (a *Command) readAndReplaceFile(pkgName, name string) error {
	buf, err := a.readAndReplace(pkgName, name)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, buf.Bytes(), 0644)
}

func (a *Command) readAndReplace(pkgName, name string) (*bytes.Buffer, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), defaultPkgName, pkgName, 1)
		buf.WriteString(line)
		buf.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buf, nil
}





