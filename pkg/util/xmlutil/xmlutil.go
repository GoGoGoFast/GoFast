package xmlutil

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
)

// Document 表示XML文档。
// Document represents an XML document.
type Document struct {
	Root string
	*xmlquery.Node
}

// ReadXML 从文件读取XML文档。
// ReadXML reads an XML document from a file.
func ReadXML(filePath string) (*Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr
		}
	}()

	doc, err := xmlquery.Parse(file)
	if err != nil {
		return nil, err
	}

	root := doc.SelectElement("root")
	if root == nil {
		return nil, errors.New("root element not found")
	}

	return &Document{Root: root.Data, Node: doc}, nil
}

// ParseXML 从字符串解析XML文档。
// ParseXML parses an XML document from a string.
func ParseXML(xmlStr string) (*Document, error) {
	doc, err := xmlquery.Parse(strings.NewReader(xmlStr))
	if err != nil {
		return nil, err
	}

	root := doc.SelectElement("root")
	if root == nil {
		return nil, errors.New("root element not found")
	}

	return &Document{Root: root.Data, Node: doc}, nil
}

// ToStr 将XML文档转换为字符串。
// ToStr converts an XML document to a string.
func ToStr(doc *Document) (string, error) {
	var buf bytes.Buffer
	if err := xmlNodeToString(doc.Node, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ToFile 将XML文档写入文件。
// ToFile writes an XML document to a file.
func ToFile(doc *Document, filePath string) error {
	xmlStr, err := ToStr(doc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, []byte(xmlStr), 0644)
}

// CreateXML 创建一个新的XML文档。
// CreateXML creates a new XML document.
func CreateXML(rootElementName string) *Document {
	root := &xmlquery.Node{Type: xmlquery.ElementNode, Data: rootElementName}
	return &Document{Root: rootElementName, Node: root}
}

// CleanInvalid 去除XML文本中的无效字符。
// CleanInvalid removes invalid characters from an XML string.
func CleanInvalid(xmlStr string) string {
	// 简单示例：仅删除非ASCII字符
	return strings.Map(func(r rune) rune {
		if r >= 0x20 && r <= 0x7E {
			return r
		}
		return -1
	}, xmlStr)
}

// GetElements 根据节点名获得子节点列表。
// GetElements gets a list of child elements by tag name.
func GetElements(doc *Document, tagName string) []*xmlquery.Node {
	return xmlquery.Find(doc.Node, "//"+tagName)
}

// GetElement 根据节点名获得第一个子节点。
// GetElement gets the first child element by tag name.
func GetElement(doc *Document, tagName string) *xmlquery.Node {
	return xmlquery.FindOne(doc.Node, "//"+tagName)
}

// ElementText 根据节点名获得第一个子节点的文本值。
// ElementText gets the text content of the first child element by tag name.
func ElementText(doc *Document, tagName string) (string, error) {
	node := GetElement(doc, tagName)
	if node == nil {
		return "", errors.New("element not found")
	}
	return node.InnerText(), nil
}

// TransElements 将NodeList转换为Element列表。
// TransElements converts a NodeList to a list of Elements.
func TransElements(nodeList []*xmlquery.Node) []*xmlquery.Node {
	return nodeList
}

// WriteObjectAsXML 将可序列化的对象转换为XML写入文件。
// WriteObjectAsXML writes a serializable object to an XML file.
func WriteObjectAsXML(obj interface{}, filePath string) error {
	output, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, output, 0644)
}

// ReadObjectFromXML 从XML中读取对象。
// ReadObjectFromXML reads an object from an XML file.
func ReadObjectFromXML(filePath string, obj interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, obj)
}

// CreateXPath 创建一个新的XPath对象。
// CreateXPath creates a new XPath object.
func CreateXPath(expr string) (*xpath.Expr, error) {
	return xpath.Compile(expr)
}

// GetByXPath 通过XPath表达式获取XML节点。
// GetByXPath gets XML nodes by an XPath expression.
func GetByXPath(doc *Document, expr string) ([]*xmlquery.Node, error) {
	nodes := xmlquery.Find(doc.Node, expr)
	return nodes, nil
}

// xmlNodeToString 将 xmlquery.Node 转换为字符串
func xmlNodeToString(node *xmlquery.Node, buf *bytes.Buffer) error {
	if node.Type == xmlquery.TextNode {
		buf.WriteString(node.Data)
		return nil
	}

	buf.WriteString("<" + node.Data)
	for _, attr := range node.Attr {
		buf.WriteString(" " + attr.Name.Local + "=\"" + attr.Value + "\"")
	}
	buf.WriteString(">")
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if err := xmlNodeToString(child, buf); err != nil {
			return err
		}
	}
	buf.WriteString("</" + node.Data + ">")
	return nil
}
