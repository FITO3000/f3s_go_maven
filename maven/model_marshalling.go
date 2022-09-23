package maven

import (
	"encoding/xml"
	"io"
)

func Marshal(model *Project) ([]byte, error) {
	return xml.MarshalIndent(model, "", "  ")
}

func Unmarshal(data []byte) (*Project, error) {
	project := NewProject()
	if err := xml.Unmarshal(data, project); err != nil {
		return nil, err
	} else {
		return project, nil
	}
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	type entry struct {
		XMLName xml.Name
		Key     string `xml:"name,attr"`
		Value   string `xml:",chardata"`
	}
	e := entry{}
	p.Entries = map[string]string{}
	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		e.Key = e.XMLName.Local
		p.Entries[e.Key] = e.Value
	}
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (p *Properties) MarshalXML(d *xml.Encoder, start xml.StartElement) (err error) {
	tokens := []xml.Token{start}

	for key, value := range p.Entries {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := d.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	return d.Flush()
}
