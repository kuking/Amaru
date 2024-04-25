package text

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNormaliseFancyUnicodeToToASCII(t *testing.T) {
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("The fox jumped over the lazy dog"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝐓𝐡𝐞 𝐟𝐨𝐱 𝐣𝐮𝐦𝐩𝐞𝐝 𝐨𝐯𝐞𝐫 𝐭𝐡𝐞 𝐥𝐚𝐳𝐲 𝐝𝐨𝐠")) //Bold
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝑻𝒉𝒆 𝒇𝒐𝒙 𝒋𝒖𝒎𝒑𝒆𝒅 𝒐𝒗𝒆𝒓 𝒕𝒉𝒆 𝒍𝒂𝒛𝒚 𝒅𝒐𝒈")) //Bold-Italic
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝒯𝒽𝑒 𝒻𝑜𝓍 𝒿𝓊𝓂𝓅𝑒𝒹 𝑜𝓋𝑒𝓇 𝓉𝒽𝑒 𝓁𝒶𝓏𝓎 𝒹𝑜𝑔")) //Script-Cursive
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝓣𝓱𝓮 𝓯𝓸𝔁 𝓳𝓾𝓶𝓹𝓮𝓭 𝓸𝓿𝓮𝓻 𝓽𝓱𝓮 𝓵𝓪𝔃𝔂 𝓭𝓸𝓰")) //Script-CursiveBold
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝙏𝙝𝙚 𝙛𝙤𝙭 𝙟𝙪𝙢𝙥𝙚𝙙 𝙤𝙫𝙚𝙧 𝙩𝙝𝙚 𝙡𝙖𝙯𝙮 𝙙𝙤𝙜")) //Monospaced
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝖳𝗁𝖾 𝖿𝗈𝗑 𝗃𝗎𝗆𝗉𝖾𝖽 𝗈𝗏𝖾𝗋 𝗍𝗁𝖾 𝗅𝖺𝗓𝗒 𝖽𝗈𝗀")) //sans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝗧𝗵𝗲 𝗳𝗼𝘅 𝗷𝘂𝗺𝗽𝗲𝗱 𝗼𝘃𝗲𝗿 𝘁𝗵𝗲 𝗹𝗮𝘇𝘆 𝗱𝗼𝗴")) //BoldSans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝑇ℎ𝑒 𝑓𝑜𝑥 𝑗𝑢𝑚𝑝𝑒𝑑 𝑜𝑣𝑒𝑟 𝑡ℎ𝑒 𝑙𝑎𝑧𝑦 𝑑𝑜𝑔")) //Sans-Serif Italic
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝙏𝙝𝙚 𝙛𝙤𝙭 𝙟𝙪𝙢𝙥𝙚𝙙 𝙤𝙫𝙚𝙧 𝙩𝙝𝙚 𝙡𝙖𝙯𝙮 𝙙𝙤𝙜")) //ItalicBoldSans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝘛𝘩𝘦 𝘧𝘰𝘹 𝘫𝘶𝘮𝘱𝘦𝘥 𝘰𝘷𝘦𝘳 𝘵𝘩𝘦 𝘭𝘢𝘻𝘺 𝘥𝘰𝘨")) //ItalicBoldSans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("Ⓣⓗⓔ ⓕⓞⓧ ⓙⓤⓜⓟⓔⓓ ⓞⓥⓔⓡ ⓣⓗⓔ ⓛⓐⓩⓨ ⓓⓞⓖ")) //Circled
	assert.Equal(t, "THE FOX JUMPED OVER THE LAZY DOG", NormaliseFancyUnicodeToASCII("🅣🅗🅔 🅕🅞🅧 🅙🅤🅜🅟🅔🅓 🅞🅥🅔🅡 🅣🅗🅔 🅛🅐🅩🅨 🅓🅞🅖")) //FilledCircled
	assert.Equal(t, "THE FOX JUMPED OVER THE LAZY DOG", NormaliseFancyUnicodeToASCII("🆃🅷🅴 🅵🅾🆇 🅹🆄🅼🅿🅴🅳 🅾🆅🅴🆁 🆃🅷🅴 🅻🅰🆉🆈 🅳🅾🅶")) //FilledCircled
	assert.Equal(t, "THE FOX JUMPED OVER THE LAZY DOG", NormaliseFancyUnicodeToASCII("🅃🄷🄴 🄵🄾🅇 🄹🅄🄼🄿🄴🄳 🄾🅅🄴🅁 🅃🄷🄴 🄻🄰🅉🅈 🄳🄾🄶"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝚃𝚑𝚎 𝚏𝚘𝚡 𝚓𝚞𝚖𝚙𝚎𝚍 𝚘𝚟𝚎𝚛 𝚝𝚑𝚎 𝚕𝚊𝚣𝚢 𝚍𝚘𝚐")) //Monospace
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("Ｔｈｅ　ｆｏｘ　ｊｕｍｐｅｄ　ｏｖｅｒ　ｔｈｅ　ｌａｚｙ　ｄｏｇ")) //full-width
	assert.Equal(t, "the fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("ᴛʜᴇ ꜰᴏx ᴊᴜᴍᴩᴇᴅ ᴏᴠᴇʀ ᴛʜᴇ ʟᴀᴢy ᴅᴏɢ"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝕋𝕙𝕖 𝕗𝕠𝕩 𝕛𝕦𝕞𝕡𝕖𝕕 𝕠𝕧𝕖𝕣 𝕥𝕙𝕖 𝕝𝕒𝕫𝕪 𝕕𝕠𝕘"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝔗𝔥𝔢 𝔣𝔬𝔵 𝔧𝔲𝔪𝔭𝔢𝔡 𝔬𝔳𝔢𝔯 𝔱𝔥𝔢 𝔩𝔞𝔷𝔶 𝔡𝔬𝔤"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("𝕿𝖍𝖊 𝖋𝖔𝖝 𝖏𝖚𝖒𝖕𝖊𝖉 𝖔𝖛𝖊𝖗 𝖙𝖍𝖊 𝖑𝖆𝖟𝖞 𝖉𝖔𝖌"))
	assert.Equal(t, "the fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("⠞⠓⠑ ⠋⠕⠭ ⠚⠥⠍⠏⠑⠙ ⠕⠧⠑⠗ ⠞⠓⠑ ⠇⠁⠵⠽ ⠙⠕⠛"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("ᵀʰᵉ ᶠᵒˣ ʲᵘᵐᵖᵉᵈ ᵒᵛᵉʳ ᵗʰᵉ ˡᵃᶻʸ ᵈᵒᵍ"))
	assert.Equal(t, "the fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("ₜₕₑ 𝒻ₒₓ ⱼᵤₘₚₑ𝒹 ₒᵥₑᵣ ₜₕₑ ₗₐ𝓏ᵧ 𝒹ₒ𝓰"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToASCII("T⃣h⃣e⃣ ⃣f⃣o⃣x⃣ ⃣j⃣u⃣m⃣p⃣e⃣d⃣ ⃣o⃣v⃣e⃣r⃣ ⃣t⃣h⃣e⃣ ⃣l⃣a⃣z⃣y⃣ ⃣d⃣o⃣g⃣"))

	assert.Equal(t, "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
		NormaliseFancyUnicodeToASCII("0123456789⓪①②③④⑤⑥⑦⑧⑨⓿❶❷❸❹❺❻❼❽❾０１２３４５６７８９𝟬𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵𝟢𝟣𝟤𝟥𝟦𝟧𝟨𝟩𝟪𝟫𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿𝟎𝟏𝟐𝟑𝟒𝟓𝟔𝟕𝟖𝟗"))
}

func TestRemoveEmojis(t *testing.T) {
	assert.Equal(t, "", RemoveEmojis("😀😃😄😁😆😅🤣😂\U0001F979🙂🙃😉😊😇🥰😍🤩😘😗☺️😚😙🥲😋😛😜🤪😝🤑🤗\U0001FAE2🤭🤫🤔\U0001FAE1🤐🤨😐️"+
		"😑😶😏😒🙄😬🤥😌😔😪😮‍💨🤤😴😷🤒🤕🤢🤮🤧\U0001FAE0🥵🥶😶‍🌫️\U0001FAE5🥴\U0001FAE8😵‍💫😵🤯🤠🥳🥸😎🤓🧐\U0001FAE4😕😟🙁☹️😮😯😲😳\U0001FAE3🥺"+
		"😦😧😨😰😥😢😭😱😖😣😞😓😩😫🥱😤😡😠🤬😈👿💀☠️💩🤡👹👺👻👽️👾🤖😺😸😹😻😼😽🙀😿😾🙈🙉🙊👋🤚🖐️✋🖖👌🤌🤏✌️🤞\U0001FAF0🤟🤘🤙👈️👉️👆️🖕👇️☝️"+
		"\U0001FAF5👍️👎️✊👊🤛🤜👏🙌👐\U0001FAF6🤲\U0001FAF3\U0001FAF4\U0001FAF1\U0001FAF2🤝\U0001FAF8\U0001FAF7🙏✍️💅🤳💪🦾🦿🦵🦶👂️🦻👃🧠🫀🫁🦷🦴👀👁️"+
		"👅👄\U0001FAE6💋👶🧒👦👧🧑👨👩🧔🧔‍♀️🧔‍♂️🧑‍🦰👨‍🦰👩‍🦰🧑‍🦱👨‍🦱👩‍🦱🧑‍🦳👨‍🦳👩‍🦳🧑‍🦲👨‍🦲👩‍🦲👱👱‍♂️👱‍♀️🧓👴👵🙍🙍‍♂️🙍‍♀️🙎🙎‍♂️🙎‍♀️🙅🙅‍♂️🙅‍♀️🙆🙆‍♂️🙆‍♀️💁💁‍♂️💁‍♀️🙋🙋‍♂️🙋‍♀️🧏🧏‍♂️🧏‍♀️🙇🙇‍♂️🙇‍♀️🤦🤦‍♂️🤦‍♀️🤷"+
		"🤷‍♂️🤷‍♀️🧑‍⚕️👨‍⚕️👩‍⚕️🧑‍🎓👨‍🎓👩‍🎓🧑‍🏫👨‍🏫👩‍🏫🧑‍⚖️👨‍⚖️👩‍⚖️🧑‍🌾👨‍🌾👩‍🌾🧑‍🍳👨‍🍳👩‍🍳🧑‍🔧👨‍🔧👩‍🔧🧑‍🏭👨‍🏭👩‍🏭🧑‍💼👨‍💼👩‍💼🧑‍🔬👨‍🔬👩‍🔬🧑‍💻👨‍💻👩‍💻🧑‍🎤👨‍🎤👩‍🎤🧑‍🎨👨‍🎨👩‍🎨🧑‍✈️👨‍✈️👩‍✈️🧑‍🚀👨‍🚀👩‍🚀🧑‍🚒👨‍🚒👩‍🚒👮👮‍♂️👮‍♀️🕵️🕵️‍♂️🕵️‍♀️💂💂‍♂️💂‍♀️🥷👷👷‍♂️👷‍♀️"+
		"\U0001FAC5🤴👸👳👳‍♂️👳‍♀️👲🧕🤵🤵‍♂️🤵‍♀️👰👰‍♂️👰‍♀️\U0001FAC4\U0001FAC3🤰🤱👩‍🍼👨‍🍼🧑‍🍼👼🎅🤶🧑‍🎄🦸🦸‍♂️🦸‍♀️🦹🦹‍♂️🦹‍♀️🧙🧙‍♂️🧙‍♀️🧚🧚‍♂️🧚‍♀️🧛🧛‍♂️🧛‍♀️🧜🧜‍♂️🧜‍♀️🧝🧝‍♂️🧝‍♀️🧞🧞‍♂️🧞‍♀️🧟🧟‍♂️🧟‍♀️"+
		"\U0001F9CC💆💆‍♂️💆‍♀️💇💇‍♂️💇‍♀️🚶🚶‍♂️🚶‍♀️🧍🧍‍♂️🧍‍♀️🧎🧎‍♂️🧎‍♀️🧑‍🦯👨‍🦯👩‍🦯🧑‍🦼👨‍🦼👩‍🦼🧑‍🦽👨‍🦽👩‍🦽🏃🏃‍♂️🏃‍♀️💃🕺🕴️👯👯‍♂️👯‍♀️🧖🧖‍♂️🧖‍♀️🧗🧗‍♂️🧗‍♀️🤺🏇⛷️🏂️🏌️🏌️‍♂️🏌️‍♀️🏄️🏄‍♂️🏄‍♀️🚣🚣‍♂️🚣‍♀️🏊️🏊‍♂️🏊‍♀️⛹️⛹️‍♂️⛹️‍♀️🏋️"+
		"🏋️‍♂️🏋️‍♀️🚴🚴‍♂️🚴‍♀️🚵🚵‍♂️🚵‍♀️🤸🤸‍♂️🤸‍♀️🤼🤼‍♂️🤼‍♀️🤽🤽‍♂️🤽‍♀️🤾🤾‍♂️🤾‍♀️🤹🤹‍♂️🤹‍♀️🧘🧘‍♂️🧘‍♀️🛀🛌🧑‍🤝‍🧑👭👫👬💏👩‍❤️‍💋‍👨👨‍❤️‍💋‍👨👩‍❤️‍💋‍👩💑👩‍❤️‍👨👨‍❤️‍👨👩‍❤️‍👩👪️👨‍👩‍👦👨‍👩‍👧👨‍👩‍👧‍👦👨‍👩‍👦‍👦👨‍👩‍👧‍👧👨‍👨‍👦👨‍👨‍👧👨‍👨‍👧‍👦👨‍👨‍👦‍👦👨‍👨‍👧‍👧👩‍👩‍👦👩‍👩‍👧👩‍👩‍👧‍👦👩‍👩‍👦‍👦👩‍👩‍👧‍👧👨‍👦👨‍👦‍👦👨‍👧👨‍👧‍👦👨‍👧‍👧👩‍👦👩‍👦‍👦"+
		"👩‍👧👩‍👧‍👦👩‍👧‍👧🗣️👤👥🫂👣"))

	assert.Equal(t, "", RemoveEmojis("🐵🐒🦍🦧🐶🐕️🦮🐕‍🦺🐩🐺🦊🦝🐱🐈️🐈‍⬛🦁🐯🐅🐆🐴🐎🦄\U0001FACF🦓🦌\U0001FACE🦬🐮🐂🐃🐄🐷🐖🐗🐽🐏🐑🐐🐪🐫"+
		"🦙🦒🐘🦣🦏🦛🐭🐁🐀🐹🐰🐇🐿️🦫🦔🦇🐻🐻‍❄️🐨🐼🦥🦦🦨🦘🦡🐾🦃🐔🐓🐣🐤🐥🐦️🐧🐦‍⬛🕊️🦅🦆\U0001FABF🦢🦉🦤🦩🦚🦜\U0001FABD🪶\U0001FAB9\U0001FABA🥚"+
		"🐸🐊🐢🦎🐍🐲🐉🦕🦖🐳🐋🐬🦭🐟️🐠🐡🦈\U0001FABC🐙🦑🦀🦞🦐\U0001FAB8🦪🐚🐌🦋🐛🐜🐝🪲🐞🦗🪳🕷️🕸️🦂🦟🪰🪱🦠🍄💐💮🏵️🌼🌻🌹🥀🌺🌷🌸\U0001FAB7"+
		"\U0001FABB🌱🪴🏕️🌲🌳🌰🌴🌵🎋🎍🌾🌿☘️🍀🍁🍂🍃🌍️🌎️🌏️🌑🌒🌓🌔🌕️🌖🌗🌘🌙🌚🌛🌜️☀️🌝🌞🪐💫⭐️🌟✨🌠☄️🌌☁️⛅️⛈️🌤️🌥️🌦️🌧️🌨️🌩️🌪️🌫️🌬️🌀🌈🌂☂️☔️"+
		"⛱️⚡️❄️☃️⛄️🏔️⛰️🗻🌋🔥💧🌊💥💦💨"))

	assert.Equal(t, "", RemoveEmojis("🍇🍈🍉🍊🍋🍌🍍🥭🍎🍏🍐🍑🍒🍓🫐🥝🍅🫒🥥🥑🍆🥔🥕🌽🌶️🫑🥒🥬🥦\U0001FADB🧄🧅\U0001FADA🍄\U0001FAD8🥜🌰"+
		"🍞🥐🥖🫓🥨🥯🥞🧇🧀🍖🍗🥩🥓🍔🍟🍕🌭🥪🌮🌯🫔🥙🧆🥚🍳🥘🍲🫕🥣🥗🍿🧈🧂🥫🍱🍘🍙🍚🍛🍜🍝🍠🍢🍣🍤🍥🥮🍡🥟🥠🥡🍦🍧🍨🍩🍪🎂🍰🧁🥧🍫🍬🍭"+
		"🍮🍯🍼🥛\U0001FAD7☕️🫖🍵🍶🍾🍷🍸️🍹🍺🍻🥂🥃🥤🧋🧃🧉🧊🥢🍽️🍴🥄🔪⚽️⚾️🥎🏀🏐🏈🏉🎾🥏🎳🏏🏑🏒🥍🏓🏸🥊🥋🥅⛳️⛸️🎣🤿🎽🎿🛷🥌🎯🪀🪁🎱🎖️🏆️"+
		"🏅🥇🥈🥉🏔️⛰️🌋🗻🏕️🏖️🏜️🏝️🏟️🏛️🏗️🧱🪨🪵🛖🏘️🏚️🏠️🏡🏢🏣🏤🏥🏦🏨🏩🏪🏫🏬🏭️🏯🏰💒🗼🗽⛪️🕌🛕🕍⛩️🕋⛲️⛺️🌁🌃🏙️🌄🌅🌆🌇🌉🗾🏞️🎠🎡🎢💈🎪🚂"+
		"🚃🚄🚅🚆🚇️🚈🚉🚊🚝🚞🚋🚌🚍️🚎🚐🚑️🚒🚓🚔️🚕🚖🚗🚘️🚙🛻🚚🚛🚜🏎️🏍️🛵🦽🦼🛺🚲️🛴🛹🛼🚏🛣️🛤️🛢️⛽️🚨🚥🚦🛑🚧⚓️⛵️🛶🚤🛳️⛴️🛥️🚢✈️🛩️🛫🛬🪂💺🚁"+
		"🚟🚠🚡🛰️🚀🛸🎆🎇🎑🗿🛎️🧳⌛️⏳️⌚️⏰⏱️⏲️🕰️🌡️🗺️🧭🎃🎄🧨🎈🎉🎊🎎\U0001FAAD🎏🎐🎀🎁🎗️🎟️🎫🔮🪄🧿🎮️🕹️🎰🎲♟️🧩🧸🪅🪆🖼️🎨🧵🪡🧶🪢👓️🕶️🥽🥼🦺"+
		"👔👕👖🧣🧤🧥🧦👗👘🥻🩱🩲🩳👙👚👛👜👝🛍️🎒🩴👞👟🥾🥿👠👡🩰👢👑👒🎩🎓️🧢🪖⛑️📿💄💍💎📢📣📯🎙️🎚️🎛️🎤🎧️📻️🎷🪗🎸🎹🎺🎻🪕\U0001FA88\U0001FA87"+
		"🥁🪘\U0001FAA9📱📲☎️📞📟️📠🔋\U0001FAAB🔌💻️🖥️🖨️⌨️🖱️🖲️💽💾💿️📀🧮🎥🎞️📽️🎬️📺️📷️📸📹️📼🔍️🔎🕯️💡🔦🏮🪔📔📕📖📗📘📙📚️📓📒📃📜📄📰🗞️📑🔖🏷️💰️"+
		"🪙💴💵💶💷💸💳️\U0001FAAA🧾✉️💌📧🧧📨📩📤️📥️📦️📫️📪️📬️📭️📮🗳️✏️✒️🖋️🖊️🖌️🖍️📝💼📁📂🗂️📅📆🗒️🗓️📇📈📉📊📋️📌📍📎🖇️📏📐✂️🗃️🗄️🗑️🔒️🔓️🔏🔐🔑🗝️"+
		"🔨🪓⛏️⚒️🛠️🗡️⚔️💣️🔫🪃🏹🛡️🪚🔧🪛🔩⚙️🗜️⚖️🦯🔗⛓️🪝🧰🧲🪜\U0001F6DD\U0001F6DE\U0001FAD9⚗️🧪🧫🧬🔬🔭📡\U0001FA7B💉🩸💊🩹🩺\U0001FA7C🚪🛗🪞🪟🛏"+
		"️🛋️🪑🪤🚽🪠🚿🛁🧼\U0001FAE7🪒\U0001FAAE🧴🧷🧹🧺🧻🪣🪥🧽🧯\U0001F6DF🛒🚬⚰️🪦⚱️🏺🪧🕳️💘💝💖💗💓💞💕💟❣️💔❤️🧡💛💚\U0001FA75💙💜"+
		"\U0001FA77🤎🖤\U0001FA76🤍❤️‍🔥❤️‍🩹💯♨️💢💬👁️‍🗨️🗨️🗯️💭💤🌐♠️♥️♦️♣️🃏🀄️🎴🎭️🔇🔈️🔉🔊🔔🔕🎼🎵🎶💹🏧🚮🚰♿️🚹️🚺️🚻🚼️🚾🛂🛃🛄🛅\U0001F6DC⚠️🚸⛔️🚫🚳🚭️"+
		"🚯🚱🚷📵🔞☢️☣️⬆️↗️➡️↘️⬇️↙️⬅️↖️↕️↔️↩️↪️⤴️⤵️🔃🔄🔙🔚🔛🔜🔝🛐⚛️🕉️✡️☸️\U0001FAAF☯️✝️☦️☪️☮️🕎🔯\U0001FAAC♈️♉️♊️♋️♌️♍️♎️♏️♐️♑️♒️♓️⛎🔀🔁🔂"+
		"▶️⏩️⏭️⏯️◀️⏪️⏮️🔼⏫🔽⏬⏸️⏹️⏺️⏏️🎦🔅🔆📶📳📴♀️♂️⚧✖️➕➖➗\U0001F7F0♾️‼️⁉️❓️❔❕❗️〰️💱💲⚕️♻️⚜️🔱📛🔰⭕️✅☑️✔️❌❎➰➿〽️✳️✴️❇️©️®️0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣🔟"+
		"🔠🔡🔢🔣🔤🅰️🆎🅱️🆑🆒🆓🆕🆖🅾️🆗🅿️🆘🆙🆚🈁🈂️🈷️🈶🈯️🉐🈹🈚️🈲🉑🈸🈴🈳㊗️㊙️🈺🈵🔴🟠🟡🟢🔵🟣🟤⚫️⚪️🟥🟧🟨🟩🟦🟪🟫⬛️⬜️◼️◻️◾️◽️▪️▫️🔶🔷🔸🔹🔺"+
		"🔻💠🔘🔳🔲🕛️🕧️🕐️🕜️🕑️🕝️🕒️🕞️🕓️🕟️🕔️🕠️🕕️🕡️🕖️🕢️🕗️🕣️🕘️🕤️🕙️🕥️🕚️🕦️"))

	assert.Equal(t, "", RemoveEmojis("🏁🚩🎌🏴🏳️🏳️‍🌈🏳️‍⚧️🏴‍☠️🇦🇨🇦🇩🇦🇪🇦🇫🇦🇬🇦🇮🇦🇱🇦🇲🇦🇴🇦🇶🇦🇷🇦🇸🇦🇹🇦🇺🇦🇼🇦🇽🇦🇿🇧🇦🇧🇧🇧🇩🇧🇪🇧🇫🇧🇬🇧🇭🇧🇮🇧🇯🇧🇱🇧🇲🇧🇳🇧🇴🇧🇶🇧🇷🇧🇸🇧🇹🇧🇻🇧🇼🇧🇾🇧🇿🇨🇦🇨🇨🇨🇩"+
		"🇨🇫🇨🇬🇨🇭🇨🇮🇨🇰🇨🇱🇨🇲🇨🇳🇨🇴🇨🇵🇨🇷🇨🇺🇨🇻🇨🇼🇨🇽🇨🇾🇨🇿🇩🇪🇩🇬🇩🇯🇩🇰🇩🇲🇩🇴🇩🇿🇪🇦🇪🇨🇪🇪🇪🇬🇪🇭🇪🇷🇪🇸🇪🇹🇪🇺🇫🇮🇫🇯🇫🇰🇫🇲🇫🇴🇫🇷🇬🇦🇬🇧🇬🇩🇬🇪🇬🇫🇬🇬🇬🇭🇬🇮🇬🇱🇬🇲🇬🇳🇬🇵🇬🇶🇬🇷🇬🇸🇬🇹🇬🇺🇬🇼🇬🇾🇭🇰🇭🇲🇭🇳🇭🇷🇭🇹🇭🇺🇮🇨🇮🇩"+
		"🇮🇪🇮🇱🇮🇲🇮🇳🇮🇴🇮🇶🇮🇷🇮🇸🇮🇹🇯🇪🇯🇲🇯🇴🇯🇵🇰🇪🇰🇬🇰🇭🇰🇮🇰🇲🇰🇳🇰🇵🇰🇷🇰🇼🇰🇾🇰🇿🇱🇦🇱🇧🇱🇨🇱🇮🇱🇰🇱🇷🇱🇸🇱🇹🇱🇺🇱🇻🇱🇾🇲🇦🇲🇨🇲🇩🇲🇪🇲🇫🇲🇬🇲🇭🇲🇰🇲🇱🇲🇲🇲🇳🇲🇴🇲🇵🇲🇶🇲🇷🇲🇸🇲🇹🇲🇺🇲🇻🇲🇼🇲🇽🇲🇾🇲🇿🇳🇦🇳🇨🇳🇪🇳🇫🇳🇬🇳🇮🇳🇱🇳🇴"+
		"🇳🇵🇳🇷🇳🇺🇳🇿🇴🇲🇵🇦🇵🇪🇵🇫🇵🇬🇵🇭🇵🇰🇵🇱🇵🇲🇵🇳🇵🇷🇵🇸🇵🇹🇵🇼🇵🇾🇶🇦🇷🇪🇷🇴🇷🇸🇷🇺🇷🇼🇸🇦🇸🇧🇸🇨🇸🇩🇸🇪🇸🇬🇸🇭🇸🇮🇸🇯🇸🇰🇸🇱🇸🇲🇸🇳🇸🇴🇸🇷🇸🇸🇸🇹🇸🇻🇸🇽🇸🇾🇸🇿🇹🇦🇹🇨🇹🇩🇹🇫🇹🇬🇹🇭🇹🇯🇹🇰🇹🇱🇹🇲🇹🇳🇹🇴🇹🇷🇹🇹🇹🇻🇹🇼🇹🇿🇺🇦🇺🇬🇺🇳"+
		"🇺🇸🇺🇾🇺🇿🇻🇦🇻🇪🇻🇬🇻🇮🇻🇳🇻🇺🇼🇫🇼🇸🇽🇰🇾🇪🇾🇹🇿🇦🇿🇲🇿🇼🇺🇲"))

}

func TestTags(t *testing.T) {
	defs := map[string]string{
		"gold": "money", "silver": "money", "bitcoin": "money",
		"cat": "animal", "dog": "animal", "elephant": "animal", "monkey": "animal",
		"red": "color", "yellow": "color", "green": "color",
	}

	assert.Equal(t, []string{}, Tags("", defs))
	assert.Equal(t, []string{}, Tags("hello world", defs))
	assert.Equal(t, []string{"animal"}, Tags("hello MONKEY world", defs))

	res := Tags("silver monkey red", defs)
	assert.Contains(t, res, "money")
	assert.Contains(t, res, "animal")
	assert.Contains(t, res, "color")
	assert.Len(t, res, 3)

	res = Tags("gold silver bitcoin BitCoin cat CaT dog elephant monkey red yellow green", defs)
	assert.Contains(t, res, "money")
	assert.Contains(t, res, "animal")
	assert.Contains(t, res, "color")
	assert.Len(t, res, 3)
}

func TestParseTagsDefinition(t *testing.T) {
	definition := `

=animal
+cat
+dog
+elephant
+monkey

# comments
=money
+bitcoin
+gold
+silver

=coLOr
+red
+green
+BLUe


`
	defs, err := ParseTagsDefinition(strings.NewReader(definition))
	assert.NoError(t, err)

	assert.Len(t, defs, 10)
	assert.Contains(t, defs, "red")

	res := Tags("gold silver bitcoin BitCoin cat CaT dog elephant monkey red yellow green", defs)
	assert.Contains(t, res, "money")
	assert.Contains(t, res, "animal")
	assert.Contains(t, res, "color")
	assert.Len(t, res, 3)
}
