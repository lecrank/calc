## Часть 2. Calc

Нужно написать калькулятор, умеющий вычислять выражение, подаваемое на STDIN.

Достаточно реализовать сложение, вычитание, умножение, деление и поддержку скобок.

### Пример работы

```bash
    $ go run calc.go "(1+2)-3"
    0

    $ go run calc.go "(1+2)*3"
    9
```

## Test Passed:
```
Running tool: /snap/bin/go test -timeout 30s -run ^TestFindUnique$ github.com/lecrank/calc/calculator

=== RUN   TestFindUnique
--- PASS: TestFindUnique (0.00s)
PASS
ok      github.com/lecrank/calc/calculator      0.003s


> Test run finished at 4/20/2023, 8:45:36 PM <
```