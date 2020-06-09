# Threads
- main thread executes the run method
- the run method executes each person's goroutine
- the run method executes the select statement within an anonimous goroutine
- each person communicates through the `onMove` and `onExit` channels to the select statement
- the run method keeps executing an update method to re-render the canvas
![](assets/threads.png)
# Libraries
* built-in
  * bufio
  * fmt
  * math
  * math/rand
  * os
  * strconv
  * strconv
  * strings
  * time
* installed
  * [pixel](https://github.com/faiface/pixel)
    * pixel/imdraw
    * pixel/pixelgl
    * pixel/text
  * [colornames](https://godoc.org/golang.org/x/image/colornames)
  * [font/basicfont](https://godoc.org/golang.org/x/image/font/basicfont)
# Structs
* coordinate
  * row (int): x coordinate, or row number
  * col (int): y coordinate, or column number
* person
  * id (int): person's id
  * peed (float32): person's slowness, the bigger the slower
  * exited (bool): wheter a person has reached the exit or not
  * path ([]coordinate): the path to the exit represented as an array of coordinates
  * position (int): number of places moved from the starting point
  * curr_position (coordinate): current coordinate where person is
# METHODS
* initializePast()
* printBuilding()
* printPast()
* printPathMatrix()
* getNumOfPeople()
* generateExits(floor [][]int)
* insertExit(floor [][]int, sideExits []int, side int, indexExit int, lenF int) bool
* distance(a coordinate, b coordinate) float64
* findClosestExit(position coordinate) int
* validate(row int, col int) bool
* searchPath(row int, col int) bool
* createWindow() \*pixelgl.Window
* drawFloor(win \*pixelgl.Window) \*imdraw.IMDraw
* drawPeople(win \*pixelgl.Window) \*imdraw.IMDraw
* printLabels(win \*pixelgl.Window)
* run()
* movePerson(p person)
* generateRandomSpeed() float32
* initiatePerson(p person, onMove chan person, onExit chan person, trapped []person)
* main()
