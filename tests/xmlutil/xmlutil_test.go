package xmlutil_test

import (
	"encoding/xml"
	"github.com/antchfx/xmlquery"
	"os"
	"testing"

	"GoFast/pkg/util/xmlutil"
	"github.com/stretchr/testify/assert"
)

func TestReadXML(t *testing.T) {
	// 创建临时 XML 文件
	tmpFile, err := os.CreateTemp("", "test.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	xmlContent := `<root><child>content</child></root>`
	_, err = tmpFile.WriteString(xmlContent)
	assert.NoError(t, err)

	doc, err := xmlutil.ReadXML(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "root", doc.Root)
}

func TestParseXML(t *testing.T) {
	xmlStr := `<root><child>content</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "root", doc.Root)
}

func TestToStr(t *testing.T) {
	xmlStr := `<root><child>content</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)

	result, err := xmlutil.ToStr(doc)
	assert.NoError(t, err)
	assert.Contains(t, result, `<root><child>content</child></root>`)
}

func TestToFile(t *testing.T) {
	doc := xmlutil.CreateXML("root")
	child := &xmlquery.Node{Type: xmlquery.ElementNode, Data: "child"}
	child.FirstChild = &xmlquery.Node{Type: xmlquery.TextNode, Data: "content"}
	doc.FirstChild = child

	tmpFile, err := os.CreateTemp("", "test.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	err = xmlutil.ToFile(doc, tmpFile.Name())
	assert.NoError(t, err)

	readDoc, err := xmlutil.ReadXML(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, readDoc)
	assert.Equal(t, "root", readDoc.Root)
}

func TestCreateXML(t *testing.T) {
	doc := xmlutil.CreateXML("root")
	assert.NotNil(t, doc)
	assert.Equal(t, "root", doc.Root)
}

func TestCleanInvalid(t *testing.T) {
	xmlStr := "valid\x00invalid"
	cleaned := xmlutil.CleanInvalid(xmlStr)
	assert.Equal(t, "validinvalid", cleaned)
}

func TestGetElements(t *testing.T) {
	xmlStr := `<root><child>content1</child><child>content2</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)

	elements := xmlutil.GetElements(doc, "child")
	assert.Len(t, elements, 2)
}

func TestGetElement(t *testing.T) {
	xmlStr := `<root><child>content1</child><child>content2</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)

	element := xmlutil.GetElement(doc, "child")
	assert.NotNil(t, element)
	assert.Equal(t, "child", element.Data)
}

func TestElementText(t *testing.T) {
	xmlStr := `<root><child>content</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)

	text, err := xmlutil.ElementText(doc, "child")
	assert.NoError(t, err)
	assert.Equal(t, "content", text)
}

func TestTransElements(t *testing.T) {
	xmlStr := `<root><child>content1</child><child>content2</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)

	elements := xmlutil.GetElements(doc, "child")
	transElements := xmlutil.TransElements(elements)
	assert.Len(t, transElements, 2)
}

func TestWriteObjectAsXML(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	obj := struct {
		XMLName xml.Name `xml:"root"`
		Child   string   `xml:"child"`
	}{
		Child: "content",
	}

	err = xmlutil.WriteObjectAsXML(obj, tmpFile.Name())
	assert.NoError(t, err)

	var readObj struct {
		XMLName xml.Name `xml:"root"`
		Child   string   `xml:"child"`
	}
	err = xmlutil.ReadObjectFromXML(tmpFile.Name(), &readObj)
	assert.NoError(t, err)
	assert.Equal(t, "content", readObj.Child)
}

func TestReadObjectFromXML(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	xmlContent := `<root><child>content</child></root>`
	_, err = tmpFile.WriteString(xmlContent)
	assert.NoError(t, err)

	var obj struct {
		XMLName xml.Name `xml:"root"`
		Child   string   `xml:"child"`
	}
	err = xmlutil.ReadObjectFromXML(tmpFile.Name(), &obj)
	assert.NoError(t, err)
	assert.Equal(t, "content", obj.Child)
}

func TestCreateXPath(t *testing.T) {
	expr, err := xmlutil.CreateXPath("//child")
	assert.NoError(t, err)
	assert.NotNil(t, expr)
}

func TestGetByXPath(t *testing.T) {
	xmlStr := `<root><child>content1</child><child>content2</child></root>`
	doc, err := xmlutil.ParseXML(xmlStr)
	assert.NoError(t, err)

	nodes, err := xmlutil.GetByXPath(doc, "//child")
	assert.NoError(t, err)
	assert.Len(t, nodes, 2)
}
