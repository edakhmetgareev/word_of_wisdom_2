package server

import "math/rand"

var quotes = []string{
	"Quote 1: Wisdom is the reward you get for a lifetime of listening when you'd have preferred to talk.",
	"Quote 2: The only true wisdom is in knowing you know nothing.",
	"Thomas Edison: Genius is one percent inspiration and ninety-nine percent perspiration.",
	"Yogi Berra: You can observe a lot just by watching.",
	"Abraham Lincoln: A house divided against itself cannot stand.",
	"Johann Wolfgang von Goethe: Difficulties increase the nearer we get to the goal.",
	"Byron Pulsifer: Fate is in your hands and no one elses",
	"Lao Tzu: Be the chief but never the lord.",
	"Carl Sandburg: Nothing happens unless first we dream.",
	"Aristotle: Well begun is half done.",
	"Yogi Berra: Life is a learning experience, only if you learn.",
	"Margaret Sangster: Self-complacency is fatal to progress.",
	"Buddha: Peace comes from within. Do not seek it without.",
	"Byron Pulsifer: What you give is what you get.",
	"Iris Murdoch: We can only learn to love by loving.",
	"Karen Clark: Life is change. Growth is optional. Choose wisely.",
	"Wayne Dyer: You'll see it when you believe it.",
}

func getRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
