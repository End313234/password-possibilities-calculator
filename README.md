# Basic examples

```go
// Should return 2
CalculatePasswordPossibilities([]string{"a", "1"}, 2, []Criteria{})

// The lack of criteria compared to the quantity of chars in the
// password is the same thing as considering that the other values
// don't have any criteria, so they can be any item in the
// possibilities array.
//
// Should return 300
CalculatePasswordPossibilities([]string{"a", "b", "c", "1", "2"}, 5, []Criteria{LETTER, LETTER_NOT_USED, NUMBER})
```
