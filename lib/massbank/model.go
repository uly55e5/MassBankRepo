package massbank

import (
	"time"
)

type Massbank struct {
	Accession          *RecordAccession     `mb:"ACCESSION"`
	Deprecated         *RecordDeprecated    `mb:"DEPRECATED"`
	RecordTitle        *RecordTitle         `mb:"RECORD_TITLE"`
	Date               *RecordDate          `mb:"DATE"`
	Authors            *RecordAuthorNames   `mb:"AUTHORS"`
	License            *RecordLicense       `mb:"LICENSE"`
	Copyright          *RecordCopyright     `mb:"COPYRIGHT"`
	Publication        *RecordPublication   `mb:"PUBLICATION"`
	Project            *RecordProject       `mb:"PROJECT"`
	Comments           []*RecordComment     `mb:"COMMENT"`
	ChNames            []*ChName            `mb:"CH$NAME"`
	ChClass            *ChCompoundClasses   `mb:"CH$COMPOUND_CLASS"`
	ChFormula          *ChFormula           `mb:"CH$FORMULA"`
	ChCdkDepict        []*CdkDepict         `mb:"CH$CDK_DEPICT"` // not for productive use
	ChMass             *ChMass              `mb:"CH$EXACT_MASS"`
	ChSmiles           *ChSmiles            `mb:"CH$SMILES"`
	ChInchi            *ChInchi             `mb:"CH$IUPAC"`
	ChLink             []*ChLink            `mb:"CH$LINK"`
	SpName             SpName               `mb:"SP$SCIENTIFIC_NAME"`
	SpLineage          SpLineage            `mb:"SP$LINEAGE"`
	SpLink             []SpLink             `mb:"SP$LINK"`
	SpSample           []SampleInformation  `mb:"SP$SAMPLE"`
	AcInstrument       AcInstrument         `mb:"AC$INSTRUMENT"`
	AcInstrumentType   AcInstrumentType     `mb:"AC$INSTRUMENT_TYPE	"`
	AcMassSpectrometry []AcMassSpectrometry `mb:"AC$MASS_SPECTROMETRY"`
	AcChromatography   []AcChromatography   `mb:"AC$CHROMATOGRAPHY"`
	AcGeneral          []AcGeneral          `mb:"AC$GENERAL"`
	MsFocusedIon       []MsFocusedIon       `mb:"MS$FOCUSED_ION"`
	MsDataProcessing   []MsDataProcessing   `mb:"MS$DATA_PROCESSING"`
	PkSplash           PkSplash             `mb:"PK$SPLASH"`
	PkAnnotation       PkAnnotation         `mb:"PK$ANNOTATION"`
	PkNumPeak          PkNumPeak            `mb:"PK$NUM_PEAK"`
	PkPeak             PkPeak               `mb:"PK$PEAK"`
}

type RecordAccession string

type RecordDeprecated struct {
	Date   time.Time
	Reason string
}

type RecordTitle struct {
	ChName           ChName
	AcInstrumentType AcInstrumentType
	MsType           MsType
}

type RecordDate struct {
	Updated time.Time
	Created time.Time
}
type RecordAuthorNames []RecordAuthorName

type RecordAuthorName struct {
	Name        string
	MarcRelator string
}

type RecordLicense string

type RecordCopyright string

type RecordPublication string

type RecordProject string

type RecordComment struct {
	Subtag  RecordSubtag
	MbTag   RecordMbTag
	Comment string
}

type RecordSubtag string

type RecordMbTag string

type ChName string

type ChCompoundClasses []ChCompoundClass
type ChCompoundClass string

type ChFormula string

type ChMass float64

type ChSmiles string

type ChInchi string

type ChLink struct {
	Database ExtDatabase
	Link     string
}

type ExtDatabase string

type CdkDepict string

type SpName string

type SpLineage []SpLineageElement

type SpLineageElement string

type SpLink struct {
	Database   string
	Identifier string
}

type SampleInformation string

type AcInstrument string

type Separation string
type Ionization string
type Analyzer string
type AcInstrumentType struct {
	Separation  Separation
	Ionization  Ionization
	IonAnalyzer Analyzer
}
type MsType string
type AcMassSpectrometry struct {
	MsType MsType
	Subtag string
	Value  string
}

type AcChromatography struct {
	Subtag string
	Value  string
}

type AcGeneral struct {
	Subtag string
	Value  string
}

type PkPeak struct {
	Header []string
	Values []PeakValue
}

type MsFocusedIon struct {
	Subtag string
	Value  string
}
type MsDataProcessing struct {
	Subtag string
	Value  string
}

type PkSplash string

type AnnotationValue []interface{}

type PkAnnotation struct {
	Header []string
	Values []AnnotationValue
}

type PkNumPeak uint

type PeakValue struct {
	mz        float64
	intensity float64
	rel       uint
}
