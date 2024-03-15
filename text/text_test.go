package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormaliseFancyUnicodeToToASCII(t *testing.T) {

	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("The fox jumped over the lazy dog"))
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝐓𝐡𝐞 𝐟𝐨𝐱 𝐣𝐮𝐦𝐩𝐞𝐝 𝐨𝐯𝐞𝐫 𝐭𝐡𝐞 𝐥𝐚𝐳𝐲 𝐝𝐨𝐠")) //Bold
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝑻𝒉𝒆 𝒇𝒐𝒙 𝒋𝒖𝒎𝒑𝒆𝒅 𝒐𝒗𝒆𝒓 𝒕𝒉𝒆 𝒍𝒂𝒛𝒚 𝒅𝒐𝒈")) //Bold-Italic
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝒯𝒽𝑒 𝒻𝑜𝓍 𝒿𝓊𝓂𝓅𝑒𝒹 𝑜𝓋𝑒𝓇 𝓉𝒽𝑒 𝓁𝒶𝓏𝓎 𝒹𝑜𝑔")) //Script-Cursive
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝓣𝓱𝓮 𝓯𝓸𝔁 𝓳𝓾𝓶𝓹𝓮𝓭 𝓸𝓿𝓮𝓻 𝓽𝓱𝓮 𝓵𝓪𝔃𝔂 𝓭𝓸𝓰")) //Script-CursiveBold
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝙏𝙝𝙚 𝙛𝙤𝙭 𝙟𝙪𝙢𝙥𝙚𝙙 𝙤𝙫𝙚𝙧 𝙩𝙝𝙚 𝙡𝙖𝙯𝙮 𝙙𝙤𝙜")) //Monospaced
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝖳𝗁𝖾 𝖿𝗈𝗑 𝗃𝗎𝗆𝗉𝖾𝖽 𝗈𝗏𝖾𝗋 𝗍𝗁𝖾 𝗅𝖺𝗓𝗒 𝖽𝗈𝗀")) //sans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝗧𝗵𝗲 𝗳𝗼𝘅 𝗷𝘂𝗺𝗽𝗲𝗱 𝗼𝘃𝗲𝗿 𝘁𝗵𝗲 𝗹𝗮𝘇𝘆 𝗱𝗼𝗴")) //BoldSans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝙏𝙝𝙚 𝙛𝙤𝙭 𝙟𝙪𝙢𝙥𝙚𝙙 𝙤𝙫𝙚𝙧 𝙩𝙝𝙚 𝙡𝙖𝙯𝙮 𝙙𝙤𝙜")) //ItalicBoldSans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝘛𝘩𝘦 𝘧𝘰𝘹 𝘫𝘶𝘮𝘱𝘦𝘥 𝘰𝘷𝘦𝘳 𝘵𝘩𝘦 𝘭𝘢𝘻𝘺 𝘥𝘰𝘨")) //ItalicBoldSans-serif
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("Ⓣⓗⓔ ⓕⓞⓧ ⓙⓤⓜⓟⓔⓓ ⓞⓥⓔⓡ ⓣⓗⓔ ⓛⓐⓩⓨ ⓓⓞⓖ")) //Circled
	assert.Equal(t, "the fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("🅣🅗🅔 🅕🅞🅧 🅙🅤🅜🅟🅔🅓 🅞🅥🅔🅡 🅣🅗🅔 🅛🅐🅩🅨 🅓🅞🅖")) //FilledCircled
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("𝚃𝚑𝚎 𝚏𝚘𝚡 𝚓𝚞𝚖𝚙𝚎𝚍 𝚘𝚟𝚎𝚛 𝚝𝚑𝚎 𝚕𝚊𝚣𝚢 𝚍𝚘𝚐")) //Monospace
	assert.Equal(t, "The fox jumped over the lazy dog", NormaliseFancyUnicodeToToASCII("Ｔｈｅ　ｆｏｘ　ｊｕｍｐｅｄ　ｏｖｅｒ　ｔｈｅ　ｌａｚｙ　ｄｏｇ")) //full-width
	assert.Equal(t, "0123456789", NormaliseFancyUnicodeToToASCII("𝟬𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵"))
	assert.Equal(t, "0123456789", NormaliseFancyUnicodeToToASCII("𝟶𝟷𝟸𝟹𝟺𝟻𝟼𝟽𝟾𝟿"))
	assert.Equal(t, "0123456789", NormaliseFancyUnicodeToToASCII("𝟘𝟙𝟚𝟛𝟜𝟝𝟞𝟟𝟠𝟡"))

	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("Thefoxjumpedoverthelazydog"))
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("Thefoxjumpedoverthelazydog"))

	//PROBLEMATICONES
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("𝔗𝔥𝔢𝔣𝔬𝔵𝔧𝔲𝔪𝔭𝔢𝔡𝔬𝔳𝔢𝔯𝔱𝔥𝔢𝔩𝔞𝔷𝔶𝔡𝔬𝔤"))//Fraktur
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("𝕋𝕙𝕖𝕗𝕠𝕩𝕛𝕦𝕞𝕡𝕖𝕕𝕠𝕧𝕖𝕣𝕥𝕙𝕖𝕝𝕒𝕫𝕪𝕕𝕠𝕘"))//Doublestruck
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("𝑇ℎ𝑒𝑓𝑜𝑥𝑗𝑢𝑚𝑝𝑒𝑑𝑜𝑣𝑒𝑟𝑡ℎ𝑒𝑙𝑎𝑧𝑦𝑑𝑜𝑔"))//Italic
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("𝘛𝘩𝘦𝘧𝘰𝘹𝘫𝘶𝘮𝘱𝘦𝘥𝘰𝘷𝘦𝘳𝘵𝘩𝘦𝘭𝘢𝘻𝘺𝘥𝘰𝘨"))//Italic
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("𝔗𝔥𝔢𝔣𝔬𝔵𝔧𝔲𝔪𝔭𝔢𝔡𝔬𝔳𝔢𝔯𝔱𝔥𝔢𝔩𝔞𝔷𝔶𝔡𝔬𝔤"))//Fraktur--PROBLEMATIC
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("ᵀʰᵉᶠᵒˣʲᵘᵐᵖᵉᵈᵒᵛᵉʳᵗʰᵉˡᵃᶻʸᵈᵒᵍ"))//Superscript
	//assert.Equal(t,"Thefoxjumpedoverthelazydog",NormaliseFancyUnicodeToToASCII("⠞⠓⠑⠋⠕⠭⠚⠥⠍⠏⠑⠙⠕⠧⠑⠗⠞⠓⠑⠇⠁⠵⠽⠙⠕⠛"))
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
