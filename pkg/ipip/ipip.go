package ipip

import (
	"fmt"
	"log"
	"os"

	"github.com/ipipdotnet/ipdb-go"
)

type IPIPFree struct {
	*ipdb.City
}

func NewIPIP(filePath string) (*IPIPFree, error) {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		log.Printf("IPIP数据库不存在，请手动下载解压后保存到本地: %s \n", filePath)
		log.Println("下载链接： https://www.ipip.net/product/ip.html")
		return nil, err
	} else {
		db, err := ipdb.NewCity(filePath)
		if err != nil {
			return nil, err
		}
		return &IPIPFree{City: db}, nil
	}
}

type Result struct {
	Country       string
	Region        string
	City          string
	Isp_domain    string
	Line          string
	Countrycode   string
	Continentcode string
}

func (r Result) String() string {
	var result string
	result = ""
	if r.City == "" {
		result = fmt.Sprintf("%s %s", r.Country, r.Region)
	} else {
		result = fmt.Sprintf("%s %s %s", r.Country, r.Region, r.City)
	}

	result = result + " " + r.Isp_domain
	result = result + " " + r.Line
	result = result + " " + r.Countrycode
	result = result + " " + r.Continentcode

	return result
}

func (db IPIPFree) Find(query string, params ...string) (result fmt.Stringer, err error) {
	info, err := db.FindInfo(query, "CN")
	if err != nil || info == nil {
		return nil, err
	} else {
		// info contains more info
		result = Result{
			Country:       info.CountryName,
			Region:        info.RegionName,
			City:          info.CityName,
			Isp_domain:    info.IspDomain,
			Line:          info.Line,
			Countrycode:   info.CountryCode,
			Continentcode: info.ContinentCode,
		}
		return
	}
}

func (db IPIPFree) Name() string {
	return "ipip"
}
