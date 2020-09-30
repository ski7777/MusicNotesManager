package main

import (
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

const (
	migrationpath  = "internal/database/migration"
	migrationspath = migrationpath + "/migrations"
	outputpath     = migrationpath + "/migrations.go"
	outputpkg      = "migration"
	outputvarname  = "migrations"
	outputvartype  = "[]*gormigrate.Migration"
	funcmigrate    = "Migrate"
	funcrollback   = "Rollback"
	importprefix   = "github.com/ski7777/MusicNotesManager/" + migrationspath
	pkgprefix      = "migration"
)

var (
	imports = []string{
		"github.com/go-gormigrate/gormigrate/v2",
	}
)

func main() {
	var output, importdata, vardata []string

	vardata = append(output, "var "+outputvarname+"="+outputvartype+"{")

	files, err := ioutil.ReadDir(os.ExpandEnv("$PWD/" + migrationspath))
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if f.IsDir() {
			vardata = append(vardata, "{")
			vardata = append(vardata, "ID:\""+f.Name()+"\",")
			vardata = append(vardata, "Migrate:"+pkgprefix+f.Name()+"."+funcmigrate+",")
			vardata = append(vardata, "Rollback:"+pkgprefix+f.Name()+"."+funcrollback+",")
			vardata = append(vardata, "},")
			imports = append(imports, importprefix+"/"+f.Name())
		}
	}
	vardata = append(vardata, "}")

	importdata = append(importdata, "import(")
	for _, i := range imports {
		importdata = append(importdata, "\""+i+"\"")
	}
	importdata = append(importdata, ")")

	output = append(output, "package "+outputpkg)
	output = append(output, importdata...)
	output = append(output, vardata...)

	code, err := format.Source([]byte(strings.Join(output, "\n")))
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(os.ExpandEnv("$PWD/"+outputpath), code, 0644)
}
