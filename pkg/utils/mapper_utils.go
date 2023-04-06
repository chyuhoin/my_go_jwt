package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine
var _cfg *OrmInfo = nil

func GetOrmEngine() *xorm.Engine {
	if engine != nil {
		return engine
	}

	info, err := parseConfig("database_config.json")
	if err != nil {
		return nil
	}
	initEngine(info)

	return engine
}

func initEngine(ormInfo *OrmInfo) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		ormInfo.UserName, ormInfo.Password, ormInfo.Host, ormInfo.Port, ormInfo.DataBase)
	var err error
	engine, err = xorm.NewEngine(ormInfo.DriverName, url)
	if err != nil {
		log.Fatal(err.Error())
	}

	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, "tb_")
	engine.SetTableMapper(tbMapper)
	engine.ShowSQL(true)
}

func parseConfig(path string) (*OrmInfo, error) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
