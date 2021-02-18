# chronicle
![alt text](logo.jpeg)
Chronicle allows client to register set of workers/ cron function to be triggered in defined interval of time.

## Installation
```
go get -u github.com/im-adarsh/chronicles
```
## How to use
```

import (
	"context"
	"fmt"
	"time"

	"github.com/im-adarsh/chronicle"
)

func main() {

	ctx := context.Background()
	c := chronicle.NewChronicle()
	c.Register(ctx, "say hello", time.Second*2, printHello)
	c.Register(ctx, "say world", time.Second*4, printWorld)
	c.Start()
}

func printHello() error {
	fmt.Println("hello")
	return nil
}

func printWorld() error {
	fmt.Println("world")
	return nil
}
```
## Output
```
Starting 2 tasks
Starting task :  say world
Starting task :  say hello
[Tick say hello] at 2021-02-18 22:08:49.342259 +0800 +08 m=+2.005345691 
hello
[Tick say hello] at 2021-02-18 22:08:51.339545 +0800 +08 m=+4.002659408 
hello
[Tick say world] at 2021-02-18 22:08:51.339548 +0800 +08 m=+4.002663066 
world
[Tick say hello] at 2021-02-18 22:08:53.341993 +0800 +08 m=+6.005134705 
hello
[Tick say world] at 2021-02-18 22:08:55.339371 +0800 +08 m=+8.002537162 
world
[Tick say hello] at 2021-02-18 22:08:55.339367 +0800 +08 m=+8.002533307 
hello
[Tick say hello] at 2021-02-18 22:08:57.341603 +0800 +08 m=+10.004793468 
hello
[Tick say world] at 2021-02-18 22:08:59.339757 +0800 +08 m=+12.002969588 
world
[Tick say hello] at 2021-02-18 22:08:59.339753 +0800 +08 m=+12.002965707 
hello

```