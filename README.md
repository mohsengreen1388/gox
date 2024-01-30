### <p style="text-align: center;">New Desktop Gui For Golang</p>

### <p style="text-align: center;">![image info](./assata/bitmap.png)</p>

<hr>

#### this is New Gui App (My goal is  make easy-use Gui I try avoid Complexity ) For Linux and Window And Mac os  andriod in the future 
 

> **_NOTE:_**  this version is beta not stable 
may have bug just test on the Ubuntu 20.4 and Window

features
--------
- Button
- Input 
- CheckBox
- RadioButton
- Icons
- MenuBar
- MenuFloat
- Notify
- TabBar
- SlideBar
- Topbar
- Image 
- Progress
- SideBar
- Board
- Modal

in th futhre
------------
- Flex
- Select
- multi windows
- New style
- Card
- multi languages (now only English Persion and Arbic is Ready Add After Test )
- Spinner
- List
- Table
- Accordions
- ....

dependency
----------
first you have to install Dependency

```go
/// this package use raylib-go
go get "github.com/gen2brain/raylib-go/raylib"
// this is temporary delete then
go get "github.com/ncruces/zenity" 
go get "golang.org/x/exp/slices"

// font text render
go get "github.com/benoitkugler/textlayout/fonts/glyphsnames"
go get "github.com/go-text/typesetting/font"
go get "github.com/go-text/typesetting/harfbuzz"
go get "github.com/go-text/typesetting/language" 
go get "github.com/lutzky/go-bidi"
go get "golang.org/x/text/unicode/bidi"

/// or inside project
go get . 
```

Init Project
``` go

package main

import (
	gox "app1/gox/widgets"
)

func main() {
    app := gox.Init("Calculator", 300, 330)

        /// here write instance New

    gox.Update(app, func() {
        
        /// here draw object
   
   }    
}

```



















