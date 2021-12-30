package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []Valute `xml:"Valute"`
}
type Valute struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func main() {

	var mass = make([]ValCurs, 90)
	buff := ValCurs{}
	date := time.Now()
	for i := 0; i < 90; i++ {
		date.Add(-24 * time.Hour)
		y, m, d := date.Date()
		s := strconv.Itoa(d) + "/" + strconv.Itoa(int(m)) + "/" + strconv.Itoa(y)
		xmlBytes, err := getXML("http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=" + s)
		log.Printf("I got bytes %s", string(xmlBytes))
		if err != nil {
			log.Printf("Failed to get XML: %v", err)
		} else {
			xml.Unmarshal(xmlBytes, &buff)
			//log.Printf("Got struct with fields %s %s %s/n", buff.Date, buff.Name, buff.Text)
			mass = append(mass, buff)
		}

	}

}
