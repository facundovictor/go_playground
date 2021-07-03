package main

import "fmt"

var columns int = 10;
var rows int = 5;
var m = [][]int{
  { 1,  2,  3,  4,  5,  6,  7,  8,  9, 10},
  {26, 27, 28, 29, 30, 31, 32, 33, 34, 11},
  {25, 44, 45, 46, 47, 48, 49, 50, 35, 12},
  {24, 43, 42, 41, 40, 39, 38, 37, 36, 13},
  {23, 22, 21, 20, 19, 18, 17, 16, 15, 14},
};
// var columns int = 3;
// var rows int = 3;
// var m = [][]int{
//   {1, 2, 3},
//   {8, 9, 4},
//   {7, 6, 5},
// };
// var columns int = 3;
// var rows int = 1;
// var m = [][]int{
//   {1, 2, 3},
// };
// var columns int = 1;
// var rows int = 3;
// var m = [][]int{
//   {1},
//   {2},
//   {3},
// };
// var columns int = 2;
// var rows int = 3;
// var m = [][]int{
//   {1, 2},
//   {6, 3},
//   {5, 4},
// };
// var columns int = 3;
// var rows int = 2;
// var m = [][]int{
//   {1, 2, 3},
//   {6, 5, 4},
// };
// var columns int = 4;
// var rows int = 4;
// var m = [][]int{
//   { 1,  2,  3, 4},
//   {12, 13, 14, 5},
//   {11, 16, 15, 6},
//   {10,  9,  8, 7},
// };


func print_circle(m [][]int, X1 int, X2 int, Y1 int, Y2 int) {
  var x, y int = X1, Y2;

  for y = Y1; y <= Y2; y = y+1 {
    fmt.Print(m[x][y],",")
  }
  y=Y2

  for x = X1 + 1; x <= X2; x = x+1 {
    fmt.Print(m[x][y],",")
  }
  x=X2

  if Y1 == Y2 || X1 == X2 { return }

  for y = Y2 - 1; y >= Y1; y = y-1 {
    fmt.Print(m[x][y],",")
  }
  y=Y1

  for x = X2 - 1; x > X1 ; x = x-1 {
    fmt.Print(m[x][y],",")
  }
}


func main() {
  var X1, X2, Y1, Y2 int;

  for
    X1,Y1,X2,Y2 = 0, 0, rows - 1, columns - 1;
    X1 <= X2 && Y1 <= Y2;
    X1,X2,Y1,Y2 = X1+1, X2-1, Y1+1, Y2-1 {

    print_circle(m, X1, X2, Y1, Y2);
  }
  fmt.Println();
}
