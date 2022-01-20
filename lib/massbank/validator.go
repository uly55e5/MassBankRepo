package massbank

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func (mb *Massbank) Validate() (bool, error) {
	return validate(mb)
}

func validate(mb interface{}) (bool, error) {
	val := reflect.ValueOf(mb)
	t := reflect.TypeOf(mb)
	var fields = []reflect.StructField{}
	if t.Kind() == reflect.Struct {
		fields = reflect.VisibleFields(t)
	} else if t.Elem().Kind() == reflect.Struct {
		fields = reflect.VisibleFields(t.Elem())
		val = val.Elem()
	}
	b, err := validateFields(fields, val, t)
	if !b || err != nil {
		return b, err
	}

	return true, nil
}

func validateFields(fields []reflect.StructField, val reflect.Value, t reflect.Type) (bool, error) {
	for _, f := range fields {
		if f.Type.Kind() == reflect.Struct {
			res, err := validate(val.FieldByName(f.Name).Interface())
			if err != nil || !res {
				return res, err
			}
		} else {
			println("Checking ", f.Name)
			field := val.FieldByName(f.Name)
			_, isOptional := f.Tag.Lookup("optional")
			if field.Kind() == reflect.Slice {
				for i := 0; i < field.Len(); i++ {
					prop := field.Index(i).Interface().(Property)
					b, err := validateProperty(prop, isOptional, t.Name())
					if !b || err != nil {
						return b, err
					}
				}
			} else {
				prop := field.Interface().(Property)
				b, err := validateProperty(prop, isOptional, t.Name())
				if !b || err != nil {
					return b, err
				}
			}
		}
	}
	return true, nil
}

func validateProperty(prop Property, optional bool, name string) (bool, error) {
	if prop != nil && (reflect.ValueOf(prop).Kind() != reflect.Ptr || !reflect.ValueOf(prop).IsNil()) {
		res, err := prop.Validate()
		return res, err
	}
	if !optional {
		println("Non optional value is nil: ", name)
		return false, nil
	}
	return true, nil
}

func (sp *StringProperty) Validate() (bool, error) {
	println("no specialized validation implemented")
	if sp.string != "" {
		return true, nil
	}
	return false, nil
}

func (sp *DefaultProperty) Validate() (bool, error) {
	println("validation not implemented")
	return true, nil
}

func (acc *RecordAccession) Validate() (bool, error) {
	re := "[A-Z]{2-3}[0-9]*"
	matched, err := regexp.MatchString(re, acc.string)
	if err != nil {
		return false, errors.New("Could not validate Accessionnumber")
	}
	if len(acc.string) == 8 || matched {
		return true, nil
	}
	return false, nil
}

func (rt *RecordTitle) Validate() (bool, error) {
	if rt.string != "" {
		return true, nil
	}
	return false, nil
}

func (rd *RecordDate) Validate() (bool, error) {
	if (rd.Updated.Before(time.Now()) || rd.Updated.Equal(time.Now())) && (rd.Created.Before(rd.Updated) || rd.Created.Equal(rd.Updated)) {
		return true, nil
	}
	return false, nil
}

func (auth *RecordAuthorNames) Validate() (bool, error) {
	for _, a := range auth.value {
		ok := true
		if a.MarcRelator != "" {
			_, ok = MarcRelators[a.MarcRelator]
		}
		if strings.TrimSpace(a.Name) == "" || !ok {
			return false, nil
		}
	}
	return true, nil
}

func (lic *RecordLicense) Validate() (bool, error) {
	return contains(licenseList, lic.string), nil
}

func (cop *RecordCopyright) Validate() (bool, error) {
	return cop.string != "", nil
}

func (pub *RecordPublication) Validate() (bool, error) {
	if !strings.Contains(pub.string, "DOI") && !strings.Contains(pub.string, "PMID") {
		println("Publication does not contain PMID or DOI")
	}
	return pub.string != "", nil
}

func (proj *RecordProject) Validate() (bool, error) {
	return proj.string != "", nil

}

func (com *RecordComment) Validate() (bool, error) {
	re := "\\[.*\\]"
	isMBTag, err := regexp.MatchString(re, com.subtag)
	if err != nil {

	}
	if isMBTag {
		tag := com.subtag[1 : len(com.subtag)-1]
		if !contains(internalCommentTagList, tag) {
			return false, nil
		}
	} else if !isUpperCase(com.subtag) {
		if !contains(commentSubtagList, com.subtag) {
			return false, nil
		}
	}
	return com.string != "", nil
}

func (name *ChName) Validate() (bool, error) {
	return name.string != "", nil
}

func isUpperCase(s string) bool {
	return s == strings.ToUpper(s)
}

func (classes *ChCompoundClasses) Validate() (bool, error) {
	for _, class := range classes.value {
		ok, err := class.Validate()
		if !ok || err != nil {
			return ok, err
		}
	}
	return true, nil
}

func (class *ChCompoundClass) Validate() (bool, error) {
	return string(*class) != "", nil
}

func (form *ChFormula) Validate() (bool, error) {
	if form.string != "" {
		re := "[A-Z\\[][A-Za-z0-9+\\-()\\]\\[]*"
		return regexp.MatchString(re, form.string)
	}
	return false, nil
}

func (mass *ChMass) Validate() (bool, error) {
	return mass.value >= 0.0, nil
}

func (smiles *ChSmiles) Validate() (bool, error) {
	if smiles.string != "" {
		re := "([^J][A-Za-z0-9@+\\-\\[\\]\\(\\)\\\\=#$]+)"
		return regexp.MatchString(re, smiles.string)
	}
	return false, nil
}

func (inchi *ChInchi) Validate() (bool, error) {
	if inchi.string != "" {
		re := "((InChI=)?[^J][0-9a-z+\\-\\(\\)\\\\\\/,]+)"
		return regexp.MatchString(re, inchi.string)
	}
	return false, nil
}

func (link *ChLink) Validate() (bool, error) {
	if !isUpperCase(link.subtag) {
		return false, nil
	} else if contains(databaseList, link.subtag) {
		return link.string != "", nil
	}
	return false, nil
}

func (instr *AcInstrument) Validate() (bool, error) {
	return instr.string != "", nil
}

func (it *AcInstrumentType) Validate() (bool, error) {
	parts := strings.Split(it.string, "-")
	var ok = false
	if len(parts) == 2 {
		ok = isIonization(parts[0]) && isAnalyzer(parts[1])
	} else if len(parts) == 3 {
		ok = isSeparation(parts[0]) && isIonization(parts[1]) && isAnalyzer(parts[2])
	}
	return ok, nil
}

func isIonization(s string) bool {
	return contains(InstrumentTypeIonizationList, s)
}

func isAnalyzer(s string) bool {
	return contains(InstrumentTypeAnalyzerList, s) || contains(TandemAnalyzerList(), s)
}

func isSeparation(s string) bool {
	return contains(InstrumentTypeSepList, s)
}

func (ms *AcMassSpectrometry) Validate() (bool, error) {
	switch ms.subtag {
	case "MS_TYPE":
		return contains(MSTypeList, ms.string), nil
	case "ION_MODE":
		return contains(MSIonModeList, ms.string), nil
	default:
		return isUpperCase(ms.subtag) && ms.string != "", nil
	}
}
