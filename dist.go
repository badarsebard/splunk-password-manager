// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/vugu/vugu/distutil"
)

func main() {
	// chmod splunk source
	cmd := exec.Command("/bin/sh", "-c", "sudo chown -R jrauen: dist/password_manager")
	cmd.Output()
	cmd = exec.Command("/bin/sh", "-c", "sudo chmod -R 777 src/splunk/*")
	cmd.Output()
	cmd = exec.Command("/bin/sh", "-c", "sudo chmod -R 777 dist/password_manager")
	cmd.Output()

	clean := flag.Bool("clean", false, "Remove dist dir before starting")
	dist := flag.String("dist", "dist", "Directory to put distribution files in")
	flag.Parse()

	start := time.Now()

	if *clean {
		os.RemoveAll(*dist)
	}

	os.MkdirAll(*dist, 0755) // create dist dir if not there

	// copy splunk app
	distutil.MustCopyDirFiltered("src/splunk", *dist, regexp.MustCompile(`[.]*`))

	// find and copy wasm_exec.js
	// distutil.MustCopyFile(distutil.MustWasmExecJsPath(), filepath.Join(*dist, "password_manager.js"))

	// check for vugugen and go get if not there
	if _, err := exec.LookPath("vugugen"); err != nil {
		fmt.Print(distutil.MustExec("go", "get", "github.com/vugu/vugu/cmd/vugugen"))
	}

	// run go generate
	fmt.Print(distutil.MustExec("go", "generate", "."))

	// run go build for wasm binary
	fmt.Print(distutil.MustEnvExec([]string{"GOOS=js", "GOARCH=wasm"}, "go", "build", "-o", filepath.Join(*dist, "password_manager/appserver/static/password_manager.wasm"), "./src/go"))

	// STATIC INDEX FILE:
	// if you are hosting with a static file server or CDN, you can write out the default index.html from simplehttp
	// req, _ := http.NewRequest("GET", "/index.html", nil)
	// outf, err := os.OpenFile(filepath.Join(*dist, "index.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	// distutil.Must(err)
	// defer outf.Close()
	// template.Must(template.New("_page_").Parse(simplehttp.DefaultPageTemplateSource)).Execute(outf, map[string]interface{}{"Request": req})

	// BUILD GO SERVER:
	// or if you are deploying a Go server (yay!) you can build that binary here
	// fmt.Print(distutil.MustExec("go", "build", "-o", filepath.Join(*dist, "server"), "."))

	log.Printf("dist.go complete in %v", time.Since(start))
}
