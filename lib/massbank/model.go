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

var lastTag string

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
	Accession          *RecordAccession      `mb2:"ACCESSION"`
	Deprecated         *RecordDeprecated     `mb2:"DEPRECATED"`
	RecordTitle        *RecordTitle          `mb2:"RECORD_TITLE"`
	Date               *RecordDate           `mb2:"DATE"`
	Authors            *RecordAuthorNames    `mb2:"AUTHORS"`
	License            *RecordLicense        `mb2:"LICENSE"`
	Copyright          *RecordCopyright      `mb2:"COPYRIGHT"`
	Publication        *RecordPublication    `mb2:"PUBLICATION"`
	Project            *RecordProject        `mb2:"PROJECT"`
	Comments           []*RecordComment      `mb2:"COMMENT"`
	ChNames            []*ChName             `mb2:"CH$NAME"`
	ChClass            *ChCompoundClasses    `mb2:"CH$COMPOUND_CLASS"`
	ChFormula          *ChFormula            `mb2:"CH$FORMULA"`
	ChCdkDepict        []*CdkDepict          `mb2:"CH$CDK_DEPICT"` // not for productive use
	ChMass             *ChMass               `mb2:"CH$EXACT_MASS"`
	ChSmiles           *ChSmiles             `mb2:"CH$SMILES"`
	ChInchi            *ChInchi              `mb2:"CH$IUPAC"`
	ChLink             []*ChLink             `mb2:"CH$LINK"`
	SpName             *SpName               `mb2:"SP$SCIENTIFIC_NAME"`
	SpLineage          *SpLineage            `mb2:"SP$LINEAGE"`
	SpLink             []*SpLink             `mb2:"SP$LINK"`
	SpSample           []*SampleInformation  `mb2:"SP$SAMPLE"`
	AcInstrument       *AcInstrument         `mb2:"AC$INSTRUMENT"`
	AcInstrumentType   *AcInstrumentType     `mb2:"AC$INSTRUMENT_TYPE"`
	AcMassSpectrometry []*AcMassSpectrometry `mb2:"AC$MASS_SPECTROMETRY"`
	AcChromatography   []*AcChromatography   `mb2:"AC$CHROMATOGRAPHY"`
	AcGeneral          []*AcGeneral          `mb2:"AC$GENERAL"`
	MsFocusedIon       []*MsFocusedIon       `mb2:"MS$FOCUSED_ION"`
	MsDataProcessing   []*MsDataProcessing   `mb2:"MS$DATA_PROCESSING"`
	PkSplash           *PkSplash             `mb2:"PK$SPLASH"`
	PkAnnotation       *PkAnnotation         `mb2:"PK$ANNOTATION"`
	PkNumPeak          *PkNumPeak            `mb2:"PK$NUM_PEAK"`
	PkPeak             *PkPeak               `mb2:"PK$PEAK"`
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
		re := regexp.MustCompile("(.*)([(.*)])?")
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

func (p *PkPeak) Parse(s string) error {
	if s != "m/z int. rel.int." {
		return errors.New("PK$ is not valid")
	}
	p.Header = strings.Split(s, " ")
	return nil
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
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "//") {
			// ignore comment
		} else if strings.HasPrefix(line, "  ") {
			if lastTag == "PK$PEAK" {
				var pv PeakValue
				if err = pv.parse(strings.TrimSpace(line)); err != nil {
					println("Could not read Peakvalue: line ", lineNum, err.Error())
				} else {
					mb.PkPeak.Values = append(mb.PkPeak.Values, pv)
				}
			} else {
				println("not implemented", line)
			}
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
				lastTag = tagname
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
	file.Close()
	return nil
}

func (p *PeakValue) parse(s string) error {
	ss := strings.Split(s, " ")
	var err error
	var rel uint64
	if p.mz, err = strconv.ParseFloat(ss[0], 32); err != nil {
		return errors.New("Could not parse mz value")
	}
	if p.intensity, err = strconv.ParseFloat(ss[1], 32); err != nil {
		return errors.New("Could not parse intensity value")
	}
	if rel, err = strconv.ParseUint(ss[2], 10, 32); err != nil {
		return errors.New("Could not parse relative intensity")
	}
	p.rel = uint(rel)
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
		tag := field.Tag.Get("mb2")
		subtag := field.Tag.Get("mb2st")
		if subtag != "" {
			subtag = ":" + subtag
		}
		TagMap[tag] = props
	}
}
