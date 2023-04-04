# auxo
![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)[![Go](https://github.com/guobinhit/auxo/actions/workflows/go.yml/badge.svg)](https://github.com/guobinhit/auxo/actions/workflows/go.yml)![issues](https://img.shields.io/github/issues/guobinhit/auxo.svg)![stars](https://img.shields.io/github/stars/guobinhit/auxo.svg)![forks](https://img.shields.io/github/forks/guobinhit/auxo.svg)![license](https://img.shields.io/github/license/guobinhit/auxo.svg)

> Auxo, goddess of summer and growth.

# Contributing

Contributions are very welcome!

If you see a problem that you'd like to see fixed, the best way to make it happen is to help out by submitting a pull request implementing it. Refer to the [CONTRIBUTING.md](../master/CONTRIBUTING.md) file for more details about the workflow.

You can also ask for problem-solving ideas and discuss in GitHub issues directly.

# INDEX

- [Overview](#overview)
- [Usage](#usage)

# Overview

In specified package, provides some practical util, as below:

- `format`: format a `time` value to `string`, such as `GetYyyyMmDdHhMmSs`,
- `parse`: parse a `string` value to `time`, such as `GetYyyyMmDdHhMmSs`.

A time toolkit written in Go.

# Usage

Firstly, download this pkg,

```go
go get github.com/guobinhit/auxo
```

Secondly, use it.

```go
import (
    "github.com/guobinhit/auxo"
    "github.com/guobinhit/auxo/format"
    "github.com/guobinhit/auxo/parse"
)

// Get a specified time by add days, such as d is 2022-04-13 10:20:30 and days is 10,
// then aDate is 2022-04-23 10:20:30
aDate := auxo.GetTimeAddDays(time.Now(), 10)

// Get a specified date format time string of common version, such as d is 2022-04-13 10:20:30,
// then aString is "2022-04-23 10:20:30"
aString := format.GetYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time string of china version, such as d is 2022-04-13 10:20:30,
// then aString2 is "2022年04月23日 10:20:30", like GetCnOfYyyyMmDdHhMmSs method, 
// GetEnOfYyyyMmDdHhMmSs return "2022/04/23 10:20:30" and GetCptOfYyyyMmDdHhMmSs return "20220423102030".
aString2 := format.GetCnOfYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time, such as dStr is "2022-04-13 10:20:30",
// then aTime is 2022-04-23 10:20:30
aTime, err := parse.GetYyyyMmDdHhMmSs("2022-04-13 10:20:30")
```

Finally, good luck guys!
