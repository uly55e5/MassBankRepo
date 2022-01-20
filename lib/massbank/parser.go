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

func (cc *ChCompoundClasses) Parse(s string) error {
	ss := strings.Split(s, ";")
	for _, s1 := range ss {
		var c = ChCompoundClass(strings.TrimSpace(s1))
		cc.value = append(cc.value, c)
	}
	return nil
}

func (mass *ChMass) Parse(s string) error {
	var err error
	mass.value, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	return nil
}

func (n *PkNumPeak) Parse(s string) error {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	n.Value = uint(val)
	return nil
}

func (p *PkPeak) Parse(s string) error {
	if s != "m/z int. rel.int." {
		return errors.New("PK$ is not valid")
	}
	p.Header = strings.Split(s, " ")
	return nil
}

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
					mb.Peak.Peak.Values = append(mb.Peak.Peak.Values, pv)
				}
			} else {
				println("not implemented", line)
			}
		} else {
			s := strings.SplitN(line, ":", 2)
			if len(s) == 2 {
				mb.addValue(strings.TrimSpace(s[0]), strings.TrimSpace(s[1]))
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

func (mb *Massbank) addValue(tagname string, value string) error {
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
	err := propInt.Parse(value)
	lastTag = tagname
	if err != nil {
		println(err.Error(), tagname)
	}
	if prop.Kind() == reflect.Slice {
		prop.Set(reflect.Append(prop, newPro))
	} else {
		prop.Set(newPro)
	}
	return nil
}
