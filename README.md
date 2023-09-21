# timekit
![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)[![Go](https://github.com/guobinhit/timekit/actions/workflows/go.yml/badge.svg)](https://github.com/guobinhit/timekit/actions/workflows/go.yml)![issues](https://img.shields.io/github/issues/guobinhit/timekit.svg)![stars](https://img.shields.io/github/stars/guobinhit/timekit.svg)![forks](https://img.shields.io/github/forks/guobinhit/timekit.svg)![license](https://img.shields.io/github/license/guobinhit/timekit.svg)

# Contributing

Contributions are very welcome!

If you see a problem that you'd like to see fixed, the best way to make it happen is to help out by submitting a pull request implementing it. Refer to the [CONTRIBUTING.md](../master/CONTRIBUTING.md) file for more details about the workflow.

You can also ask for problem-solving ideas and discuss in GitHub issues directly.

# Usage

Firstly, download this pkg,

```go
go get github.com/guobinhit/timekit
```

Secondly, use it.

```go
import (
    "github.com/guobinhit/timekit"
)

// Get a specified time by add days, such as d is 2022-04-13 10:20:30 and days is 10,
// then aDate is 2022-04-23 10:20:30
aDate := timekit.GetTimeAddDays(time.Now(), 10)

// Get a specified date format time string of common version, such as d is 2022-04-13 10:20:30,
// then aString is "2022-04-23 10:20:30"
aString := timekit.FormatYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time string of china version, such as d is 2022-04-13 10:20:30,
// then aString2 is "2022年04月23日 10:20:30", like FormatCnOfYyyyMmDdHhMmSs method, 
// FormatEnOfYyyyMmDdHhMmSs return "2022/04/23 10:20:30" and FormatCptOfYyyyMmDdHhMmSs return "20220423102030".
aString2 := timekit.FormatCnOfYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time, such as dStr is "2022-04-13 10:20:30",
// then aTime is 2022-04-23 10:20:30
aTime, err := timekit.ParseYyyyMmDdHhMmSs("2022-04-13 10:20:30")
```

Finally, good luck guys!
