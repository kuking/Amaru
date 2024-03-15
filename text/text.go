package text

import (
	"fmt"
	"github.com/abadojack/whatlanggo"
	"github.com/kljensen/snowball"
	"github.com/rivo/uniseg"
	"net/url"
	"regexp"
	"strings"
)

var patterns []string
var stopWordsReplacer *strings.Replacer

func init() {
	patterns = []string{
		"abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", // Reference
		"ⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩ ⒶⒷⒸⒹⒺⒻⒼⒽⒾⒿⓀⓁⓂⓃⓄⓅⓆⓇⓈⓉⓊⓋⓌⓍⓎⓏ⓪①②③④⑤⑥⑦⑧⑨",
		"🅐🅑🅒🅓🅔🅕🅖🅗🅘🅙🅚🅛🅜🅝🅞🅟🅠🅡🅢🅣🅤🅥🅦🅧🅨🅩 🅐🅑🅒🅓🅔🅕🅖🅗🅘🅙🅚🅛🅜🅝🅞🅟🅠🅡🅢🅣🅤🅥🅦🅧🅨🅩⓿❶❷❸❹❺❻❼❽❾",
		"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ　ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ０１２３４５６７８９",
		"ᴀʙᴄᴅᴇꜰɢʜɪᴊᴋʟᴍɴᴏᴩqʀꜱᴛᴜᴠᴡxyᴢ ᴀʙᴄᴅᴇꜰɢʜɪᴊᴋʟᴍɴᴏᴩQʀꜱᴛᴜᴠᴡxYᴢ0123456789",
		"𝗮𝗯𝗰𝗱𝗲𝗳𝗴𝗵𝗶𝗷𝗸𝗹𝗺𝗻𝗼𝗽𝗾𝗿𝘀𝘁𝘂𝘃𝘄𝘅𝘆𝘇 𝗔𝗕𝗖𝗗𝗘𝗙𝗚𝗛𝗜𝗝𝗞𝗟𝗠𝗡𝗢𝗣𝗤𝗥𝗦𝗧𝗨𝗩𝗪𝗫𝗬𝗭𝟬𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵",
		"𝙖𝙗𝙘𝙙𝙚𝙛𝙜𝙝𝙞𝙟𝙠𝙡𝙢𝙣𝙤𝙥𝙦𝙧𝙨𝙩𝙪𝙫𝙬𝙭𝙮𝙯 𝘼𝘽𝘾𝘿𝙀𝙁𝙂𝙃𝙄𝙅𝙆𝙇𝙈𝙉𝙊𝙋𝙌𝙍𝙎𝙏𝙐𝙑𝙒𝙓𝙔𝙕0123456789",
		"𝘢𝘣𝘤𝘥𝘦𝘧𝘨𝘩𝘪𝘫𝘬𝘭𝘮𝘯𝘰𝘱𝘲𝘳𝘴𝘵𝘶𝘷𝘸𝘹𝘺𝘻 𝘈𝘉𝘊𝘋𝘌𝘍𝘎𝘏𝘐𝘑𝘒𝘓𝘔𝘕𝘖𝘗𝘘𝘙𝘚𝘛𝘜𝘝𝘞𝘟𝘠𝘡0123456789",
		"𝑎𝑏𝑐𝑑𝑒𝑓𝑔ℎ𝑖𝑗𝑘𝑙𝑚𝑛𝑜𝑝𝑞𝑟𝑠𝑡𝑢𝑣𝑤𝑥𝑦𝑧 𝐴𝐵𝐶𝐷𝐸𝐹𝐺𝐻𝐼𝐽𝐾𝐿𝑀𝑁𝑂𝑃𝑄𝑅𝑆𝑇𝑈𝑉𝑊𝑋𝑌𝑍0123456789",
		"𝖺𝖻𝖼𝖽𝖾𝖿𝗀𝗁𝗂𝗃𝗄𝗅𝗆𝗇𝗈𝗉𝗊𝗋𝗌𝗍𝗎𝗏𝗐𝗑𝗒𝗓 𝖠𝖡𝖢𝖣𝖤𝖥𝖦𝖧𝖨𝖩𝖪𝖫𝖬𝖭𝖮𝖯𝖰𝖱𝖲𝖳𝖴𝖵𝖶𝖷𝖸𝖹𝟢𝟣𝟤𝟥𝟦𝟧𝟨𝟩𝟪𝟫",
		"𝕒𝕓𝕔𝕕𝕖𝕗𝕘𝕙𝕚𝕛𝕜𝕝𝕞𝕟𝕠𝕡𝕢𝕣𝕤𝕥𝕦𝕧𝕨𝕩𝕪𝕫 𝔸𝔹ℂ𝔻𝔼𝔽𝔾ℍ𝕀𝕁𝕂𝕃𝕄ℕ𝕆ℙℚℝ𝕊𝕋𝕌𝕍𝕎𝕏𝕐ℤ𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡",
		"𝚊𝚋𝚌𝚍𝚎𝚏𝚐𝚑𝚒𝚓𝚔𝚕𝚖𝚗𝚘𝚙𝚚𝚛𝚜𝚝𝚞𝚟𝚠𝚡𝚢𝚣 𝙰𝙱𝙲𝙳𝙴𝙵𝙶𝙷𝙸𝙹𝙺𝙻𝙼𝙽𝙾𝙿𝚀𝚁𝚂𝚃𝚄𝚅𝚆𝚇𝚈𝚉𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿",
		"𝐚𝐛𝐜𝐝𝐞𝐟𝐠𝐡𝐢𝐣𝐤𝐥𝐦𝐧𝐨𝐩𝐪𝐫𝐬𝐭𝐮𝐯𝐰𝐱𝐲𝐳 𝐀𝐁𝐂𝐃𝐄𝐅𝐆𝐇𝐈𝐉𝐊𝐋𝐌𝐍𝐎𝐏𝐐𝐑𝐒𝐓𝐔𝐕𝐖𝐗𝐘𝐙𝟎𝟏𝟐𝟑𝟒𝟓𝟔𝟕𝟖𝟗",
		"𝒂𝒃𝒄𝒅𝒆𝒇𝒈𝒉𝒊𝒋𝒌𝒍𝒎𝒏𝒐𝒑𝒒𝒓𝒔𝒕𝒖𝒗𝒘𝒙𝒚𝒛 𝑨𝑩𝑪𝑫𝑬𝑭𝑮𝑯𝑰𝑱𝑲𝑳𝑴𝑵𝑶𝑷𝑸𝑹𝑺𝑻𝑼𝑽𝑾𝑿𝒀𝒁0123456789",
		"𝔞𝔟𝔠𝔡𝔢𝔣𝔤𝔥𝔦𝔧𝔨𝔩𝔪𝔫𝔬𝔭𝔮𝔯𝔰𝔱𝔲𝔳𝔴𝔵𝔶𝔷 𝔄𝔅ℭ𝔇𝔈𝔉𝔊ℌℑ𝔍𝔎𝔏𝔐𝔑𝔒𝔓𝔔ℜ𝔖𝔗𝔘𝔙𝔚𝔛𝔜ℨ0123456789",
		"𝖆𝖇𝖈𝖉𝖊𝖋𝖌𝖍𝖎𝖏𝖐𝖑𝖒𝖓𝖔𝖕𝖖𝖗𝖘𝖙𝖚𝖛𝖜𝖝𝖞𝖟 𝕬𝕭𝕮𝕯𝕰𝕱𝕲𝕳𝕴𝕵𝕶𝕷𝕸𝕹𝕺𝕻𝕼𝕽𝕾𝕿𝖀𝖁𝖂𝖃𝖄𝖅0123456789",
		"⠁⠃⠉⠙⠑⠋⠛⠓⠊⠚⠅⠇⠍⠝⠕⠏⠟⠗⠎⠞⠥⠧⠺⠭⠽⠵ ⠁⠃⠉⠙⠑⠋⠛⠓⠊⠚⠅⠇⠍⠝⠕⠏⠟⠗⠎⠞⠥⠧⠺⠭⠽⠵⠚⠁⠃⠉⠙⠑⠋⠛⠓⠊",
		"ᵃᵇᶜᵈᵉᶠᵍʰᶦʲᵏˡᵐⁿᵒᵖᵠʳˢᵗᵘᵛʷˣʸᶻ ᴬᴮᶜᴰᴱᶠᴳᴴᴵᴶᴷᴸᴹᴺᴼᴾᵠᴿˢᵀᵁⱽᵂˣʸᶻ⁰¹²³⁴⁵⁶⁷⁸⁹",
		"𝒶𝒷𝒸𝒹𝑒𝒻𝑔𝒽𝒾𝒿𝓀𝓁𝓂𝓃𝑜𝓅𝓆𝓇𝓈𝓉𝓊𝓋𝓌𝓍𝓎𝓏 𝒜𝐵𝒞𝒟𝐸𝐹𝒢𝐻𝐼𝒥𝒦𝐿𝑀𝒩𝒪𝒫𝒬𝑅𝒮𝒯𝒰𝒱𝒲𝒳𝒴𝒵𝟢𝟣𝟤𝟥𝟦𝟧𝟨𝟩𝟪𝟫",
		"𝓪𝓫𝓬𝓭𝓮𝓯𝓰𝓱𝓲𝓳𝓴𝓵𝓶𝓷𝓸𝓹𝓺𝓻𝓼𝓽𝓾𝓿𝔀𝔁𝔂𝔃 𝓐𝓑𝓒𝓓𝓔𝓕𝓖𝓗𝓘𝓙𝓚𝓛𝓜𝓝𝓞𝓟𝓠𝓡𝓢𝓣𝓤𝓥𝓦𝓧𝓨𝓩0123456789",
		"🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉 🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣",
		"🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉 🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣",
		"a⃣b⃣c⃣d⃣e⃣f⃣g⃣h⃣i⃣j⃣k⃣l⃣m⃣n⃣o⃣p⃣q⃣r⃣s⃣t⃣u⃣v⃣w⃣x⃣y⃣z⃣ ⃣A⃣B⃣C⃣D⃣E⃣F⃣G⃣H⃣I⃣J⃣K⃣L⃣M⃣N⃣O⃣P⃣Q⃣R⃣S⃣T⃣U⃣V⃣W⃣X⃣Y⃣Z⃣0⃣1⃣2⃣3⃣4⃣5⃣6⃣7⃣8⃣9⃣",
	}

	var replacements []string
	for _, r := range ".,:;[]{}()|!?`''\"\\/=><+-_*~@£$€%^&#…︙＊⋆˘ʕ⊙ᴥ⊙ʔⁱᴗ͈ˬᴗ͈ʚɞ⊹" {
		replacements = append(replacements, string(r), " ")
	}
	stopWordsReplacer = strings.NewReplacer(replacements...)
}

