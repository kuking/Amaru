package util

import (
	"strings"
)

var patterns []string

func init() {
	patterns = []string{
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", // Reference
		"𝕒𝕓𝕔𝕕𝕖𝕗𝕘𝕙𝕚𝕛𝕜𝕝𝕞𝕟𝕠𝕡𝕢𝕣𝕤𝕥𝕦𝕧𝕨𝕩𝕪𝕫𝔸𝔹ℂ𝔻𝔼𝔽𝔾ℍ𝕀𝕁𝕂𝕃𝕄ℕ𝕆ℙℚℝ𝕊𝕋𝕌𝕍𝕎𝕏𝕐ℤ𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡", // Double-struck
		"𝔞𝔟𝔠𝔡𝔢𝔣𝔤𝔥𝔦𝔧𝔨𝔩𝔪𝔫𝔬𝔭𝔮𝔯𝔰𝔱𝔲𝔳𝔴𝔵𝔶𝔷𝔄𝔅ℭ𝔇𝔈𝔉𝔊ℌℑ𝔍𝔎𝔏𝔐𝔑𝔒𝔓𝔔ℜ𝔖𝔗𝔘𝔙𝔚𝔛𝔜ℨ0123456789", // Fraktur
		//"𝖆𝖇𝖈𝖉𝖊𝖋𝖌𝖍𝖎𝖏𝖐𝖑𝖒𝖓𝖔𝖕𝖖𝖗𝖘𝖙𝖚𝖛𝖜𝖝𝖞𝖟 𝕬𝕭𝕮𝕯𝕰𝕱𝕲𝕳𝕴𝕵𝕶𝕷𝕸𝕹𝕺𝕻𝕼𝕽𝕾𝕿𝖀𝖁𝖂𝖃𝖄𝖅 0123456789", // Fraktur Bold
		//"ᵃᵇᶜᵈᵉᶠᵍʰᶦʲᵏˡᵐⁿᵒᵖᵠʳˢᵗᵘᵛʷˣʸᶻᴬᴮᶜᴰᴱᶠᴳᴴᴵᴶᴷᴸᴹᴺᴼᴾᵠᴿˢᵀᵁⱽᵂˣʸᶻ⁰¹²³⁴⁵⁶⁷⁸⁹", // superscript
		//"ᴀʙᴄᴅᴇꜰɢʜɪᴊᴋʟᴍɴᴏᴩqʀꜱᴛᴜᴠᴡxyᴢᴀʙᴄᴅᴇꜰɢʜɪᴊᴋʟᴍɴᴏᴩQʀꜱᴛᴜᴠᴡxYᴢ0123456789", // small-capitals
		//"🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮 🇯 🇰 🇱 🇲 🇳 🇴 🇵 🇶 🇷 🇸 🇹 🇺 🇻 🇼 🇽 🇾 🇿 🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮 🇯 🇰 🇱 🇲 🇳 🇴 🇵 🇶 🇷 🇸 🇹 🇺 🇻 🇼 🇽 🇾 🇿", // regionals
		//"🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉🄰🄱🄲🄳🄴🄵🄶🄷🄸🄹🄺🄻🄼🄽🄾🄿🅀🅁🅂🅃🅄🅅🅆🅇🅈🅉",                                                    // squared
		//"🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉🅰🅱🅲🅳🅴🅵🅶🅷🅸🅹🅺🅻🅼🅽🅾🅿🆀🆁🆂🆃🆄🆅🆆🆇🆈🆉",                                                    // filled-squared
		//"⠁⠃⠉⠙⠑⠋⠛⠓⠊⠚⠅⠇⠍⠝⠕⠏⠟⠗⠎⠞⠥⠧⠺⠭⠽⠵ ⠁⠃⠉⠙⠑⠋⠛⠓⠊⠚⠅⠇⠍⠝⠕⠏⠟⠗⠎⠞⠥⠧⠺⠭⠽⠵ ⠚⠁⠃⠉⠙⠑⠋⠛⠓⠊", // Braille

	}
}

