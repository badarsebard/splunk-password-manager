package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

import "encoding/json"

import "io/ioutil"
import "net/http"
import "net/url"
import "log"

type Root struct {
	Data		[]Credential	`vugu:"data"`
	Rerender	bool
}

type Credential struct {
	app		string
	realm		string
	username	string
	password	string
	clear_password	string
	ShowPassword	bool	`vugu:"data"`
}

func (c *Root) BeforeBuild() {
	if c.Rerender {
		return
	}
	res, err := http.PostForm("https://localhost:8089/services/auth/login", url.Values{"username": {"admin"}, "password": {"changeme1"}, "output_mode": {"json"}})
	if err != nil {
		log.Printf("Error fetch()ing: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)

	session_key := data["sessionKey"].(string)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://localhost:8089/servicesNS/-/-/storage/passwords", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}
	q := req.URL.Query()
	q.Add("output_mode", "json")
	req.URL.RawQuery = url.Values{"output_mode": {"json"}}.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Splunk %s", session_key))
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching()ing: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &data)
	entries := data["entry"].([]interface{})
	credentials := []Credential{}
	for i := 0; i < len(entries); i++ {
		entry := entries[i].(interface{}).(map[string]interface{})
		acl := entry["acl"].(map[string]interface{})
		app := acl["app"].(string)
		content := entry["content"].(map[string]interface{})
		realm := content["realm"].(string)
		username := content["username"].(string)
		password := content["password"].(string)
		clear_password := content["clear_password"].(string)
		credentials = append(credentials, Credential{app, realm, username, password, clear_password, false})
	}
	c.Data = credentials
	c.Rerender = true
}
func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "table", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "tbody", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "tr", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				vgn.SetInnerHTML(vugu.HTML("\n            \x3Cth\x3EApp\x3C/th\x3E\n            \x3Cth\x3ERealm\x3C/th\x3E\n            \x3Cth\x3EUsername\x3C/th\x3E\n            \x3Cth\x3EPassword\x3C/th\x3E\n        "))
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
				vgparent.AppendChild(vgn)
				for i := 0; i < len(c.Data); i++ {
					var vgiterkey interface{} = i
					_ = vgiterkey
					i := i
					_ = i
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "tr", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(c.Data[i].app)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(c.Data[i].realm)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						vgn.SetInnerHTML(c.Data[i].username)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						if !c.Data[i].ShowPassword {
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.Data[i].password)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						if c.Data[i].ShowPassword {
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							vgn.SetInnerHTML(c.Data[i].clear_password)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						if !c.Data[i].ShowPassword {
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "button", Attr: []vugu.VGAttribute(nil)}
								vgparent.AppendChild(vgn)
								vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
									EventType:	"click",
									Func:		func(event vugu.DOMEvent) { c.Data[i].ShowPassword = !c.Data[i].ShowPassword },
									// TODO: implement capture, etc. mostly need to decide syntax
								})
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Reveal"}
									vgparent.AppendChild(vgn)
								}
							}
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
						vgparent.AppendChild(vgn)
						if c.Data[i].ShowPassword {
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
							vgparent.AppendChild(vgn)
							{
								vgparent := vgn
								_ = vgparent
								vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "button", Attr: []vugu.VGAttribute(nil)}
								vgparent.AppendChild(vgn)
								vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
									EventType:	"click",
									Func:		func(event vugu.DOMEvent) { c.Data[i].ShowPassword = !c.Data[i].ShowPassword },
									// TODO: implement capture, etc. mostly need to decide syntax
								})
								{
									vgparent := vgn
									_ = vgparent
									vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Hide"}
									vgparent.AppendChild(vgn)
								}
							}
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
				vgparent.AppendChild(vgn)
			}
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n"}
		vgparent.AppendChild(vgn)
	}
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value
