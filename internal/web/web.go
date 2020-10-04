package web

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/markbates/pkger"
	"github.com/ski7777/MusicNotesManager/internal/mnm"
	"io"
	"mime"
	"net/http"
	"os"
	pathlib "path"
	"strings"
)

const (
	PKG    = "github.com/ski7777/MusicNotesManager"
	HTDOCS = "/internal/web/htdocs"
	PREFIX = PKG + ":" + HTDOCS
)

type web struct {
	echo *echo.Echo
	mnm  *mnm.MusicNotesManager
}

func (w *web) Start(addr string) error {
	return w.echo.Start(addr)
}

func NewWeb(mnm *mnm.MusicNotesManager) (*web, error) {
	w := new(web)
	w.mnm = mnm
	w.echo = echo.New()
	if err := pkger.Walk(HTDOCS, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		w.echo.GET(strings.TrimPrefix(path, PREFIX), func(c echo.Context) error {
			file, err := pkger.Open(path)
			if err != nil {
				return err
			}
			var data bytes.Buffer
			_, err = io.Copy(&data, file)
			if err != nil {
				return err
			}
			return c.Blob(http.StatusOK, mime.TypeByExtension(pathlib.Ext(path)), data.Bytes())
		})
		return nil
	}); err != nil {
		return nil, err
	}
	return w, nil
}
