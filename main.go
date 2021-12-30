package main

import (
	"net/http"
	//"strings"
	"log"
	"fmt"
	"encoding/xml"
	"io/ioutil"
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
  
  


  func main(){

	var mass = make([]ValCurs,90)
	var buff ValCurs

	for i := 0; i < 90; i++{
		
		xmlBytes, err := getXML("http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=11/11/2020")
		//log.Printf("I got bytes %s", string(xmlBytes))
		if err != nil {
			log.Printf("Failed to get XML: %v", err)
		  } else {
			xml.Unmarshal(xmlBytes, &buff)
			log.Printf("Got struct with fields %s %s %s/n", buff.Date , buff.Name, buff.Text )
			mass = append(mass, buff)
		  }

	}

  }