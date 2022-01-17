package massbank

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "2006.01.02"

type Property interface {
	Validate() bool
	Output() string
	Parse(string) error
}

type DefaultProperty struct {
}

type StringProperty struct {
	string
	DefaultProperty
}

type SubtagProperty struct {
	StringProperty
	subtag string
}

func (p *DefaultProperty) Validate() bool {
	return false
}

func (p *DefaultProperty) Output() string {
	return "not implemented"
}

func (p *DefaultProperty) Parse(string) error {
	return errors.New("not implemented")
}

func (p *StringProperty) Parse(s string) error {
	p.string = s
	return nil
}

func (p *SubtagProperty) Parse(s string) error {
	ss := strings.SplitN(s, " ", 2)
	p.subtag = ss[0]
	p.string = ss[1]
	return nil
}

type tagProperties struct {
	Type  reflect.Type
	Name  string
	Index []int
}

var TagMap = map[string]tagProperties{}

type Massbank struct {
	Accession          *RecordAccession      `mb:"ACCESSION"`
	Deprecated         *RecordDeprecated     `mb:"DEPRECATED"`
	RecordTitle        *RecordTitle          `mb:"RECORD_TITLE"`
	Date               *RecordDate           `mb:"DATE"`
	Authors            *RecordAuthorNames    `mb:"AUTHORS"`
	License            *RecordLicense        `mb:"LICENSE"`
	Copyright          *RecordCopyright      `mb:"COPYRIGHT"`
	Publication        *RecordPublication    `mb:"PUBLICATION"`
	Project            *RecordProject        `mb:"PROJECT"`
	Comments           []*RecordComment      `mb:"COMMENT"`
	ChNames            []*ChName             `mb:"CH$NAME"`
	ChClass            *ChCompoundClasses    `mb:"CH$COMPOUND_CLASS"`
	ChFormula          *ChFormula            `mb:"CH$FORMULA"`
	ChCdkDepict        []*CdkDepict          `mb:"CH$CDK_DEPICT"` // not for productive use
	ChMass             *ChMass               `mb:"CH$EXACT_MASS"`
	ChSmiles           *ChSmiles             `mb:"CH$SMILES"`
	ChInchi            *ChInchi              `mb:"CH$IUPAC"`
	ChLink             []*ChLink             `mb:"CH$LINK"`
	SpName             *SpName               `mb:"SP$SCIENTIFIC_NAME"`
	SpLineage          *SpLineage            `mb:"SP$LINEAGE"`
	SpLink             []*SpLink             `mb:"SP$LINK"`
	SpSample           []*SampleInformation  `mb:"SP$SAMPLE"`
	AcInstrument       *AcInstrument         `mb:"AC$INSTRUMENT"`
	AcInstrumentType   *AcInstrumentType     `mb:"AC$INSTRUMENT_TYPE"`
	AcMassSpectrometry []*AcMassSpectrometry `mb:"AC$MASS_SPECTROMETRY"`
	AcChromatography   []*AcChromatography   `mb:"AC$CHROMATOGRAPHY"`
	AcGeneral          []*AcGeneral          `mb:"AC$GENERAL"`
	MsFocusedIon       []*MsFocusedIon       `mb:"MS$FOCUSED_ION"`
	MsDataProcessing   []*MsDataProcessing   `mb:"MS$DATA_PROCESSING"`
	PkSplash           *PkSplash             `mb:"PK$SPLASH"`
	PkAnnotation       *PkAnnotation         `mb:"PK$ANNOTATION"`
	PkNumPeak          *PkNumPeak            `mb:"PK$NUM_PEAK"`
	PkPeak             *PkPeak               `mb:"PK$PEAK"`
}

type RecordAccession struct {
	StringProperty
}

type RecordDeprecated struct {
	Date   time.Time
	Reason string
}

type RecordTitle struct {
	StringProperty
}

type RecordDate struct {
	DefaultProperty
	Updated time.Time
	Created time.Time
}

func (d *RecordDate) Parse(s string) error {
	var err error
	ss := strings.SplitN(s, " ", 2)
	if len(ss) > 1 {
		re := regexp.MustCompile("\\(Cretated (.*)\\)")
		ss2 := re.FindStringSubmatch(ss[1])
		if len(ss2) == 2 {
			if d.Created, err = time.Parse(dateFormat, ss2[1]); err != nil {
				return err
			} else {
				return errors.New("Format error in Date")
			}
		}
	} else {
		if d.Created, err = time.Parse(dateFormat, ss[0]); err != nil {
			return err
		}
	}
	d.Updated, err = time.Parse(dateFormat, ss[0])
	if err != nil {
		return err
	}
	return nil

}

type RecordAuthorNames struct {
	DefaultProperty
	value []RecordAuthorName
}

