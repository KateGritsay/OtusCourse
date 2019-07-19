//Реализовать утилиту envdir на Go.
//
//Эта утилита позволяет запускать программы получая переменные окружения из определенной директории.
//Пример использования:
//```
//go-envdir /path/to/env/dir some_prog
//```
//Если в директории /path/to/env/dir содержатся файлы
//* `A_ENV` с содержимым `123`
//* `B_VAR` с содержимым `another_val`
//То программа `some_prog` должать быть запущена с переменными окружения `A_ENV=123 B_VAR=another_val`

package envdir

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Envdir(in io.Reader, out io.Writer, errWriten io.Writer, args []string) error {
	if len(args) != 2{
		return errors.New("It has to have two arguments: fileName and path")
	}
	path := args[0]
	progName := args[1]

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	cmd := exec.Command(progName)
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = errWriten

	err = cmd.Run()
	if err != nil {
		fmt.Errorf("Command finished with error: %v", err)
	}

	return nil
}

func envir (path string, files []os.FileInfo) []string{

	envir := make ([]string,len(files))

	for _, file := range files {
		if file.IsDir(){
			continue
		}

	fileName := filepath.Join(path, file.Name())
	content, err := ioutil.ReadFile(fileName)
		if err != nil {
			continue
		}

		var b strings.Builder
		b.WriteString(file.Name())
		b.WriteString(string(content))

	envir = append(envir,b.String())

	}

	return envir

}

func main (){

	err := Envdir(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}




}


