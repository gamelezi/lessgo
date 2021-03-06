package base

import (
	"github.com/lessos/lessgo/utils"
	"io/ioutil"
	"os"
)

type DataSet struct {
	DbName  string   `json:"dbname"`
	Engine  string   `json:"engine,omitempty"`
	Charset string   `json:"charset,omitempty"`
	Version int      `json:"version,omitempty"`
	Tables  []*Table `json:"tables"`
}

func LoadDataSetFromFile(file string) (DataSet, error) {

	var ds DataSet
	var err error

	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		return ds, err
	}

	fp, err := os.Open(file)
	if err != nil {
		return ds, err
	}
	defer fp.Close()

	cfg, err := ioutil.ReadAll(fp)
	if err != nil {
		return ds, err
	}

	return LoadDataSetFromString(string(cfg))
}

func LoadDataSetFromString(json string) (DataSet, error) {

	var ds DataSet

	err := utils.JsonDecode(json, &ds)

	return ds, err
}
