package main

import(
"os"
"log"
"os/exec"
"fmt"
"path"
)

func main() {
	file, _ := os.Getwd()
//	log.Println("current path:", file)
//	fmt.Println(os.Args[0])
//	file, _ = exec.LookPath(os.Args[0])
	abs :=fmt.Sprintf(`%s/metronic/templates/admin/extra_profile.html`,file)
	fmt.Println(abs)
	file, _ = exec.LookPath(fmt.Sprintf(`%smetronic/templates/admin/extra_profile.html`,file))
	fmt.Printf("%T",os.Args[0])
	fmt.Println(file)
	log.Println("exec path:", file)
	filename := path.Dir(abs)
	os.MkdirAll(filename,0777)
	dir,_ := path.Split(file)
	log.Println("exec folder relative path:", dir)
	os.OpenFile(file,os.O_CREATE,0777)
	os.Chdir(dir)
	wd, _ := os.Getwd()
	log.Println("exec folder absolute path:", wd)
}
