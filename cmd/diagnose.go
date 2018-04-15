package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/Sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"os"
	"path"
	"bufio"
	"io"
	"fmt"
)

const (
	INVALID_CLUSTER_DUMP_FILE = "Please input a valid cluster dump file."
	DIAGNOSE_ERROR            = "Failed to diagnose your cluster dump"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "diagnose will split cluster dump to multi files and check the key words in file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		if len(args) != 1 {
			log.Errorf(INVALID_CLUSTER_DUMP_FILE)
			return
		}

		var absolute_path string
		file_path := args[0]
		if strings.HasPrefix(file_path, "/") || strings.HasPrefix(file_path, "\\") {
			absolute_path = file_path
		} else {
			pwd, err := os.Getwd()
			if err != nil {
				log.Errorf(INVALID_CLUSTER_DUMP_FILE+",Because of %s", err.Error())
				return
			}
			absolute_path = path.Join(pwd, file_path)
		}

		if _, err := os.Stat(absolute_path); os.IsNotExist(err) == true {
			log.Errorf(INVALID_CLUSTER_DUMP_FILE+",Because of %s", err.Error())
			return
		}
		err := DiagnoseClusterDump(absolute_path)
		if err != nil {
			log.Errorf(DIAGNOSE_ERROR+",Because of %s", err.Error())
			return
		}
	},
}

func DiagnoseClusterDump(file_path string) (err error) {
	file, err := os.Open(file_path)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var line string

	is_json_line := false
	is_log_line := false
	cache_lines := make([]string, 0)
	i := 0
	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		if is_json_line == false || is_log_line == false {
			if strings.HasPrefix(line, "{") {
				is_json_line = true
			}
			if strings.HasPrefix(line, "==== START logs") {
				is_log_line = true
			}
			cache_lines = append(cache_lines, line)
		}

		if strings.HasPrefix(line, "}") || strings.HasPrefix(line, "==== END logs") {
			if is_json_line == true {
				JsonLinesHandler(cache_lines)
			}
			if is_log_line == true {
				LogLinesHandler(cache_lines)
			}
			is_json_line = false
			is_log_line = false
			i++
			cache_lines = make([]string, 0)
		}

	}

	if err != io.EOF {
		return
	}

	return nil
}

func JsonLinesHandler(lines []string) {

}

func LogLinesHandler(lines []string) {
	fmt.Println(lines[0], lines[len(lines)-1])
}
