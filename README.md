# Consistent Hashing in Go

This is a basic implementation of consistent hashing in Go. Consistent hashing is a technique used in distributed computing to evenly distribute data across multiple nodes in a cluster.

## Installation

To use this package in your Go project, you can install it using `go get`:

```bash
go get github.com/irononet/consistent_hashing
```

## Usage

Here's an example of how to use this package:

```go
package main

import (
	"fmt"
	"github.com/irononet/consistent_hashing"
)

func main() {
	// Create a new ConsistentHash instance with 3 replicas and 3 nodes.
	chash := consistent_hashing.NewConsistentHash(3, []string{"node1", "node2", "node3"})

	// Get the node that each key belongs to.
	node := chash.GetNode("my_key")

	fmt.Printf("my_key belongs to node %s\n", node)
}
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub. If you want to contribute code, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
