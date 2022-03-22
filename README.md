# go-moneroocean

go-moneroocean is a Go client library for accessing the public facing API of MoneroOcean.

## Table of Contents
* [Installation](#installation)
* [Usage](#usage)
* [To-do](#todo)
* [Contributing](#contributing)
* [License](#license)

### Installation

go-moneroocean is compatible with modern Go releases in module mode, to install the latest version:
```shell
go get github.com/thelolagemann/go-moneroocean
```

Or to install a specific version:
```shell
go get github.com/thelolagemann/go-moneroocean@x.y.z
```

### Usage
<sup>*Please note: error handling has been omitted from examples for the sake of brevity*</sup>

```go
package main

import (
	"fmt"
	mo "github.com/thelolagemann/go-moneroocean"
)

func main() {
	stats, _ := mo.Stats("ADDRESS")
	fmt.Println(stats.AmountDue)
	
	// Outputs:
	// Your XMR amount due on MoneroOcean
}
```

### To-do
- [ ] improve upon documentation
- [ ] add more examples
- [ ] more thorough testing

### Contributing

All contributions are welcome! For any bug reports/feature requests, feel free to open an issue or submit a pull request. 

### License
This project is licensed under the MIT license - see the [LICENSE](https://github.com/thelolagemann/go-moneroocean/blob/main/LICENSE) file for details.