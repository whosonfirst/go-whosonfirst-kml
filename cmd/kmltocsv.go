package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type KML struct {
	XMLName       xml.Name `xml:"kml"`
	Namespace     string   `xml:"xmlns,attr"`
	AtomNamespace string   `xml:"xmlns:atom,attr"`
	GxNamespace   string   `xml:"xmlns:gx,attr"`
	KmlNamespace  string   `xml:"xmlns:kml,attr"`
	Document      Document `xml:"Document"`
}

type Document struct {
	Name  string `xml:"name"`
	Style Style  `xml:"Style,omitempty"`
	// StyleMap
	Placemark []Placemark `xml:"Placemark"`
	Folder    Folder      `xml:"Folder"`
}

type Folder struct {
	Name      string      `xml:"name"`
	Placemark []Placemark `xml:"Placemark"`
}

type Placemark struct {
	Name        string       `xml:"name"`
	Description Description  `xml:"description"`
	StyleUrl    string       `xml:"styleUrl,omitempty"`
	Point       Point        `xml:"Point,omitempty"`
	LineString  []LineString `xml:"LineString,omitempty"`
}

type Point struct {
	Coordinates string `xml:"coordinates,omitempty"`
}

func (pt Point) ToLatLon() (float64, float64, error) {

	lat := 0.0
	lon := 0.0

	coords := pt.Coordinates

	coords = strings.TrimSpace(coords)
	parts := strings.Split(coords, ",")

	str_lon := parts[0]
	str_lat := parts[1]

	var err error

	lon, err = strconv.ParseFloat(str_lon, 64)

	if err != nil {
		return lat, lon, err
	}

	lat, err = strconv.ParseFloat(str_lat, 64)

	if err != nil {
		return lat, lon, err
	}

	return lat, lon, err
}

type Description struct {
	Data string `xml:",cdata"`
}

type LineString struct {
	Extrude      int    `xml:"extrude,omitempty"`
	AltitudeMode string `xml:"altitudeMode,omitempty"`
	Coordinates  string `xml:"coordinates,omitempty"`
}

type Icon struct {
	Href  string `xml:"href"`
	Scale string `xml:"scale"`
}

type IconStyle struct {
	Id   string `xml:"id,attr"`
	Icon Icon   `xml:"Icon"`
}

type Style struct {
	Id        string    `xml:"id,attr"`
	IconStyle IconStyle `xml:"IconStyle"`
}

func main() {

	flag.Parse()

	args := flag.Args()

	header := []string{"latitude", "longitude", "name", "description"}

	writer := csv.NewWriter(os.Stdout)
	writer.Write(header)

	for _, path := range args {

		body, err := ioutil.ReadFile(path)

		if err != nil {
			log.Fatal(err)
		}

		var k KML

		err = xml.Unmarshal(body, &k)

		if err != nil {
			log.Fatal(err)
		}

		for _, p := range k.Document.Folder.Placemark {

			if p.Point.Coordinates == "" {
				continue
			}

			lat, lon, err := p.Point.ToLatLon()

			if err != nil {
				log.Fatal(err)
			}

			row := []string{
				strconv.FormatFloat(lat, 'f', -1, 32),
				strconv.FormatFloat(lon, 'f', -1, 32),
				p.Name,
				p.Description.Data,
			}

			writer.Write(row)
		}

		writer.Flush()
	}
}
