// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/xml"
	"fmt"
)

type Ht struct {
	Criteria struct {
		Nodes []struct {
			XMLName xml.Name
			Content string `xml:",chardata"`
		} `xml:",any"`
	}
	Session struct {
		ID  string `xml:"id,attr"`
		Htd struct {
			ID    string `xml:"ID,attr"`
			DataX []struct {
				XMLName xml.Name
				Info    struct {
					Count int `xml:"Count,attr"`
				}
				DataNodes []struct {
					XMLName xml.Name
					ID      string `xml:"Id,attr"`
					Type    string
				} `xml:",any"`
			} `xml:",any"`
		}
	} `xml:"session"`
}

func main() {
	var ht Ht
	if err := xml.Unmarshal([]byte(src), &ht); err != nil {
		panic(err)
	}

	result, err := xml.MarshalIndent(ht, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}

const src = `<Ht>
<Criteria>
    <ADNode_1>2</ADNode_1>
    <CDNode_1>2</CDNode_1>
    <IFNode_1>0</IFNode_1>
    <ADNode_2>2</ADNode_2>
    <CDNode_2>0</CDNode_2>
    <IFNode_2>0</IFNode_2>
    <ADNode_3>0</ADNode_3>
    <CDNode_3>0</CDNode_3>
    <IFNode_3>0</IFNode_3>
</Criteria>
<session id="1056134770841202228344907">
    <Htd ID="21170">
        <Data_1>
            <Info Count="2"></Info>
            <Data Id="iV29308/B2/R1">
                <Type>TR1</Type>
            </Data>
            <Data Id="iV29308/B3/R1">
                <Type>TR1</Type>
            </Data>
            <Data Id="iV29308/B4/R1">
                <Type>TR1</Type>
            </Data>
            <Data Id="iV29308/B6/R1">
                <Type>TR1</Type>
            </Data>
        </Data_1>
        <Data_2>
            <Info Count="2"></Info>
            <Data Id="iV29308/B2/R1">
                <Type>TR2</Type>
            </Data>
            <Data Id="iV29308/B3/R1">
                <Type>TR2</Type>
            </Data>
            <Data Id="iV29308/B4/R1">
                <Type>TR2</Type>
            </Data>
            <Data Id="iV29308/B6/R1">
                <Type>TR3</Type>
            </Data>
        </Data_2>
    </Htd>
</session>
</Ht>`
