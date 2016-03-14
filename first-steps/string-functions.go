package main

import s "strings"
import "fmt"

var p = fmt.Println

const sentence = `
In The Restaurant at the End of the Universe (published in 1980), Zaphod is
separated from the others and finds he is part of a conspiracy to uncover who
really runs the Universe. Zaphod meets Zarniwoop, a conspirator and editor for
The Guide, who knows where to find the secret ruler. Zaphod becomes briefly
reunited with the others for a trip to Milliways, the restaurant of the title.

Zaphod and Ford decide to steal a ship from there, which turns out to be a
stunt ship pre-programmed to plunge into a star as a special effect in a stage
show. Unable to change course, the main characters get Marvin to run the
teleporter they find in the ship, which is working other than having no
automatic control (someone must remain behind to operate it), and Marvin
seemingly sacrifices himself. Zaphod and Trillian discover that the Universe
is in the safe hands of a simple man living on a remote planet in a wooden
shack with his cat.`

func main() {

	p("Contains:  ", s.Contains(sentence, "Zaphod"))
	p("Contains:  ", s.Contains(s.ToLower(sentence), "zaphod"))

	p("Count:     ", s.Count(sentence, "it"))

	p("HasPrefix: ", s.HasPrefix("hitchhiker", "hitch"))
	p("HasSuffix: ", s.HasSuffix("hitchhiker", "er"))

	p("Index:     ", s.Index(sentence, "Trillian"))
	p("Join:      ", s.Join([]string{"h", "i", "t", "c", "h"}, "\\"))
	p("Repeat:    ", s.Repeat("cat ", 3))

	p("Replace:   ", s.Replace("test", "e", "3", 1))
	p("Replace:   ", s.Replace("test", "t", "7", 2))

	p("Split:     ", s.Split("h i t c h", " "))

	p("ToLower:   ", s.ToLower("Hitch"))
	p("ToUpper:   ", s.ToUpper("hitch"))
	p()
	p("Len: ", len("htich"))
	p("Char:", "hitch"[3])

}
