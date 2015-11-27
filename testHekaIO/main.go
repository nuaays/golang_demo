package main

import (
	"fmt"
	"os"
	"time"
)

func make_payload() string {
	var payload string
	for {
		for i := 0; i < 100; i++ {
			payload = fmt.Sprint(time.Now().Unix())
		}
		payload = fmt.Sprint(payload, "\n")
	}

	return payload
}

func make_log(files []string) {
	for i := 0; i < len(files); i++ {
		go func() {
			fd, _ := os.Open(files[i])
			payload := make_payload()
			fd.WriteString(string(payload))
		}()
	}
}

func main() {
	var hekad_conf string = "/etc/hekad.toml"
	var file_num int = 800
	// generate config file
	var file_name_pattern string = "access\\.log"
	files := make([]string, 0)
	for i := 0; i < file_num; i++ {
		fname := fmt.Sprint(file_name_pattern, ".", i)
		files = append(files, fname)
	}

	var conf string
	for i := 0; i < len(files); i++ {
		logstreamer_template := fmt.Sprintf("[test_log%d]\ntype = \"LogstreamerInput\"\nlog_directory = \"/var/log/nginx\"\nfile_match = '%s\\.%d'\njournal_directory = \"/var/heka\"\n\n", i, file_name_pattern, i)
		conf += logstreamer_template
	}
	conf += fmt.Sprintf("[PayloadEncoder]\n\n")
	conf += fmt.Sprintf("[LogOutput]\nencoder=\"PayloadEncoder\"\nmessage_matcher = \"TRUE\"\n\n")

	// write config file
	f, err := os.Open(hekad_conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(conf)
	f.Close()

}
