package codes

import (
	"fmt"
	"sort"
)

func init() {
	CountriesCodes = make([]string, 0)
	for k := range Countries {
		CountriesCodes = append(CountriesCodes, k)
	}
	sort.Strings(CountriesCodes)
}

var CountriesCodes []string

var Countries = map[string]string{
	"AD": "Andorra",
	"AE": "Utd.Arab Emir.",
	"AF": "Afghanistan",
	"AG": "Antigua/Barbuda",
	"AI": "Anguilla",
	"AL": "Albania",
	"AM": "Armenia",
	"AN": "Dutch Antilles",
	"AO": "Angola",
	"AQ": "Antarctica",
	"AR": "Argentina",
	"AS": "Samoa, American",
	"AT": "Austria",
	"AU": "Australia",
	"AW": "Aruba",
	"AX": "Aland Islands",
	"AZ": "Azerbaijan",
	"BA": "Bosnia-Herz.",
	"BB": "Barbados",
	"BD": "Bangladesh",
	"BE": "Belgium",
	"BF": "Burkina-Faso",
	"BG": "Bulgaria",
	"BH": "Bahrain",
	"BI": "Burundi",
	"BJ": "Benin",
	"BL": "St. Barthélemy",
	"BM": "Bermuda",
	"BN": "Brunei Dar-es-S",
	"BO": "Bolivia",
	"BQ": "Bonaire",
	"BR": "Brazil",
	"BS": "Bahamas",
	"BT": "Bhutan",
	"BV": "Bouvet Island",
	"BW": "Botswana",
	"BY": "Belarus",
	"BZ": "Belize",
	"CA": "Canada",
	"CC": "Coconut Islands",
	"CD": "Republic Congo",
	"CF": "Central Afr.Rep",
	"CG": "Congo",
	"CH": "Switzerland",
	"CI": "Ivory Coast",
	"CK": "Cook Islands",
	"CL": "Chile",
	"CM": "Cameroon",
	"CN": "China",
	"CO": "Colombia",
	"CR": "Costa Rica",
	"CU": "Cuba",
	"CV": "Cape Verde",
	"CW": "Curacao",
	"CX": "Christmas Islnd",
	"CY": "Cyprus",
	"CZ": "Czech Republic",
	"DE": "Germany",
	"DJ": "Djibouti",
	"DK": "Denmark",
	"DM": "Dominica",
	"DO": "Dominican Rep.",
	"DZ": "Algeria",
	"EC": "Ecuador",
	"EE": "Estonia",
	"EG": "Egypt",
	"EH": "West Sahara",
	"ER": "Eritrea",
	"ES": "Spain",
	"ET": "Ethiopia",
	"FI": "Finland",
	"FJ": "Fiji",
	"FK": "Falkland Islnds",
	"FM": "Micronesia",
	"FO": "Faroe Islands",
	"FR": "France",
	"GA": "Gabon",
	"GB": "United Kingdom",
	"GD": "Grenada",
	"GE": "Georgia",
	"GF": "French Guayana",
	"GG": "Guernsey",
	"GH": "Ghana",
	"GI": "Gibraltar",
	"GL": "Greenland",
	"GM": "Gambia",
	"GN": "Guinea",
	"GP": "Guadeloupe",
	"GQ": "Equatorial Gui.",
	"GR": "Greece",
	"GS": "S. Sandwich Ins",
	"GT": "Guatemala",
	"GU": "Guam",
	"GW": "Guinea-Bissau",
	"GY": "Guyana",
	"HK": "Hong Kong",
	"HM": "Heard/McDon.Isl",
	"HN": "Honduras",
	"HR": "Croatia",
	"HT": "Haiti",
	"HU": "Hungary",
	"ID": "Indonesia",
	"IE": "Ireland",
	"IL": "Israel",
	"IM": "Isle of Man",
	"IN": "India",
	"IO": "Brit.Ind.Oc.Ter",
	"IQ": "Iraq",
	"IR": "Iran",
	"IS": "Iceland",
	"IT": "Italy",
	"JE": "Jersey",
	"JM": "Jamaica",
	"JO": "Jordan",
	"JP": "Japan",
	"KE": "Kenya",
	"KG": "Kyrgyzstan",
	"KH": "Cambodia",
	"KI": "Kiribati",
	"KM": "Comoros",
	"KN": "St Kitts&Nevis",
	"KP": "North Korea",
	"KR": "South Korea",
	"KW": "Kuwait",
	"KY": "Cayman Islands",
	"KZ": "Kazakhstan",
	"LA": "Laos",
	"LB": "Lebanon",
	"LC": "St. Lucia",
	"LI": "Liechtenstein",
	"LK": "Sri Lanka",
	"LR": "Liberia",
	"LS": "Lesotho",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"LV": "Latvia",
	"LY": "Libya",
	"MA": "Morocco",
	"MC": "Monaco",
	"MD": "Moldova",
	"ME": "Montenegro",
	"MF": "St. Martin",
	"MG": "Madagascar",
	"MH": "Marshall Islnds",
	"MK": "Macedonia",
	"ML": "Mali",
	"MM": "Myanmar",
	"MN": "Mongolia",
	"MO": "Macau",
	"MP": "N.Mariana Islnd",
	"MQ": "Martinique",
	"MR": "Mauretania",
	"MS": "Montserrat",
	"MT": "Malta",
	"MU": "Mauritius",
	"MV": "Maldives",
	"MW": "Malawi",
	"MX": "Mexico",
	"MY": "Malaysia",
	"MZ": "Mozambique",
	"NA": "Namibia",
	"NC": "New Caledonia",
	"NE": "Niger",
	"NF": "Norfolk Island",
	"NG": "Nigeria",
	"NI": "Nicaragua",
	"NL": "Netherlands",
	"NO": "Norway",
	"NP": "Nepal",
	"NR": "Nauru",
	"NU": "Niue Islands",
	"NZ": "New Zealand",
	"OM": "Oman",
	"PA": "Panama",
	"PE": "Peru",
	"PF": "Frenc.Polynesia",
	"PG": "Papua Nw Guinea",
	"PH": "Philippines",
	"PK": "Pakistan",
	"PL": "Poland",
	"PM": "St.Pier,Miquel.",
	"PN": "Pitcairn Islnds",
	"PR": "Puerto Rico",
	"PS": "Palestine State",
	"PT": "Portugal",
	"PW": "Palau",
	"PY": "Paraguay",
	"QA": "Qatar",
	"RE": "Reunion",
	"RO": "Romania",
	"RS": "Serbia",
	"RU": "Russian Fed.",
	"RW": "Rwanda",
	"SA": "Saudi Arabia",
	"SB": "Solomon Islands",
	"SC": "Seychelles",
	"SD": "Sudan",
	"SE": "Sweden",
	"SG": "Singapore",
	"SH": "St. Helena",
	"SI": "Slovenia",
	"SJ": "Svalbard",
	"SK": "Slovak Republic",
	"SL": "Sierra Leone",
	"SM": "San Marino",
	"SN": "Senegal",
	"SO": "Somalia",
	"SR": "Suriname",
	"SS": "South Sudan",
	"ST": "S.Tome,Principe",
	"SV": "El Salvador",
	"SX": "Sint Maarten",
	"SY": "Syria",
	"SZ": "Swaziland",
	"TC": "Turksh Caicosin",
	"TD": "Chad",
	"TF": "French S.Territ",
	"TG": "Togo",
	"TH": "Thailand",
	"TJ": "Tajikistan",
	"TK": "Tokelau Islands",
	"TL": "Timor-Leste",
	"TM": "Turkmenistan",
	"TN": "Tunisia",
	"TO": "Tonga",
	"TP": "USE TL",
	"TR": "Turkey",
	"TT": "Trinidad,Tobago",
	"TV": "Tuvalu",
	"TW": "Taiwan",
	"TZ": "Tanzania",
	"UA": "Ukraine",
	"UG": "Uganda",
	"UM": "Minor Outl.Ins.",
	"US": "USA",
	"UY": "Uruguay",
	"UZ": "Uzbekistan",
	"VA": "Vatican City",
	"VC": "St. Vincent",
	"VE": "Venezuela",
	"VG": "Brit.Virgin Is.",
	"VI": "Amer.Virgin Is.",
	"VN": "Vietnam",
	"VU": "Vanuatu",
	"WF": "Wallis,Futuna",
	"WS": "Western Samoa",
	"XK": "USE RS-KL",
	"YE": "Yemen",
	"YT": "Mayotte",
	"ZA": "South Africa",
	"ZM": "Zambia",
	"ZW": "Zimbabwe",
}

func CountriesToMarkdown() string {

	res := "Code | Country name\n" +
		":---:|:------------\n"

	for _, i := range CountriesCodes {
		res += fmt.Sprintf("%-5v| %s\n", i, Countries[i])
	}

	return res
}
