package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormaliseFancyUnicodeToToASCII(t *testing.T) {

	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("The fox jumped over the lazy dog"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝐓𝐡𝐞 𝐟𝐨𝐱 𝐣𝐮𝐦𝐩𝐞𝐝 𝐨𝐯𝐞𝐫 𝐭𝐡𝐞 𝐥𝐚𝐳𝐲 𝐝𝐨𝐠")) //Bold
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝑻𝒉𝒆 𝒇𝒐𝒙 𝒋𝒖𝒎𝒑𝒆𝒅 𝒐𝒗𝒆𝒓 𝒕𝒉𝒆 𝒍𝒂𝒛𝒚 𝒅𝒐𝒈")) // Bold-Italic
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝒯𝒽𝑒 𝒻𝑜𝓍 𝒿𝓊𝓂𝓅𝑒𝒹 𝑜𝓋𝑒𝓇 𝓉𝒽𝑒 𝓁𝒶𝓏𝓎 𝒹𝑜𝑔")) // Script-Cursive
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝓣𝓱𝓮 𝓯𝓸𝔁 𝓳𝓾𝓶𝓹𝓮𝓭 𝓸𝓿𝓮𝓻 𝓽𝓱𝓮 𝓵𝓪𝔃𝔂 𝓭𝓸𝓰")) // Script-Cursive Bold
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝙏𝙝𝙚 𝙛𝙤𝙭 𝙟𝙪𝙢𝙥𝙚𝙙 𝙤𝙫𝙚𝙧 𝙩𝙝𝙚 𝙡𝙖𝙯𝙮 𝙙𝙤𝙜")) // Monospaced
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝖳𝗁𝖾 𝖿𝗈𝗑 𝗃𝗎𝗆𝗉𝖾𝖽 𝗈𝗏𝖾𝗋 𝗍𝗁𝖾 𝗅𝖺𝗓𝗒 𝖽𝗈𝗀")) // sans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝗧𝗵𝗲 𝗳𝗼𝘅 𝗷𝘂𝗺𝗽𝗲𝗱 𝗼𝘃𝗲𝗿 𝘁𝗵𝗲 𝗹𝗮𝘇𝘆 𝗱𝗼𝗴")) // Bold Sans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝙏𝙝𝙚 𝙛𝙤𝙭 𝙟𝙪𝙢𝙥𝙚𝙙 𝙤𝙫𝙚𝙧 𝙩𝙝𝙚 𝙡𝙖𝙯𝙮 𝙙𝙤𝙜")) // Italic bold Sans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("Ⓣⓗⓔ ⓕⓞⓧ ⓙⓤⓜⓟⓔⓓ ⓞⓥⓔⓡ ⓣⓗⓔ ⓛⓐⓩⓨ ⓓⓞⓖ")) // Circled
	assert.Equal(t, "the fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("🅣🅗🅔 🅕🅞🅧 🅙🅤🅜🅟🅔🅓 🅞🅥🅔🅡 🅣🅗🅔 🅛🅐🅩🅨 🅓🅞🅖")) // Filled Circled
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝚃𝚑𝚎 𝚏𝚘𝚡 𝚓𝚞𝚖𝚙𝚎𝚍 𝚘𝚟𝚎𝚛 𝚝𝚑𝚎 𝚕𝚊𝚣𝚢 𝚍𝚘𝚐")) // Monospace
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("Ｔｈｅ　ｆｏｘ　ｊｕｍｐｅｄ　ｏｖｅｒ　ｔｈｅ　ｌａｚｙ　ｄｏｇ")) // full-width
	assert.Equal(t, "0123456789", NormaliseFancyUnicodeToToASCII("𝟬𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵"))
	assert.Equal(t, "0123456789", NormaliseFancyUnicodeToToASCII("𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿"))
	assert.Equal(t, "0123456789", NormaliseFancyUnicodeToToASCII("𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡"))

	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("The fox jumped over the lazy dog"))
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("The fox jumped over the lazy dog"))

	// PROBLEMATIC ONES
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝔗𝔥𝔢 𝔣𝔬𝔵 𝔧𝔲𝔪𝔭𝔢𝔡 𝔬𝔳𝔢𝔯 𝔱𝔥𝔢 𝔩𝔞𝔷𝔶 𝔡𝔬𝔤")) // Fraktur
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝕋𝕙𝕖 𝕗𝕠𝕩 𝕛𝕦𝕞𝕡𝕖𝕕 𝕠𝕧𝕖𝕣 𝕥𝕙𝕖 𝕝𝕒𝕫𝕪 𝕕𝕠𝕘")) // Double struck
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝑇ℎ𝑒 𝑓𝑜𝑥 𝑗𝑢𝑚𝑝𝑒𝑑 𝑜𝑣𝑒𝑟 𝑡ℎ𝑒 𝑙𝑎𝑧𝑦 𝑑𝑜𝑔")) // Italic
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝘛𝘩𝘦 𝘧𝘰𝘹 𝘫𝘶𝘮𝘱𝘦𝘥 𝘰𝘷𝘦𝘳 𝘵𝘩𝘦 𝘭𝘢𝘻𝘺 𝘥𝘰𝘨")) // Italic
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝔗𝔥𝔢 𝔣𝔬𝔵 𝔧𝔲𝔪𝔭𝔢𝔡 𝔬𝔳𝔢𝔯 𝔱𝔥𝔢 𝔩𝔞𝔷𝔶 𝔡𝔬𝔤")) // Fraktur -- PROBLEMATIC
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("ᵀʰᵉ ᶠᵒˣ ʲᵘᵐᵖᵉᵈ ᵒᵛᵉʳ ᵗʰᵉ ˡᵃᶻʸ ᵈᵒᵍ")) // Superscript
	//assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("⠞⠓⠑ ⠋⠕⠭ ⠚⠥⠍⠏⠑⠙ ⠕⠧⠑⠗ ⠞⠓⠑ ⠇⠁⠵⠽ ⠙⠕⠛"))
}
