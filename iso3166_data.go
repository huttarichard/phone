package phone

import (
	c "github.com/huttarichard/phone/countries"
)

var (
	// ISO3166Data ...
	ISO3166Data = []*Data{
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUnitedStates],
			MobileBeginsWith:   []string{"201", "202", "203", "205", "206", "207", "208", "209", "210", "212", "213", "214", "215", "216", "217", "218", "219", "224", "225", "227", "228", "229", "231", "234", "239", "240", "248", "251", "252", "253", "254", "256", "260", "262", "267", "269", "270", "272", "274", "276", "278", "281", "283", "301", "302", "303", "304", "305", "307", "308", "309", "310", "312", "313", "314", "315", "316", "317", "318", "319", "320", "321", "323", "325", "327", "330", "331", "334", "336", "337", "339", "341", "346", "347", "351", "352", "360", "361", "364", "369", "380", "385", "386", "401", "402", "404", "405", "406", "407", "408", "409", "410", "412", "413", "414", "415", "417", "419", "423", "424", "425", "430", "432", "434", "435", "440", "442", "443", "445", "447", "458", "464", "469", "470", "475", "478", "479", "480", "484", "501", "502", "503", "504", "505", "507", "508", "509", "510", "512", "513", "515", "516", "517", "518", "520", "530", "531", "534", "539", "540", "541", "551", "557", "559", "561", "562", "563", "564", "567", "570", "571", "573", "574", "575", "580", "582", "585", "586", "601", "602", "603", "605", "606", "607", "608", "609", "610", "612", "614", "615", "616", "617", "618", "619", "620", "623", "626", "627", "628", "630", "631", "636", "641", "646", "650", "651", "657", "659", "660", "661", "662", "667", "669", "678", "679", "681", "682", "689", "701", "702", "703", "704", "706", "707", "708", "712", "713", "714", "715", "716", "717", "718", "719", "720", "724", "725", "727", "730", "731", "732", "734", "737", "740", "747", "752", "754", "757", "760", "762", "763", "764", "765", "769", "770", "772", "773", "774", "775", "779", "781", "785", "786", "801", "802", "803", "804", "805", "806", "808", "810", "812", "813", "814", "815", "816", "817", "818", "828", "830", "831", "832", "835", "843", "845", "847", "848", "850", "856", "857", "858", "859", "860", "862", "863", "864", "865", "870", "872", "878", "901", "903", "904", "906", "907", "908", "909", "910", "912", "913", "914", "915", "916", "917", "918", "919", "920", "925", "927", "928", "929", "931", "935", "936", "937", "938", "940", "941", "947", "949", "951", "952", "954", "956", "957", "959", "970", "971", "972", "973", "975", "978", "979", "980", "984", "985", "989"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAruba],
			MobileBeginsWith:   []string{"5", "6", "7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAfghanistan],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAngola],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAnguilla],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAlandIslands],
			MobileBeginsWith:   []string{"18"},
			PhoneNumberLengths: []int{6, 7, 8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAlbania],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAndorra],
			MobileBeginsWith:   []string{"3", "4", "6"},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUnitedArabEmirates],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameArgentina],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameArmenia],
			MobileBeginsWith:   []string{"5", "7", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAmericanSamoa],
			MobileBeginsWith:   []string{"733", "258"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAntiguaandBarbuda],
			MobileBeginsWith:   []string{"4", "7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAustralia],
			MobileBeginsWith:   []string{"4"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAustria],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{10, 12},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAzerbaijan],
			MobileBeginsWith:   []string{"4", "5", "6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBurundi],
			MobileBeginsWith:   []string{"7", "29"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBelgium],
			MobileBeginsWith:   []string{"4"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBenin],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBurkinaFaso],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBangladesh],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{8, 9, 10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBulgaria],
			MobileBeginsWith:   []string{"87", "88", "89", "98", "99", "43"},
			PhoneNumberLengths: []int{8, 9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBahrain],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBahamas],
			MobileBeginsWith:   []string{"3", "4", "5", "6", "7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBosniaandHerzegovina],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBelarus],
			MobileBeginsWith:   []string{"25", "29", "33", "44"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBelize],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBermuda],
			MobileBeginsWith:   []string{"3", "5", "7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBolivia],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBrazil],
			MobileBeginsWith:   []string{"119", "129", "139", "149", "159", "169", "179", "189", "199", "219", "229", "249", "279", "289", "31", "32", "34", "38", "41", "43", "44", "45", "47", "48", "51", "53", "54", "55", "61", "62", "65", "67", "68", "69", "71", "73", "75", "77", "79", "81", "82", "83", "84", "85", "86", "91", "92", "95", "96", "98"},
			PhoneNumberLengths: []int{10, 11},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBarbados],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBruneiDarussalam],
			MobileBeginsWith:   []string{"7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBhutan],
			MobileBeginsWith:   []string{"17"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameBotswana],
			MobileBeginsWith:   []string{"71", "72", "73", "74", "75", "76"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCentralAfricanRepublic],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCanada],
			MobileBeginsWith:   []string{"204", "226", "236", "249", "250", "289", "306", "343", "365", "403", "416", "418", "431", "437", "438", "450", "506", "514", "519", "579", "581", "587", "600", "604", "613", "639", "647", "705", "709", "778", "780", "807", "819", "867", "873", "902", "905"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSwitzerland],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameChile],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameChina],
			MobileBeginsWith:   []string{"13", "14", "15", "18"},
			PhoneNumberLengths: []int{11},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCoteDIvoire],
			MobileBeginsWith:   []string{"0", "4", "5", "6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCameroon],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCongo],
			MobileBeginsWith:   []string{"8", "9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCongo],
			MobileBeginsWith:   []string{"0"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCookIslands],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{5},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameColombia],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameComoros],
			MobileBeginsWith:   []string{"3", "76"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCapeVerde],
			MobileBeginsWith:   []string{"5", "9"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCostaRica],
			MobileBeginsWith:   []string{"5", "6", "7", "8"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCuba],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCaymanIslands],
			MobileBeginsWith:   []string{"345"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCyprus],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCzechRepublic],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGermany],
			MobileBeginsWith:   []string{"15", "16", "17"},
			PhoneNumberLengths: []int{10, 11},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameDjibouti],
			MobileBeginsWith:   []string{"77"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameDominica],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameDenmark],
			MobileBeginsWith:   []string{"2", "30", "31", "40", "41", "42", "50", "51", "52", "53", "60", "61", "71", "81", "91", "92", "93"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameDominicanRepublic],
			MobileBeginsWith:   []string{"809", "829", "849"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameAlgeria],
			MobileBeginsWith:   []string{"5", "6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameEcuador],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameEgypt],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameEritrea],
			MobileBeginsWith:   []string{"1", "7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSpain],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameEstonia],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameEthiopia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFinland],
			MobileBeginsWith:   []string{"4", "5"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFiji],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFalklandIslands],
			MobileBeginsWith:   []string{"5", "6"},
			PhoneNumberLengths: []int{5},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFrance],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFaroeIslands],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMicronesiaFederatedStatesOf],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGabon],
			MobileBeginsWith:   []string{"05", "06", "07"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUnitedKingdom],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGeorgia],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGhana],
			MobileBeginsWith:   []string{"2", "5"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGibraltar],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGuinea],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGuadeloupe],
			MobileBeginsWith:   []string{"690"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGambia],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGuineaBissau],
			MobileBeginsWith:   []string{"5", "6", "7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameEquatorialGuinea],
			MobileBeginsWith:   []string{"222", "551"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGreece],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGrenada],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGreenland],
			MobileBeginsWith:   []string{"4", "5"},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGuatemala],
			MobileBeginsWith:   []string{"3", "4", "5"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFrenchGuiana],
			MobileBeginsWith:   []string{"694"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGuam],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameGuyana],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameHongKong],
			MobileBeginsWith:   []string{"5", "6", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameHonduras],
			MobileBeginsWith:   []string{"3", "7", "8", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCroatia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameHaiti],
			MobileBeginsWith:   []string{"3", "4"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameHungary],
			MobileBeginsWith:   []string{"20", "30", "31", "70"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIndonesia],
			MobileBeginsWith:   []string{"8"},
			PhoneNumberLengths: []int{9, 10, 11},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIndia],
			MobileBeginsWith:   []string{"7", "8", "9"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIreland],
			MobileBeginsWith:   []string{"82", "83", "84", "85", "86", "87", "88", "89"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIran],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIraq],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIceland],
			MobileBeginsWith:   []string{"6", "7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameIsrael],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameItaly],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameJamaica],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameJordan],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameJapan],
			MobileBeginsWith:   []string{"70", "80", "90"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameKazakhstan],
			MobileBeginsWith:   []string{"70", "77"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameKenya],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameKyrgyzstan],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameCambodia],
			MobileBeginsWith:   []string{"1", "6", "7", "8", "9"},
			PhoneNumberLengths: []int{8, 9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameKiribati],
			MobileBeginsWith:   []string{"9", "30"},
			PhoneNumberLengths: []int{5},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaintKittsAndNevis],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameKorea],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{9, 10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameKuwait],
			MobileBeginsWith:   []string{"5", "6", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLao],
			MobileBeginsWith:   []string{"20"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLebanon],
			MobileBeginsWith:   []string{"3", "7"},
			PhoneNumberLengths: []int{7, 8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLiberia],
			MobileBeginsWith:   []string{"4", "5", "6", "7"},
			PhoneNumberLengths: []int{7, 8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLibyanArabJamahiriya],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaintLucia],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLiechtenstein],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSriLanka],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLesotho],
			MobileBeginsWith:   []string{"5", "6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLithuania],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLuxembourg],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameLatvia],
			MobileBeginsWith:   []string{"2"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMacao],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMorocco],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMonaco],
			MobileBeginsWith:   []string{"4", "6"},
			PhoneNumberLengths: []int{8, 9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMoldova],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMadagascar],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMaldives],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMexico],
			MobileBeginsWith:   []string{""},
			PhoneNumberLengths: []int{10, 11},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMarshallIslands],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMacedonia],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMali],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMalta],
			MobileBeginsWith:   []string{"79", "99"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMyanmar],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMontenegro],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMongolia],
			MobileBeginsWith:   []string{"5", "8", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNorthernMarianaIslands],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMozambique],
			MobileBeginsWith:   []string{"8"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMauritania],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMontserrat],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMartinique],
			MobileBeginsWith:   []string{"696"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMauritius],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMalawi],
			MobileBeginsWith:   []string{"77", "88", "99"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMalaysia],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{9, 10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameMayotte],
			MobileBeginsWith:   []string{"639"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNamibia],
			MobileBeginsWith:   []string{"60", "81", "82", "85"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNewCaledonia],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNiger],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNorfolkIsland],
			MobileBeginsWith:   []string{"5", "8"},
			PhoneNumberLengths: []int{5},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNigeria],
			MobileBeginsWith:   []string{"70", "80", "81"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNicaragua],
			MobileBeginsWith:   []string{"8"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNiue],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{4},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNetherlands],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNorway],
			MobileBeginsWith:   []string{"4", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNepal],
			MobileBeginsWith:   []string{"97", "98"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNauru],
			MobileBeginsWith:   []string{"555"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameNewZealand],
			MobileBeginsWith:   []string{"2"},
			PhoneNumberLengths: []int{8, 9, 10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameOman],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePakistan],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePanama],
			MobileBeginsWith:   []string{"5", "6"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePeru],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePhilippines],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePalau],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePapuaNewGuinea],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePoland],
			MobileBeginsWith:   []string{"5", "6", "7", "8"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePuertoRico],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePortugal],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameParaguay],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNamePalestinia],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameFrenchPolynesia],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameQatar],
			MobileBeginsWith:   []string{"33", "55", "66", "77"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameReunion],
			MobileBeginsWith:   []string{"692", "693"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameRomania],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameRussianFederation],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameRwanda],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaudiArabia],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSudan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSenegal],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSingapore],
			MobileBeginsWith:   []string{"8", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaintHelena],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{4},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSvalbardAndJanMayen],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSolomonIslands],
			MobileBeginsWith:   []string{"7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSierraLeone],
			MobileBeginsWith:   []string{"21", "25", "30", "33", "34", "40", "44", "50", "55", "76", "77", "78", "79", "88"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameElSalvador],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSanMarino],
			MobileBeginsWith:   []string{"3", "6"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSomalia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaintPierreAndMiquelon],
			MobileBeginsWith:   []string{"55"},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSerbia],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8, 9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaoTomeandPrincipe],
			MobileBeginsWith:   []string{"98", "99"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSuriname],
			MobileBeginsWith:   []string{"6", "7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSlovakia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSlovenia],
			MobileBeginsWith:   []string{"3", "4", "5", "6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSweden],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSeychelles],
			MobileBeginsWith:   []string{"2"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSyrianArabRepublic],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTurksAndCaicosIslands],
			MobileBeginsWith:   []string{"2", "3", "4"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameChad],
			MobileBeginsWith:   []string{"6", "7", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTogo],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameThailand],
			MobileBeginsWith:   []string{"8", "9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTajikistan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTokelau],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{4},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTurkmenistan],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTimorLeste],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTonga],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{5},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTrinidadAndTobago],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTunisia],
			MobileBeginsWith:   []string{"2", "9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTurkey],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTuvalu],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{5},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTaiwan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameTanzania],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUganda],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUkraine],
			MobileBeginsWith:   []string{"39", "50", "63", "66", "67", "68", "9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUruguay],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameUzbekistan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSaintVincentAndTheGrenedines],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameVenezuela],
			MobileBeginsWith:   []string{"4"},
			PhoneNumberLengths: []int{10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameVirginIslandsBritish],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameVirginIslandsUS],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameVietNam],
			MobileBeginsWith:   []string{"9", "1"},
			PhoneNumberLengths: []int{9, 10},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameVanuatu],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameWallisAndFutuna],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSamoa],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{7},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameYemen],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameSouthAfrica],
			MobileBeginsWith:   []string{"6", "7", "8"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameZambia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&Data{
			CountryData:        c.ISO3166CountriesData[c.CountryNameZimbabwe],
			MobileBeginsWith:   []string{"71", "73", "77"},
			PhoneNumberLengths: []int{9},
		},
	}
)
