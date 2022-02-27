package list

import (
	"I2Oprobe/internal/log"
	"I2Oprobe/internal/model"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


const (
	sourceConfigErr = "You must write with this type" +
		"Address(string)_port(int)_timeout(int)_concurrence(int)_counts(int)_threshold(int)_shopname(string)." +
		" Please check the source file"
)

type sourFile struct {
	path string
	texts []string
}

func NewSourFile(path string) SourceFile {
	return &sourFile{
		path: path,
	}
}

func (sf *sourFile) Read()  {
	fmt.Println(sf.path)
	f, err := os.Open(sf.path)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer f.Close()

	rds := bufio.NewReader(f)
	for {
		line, err := rds.ReadString('\n')
		if err != nil || io.EOF == err{
			log.Debug(err)
			break
		}
		sf.texts = append(sf.texts, line)
	}
}


func (sf *sourFile) ToChan() (ds []interface{}) {
	// nil slice
	if len(sf.texts) == 0 {
		log.Warn("No source in probelist.txt. Please check")
		panic("No source in probelist.txt. Please check")
	}
	for _, val := range sf.texts {
		address, port, timeout, concur, counts, threshold, shopname := strSplitTools(val)

		ds = append(ds, model.Source{
			Address:     address,
			Port:        port,
			Timeout:     timeout,
			Concurrence: concur,
			Counts:      counts,
			Threshold:   threshold,
			Shopname:    shopname,
		})
	}
	return ds
}

func strSplitTools(s string) (Address string, port, timeout, concur, counts int, threshold int, Shopname string) {
	res := strings.Split(s, "_")

	if len(res) != 7 {
		log.Error(sourceConfigErr)
		panic(sourceConfigErr)
	}

	//Port
	port, err :=strconv.Atoi(res[1])
	if err != nil {
		log.Error(sourceConfigErr)
		panic(sourceConfigErr)
	}

	//timeout
	timeout, err =strconv.Atoi(res[2])
	if err != nil {
		log.Error(sourceConfigErr)
		panic(sourceConfigErr)
	}

	//concurrence
	concur, err =strconv.Atoi(res[3])
	if err != nil {
		log.Error(sourceConfigErr)
		panic(sourceConfigErr)
	}

	//counts
	counts, err =strconv.Atoi(res[4])
	if err != nil {
		log.Error(sourceConfigErr)
		panic(sourceConfigErr)
	}
	//threshold
	threshold, err =strconv.Atoi(res[5])
	if err != nil {
		log.Error(sourceConfigErr)
		panic(sourceConfigErr)
	}

	//

	return res[0], port, timeout, concur, counts, threshold, res[6]
}
