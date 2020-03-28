package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type inforesult struct {
	Country   string
	Cases     string
	Deaths    string
	Recovered string
}

type infoCountry struct {
	Selected  string
	Countries []string
}

//Myscraper struc for scrap web
type Myscraper struct {
	doc       *goquery.Document
	timesaved int64
}

func (i *Myscraper) getNewDoc() {
	res, err := http.Get("https://www.worldometers.info/coronavirus/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var errreader error
	i.doc, errreader = goquery.NewDocumentFromReader(res.Body)
	if errreader != nil {
		log.Fatal(err)
	}

}

func (i *Myscraper) getDOC() *goquery.Document {
	secondNowTime := time.Now().Unix()
	if i.doc == nil || (secondNowTime) > (i.timesaved+600) {
		i.timesaved = secondNowTime
		i.getNewDoc()
	}

	return i.doc
}

//GetListCountry get list country available
func (i *Myscraper) GetListCountry() []byte {
	table := i.getDOC().Find("#main_table_countries_today")
	tr := table.Find("tr")

	infoCountryV := infoCountry{}

	listcountry := make([]string, 0)

	tr.Each(func(x int, curtr *goquery.Selection) {
		country := curtr.Find("td").Eq(0).Text()
		if country == "" {
			return
		}
		listcountry = append(listcountry, country)
	})

	infoCountryV.Countries = listcountry
	infoCountryV.Selected = getSelectedCountry()

	b, _ := json.Marshal(infoCountryV)

	return b
}

func (i Myscraper) getInfobyCountry(countryselect string) []byte {

	saveSelectedCountr(countryselect)
	inforesultV := inforesult{}

	table := i.getDOC().Find("#main_table_countries_today")
	tr := table.Find("tr")
	tr.Each(func(x int, curtr *goquery.Selection) {
		listtd := curtr.Find("td")
		country := listtd.Eq(0).Text()
		if country != countryselect {
			return
		}

		inforesultV.Country = country
		inforesultV.Cases = listtd.Eq(1).Text()
		inforesultV.Deaths = listtd.Eq(3).Text()
		inforesultV.Recovered = listtd.Eq(5).Text()

	})

	b, _ := json.Marshal(inforesultV)
	return b
}
