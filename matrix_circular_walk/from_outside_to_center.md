# Matrix - clockwise spiral traverse

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

1. From left (y=0) to right (y=columns - 1) the first row (x=0).

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 1,2,3

2. From top (x=1 to avoid repeating) to bottom (x=rows - 1) the last column (y= columns - 1)

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 4,5

3. From bottom right (x = rows - 1, y = columns - 2 to avoid repeating) to bottom left (y = 0).

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 6,7

4. From bottom left (x = rows - 2 to avoid repeating, y=0) to top left (x = 1, to avoid repeating).

Ex for general case:
```go
{
  {1,2,3},
  {8,9,4},
  {7,6,5},
}
```
Output: 8

And that is it, what is pending is the inner circle, which is a base case.

#### Implementation

##### Traverse a single layer

For building the function `print_circle`, expecting the bi dimensional matrix `m`, the first row, the last row, the first column and the last column to print, for a general case that does a complete loop, could look like (won't compile):
```go
func print_circle(m[][], X1, X2, Y1, Y2) {
  for y = Y1; y <= Y2; y = y+1
    print(m[x][y])

  for x = X1 + 1; x <= X2; x = x+1
    print(m[x][y])

  for y = Y2 - 1; y >= Y1; y = y-1
    print(m[x][y])

  for x = X2 - 1; x > X1 ; x = x-1
    print(m[x][y])
}
```

But, what about the base cases?

If the amount of columns is 1, or the amount of rows is 1, it is not a complete loop. Taking a deeper look:

An uni dimensional array with 1 row and N columns can have the inputs X1=X2=0, Y1=0, Y2=Columns - 1. The function would traverse it with the first loop, the second loop would skip it as X1 + 1 > X2. But, the 3th loop would try to print the array in the opposite direction, starting from the element before the last. all the way left to the 2nd element. Finally, the 4th loop, would try going from -1 to 0.

Similarly, an uni dimensional array with 1 column and N rows have the same problems as the inputs would be X1=0, X2=Rows -1 and Y1=Y2=0. The first loop would print the first row with a single element, the second loop would print the whole array, the 3rd loop would go from -1 to 0, and the 4th loop would try to print from bottom to top.

To avoid this problem, we can just check if the loop to print is uni dimensional and skip the latest 2 loops:
```go
func print_circle(m[][], X1, X2, Y1, Y2) {
  for y = Y1; y <= Y2; y = y+1
    print(m[x][y])

  for x = X1 + 1; x <= X2; x = x+1
    print(m[x][y])

  if Y1 == Y2 || X1 == X2 { return } // <=== Check if m is uni dimesional and skip

  for y = Y2 - 1; y >= Y1; y = y-1
    print(m[x][y])

  for x = X2 - 1; x > X1 ; x = x-1
    print(m[x][y])
}
```

##### Traverse a multiple layers

The layers would start from outside to inside in a clockwise.

Using again X1,X2,Y1,Y2 are initialized as:
```go
X1 = 0            // <=== Start vertical position
Y1 = 0            // <=== Start horizontal position
X2 = rows - 1     // <=== End vertical position
Y2 = columns - 1  // <=== End horizontal position
```

2nd iteration would continue with
```go
X1 = 2
Y1 = 2
X2 = rows - 2
Y2 = rows - 2
```

3rd iteration would be similar, so we can make a general loop to traverse the layers as:
```go
for
  X1,Y1,X2,Y2 = 0, 0, rows - 1, columns - 1;
  X1 <= X2 && Y1 <= Y2;
  X1++,X2--,Y1++,Y2--
{
  print_circle(m, X1, X2, Y1, Y2);
}
```
