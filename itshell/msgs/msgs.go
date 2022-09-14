package msgs

const (
	Welcome string = `
Welcome to Dockeroller!
Enter your desired number:
1 - help
2 - gates
Feel free to enter 0 to get back to the main menu.
`

	Help = `
Dockeroller is an open-source project made for fun
But it appreas that it has many real-world use-cases and it is a part of ChatOps world!
You can find more information on it's github page!
WE WON'T STORE ANY OF YOUR DATA, 
It's an open-source project that you can check, make sure and build it yourself.
As long as I know, this project is safe, but use it on your own risk.
This shell should remain open as long as you want this app to be working.
`
	// TODO: Gates can't be a constant, it should be a variable.
	Gates = `
Here are the avaiable gates:
1 - Telegram
2 - API     
You can turn them on by entering their number. (or 0 for main menu)
`

	Telegram = `
You should activate telegram by making a bot and enter its token here.
Also, you should provide your telegram username,
Additional security layers are not needed as long as you keep your telegram account safe,
By the way, we're working on password mechanism.
`

	API = `
You should activate API by defining a superadmin password and a port
`
)
