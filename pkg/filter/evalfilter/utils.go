package	evalfilter


import (
	"io/fs"
	"embed"
	"path/filepath"
	"strings"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
)

func init() {
	registerFileScripts()
}


func filterFiles(efs embed.FS, root, ext string) []string {
	var a []string
	fs.WalkDir(efs, root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}

		return nil
	})
	return a
}

func trimFileExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func registerFileScripts() {
	for _, f := range filterFiles(embededScripts, "scripts", ".ef") {
		//fmt.Printf("Registering %s\n", trimFileExtension(filepath.Base(f)))
		filter.RegisterFilterImpl(trimFileExtension(filepath.Base(f)), EvalFilterFactory{FilterTypeEmbeded, f})
	}
}
