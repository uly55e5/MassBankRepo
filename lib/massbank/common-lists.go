package massbank

var licenseList = []string{
	"CC0",
	"CC BY",
	"CC BY-NC",
	"CC BY-NC-SA",
	"CC BY-SA",
}

var internalCommentTagList = []string{
	"MS2",
	"Merging",
	"Merged",
	"Mass spectrometry",
	"Chromatography",
	"Profile",
}

var commentSubtagList = []string{
	"CONFIDENCE",
	"INTERNAL_ID",
}

var databaseList = []string{
	"CAS",
	"CAYMAN",
	"CHEBI",
	"CHEMBL",
	"CHEMPDB",
	"CHEMSPIDER",
	"COMPTOX",
	"HMDB",
	"INCHIKEY",
	"KAPPAVIEW",
	"KEGG",
	"KNAPSACK",
	"LIPIDBANK",
	"LIPIDMAPS",
	"NIKKAJI",
	"PUBCHEM",
	"ZINC",
}

var InstrumentTypeSepList = []string{
	"CE",
	"GC",
	"LC",
}

var InstrumentTypeIonizationList = []string{
	"APCI",
	"APPI",
	"EI",
	"ESI",
	"FAB",
	"MALDI",
	"FD",
	"CI",
	"FI",
	"SIMS",
}

var InstrumentTypeAnalyzerList = []string{
	"B",
	"E",
	"FT",
	"IT",
	"Q",
	"TOF",
}

var tandemAnalyzerList []string

func TandemAnalyzerList() []string {
	if len(tandemAnalyzerList) == 0 {
		for _, s := range InstrumentTypeAnalyzerList {
			for _, ss := range InstrumentTypeAnalyzerList {
				tandemAnalyzerList = append(tandemAnalyzerList, s+ss)
			}
		}
	}
	return tandemAnalyzerList
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var MSTypeList = []string{
	"MS",
	"MS2",
	"MS3",
	"MS4",
}

var MSIonModeList = []string{
	"POSITIVE",
	"NEGATIVE",
}
