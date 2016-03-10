package main

import (
	"fmt"
	i "github.com/retep-mathwizard/utils/src/input"
	c "github.com/skilstak/go/colors"
	"math/rand"
)

func main() {

	fmt.Println(c.Clear + c.Multi("Welcome to the Magic Eightball"))
	for {
		promt := i.StringInput(c.Yellow + "What is your question  > ")
		if promt == "love" {
			fmt.Println(c.Rc() + "I dont deal with peoples love lives")
		} else if promt == "death" {
			fmt.Println(c.Rc() + "You are morbid")
		} else {
			rand.Seed(42)
			answers := []string{
				"It is certain",
				"It is decidedly so",
				"Without a doubt",
				"Yes definitely",
				"You may rely on it",
				"As I see it yes",
				"Most likely",
				"Outlook good",
				"Yes",
				"Signs point to yes",
				"Reply hazy try again",
				"Ask again later",
				"Better not tell you now",
				"Cannot predict now",
				"Concentrate and ask again",
				"Don't count on it",
				"My reply is no",
				"My sources say no",
				"Outlook not so good",
				"Very doubtful",
			}
			fmt.Println(c.Rc() + "Magic 8-Ball says:" + answers[rand.Intn(len(answers))])
		}
	}

}
