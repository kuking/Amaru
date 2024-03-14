package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormaliseFancyUnicodeToToASCII(t *testing.T) {

	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("Thefoxjumpedoverthelazydog"))
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝐓𝐡𝐞𝐟𝐨𝐱𝐣𝐮𝐦𝐩𝐞𝐝𝐨𝐯𝐞𝐫𝐭𝐡𝐞𝐥𝐚𝐳𝐲𝐝𝐨𝐠"))       //Bold
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝑻𝒉𝒆𝒇𝒐𝒙𝒋𝒖𝒎𝒑𝒆𝒅𝒐𝒗𝒆𝒓𝒕𝒉𝒆𝒍𝒂𝒛𝒚𝒅𝒐𝒈"))       //Bold-Italic
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝒯𝒽𝑒𝒻𝑜𝓍𝒿𝓊𝓂𝓅𝑒𝒹𝑜𝓋𝑒𝓇𝓉𝒽𝑒𝓁𝒶𝓏𝓎𝒹𝑜𝑔"))       //Script-Cursive
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝓣𝓱𝓮𝓯𝓸𝔁𝓳𝓾𝓶𝓹𝓮𝓭𝓸𝓿𝓮𝓻𝓽𝓱𝓮𝓵𝓪𝔃𝔂𝓭𝓸𝓰"))       //Script-CursiveBold
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝙏𝙝𝙚𝙛𝙤𝙭𝙟𝙪𝙢𝙥𝙚𝙙𝙤𝙫𝙚𝙧𝙩𝙝𝙚𝙡𝙖𝙯𝙮𝙙𝙤𝙜"))       //Monospaced
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝖳𝗁𝖾𝖿𝗈𝗑𝗃𝗎𝗆𝗉𝖾𝖽𝗈𝗏𝖾𝗋𝗍𝗁𝖾𝗅𝖺𝗓𝗒𝖽𝗈𝗀"))       //sans-serif
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝗧𝗵𝗲𝗳𝗼𝘅𝗷𝘂𝗺𝗽𝗲𝗱𝗼𝘃𝗲𝗿𝘁𝗵𝗲𝗹𝗮𝘇𝘆𝗱𝗼𝗴"))       //BoldSans-serif
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝙏𝙝𝙚𝙛𝙤𝙭𝙟𝙪𝙢𝙥𝙚𝙙𝙤𝙫𝙚𝙧𝙩𝙝𝙚𝙡𝙖𝙯𝙮𝙙𝙤𝙜"))       //ItalicboldSans-serif
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("Ⓣⓗⓔⓕⓞⓧⓙⓤⓜⓟⓔⓓⓞⓥⓔⓡⓣⓗⓔⓛⓐⓩⓨⓓⓞⓖ"))       //Circled
	assert.Equal(t, "thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("🅣🅗🅔🅕🅞🅧🅙🅤🅜🅟🅔🅓🅞🅥🅔🅡🅣🅗🅔🅛🅐🅩🅨🅓🅞🅖"))       //FilledCircled
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("𝚃𝚑𝚎𝚏𝚘𝚡𝚓𝚞𝚖𝚙𝚎𝚍𝚘𝚟𝚎𝚛𝚝𝚑𝚎𝚕𝚊𝚣𝚢𝚍𝚘𝚐"))       //Monospace
	assert.Equal(t, "Thefoxjumpedoverthelazydog", NormaliseFancyUnicodeToToASCII("Ｔｈｅ　ｆｏｘ　ｊｕｍｐｅｄ　ｏｖｅｒ　ｔｈｅ　ｌａｚｙ　ｄｏｇ")) //full-width
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
