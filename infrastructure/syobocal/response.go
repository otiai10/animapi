package syobocal

import "encoding/xml"

type TitleLookupResponse struct {
	XMLName    xml.Name   `xml:"TitleLookupResponse"`
	Result     Result     `xml:"Result"`
	TitleItems TitleItems `xml:"TitleItems"`
}
type Result struct {
	XMLName xml.Name `xml:"Result"`
	Code    int      `xml:"Code"`
	Message string   `xml:"Message"`
}
type TitleItems struct {
	XMLName   xml.Name `xml:"TitleItems"`
	TitleItem []TitleItem
}
type TitleItem struct {
	XMLName       xml.Name `xml:"TitleItem"`
	TID           string   `xml:"TID"`
	LastUpdate    string   `xml:"LastUpdate"`
	Title         string   `xml:"Title"`
	ShortTitle    string   `xml:"ShortTitle"`
	TitleYomi     string   `xml:"TitleYomi"`
	Comment       string   `xml:"Comment"`
	FirstYear     string   `xml:"FirstYear"`
	FirstMonth    string   `xml:"FirstMonth"`
	FirstEndYear  string   `xml:"FirstEndYear"`
	FirstEndMonth string   `xml:"FirstEndMonth"`
	Cat           string   `xml:"Cat"`
}
