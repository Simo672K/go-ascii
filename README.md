# ASCII art built with go

An ASCII art cli application that converts images to ascii art

```bash
$ go run main.go -density=<number> <image name>
```

where density is a number that specify number of pixels for character, default is 5 which mean for every 5x5 pixels replace it with a character.

PS: currently only working with .png images
