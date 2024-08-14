## here is a small documentation(or whatever interesting) for personal reference :D

- [go docs strings](https://go.dev/blog/strings)
- [There Ain’t No Such Thing As Plain Text.](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- [pipelining bash commands](https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go#comment59574308_10781582)
- [more on pipelining](https://www.educative.io/answers/how-to-pipe-several-commands-in-go)
- [quich tour of Unicode](https://pro.arcgis.com/en/pro-app/latest/help/data/geodatabases/overview/a-quick-tour-of-unicode.htm)

## Code points, characters, and runes

- strings and characters aren't the same, strings hold bytes and not characters,
rune is basically an alias for int32, it makes the program clear whether when the int value represents a _code point_

## What is a Code point ?

- Unicode provides a unique number (a code point). Character encoding defines each character-code point.

## Conclusion 

- Know what character encoding you are using, as it may not be easy to convert between one encoding to another, using Unicode is usually the best idea as it can interpret nearly every known character.
