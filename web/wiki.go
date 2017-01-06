package main

  import (
  	"fmt"
  	"io/ioutil"
  )

  type page struct {
  	title	string
  	body	[]byte
  }
  func (p *page) save() string {
  	filename := p.title + ".txt"
  	err := ioutil.WriteFile(filename, p.body, 0600)
    if err != nil{
      return err.Error()
    }else{
      return ""
    }
  }
  func loadPage(title string) (*page, string) {
  	filename := title + ".txt"
  	body, err := ioutil.ReadFile(filename)
  	if err != nil {
  		return nil, err.Error()
  	}
  	return &page{title: title, body: body}, ""
  }
  func main() {
  	p1 := &page{title: "TestPage", body: []byte("This is a sample page.")}
  	p1.save()
  	p2, _ := loadPage("TestPage")
  	fmt.Println(string(p2.body))
  }