func (names *RecordAuthorNames) Parse(s string) error {
	ss := strings.Split(s, ",")
	for _, s1 := range ss {
		re := regexp.MustCompile("(.*)(\\[(.*)\\])?")
		ss1 := re.FindStringSubmatch(s1)
		marc := ""
		if len(ss1) > 2 {
			marc = ss1[2]
		}
		if len(ss1) > 1 {
			names.value = append(names.value, RecordAuthorName{ss1[1], marc})
		}
	}
	return nil
}

type RecordAuthorName struct {
	Name        string
	MarcRelator string
}

type RecordLicense struct {
	StringProperty
}

type RecordCopyright struct {
	StringProperty
}

type RecordPublication struct {
	StringProperty
}

type RecordProject struct {
	StringProperty
}

type RecordComment struct {
	SubtagProperty
}

type RecordSubtag string

type RecordMbTag string

type ChName struct {
	StringProperty
}

type ChCompoundClasses struct {
	DefaultProperty
	value []ChCompoundClass
}

func (cc *ChCompoundClasses) Parse(s string) error {
	ss := strings.Split(s, ";")
	for _, s1 := range ss {
		var c = ChCompoundClass(strings.TrimSpace(s1))
		cc.value = append(cc.value, c)
	}
	return nil
}

type ChCompoundClass string

type ChFormula struct {
	StringProperty
}

type ChMass struct {
	DefaultProperty
	value float64
}

func (mass *ChMass) Parse(s string) error {
	var err error
	mass.value, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	return nil
}

type ChSmiles struct {
	StringProperty
}

type ChInchi struct {
	StringProperty
}

type ChLink struct {
	SubtagProperty
}

type ExtDatabase struct {
	StringProperty
}

type CdkDepict struct {
	StringProperty
}

type SpName struct {
	StringProperty
}

type SpLineage struct {
	DefaultProperty
	value []SpLineageElement
}

type SpLineageElement struct {
	StringProperty
}

type SpLink struct {
	DefaultProperty
	Database   string
	Identifier string
}

type SampleInformation struct {
	StringProperty
}

type AcInstrument struct {
	StringProperty
}

type Separation string
type Ionization string
type Analyzer string
type AcInstrumentType struct {
	StringProperty
}
type MsType string
type AcMassSpectrometry struct {
	SubtagProperty
}

type AcChromatography struct {
	SubtagProperty
}

type AcGeneral struct {
	SubtagProperty
}

type PkPeak struct {
	DefaultProperty
	Header []string
	Values []PeakValue
}

type MsFocusedIon struct {
	SubtagProperty
}
type MsDataProcessing struct {
	SubtagProperty
}

type PkSplash struct {
	StringProperty
}

type AnnotationValue []interface{}

type PkAnnotation struct {
	DefaultProperty
	Header []string
	Values []AnnotationValue
}

type PkNumPeak struct {
	DefaultProperty
	Value uint
}

func (n *PkNumPeak) Parse(s string) error {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	n.Value = uint(val)
	return nil
}

type PeakValue struct {
	mz        float64
	intensity float64
	rel       uint
}

func (p *PeakValue) Parse(s string) error {

}

type TagValue struct {
	tag    string
	values []string
}

type TagValues []TagValue

func (mb *Massbank) ParseFile(fileName string) error {
	if len(TagMap) == 0 {
		buildTags()
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "//") {
			// ignore comment
		} else if strings.HasPrefix(line, "  ") {
			// multiline value
			println("not implemented", line)
		} else {
			s := strings.SplitN(line, ":", 2)
			if len(s) == 2 {
				tagname := s[0]
				value := strings.TrimSpace(s[1])
				tagInfo := TagMap[tagname]
				index := tagInfo.Index
				mb2 := reflect.ValueOf(mb)
				mb3 := reflect.Indirect(mb2)
				prop := mb3.FieldByIndex(index)
				prop2 := prop.Type().Elem()
				if prop.Kind() == reflect.Slice {
					prop2 = prop2.Elem()
				}
				newPro := reflect.New(prop2)
				newInterf := newPro.Interface()
				propInt := newInterf.(Property)
				err = propInt.Parse(value)
				if err != nil {
					println(err.Error(), line)
				}
				if prop.Kind() == reflect.Slice {
					prop.Set(reflect.Append(prop, newPro))
				} else {
					prop.Set(newPro)
				}
			} else {
				println("The line is not a valid massbank tag line: \n", line)
			}
		}
		lineNum++
	}
	return nil
}

// Build an array with type information and tag strings for parsing
func buildTags() {
	var mb = Massbank{}
	for _, field := range reflect.VisibleFields(reflect.TypeOf(mb)) {
		var props = tagProperties{}
		props.Name = field.Name
		props.Type = field.Type
		props.Index = field.Index
		TagMap[field.Tag.Get("mb")] = props
	}
}
