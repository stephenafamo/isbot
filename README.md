# IsBot

Detect bots/crawlers/spiders using the user agent string.

Using a list downloaded from <https://user-agents.net>, the regexes in this package hit a 90% detection rate.  
If nothing matches, it also checks if the user agent is present in the list.

## Usage

```go
import "github.com/stephenafamo/isbot"

func main() {
    userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.0.0 Safari/537.36 Chrome-Lighthouse"
    isBot := isbot.Check(userAgent)
    if isBot {
        // do something
    }
}
```

## Sources

The sources for detection are:

* <https://github.com/monperrus/crawler-user-agents>
* <https://user-agents.net>
* A manual list

## Contributing

Run `go generate` to refresh the lists from the sources.

If something is misisng, send in a pull request.
