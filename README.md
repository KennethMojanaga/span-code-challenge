# span-code-challenge

## Usage
Assuming that you've installed `GO`, and also set your `GOPATH`, you can run the code by following the steps below
- Change direcory to the root of `span-code-challenge` directory
- Run `go install`. This will generate a binary name `span-code-challenge` in the `bin` folder of your `GOPATH` directory
- Then you can proceed to run the binary by providing either `-input-string` or `-input-file` flags. E.g.
```
./span-code-challenge -input-string "Team A 3, Team B 4\nTeam-C 1, Team D 2"
```

or

```
./span-code-challenge -input-file "./data.txt"
```

Where `data.txt` is a file containing the input data.