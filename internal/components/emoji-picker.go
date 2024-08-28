package components

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// EmojiPickerProps ...
type EmojiPickerProps struct {
	ClassNames htmx.ClassNames
}

// EmojiPicker ...
func EmojiPicker(props EmojiPickerProps) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			props.ClassNames,
		),
		alpine.XData(`{
            open: false,
            search: '',
            input: '🙂',
            emojis: {
                'tractor, farm, machine, agriculture': '🚜',
                'leaf, plant, nature, green, agriculture, ecology': '🌿',
                'corn, field, agricuultre, vegetable, plant, nature, green, ecology': '🌽',
                'fish, sea, ocan, swimming, water': '🐟',
                'home, house, building, apartment, residence': '🏠',
                'university, official, building, columns, institution': '🏦',
                'school, education, student, learn, diploma': '🏫',
                'education, school, student, learn, diploma': '🎓',
                'child, children, young': '🧒',
                'book, paper, knowledge, reading, library, books, literature': '📖',
                'scroll, paper, document, page, book': '📜',
                'contract, bookmark, tab, sheet, signature': '📑',
                'pencil, write, edit, paper, memo, note': '✏️',
                'pen, write, paper, memo, note, fountain pen': '✒️',
                'military, army, soldier, war, helmet': '🪖',
                'tool, measure, scale, ruler, law, regulation, enforcement': '⚖️',
                'police, cop, urgence, security, law, enforcement, arrest, criminal, law enforcement': '🚓',
                'shield, protection, security, safety, defense': '🛡️',
                'urgence, police, fire, light, warning, danger': '🚨',
                'bomb, explode, explosion, bang, blast, grenade': '💣',
                'fire, flame, hot, heat, blaze, brigade': '🔥',
                'thermometer, hot, temperature, warm, ill, illness, fever': '🌡️',
                'money, bag, dollar, coin': '💰',
                'money, purse, wallet, bag, dollar, euro': '👛',
                'credit, bank, money, loan, bill, payment, credit card': '💶',
                'chart, graph, analytics, statistics, data, report': '📊',
                'money, dollar, currency, payment, bank, banknote, exchange, cash': '💱',
                'money, dollar, currency, payment, bank, banknote, exchange, cash': '💵',
                'money, dollar, currency, payment, bank, banknote, exchange, cash, fly': '💸',
                'shopping, buy, purchase, cart, buy': '🛒',
                'shopping, buy, purchase, shopping cart': '🛍️',
                'travel, luggage, bag, suitcase, bag': '🧳',
                'film, movie, motion, cinema, theater, culture': '🎬',
                'computer, laptop, digital, keyboard, monitor, screen': '💻',
                'lightning, bolt, electricity, science': '⚡',
                'light, bulb, electric, electricity': '💡',
                'flashlight, light, lamp': '🔦',
                'rocket, launch, space, ship, plane, space, start up': '🚀',
                'hospital, medical, center, care, health, sickness, disease, illness': '🏥',
                'clothing, lab, coat, science, laboratory': '🥼',
                'factory, building, manufacturing, production, construction, polution': '🏭',
                'globe, world, earth, planet, map, travel, space': '🌍',
                'location, map, pin, marker, navigation, aid': '📍',
                'europe, european union, flag, country, nation, place, location, geography, globe': '🇪🇺',
                'custom, border, control, security, safety, protection': '🛂',
                'bus, car, transportation, transportation vehicle, trolly': '🚎',
                'alarm, clock, morning, ring, wake up': '⏰',
                'clock, time, timer, watch, stopwatch': '⏱',
                'truck, transportation, delivery, road, vehicule': '🚚',
                'truck, transportation, delivery, road, vehicule': '🚛',
                'key, lock, password, secure': '🔑',
                'trophy, award, cup, competition, game, sport, winner': '🏆',
                'win, medal, gold, silver, bronze, rank, trophy, sport, competition, game, award': '🏅',
                'flex, muscle, body, workout, exercise': '💪',
                'congratulations, party, popper, confetti, celebration': '🎉',
                'ticket, prize, gift, award, prize, gift, admission': '🎟',
                'star, gold, yellow, sky, space, night, evening, dusk': '⭐️',
                'star, astronomy, sparkle, sparkles, magic': '✨',
                'heart, like, favorite, love': '❤️',
                'handshake, agreement, hands': '🤝‍',
                'eye, vision, look, see': '👀',
                'megaphone, announcement, broadcast, public, speaking': '📣',
                'dice, game, chance, roll, random, target, center': '🎯',
                'gift, present, package, box, celebrate, birthday, party': '🎁',
                'balloon, celebration,party, birthday,': '🎈',
                'hourglass, time, timer, watch, stopwatch': '⏳',
                'clap, applause, bravo, hand, gesture, wave, hand clapping': '👏',
                'clown, face, funny, lol, party, hat': '🥳',
                'face, happy, joy, heart, love, emotion, smiley': '🥰',
                'sunglasses, cool, smile, smiley': '😎',
                'laughing, lol, smile, smiley': '😂',
                'open hands, smiley, hug, love, care': '🤗',
                'smiley, face, happy, joy, emotion, smiley': '🙂',
            },
            toggle() {
                if (this.open) {
                    return this.close()
                }

                this.$refs.button.focus()

                this.open = true
            },
            close(focusAfter) {
                if (! this.open) return

                this.open = false

                focusAfter && focusAfter.focus()
            },
            get filteredEmojis() {
                return Object.keys(this.emojis)
                .filter(key => key.includes(this.search))
                .reduce((obj, key) => {
                  obj[key] = this.emojis[key];
                  return obj;
                }, {})
            }
        }
      `),
		alpine.XOn("keydown.escape.prevent.stop", "close($refs.button)"),
		alpine.XOn("focusin.window", "!$refs.panel.contains($event.target) && close()"),
		htmx.ID("dropdown-button"),
		// Hidden input to track the selected emoji
		htmx.Input(
			htmx.Attribute("type", "hidden"),
			htmx.Attribute("name", "emoji"),
			htmx.ID("emoji"),
			alpine.XModel("input"),
		),
		dropdowns.Dropdown(
			dropdowns.DropdownProps{
				ClassNames: htmx.ClassNames{
					"dropdown-right": true,
				},
			},
			dropdowns.DropdownButton(
				dropdowns.DropdownButtonProps{},
				alpine.XRef("button"),
				alpine.XOn("click", "toggle()"),
				htmx.Attribute(":aria-expanded", "open"),
				htmx.Attribute(":aria-controls", "$id('dropdown-button')"),
				htmx.Attribute("type", "button"),
				htmx.Span(
					alpine.XText("input"),
				),
			),
			dropdowns.DropdownMenuItems(
				dropdowns.DropdownMenuItemsProps{
					ClassNames: htmx.ClassNames{
						"menu":      false,
						"w-52":      false,
						"w-96":      true,
						"w-full":    false,
						"z-[1]":     false,
						"z-[99999]": true,
					},
				},
				alpine.XRef("panel"),
				alpine.XShow("open"),
				alpine.XTransition("origin.top.left"),
				alpine.XOn("click.outside", "close($refs.button)"),
				htmx.Attribute(":id", "$id('dropdown-button')"),
				// Emoji search text input
				forms.TextInputBordered(
					forms.TextInputProps{
						ClassNames:  htmx.ClassNames{},
						Placeholder: "Search an emoji...",
					},
					htmx.Attribute("type", "search"),
					alpine.XModel("search"),
				),
				htmx.Template(
					alpine.XFor("(emoji, keywords) in filteredEmojis"),
					htmx.Attribute(":key", "emoji"),
					buttons.Button(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								tailwind.CursorPointer: true,
								tailwind.M1:            true,
							},
						},
						alpine.XOn("click", "input = emoji; toggle();"),
						htmx.Span(
							htmx.Class("inline-block"),
							htmx.Class("w-5"),
							htmx.Class("h-5"),
							alpine.XText("emoji"),
						),
					),
				),
			),
		),
	)
}
