# Marker

Marker is a small package for styling text in the terminal using ANSI codes.

## Color and Background Color

Color and background colors can be set using the chainable Color and Background functions.

```go
style := marker.Style().Background(marker.Magenta).Color(marker.Blue)
```

This will generate a style object which can be used as a terminal code.

```go
fmt.Printf("%sThis text should be Blue with a Magenta background. Hideous!", style)
```

`style` is of an unexported type which fulfills the fmt.String() interface.

## Effects

Effects can also be set, and are also chainable!

```go
style := marker.Style().Effect(
	marker.Underline,
).Background(
	marker.Red,
).Color(
	marker.BrtGreen,
)
```

This style will turn into a the terminal code for text that is Bright Green with a Red background and which is underlined.

## Write

The examples above will cause the cursor to change indefinitely into the horrid styles that it was given. To fix that, you could pass a reset code at the end to reset the cursor to its normal style.

```go
style := marker.Style().Color(marker.Green)
fmt.Print("%sThis text is green!%s", style, marker.Reset())
```

The package provides a convenient write method to make this a little easier, though.

```
marker.Style().Color(marker.Green).Write("This text is green!")
```

This will append the message you pass it with the given styles, and will pop a reset on the end for you.