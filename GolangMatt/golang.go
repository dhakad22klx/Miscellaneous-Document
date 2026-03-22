//---------- Slices --------//

1. Slices has descriptor (pointer - pointer to array , length, capacity).
2. 



//------- Formatted and File I/O --------//

1. %p : to print the memory address where element lies
2. 

//---- Functions ------//


 
 -->Parameter Passing :- 

Passed by value : number, bools, array, structs.

Passed by reference : (everything a passed by values, its is descriptor that plays role here.).

1. Things passed by pointers ( &x)
2. Strings (but they are immutable)
3. Maps
4. Channels
5. Slices

 -->Return Values :-

1. Go functions allows multiple return values.


 -->Defer Statement :-
 
 Defer statements are functional scoped, they executes in LIFO order.
 If there is defer in nested function then it is scoped to nested function.
 
 Parameter gets copied at a defer statement(not as reference)
 A defer statement happens before the return is done.





//--------- Go class 9 : Closures ---------//


 Closure is not a function, it is what begin returned. It is actually a runtime thing.


//--------- Go class 10 : Slices in more detail ---------//

 Slice : Descriptor( Length , Capacity, pointer). Pointer points to some sequence of memory.
 Slice ( and Map) encoding differently in Json when nil.


//--------- Go class 12 : Structs, Struct tags, and Json  ---------//

You can not store address of map(&m) because map under the hood arranges itself if we do some 
insert operations.
Maps return a copy of the value, not a reference
When we are dealing with map of structs, we generally create map of pointers to structs.
 

Two structs types are compatible if : 
- fields have the same types and names
- in the same order 
- and the same tags(*) 

Anonymous structs are structurally compatible.

Types with different user-declared names are never compatible.

Names structs types are convertible (if they are structurally compatible).

Tag differences dont prevent type conversion.

Structs are passed by values unless  a pointer is used.

 "dot" notation works well with values but works on pointers too , equivalent to (*a).copies++

 A struct with no fields is useful,It takes up no space.

 Structs Tags are mostly used for conversion of data in the program to an external format json, xml etc.
 If field name in structs starts with lower case letter then it is not exported.
 

//--------- Go class 14 : References and Values semantics   ---------//

Any struct with mutex must be passed by reference(if they are copied they dont work).

If things to be shared then always pass a pointer.

Stack allocation is more efficient and go tries to allocate on stack(accessing a variable directly is more efficient than following a pointer).
Accessing a dense sequence of data is more efficient  than sparse data(an array is faster than linked list).

It means that go prefers to allocate on stack but sometimes cant.(9:02 in lecture).

Value returned by range is always a copy, we have to use index if we want to mutate Element.
 
Anytime function mutates a slice thats passed in , we must return a copy.(thats because slice backing array may be reallocated to grow).

Keeping a pointer to an element of slice is risky(as if you mutate slice by appending some elements it address may change due to reallocation).

Good to watch example in minute 17-18 of video. 

//--------- Go class 15 : Newtworking with HTTP   ---------//



//--------- Go class 17 : OOPs ---------//

 An interface specifies abstract behaviour in terms of methods.
 Concrete type offer methods that satisfy interfaces.
 Methods are type bound functions.A method is a special type of function , it has receiver parameter before the function name parameter.
 Only methods may be used to satisfy the interfaces.
 An interface specifies required behaviour as a method set


//--------- Go class 18 : Methods and interfaces ---------//

A function can be either pointer receiver or value receiver, but it can not be both.

Composition is union(and) of interfaces.

All the methods for a given type must be declared in the same package where the type is declared.This allows a package importing the type to know all the methods at compoile time.




//--------- Go class 19 : Composition  ---------//
Struct Composition- Building structs from other structs (feild and methods of embedded structs are promoted).
A struct can embed a pointer to another type; promotion of its feild and methods works in the same way.
Composition is not a inheritence.
We can embed a type into another type even if it is not a struct.
If fields are get promoted then so do the methods.

sorting : At 14 miniute (Three function on type are Len(), Less(i,j), Swap(i,j))

Nothing in Go prevents calling a method with nil receiver.

//--------- Go class 20  : Methods and interfaces in details ---------//
Nil interface : An interface variable in nil until initialized.

We called error a special type, but its really an interface

type error interface{
	func Error() string{}
}

Pointer methods may be called on non-pointers and vice-versa, Go will automatically use * and & as needed.
A method requiring a pointer receiver may only be called on addressable object (not literals).

If one method of a type take pointer receiver, then all its methods should take pointers.
And, in general objects of that type are probably not safe to copy.

Method Value : 
//--------- Go class 22 : Concurrency ---------//
 
Execution happens in some non-deterministic order 
-Different behaviour on different run even with the same input.
-Executes independently.

Concurrency does not necessarily make the program faster, parallelism does. 
We need Concurrency in order to enable parallelism. We have concurrent program, but parallelism happens at runtime if cpu has multiple cores.











  