func NormaliseFancyUnicodeToToASCII(input string) string {
	var result strings.Builder
	for _, char := range input {
		var base rune
		var reloc rune
		switch {
		case '𝐀' <= char && char <= '𝐙': // Bold
			base = '𝐀'
			reloc = 'A'
		case '𝐚' <= char && char <= '𝐳':
			base = '𝐚'
			reloc = 'a'
		case '𝐴' <= char && char <= '𝑍': // Italic
			base = '𝐴'
			reloc = 'A'
		case '𝑎' <= char && char <= '𝑧':
			base = '𝑎'
			reloc = 'a'
		// Handle bold-italic
		case '𝑨' <= char && char <= '𝒁': // Bold Italic
			base = '𝑨'
			reloc = 'A'
		case '𝒂' <= char && char <= '𝒛':
			base = '𝒂'
			reloc = 'a'
		case '𝖠' <= char && char <= '𝖹': // Sans-serif
			base = '𝖠'
			reloc = 'A'
		case '𝖺' <= char && char <= '𝗓':
			base = '𝖺'
			reloc = 'a'
		case '𝗔' <= char && char <= '𝗭': // Sans-serif bold
			base = '𝗔'
			reloc = 'A'
		case '𝗮' <= char && char <= '𝘇':
			base = '𝗮'
			reloc = 'a'
		case '𝘼' <= char && char <= '𝙕': // Sans-serif bold italic
			base = '𝘼'
			reloc = 'A'
		case '𝙖' <= char && char <= '𝙯':
			base = '𝙖'
			reloc = 'a'
		// Monospace uppercase
		case '𝙰' <= char && char <= '𝚉':
			base = '𝙰'
			reloc = 'A'
		// Monospace lowercase
		case '𝚊' <= char && char <= '𝚣':
			base = '𝚊'
			reloc = 'a'
		// Script cursive
		case '𝒶' <= char && char <= '𝓏':
			base = '𝒶'
			reloc = 'a'
		case '𝒜' <= char && char <= '𝒵':
			base = '𝒜'
			reloc = 'A'
		// Script cursive bold
		case '𝓪' <= char && char <= '𝔃':
			base = '𝓪'
			reloc = 'a'
		case '𝓐' <= char && char <= '𝓩':
			base = '𝓐'
			reloc = 'A'
		//Monospace Normal ????
		case 'a' <= char && char <= 'z':
			base = 'a'
			reloc = 'a'
		// circle
		case 'ⓐ' <= char && char <= 'ⓩ':
			base = 'ⓐ'
			reloc = 'a'
		case 'Ⓐ' <= char && char <= 'Ⓩ':
			base = 'Ⓐ'
			reloc = 'A'
		// filled circle
		case '🅐' <= char && char <= '🅩':
			base = '🅐'
			reloc = 'a'
		// double struck
		//case '𝕒' <= char && char <= '𝕫':
		//	base = '𝕒'
		//	reloc = 'a'
		//case '𝔸' <= char && char <= 'ℤ':
		//	base = '𝔸'
		//	reloc = 'A'
		// full-width
		case 'ａ' <= char && char <= 'ｚ':
			base = 'ａ'
			reloc = 'a'
		case 'Ａ' <= char && char <= 'Ｚ':
			base = 'Ａ'
			reloc = 'A'
		case '　' == char:
			base = '　'
			reloc = ' '
		// Mathematical block numbers bold
		case '𝟬' <= char && char <= '𝟵':
			base = '𝟬'
			reloc = '0'
		// Mathematical block numbers
		case '𝟶' <= char && char <= '𝟿':
			base = '𝟶'
			reloc = '0'
		// Mathematical double barred numbers
		case '𝟘' <= char && char <= '𝟡':
			base = '𝟘'
			reloc = '0'

			// Fraktur
		//case '𝔞' <= char && char <= '𝔷':
		//	base = '𝔞'
		//	reloc = 'a'
		//case '𝔄' <= char && char <= 'ℨ':
		//	base = '𝔄'
		//	reloc = 'A'

		default:
			base = 0
		}

		if base != 0 {
			offset := char - base
			result.WriteRune(reloc + offset)
		} else {
			done := false
			for n := 1; n < len(patterns); n++ {
				for idx, rv := range patterns[n] {
					if rv == char {
						result.WriteRune(getNthRune(patterns[0], idx))
						done = true
						break
					}
				}
				if done {
					break
				}
			}

			result.WriteRune(char) // Non-alphabetic characters are left as-is
		}
	}
	return result.String()
}

func getNthRune(s string, n int) rune {
	for i, r := range s {
		if i == n {
			return r
		}
	}
	return '?'
}
