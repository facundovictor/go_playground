# Matrix - clockwise spiral traverse (from center to outside)

### Scenarios

#### Single element matrix

Max X=1, Y=1
```go
{
  {1},
}
```

#### Uni-dimensional

Horizontal. Max X=1, Y=2
```go
{
  {1,2},
}
```

Vertical, Max X=2, Y=1
```go
{
  {1},
  {2},
}
```

#### Bi-dimensional

Square. Max X=2, Y=2
```go
{
  {1,2},
  {4,3},
}
```

Non-uniform, horizontal. Max X=2, Y=3
```go
{
  {1,2,3},
  {6,5,4},
}
```

Non-uniform, vertical. Max X=2, Y=3
```go
{
  {1,2},
  {6,3},
  {5,4},
}
```

General case. Max X>2, Y>2
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```

### How to print a circle layer

This is slightly more triky than the other algorithm, but still simple enough.

1. From top left (y=0, x=1 leaving x=0 for the last element to visit) to the bottom left (x=rows -1).

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 8,7

2. From bottom left (y=1 to avoid repeating) to bottom right (y=columns - 1)

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 6,5

3. from bottom right (x = rows - 1, y = columns - 1) to top right (x = 0)

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 4,3

4. From top right (x = 0, y = columns - 2 to avoid repeating) to top left (y = 0

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 2,1

#### Implementation

##### Traverse a single layer

For building the function `print_circle`, expecting the bi dimensional matrix `m`, the first row, the last row, the first column and the last column to print, for a general case that does a complete loop, could look like (won't compile):
```go
func print_circle(m[][], X1, X2, Y1, Y2) {
  y = Y1

  for x = X1 + 1; x <= X2; x = x+1
    print(m[x][y])

  for y = Y1 + 1; y <= Y2 - 1; y = y+1
    print(m[x][y])

  for x = X2; x >= X1 ; x = x-1
    print(m[x][y])

  for y = Y2 - 1; y >= Y1; y = y-1
    print(m[x][y])
}
```

But, what about the base cases?

If the amount of columns is 1, or the amount of rows is 1, it is not a complete loop, but a single dimensional array. Taking a deeper look:

An uni dimensional array with 1 row and N columns would have the inputs X1=X2=0, Y1=0, Y2=Columns - 1. The 1st loop in the function would be skipped as it would evaluate 1 <= 0. The 2nd loop would print the array left to right, which is undesired. The 3rd loop would print the last element, and finally the 4th loop would print the remaining part of the array from right to left.

Similarly, an uni dimensional array with 1 column and N rows have the same problems as the inputs would be X1=0, X2=Rows -1 and Y1=Y2=0. The first loop would print the first column from top to bottom which is undesired (the idea is to print it backwards), the 2nd loop would be skipped with the condition 1 <= 0. The 3rd loop would print the array from bottom top as desired. The 4th loop would be skipped as well.

To avoid this problem, we can just check if the array to print is uni dimensional and skip the first 2 loops:
```go
func print_circle(m[][], X1, X2, Y1, Y2) {
  y = Y1

  if Y1 != Y2 && X1 != X2 {
    for x = X1 + 1; x <= X2; x = x+1
      print(m[x][y])

    for y = Y1 + 1; y <= Y2 - 1; y = y+1
      print(m[x][y])
  }

  for x = X2; x >= X1 ; x = x-1
    print(m[x][y])

  for y = Y2 - 1; y >= Y1; y = y-1
    print(m[x][y])
}
```

##### Traverse a multiple layers

The layers would start from inside to outside in an opposite direction compared to the clock.

In this case, it is clear that the traverse would end with the most external layer:
```go
X1 = 0
Y1 = 0
X2 = rows - 1
Y2 = columns - 1
```

Thus, it is reasonable that on each layer, the next one is calculated by:
```go
X1 = X1 + 1
Y1 = Y1 + 1
X2 = X2 - 1
Y2 = Y2 - 1
```

But, how to get the initial values?

A bi dimensional array with equal amount of rows and columns, can have the central layer:
- being a single element for uneven sizes:
```go
{
  {1,2,3},
  {8,9,4},  ==> {9}
  {7,6,5},
}
```
- being a simple 4 element perfect square for pair sizes:
```go
{
  { 1,  2,  3, 4},
  {12, 13, 14, 5},  \__ {13, 14},
  {11, 16, 15, 6},  /   {16, 15},
  {10,  9,  8, 7},
}
```

Thus, the initial values, for an square matrix can be calculated with:
```go
X1 = (rows - 1) / 2
Y1 = (columns - 1) / 2
X2 = rows - 1 - X1
Y2 = columns - 1 - Y2
```

What is the problem with a non-square matrix?

A bi dimensional array with rows != columns can be tricky, as having 5 columns and 3 rows, would have a single dimension array as central layer. Ex:
```go
{
  {  1,  2,  3,  4, 5},
  { 12, 13, 14, 15, 6},  ==> {13, 14, 15}
  { 11, 10,  9,  8, 7},
}
```
Where the initial values in this example seems to be:
```go
X1 = 1 = (rows - 1) / 2 = 1. Which is just like before.
Y1 = 1 != (columns - 1) / 2 = 2. In this case we want the same value as X1.
X2 = 1 = rows - 1 - X1. Same as before.
Y2 = 3 = columns - 1 - Y2. Same as before.
```

Similarly, for bi dimensional arrays with more rows than columns, which would have a vertical single dimensional array as central layer. Ex:
```go
{
  {  1,  2,  3},
  { 12, 13,  4},  \    {13},
  { 11, 14,  5},  |==> {14},
  { 10, 15,  6},  /    {15}
  {  9,  8,  7},
}
```
Where the initial values in this example seems to be:
```go
X1 = 1 != (rows - 1) / 2 = 2. In this case we want the same value as Y1.
Y1 = 1 = (columns - 1) / 2 = 1. Same as before.
X2 = 3 = rows - 1 - X1. Same as before.
Y2 = 1 = columns - 1 - Y2. Same as before.
```

For these cases, we need to know the minor dimension to help calculating the boundaries and size of the central layer.
```go
min_dimension = Min(rows, columns)
X1 = (min_dimension - 1) / 2
Y1 = (min_dimension - 1) / 2
X2 = rows - 1 - X1
Y2 = columns - 1 - Y2
```

Thus, the general loop to traverse the layers is:
```go
start = (min(rows, columns) - 1) / 2

for
  X1,Y1,X2,Y2 = start, start, (rows - 1) - start, (columns - 1) - start;
  X1 >= 0 && Y1 >= 0 && X2 < rows && Y2 < columns;
  X1--,X2++,Y1--,Y2++
{
  print_circle(m, X1, X2, Y1, Y2);
}
```
