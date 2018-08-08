# SATsolver

A naive SAT Solver written in Go

## Example

Given following CNF
```
A B C
-B -C
```
You can parse it like:

```Go
p := parser.Parser{FilePath: "./cnfTestFile.txt"}
parsedDomain := p.ParseFile()
fmt.Println(parsedDomain.solve())
//Expected Output: map[A:1 C:1 B:-1]
```

See the tests for further examples
