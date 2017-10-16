//Package sbotFuncs contains functions for slack bots
package sbotFuncs

import (
  "math/rand"
  "time"
)

//Return random joke
func GetJoke() string {
  jokes := []string{
    "What do you call a pirate droid?\n ```Argh2-D2```",
    "What do Gungans put things in?\n ```Jar Jars```",
    "Who tries to be a Jedi?\n ```Obi-Wannabe```",
    "What do Whipids say when they kiss?\n ```Ouch```",
    "What do Jedi use to view PDF files?\n ```Adobe Wan Kenobi```",
    "Why did the smuggler cross the spacelanes?\n ```To get to the other side```",
    "Why should you never tell jokes on the Falcon?\n ```The ship might crack up```",
    "What do you call a potato that has turned to the Dark side?\n ```Vader Tots```",
    "What do you call a person who brings a rancor its dinner?\n ```The appetizer```",
    "Why shouldn’t you ask Yoda for money?\n ```Because he’s always a little short```",
    "Why do Twi’leks like to flip coins?\n ```So that they can say, 'Heads or tails!'```",
    "Why is a droid mechanic never lonely?\n ```Because he’s always making new friends!```",
    "Why is a droid mechanic never lonely?\n ```Because he’s always making new friends```",
    "What does Yoda say to encourage a Padawan before a test?\n ```Do well, you will do!```",
    "What happens when a red and white X-Wing crashes into green water?\n ```It gets wet```",
    "Why does Princess Leia keep her hair tied up in buns?\n ```So it doesn’t Hang Solow!```",
    "Why didn’t Luke Skywalker cross the road?\n ```Because he got a ticket for Skywalking```",
    "Why is the Millenium Falcon so slow?\n ```Because it takes a millenium to go anywhere```",
    "Why does Leia wear buns on her head?\n ```In case she gets hungry in a Senate meeting```",
    "Why did the crazy Angrallian Toobir cross the nebula?\n ```To get to the other dementia```",
    "As a Disney character what song would Vader sing?\n ```'When You Wish Upon A Death Star'```",
    "What time is it when an AT-AT steps on your chronometer?\n ```Time to get a new chronometer```",
    "Where does Princess Leia go shopping for clothing and such?\n ```At the Darth Maul, of course```",
    "What do you call Chewbacca when he has chocolate stuck in his hair?\n ```Chocolate Chip Wookiee```",
    "What do you call the website Chewbacca started that gives out Empire secrets?/n ```Wookieeleaks```",
    "Which Star Wars character uses meat for a weapon instead of a Lightsaber?\n ```Obi Wan Baloney```",
    "Why did Kit Fisto storm out of the sushi restaurant?\n ```Because they were serving Mon Calamari```",
    "When did Anakin’s Jedi masters know he was leaning towards the dark side?\n ```In the Sith Grade```",
    "How many Sith does it take to screw in a hyperdrive?\n ```Two, but I don’t know how they got in it```",
    "What goes, “Ha, ha, ha, haaaa…. AGGGHHHH! Thump”?\n ```An Imperial Officer laughing at Darth Vader```",
    "What do you call it when only one Star Wars character gives you a round of applase?\n ```A Hand Solo!```",
    "How is Ducktape like the Force?\n ```It has a Dark Side, a Light side and it binds the galaxy together```",
    "Where does Qui-Gon keep his jam? A: In a Jar-Jar. Why did Padme Amidal keep her Boots on?\n ```Because they were too BOOT-iful!```",
    "What’s the differance between an ATAT and a stormtrooper?\n ```One’s an Imperial```",
    "How many stormtroopers does it take to replace a lightbulb?\n ```Two; one to screw the bulb in, the other to shoot him and take the credit```",
    "What do you call a bounty hunter from the South?\n ```Bubba Fett```",
    "Why did the Stormtrooper start jumping up and down?\n ```He stepped on Ant-hillies```",
  }
  rand.Seed(time.Now().UTC().UnixNano())
  random := rand.Intn(len(jokes))
  return jokes[random]
}
