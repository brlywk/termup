package main

import "github.com/brlywk/termup/cmd"

/*
Workflow:
1. Check for homebrew, install if necessary
2. Check for git, install if necessary
3. Check for chezmoi, install if necessary
4. Check if chezmoi config is in ~/.config/chezmoi, clone Repo if necessary
5. Run setup process for chezmoi
6. If all of the above is already done, offer theme / font change options
7. ...
8. Profit!
*/

func main() {
	cmd.Execute()
}
