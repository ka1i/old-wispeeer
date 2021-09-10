package usage

import "fmt"

// Wispeeer Usage ...
func Usage() {
	fmt.Println(`Usage: wispeeer [-i <alias>] [-n <title>] -[gsdhv]

If you are using it for the first time,
first execute the command "wispeeer init <Blog directory>" to initialize the blog.

     ------- < Commands Arguments > -------
optional:
  -i, init          Create a new blog. (e.g wispeeer -i <alias>)
  -n, new           Create a new post or page. (e.g wispeeer -n [post] <title>)
  -g, generate      Generate static files. (e.g wispeeer -g)
  -s, server        Start the server. (e.g wispeeer -s)
  -d, deploy        Deploy your website. (e.g wispeeer -d)

  -h, help          Show this help message. 
  -v, version       Show the app version. 

For more help, you can use 'wispeeer help' for the detailed information
or you can check the docs: https://github.com/Wispeeer/wispeeer`)
}
