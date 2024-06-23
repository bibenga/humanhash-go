// original code is https://github.com/django-q2/django-q2/blob/master/django_q/humanhash.py
package humanhashgo

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var DefaultWordlist = []string{
	"ack",
	"alabama",
	"alanine",
	"alaska",
	"alpha",
	"angel",
	"apart",
	"april",
	"arizona",
	"arkansas",
	"artist",
	"asparagus",
	"aspen",
	"august",
	"autumn",
	"avocado",
	"bacon",
	"bakerloo",
	"batman",
	"beer",
	"berlin",
	"beryllium",
	"black",
	"blossom",
	"blue",
	"bluebird",
	"bravo",
	"bulldog",
	"burger",
	"butter",
	"california",
	"carbon",
	"cardinal",
	"carolina",
	"carpet",
	"cat",
	"ceiling",
	"charlie",
	"chicken",
	"coffee",
	"cola",
	"cold",
	"colorado",
	"comet",
	"connecticut",
	"crazy",
	"cup",
	"dakota",
	"december",
	"delaware",
	"delta",
	"diet",
	"don",
	"double",
	"early",
	"earth",
	"east",
	"echo",
	"edward",
	"eight",
	"eighteen",
	"eleven",
	"emma",
	"enemy",
	"equal",
	"failed",
	"fanta",
	"fifteen",
	"fillet",
	"finch",
	"fish",
	"five",
	"fix",
	"floor",
	"florida",
	"football",
	"four",
	"fourteen",
	"foxtrot",
	"freddie",
	"friend",
	"fruit",
	"gee",
	"georgia",
	"glucose",
	"golf",
	"green",
	"grey",
	"hamper",
	"happy",
	"harry",
	"hawaii",
	"helium",
	"high",
	"hot",
	"hotel",
	"hydrogen",
	"idaho",
	"illinois",
	"india",
	"indigo",
	"ink",
	"iowa",
	"island",
	"item",
	"jersey",
	"jig",
	"johnny",
	"juliet",
	"july",
	"jupiter",
	"kansas",
	"kentucky",
	"kilo",
	"king",
	"kitten",
	"lactose",
	"lake",
	"lamp",
	"lemon",
	"leopard",
	"lima",
	"lion",
	"lithium",
	"london",
	"louisiana",
	"low",
	"magazine",
	"magnesium",
	"maine",
	"mango",
	"march",
	"mars",
	"maryland",
	"massachusetts",
	"may",
	"mexico",
	"michigan",
	"mike",
	"minnesota",
	"mirror",
	"mississippi",
	"missouri",
	"mobile",
	"mockingbird",
	"monkey",
	"montana",
	"moon",
	"mountain",
	"muppet",
	"music",
	"nebraska",
	"neptune",
	"network",
	"nevada",
	"nine",
	"nineteen",
	"nitrogen",
	"north",
	"november",
	"nuts",
	"october",
	"ohio",
	"oklahoma",
	"one",
	"orange",
	"oranges",
	"oregon",
	"oscar",
	"oven",
	"oxygen",
	"papa",
	"paris",
	"pasta",
	"pennsylvania",
	"pip",
	"pizza",
	"pluto",
	"potato",
	"princess",
	"purple",
	"quebec",
	"queen",
	"quiet",
	"red",
	"river",
	"robert",
	"robin",
	"romeo",
	"rugby",
	"sad",
	"salami",
	"saturn",
	"september",
	"seven",
	"seventeen",
	"shade",
	"sierra",
	"single",
	"sink",
	"six",
	"sixteen",
	"skylark",
	"snake",
	"social",
	"sodium",
	"solar",
	"south",
	"spaghetti",
	"speaker",
	"spring",
	"stairway",
	"steak",
	"stream",
	"summer",
	"sweet",
	"table",
	"tango",
	"ten",
	"tennessee",
	"tennis",
	"texas",
	"thirteen",
	"three",
	"timing",
	"triple",
	"twelve",
	"twenty",
	"two",
	"uncle",
	"undress",
	"uniform",
	"uranus",
	"utah",
	"vegan",
	"venus",
	"vermont",
	"victor",
	"video",
	"violet",
	"virginia",
	"washington",
	"west",
	"whiskey",
	"white",
	"william",
	"winner",
	"winter",
	"wisconsin",
	"wolfram",
	"wyoming",
	"xray",
	"yankee",
	"yellow",
	"zebra",
	"zulu",
}

type HumanHasher struct {
	Wordlist  []string
	Words     int
	Separator string
}

func (h *HumanHasher) Humanize(data []byte) (string, error) {
	compressed, err := h.compress(data)
	if err != nil {
		return "", err
	}

	var words []string = make([]string, h.Words)
	for i, d := range compressed {
		words[i] = h.Wordlist[d]
	}
	return strings.Join(words, h.Separator), nil
}

func (h *HumanHasher) compress(data []byte) ([]byte, error) {
	length := len(data)
	if h.Words > length {
		return nil, errors.New("fewer input bytes than requested output")
	}

	seg_size := length / h.Words
	var checksums []byte = make([]byte, h.Words)
	for i := 0; i < h.Words; i++ {
		b := i * seg_size
		e := (i + 1) * seg_size
		if e > length {
			e = length
		}
		if i == h.Words-1 {
			e = length
		}
		segment := data[b:e]
		c := byte(0)
		for _, d := range segment {
			c = c ^ d
		}
		checksums[i] = c
	}
	return checksums, nil
}

func (h *HumanHasher) Uuid(value uuid.UUID) (string, error) {
	humanized, err := h.Humanize(value[:])
	if err != nil {
		return "", err
	}
	return humanized, err
}

func (h *HumanHasher) NewUuid() (uuid.UUID, string, error) {
	value, err := uuid.NewRandom()
	if err != nil {
		return value, "", err
	}
	humanized, err := h.Humanize(value[:])
	if err != nil {
		return value, "", err
	}
	return value, humanized, err
}

var DefaultHasher = HumanHasher{
	Wordlist:  DefaultWordlist,
	Words:     4,
	Separator: "-",
}
var Humanize = DefaultHasher.Humanize
var NewUuid = DefaultHasher.NewUuid
