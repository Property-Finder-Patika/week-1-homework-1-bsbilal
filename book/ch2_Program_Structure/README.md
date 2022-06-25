#-
The letters of acronyms and initialisms like ASCII and HTML are always rendered in the same cas e, so a function might be called html-Escape, HTMLEscape, or escapeHTML, but not escapeHtml.
var-const-type and func.

#-
A set of var iables can also be initialize d by cal lingafunction that retur ns multiple values:
var f, err = os.Open(name) // os.Open returns a file and an error

#-
One subtle but imp ortant point:ashort var iable decl arat ion does not necessarily decl are al l the
var iables on its left-hand side. If som e of them were already declare d in the same lexic al block
(§2.7), then the short var iable decl arat ion acts like an assig nment to those var iables.

![image](https://user-images.githubusercontent.com/9121424/175786025-d58b256a-e2d2-47e4-9dff-00fd1f22e751.png)
#- 
The zero value for a point er of any typ e is nil. The test p != nil is true if p points toavar i-
able. Point ers are comparable; two point ers are equ al if and only if the y point to the same
var iable or bot h are nil.
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

 

flag******



// pointers
#-
struct{} or [0]int, may, depending on the implementation,
have the same address.
##
// type definations
fmt.Println(c == Celsius(f)) // "true"!


