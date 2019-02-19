package main

import (
	"fmt"
	"strings"

	"github.com/prosline/link"
)

var exampleHTML = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://twitter.com/marchito8">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/prosline">
      My short portfolio on GO in  <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v \n", links)
}
