// +-------------------------------------------------------------------------
// | Copyright (C) 2017 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

// +build ignore

// download snips binrary.
//
// Examples:
//
//	go run get-snips.go
//	go run get-snips.go -os=windows -output=snips.exe
//
package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

var (
	flagOS       = flag.String("os", runtime.GOOS, "Specify the os (windows|darwin|linux)")
	flagFileName = flag.String("output", "snips-v0.0.9.exe", "Specify the output filename")
)

func main() {
	flag.Parse()

	var snipsURL string
	switch *flagOS {
	case "windows", "darwin", "linux":
		const baseURL = "https://github.com/yunify/snips/releases/"
		snipsURL = fmt.Sprintf("%s/download/v0.0.9/snips-v0.0.9-%s_amd64.tar.gz", baseURL, *flagOS)
	default:
		log.Fatalf("invalid os: %v", *flagOS)
	}

	data, err := fetchTarGzFile(snipsURL)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(*flagFileName, data, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("download snips-v0.0.9 ok")
}

func fetchTarGzFile(downloadURL string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(downloadURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return readTarGz(resp.Body)
}

func readTarGz(r io.Reader) ([]byte, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	// tar read
	tr := tar.NewReader(gr)
	for {
		if _, err := tr.Next(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, tr); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	return nil, nil
}
