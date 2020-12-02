package main

import (
	"flag"
	"github.com/swatkat/gotrntmetainfoparser"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

func main() {
	td := flag.String("td", ".", "Torrents dir")
	dd := flag.String("dd", ".", "Data dir")
	clean := flag.Bool("delete", false, "Delete missing files")
	flag.Parse()
	files, err := ioutil.ReadDir(*dd)
	if err != nil {
		log.Panic(err)
	}
	torrents, err := ioutil.ReadDir(*td)
	if err != nil {
		log.Panic(err)
	}
	mi := gotrntmetainfoparser.MetaInfo{}
	tf := make(map[string]interface{})
	for _, t := range torrents {
		if strings.HasSuffix(t.Name(), ".torrent") {
			if (! mi.ReadTorrentMetaInfoFile(path.Join(*td, t.Name()))) {
				log.Printf("Error reading %s", t.Name())
			} else {
				tf[strings.Trim(mi.Info.Name, " \t")] = 1
				log.Printf("T: %s\n", mi.Info.Name)
			}
		}
	}

	for _, f := range files {
		if _, ok := tf[f.Name()]; ! ok {
			if (*clean) {
				log.Printf("Removing %s\n", f.Name())
			} else {
				log.Printf("Missing %s\n", f.Name())
			}
		}
	}



}