func NormaliseFancyUnicodeToASCII(input string) string {
	var result strings.Builder
	gin := uniseg.NewGraphemes(input)
	for gin.Next() {
		done := false
		for p := 0; p < len(patterns); p++ {
			pin := uniseg.NewGraphemes(patterns[p])
			n := 0
			for pin.Next() {
				if gin.Str() == pin.Str() {
					result.WriteString(getNthGraphene(patterns[0], n))
					done = true
					break
				}
				n++
			}
			if done {
				break
			}
		}
		if !done {
			result.WriteString(gin.Str())
		}
	}
	return result.String()
}

func getNthGraphene(s string, n int) string {
	g := uniseg.NewGraphemes(s)
	c := 0
	for g.Next() {
		if c == n {
			return g.Str()
		}
		c++
	}
	return ""
}

func RemoveEmojis(str string) string {
	regex := `(?U)[\x{1F600}-\x{1F64F}]|` + // Emoticons
		`[\x{1F300}-\x{1F5FF}]|` + // Miscellaneous symbols and pictographs
		`[\x{1F680}-\x{1F6FF}]|` + // Transport and map symbols
		`[\x{1F700}-\x{1F77F}]|` + // Alchemical symbols
		`[\x{1F780}-\x{1F7FF}]|` + // Geometric Shapes Extended
		`[\x{1F800}-\x{1F8FF}]|` + // Supplemental Arrows-C
		`[\x{1F900}-\x{1F9FF}]|` + // Supplemental Symbols and Pictographs
		`[\x{1FA00}-\x{1FA6F}]|` + // Chess Symbols
		`[\x{1FA70}-\x{1FAFF}]|` + // Symbols and Pictographs Extended-A
		`[\x{0080}-\x{00FF}]|` + // Latin-1 supplements
		`[\x{2000}-\x{206F}]|` + // General Puntuations
		`[\x{2190}-\x{21FF}]|` + // Arrows
		`[\x{2300}-\x{23FF}]|` + //  Miscellaneous Technical block
		`[\x{25A0}-\x{25FF}]|` + //  Geometric Shapes
		`[\x{2600}-\x{26FF}]|` + // Miscellaneous symbols blocks
		`[\x{2B00}-\x{2BFF}]|` + // Miscellaneous Symbols and Arrow blocks
		`[\x{2700}-\x{27BF}]|` + // Dingbats
		`[\x{2900}-\x{297F}]|` + //Arrows B-Block
		`[\x{3000}-\x{303F}]|` + // CJF Symbols and puntuations
		`[\x{3200}-\x{32FF}]|` + // Enclosed CJK Letters
		`\x{200D}|` + // Zero Width Joiner
		`\x{FE0F}|` + // Variation Selector-16
		`[\x{0030}-\x{0039}]\x{FE0F}?\x{20E3}|` + // Keycap sequences (0️⃣ to 9️⃣)
		`[\x{1F000}-\x{1F02F}]|` + // Mahjong tiles
		`[\x{1F170}-\x{1F1FF}]|` + // Enclosed Alphanumeric Supplement
		`[\x{1F1E6}-\x{1F1FF}][\x{1F1E6}-\x{1F1FF}]|` + // Country Flags
		`[\x{1F200}-\x{1F2FF}]|` + //  Enclosed Ideographic Supplement
		`[\x{1F0A0}-\x{1F0FF}]` // Playing Cards

	emojiRx := regexp.MustCompile(regex)
	return emojiRx.ReplaceAllString(str, "")
}

func ReplaceStopWords(str string) string {
	return stopWordsReplacer.Replace(str)
}

func Stems(text string) []string {
	var res []string

	var lang string
	info := whatlanggo.Detect(text)
	switch info.Lang.Iso6391() {
	case "es":
		lang = "spanish"
	case "fr":
		lang = "french"
	case "hu":
		lang = "hungarian"
	case "ro":
		lang = "spanish" // rumanian as spanish
	case "ru":
		lang = "russian"
	case "sv":
		lang = "swedish"
	default:
		lang = "english"
	}
	if !info.IsReliable() {
		lang = "english"
	}

	for _, word := range strings.Fields(text) {
		if IsURL(word) {
			continue
		}
		if len(word) == 0 {
			continue
		}
		if stemmed, err := snowball.Stem(word, lang, true); err == nil {
			res = append(res, stemmed)
		} else {
			fmt.Print(err)
		}
	}
	return res
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
