- [Bits in the hardware](#bits-in-the-hardware)
- [Arithmetic Overflows and Underflows](#arithmetic-overflows-and-underflows)
- [Bitwise Operations](#bitwise-operations)
- [Bitwise Operators](#bitwise-operators)
- [Techniques](#techniques)
  * [Zeroing Registers](#zeroing-registers)
  * [Bit Masking](#bit-masking)
  * [Reading Bits](#reading-bits)
  * [Clearing Bits](#clearing-bits)
  * [Clearing a Range of Bits](#clearing-a-range-of-bits)
  * [Replacing Bits](#replacing-bits)
  * [Setting Bits](#setting-bits)
  * [Toggling bits](#toggling-bits)
  * [Counting Bits](#counting-bits)
  * [Converting from Decimal to Binary](#converting-from-decimal-to-binary)
## Bits in the hardware

At the lowest level computers have circuits, and through these circuits flows electricity. A bit in the context of digital computing can be thought of as representing either the absence or presence of an electrical current through a circuit.

The presence of an electrical current can be determined based on the voltage. For example 0 volts could indicate ("off") or a 0 bit, and 5 volts could indicate ("on") or a 1 bit. The actual voltage used in real circuits varies depending on the tech, these values were used for the purpose of explanation.

Transistors are the main component behind bits, acting as switches, allowing or blocking the flow of an electrical current, by amplifying, or cutting of voltage, based on wether the presence or absence of an electrical signal should exist. 

They are the basis behind logic gates, and really the entirety of all modern day computation. The state of a transistor (on or off) can represent a bit (0 or 1).

Transistors can be arranged in very complex ways within integrated circuits to form not only the basis for computer processors, but also memory chips, enabling the manipulation and storage of digital information in the form of bits.

There are other components aswell, such as capacitors, and inductors that fine-tune the electrical signals, but they're not as pertinent to the topic of bits.

CPU's happen to have billions of transistors. Modern day processors, like the i7 can have upward of 3 billion transistors.

Bits, are of course the smallest unit of information possible on a computer. A single byte is equivalent to 8 bits.  But bits also represent digits in binary, the base 2 number system. Binary is different then the decimal base 10 number system that we're used to. All machine instructions are comprised of binary when executed. 

Another numerical system often used in computer science is hexadecimal, a base 16 number system, which is actually a more concise way of expressing binary numbers, due to the fact that 16 is a power of 2.

But perhaps we're getting a bit ahead of ourselves, we're here to talk about bit manipulation, which is all about manipulating data at the binary level.

Before we can talk about bit manipulation we do need to talk a bit more about binary. The truth is we probably could've had a different number system looking back on it, but we settled on the number 10 likely because we have 10 fingers. If you've never counted in any other base besides 10, thinking about other number systems can be quite strange at first. Before we talk about binary lets talk about the number system we're used to.

Each digit in our number system is equivalent to an exponent of its base. The further left the digit the higher the exponential value. 10 in any number system is equal to its base. In hexadecimal 10 is equivalent to 16, in binary 10 is 2, and in the decimal system 10 is well 10. The first digit following the first is equivalent to the base to the power of 1, the further left you go, it becomes the base to the power of 2, and 3 and 4, and so on. 100 is just 10 to the 2nd power, 1,000,000 is just 10 to 6th power,  which is why scientific notation is such a neat way of expressing large values. With this information we can actually calculate the value of any arrangement of digits in any number system. 

In base 2, we can only have a 1 or a 0, which means 100, is equivalent to 2 to the power of 2, because of 2 0s, which is 4. In base 16, 100 would be 16 to the power of 2, which is 256. in an 8 bit number, the highest possible value is 255, but why exactly. Well lets look at our decimal system. What is the highest value that can fit within 2 digits, well 99, which is 100 - 1, a 1 with 8 zero after it conforming to the rules that we've established where every digital exponentially increases the value by the base, this number would be 2 to the power of 8, which is 256, but 1 with 8 zeros after it is actually 9 digits, so we subtract 1 and we get the value that can fit within 8 digits, which is whatever number can be created if all bits are 1. So you can now calculate the highest possible number within any number of bits, but do so by actually understanding why, well if my explanation made sense.

But what about negative numbers, how do those get expressed? Well, we can fit 256 values within 8 bits, but those values don't necessary all need to be positive. The way we express negative numbers in binary is with something known as 2s compliment.

Now before 2s compliment, there was 1's compliment, which can found by inverting all bits in a number. 2's compliments differs slightly in that you add 1 after inverting all the bits. Now the reason for using 2s compliment instead of 1s compliment has to do with the inefficiencies of 1s compliment. For one 0 can be expressed in two ways as all 0s or all 1s, but this issue doesn't exist in 2s compliment, because you would end up with the number -1, and this means that there is only 1 representation of 0 in twos compliment. Another issue is end around carry, which has something to do with overflow and the most significant bit, but I don't quite understand it. 

Point is, no one uses that crap, we use 2s compliment, which as stated before allows you to find the negative or positive representation of any number, by flipping the bits and adding 1. Take a 4 bit number 0100, which is 4, if you want -4, its simply 1100, invert and add 1, and if you want the positive representation, you can repeat the process to return to the positive value. Twos compliment is actually that simple. The most significant bit is the signed bit, which will let you know if a number is negative or positive. 1 indicates that a number is negative, and 0 indicates that the number is positive.

The only thing to keep in mind is that because you have negative numbers, you can only represent half the number of positive values as opposed to unsigned numbers. In an unsigned 8 bit integer the highest number you can express is 255, but a signed 8 bit number can only express -128 to 127, this is still 256 values, but half of the values are negative.

Now lets talk a little about binary addition and subtraction.

Binary addition works bit by bit, starting from the rightmost bit (the least significant bit), this should seem familiar since its essentially the same way as demical addition with a few changes. The rules are straightforward:
- **0 + 0** gives **0**.
- **0 + 1** or **1 + 0** gives **1**.
- **1 + 1** gives **0** and produces a carry of **1** to the next higher bit.

For example, if we add `00001111` and `00000001`, it results in `00010000`. This is similar to decimal addition but with only two digits `(0 and 1)`.

So what if we run out of bits, what happens? 

## Arithmetic Overflows and Underflows
Well, lets try this, in the following code, `1` has been added to a 8-bit unsigned integer with a value of `255`, and it turns the value into `0`. This is known as an overflow, because once we run out of bits, the integer wraps around to `0`.
```go
package main

import "fmt"

func main() {
	var num uint8 = 255
	fmt.Println(num + 1)
}
```
The opposite is true aswell, in the following code we've subtracted `1` from an 8-bit unsigned integer with a value of `0`, and it turns the value into `11111111` or `255`. This is known as an underflow, and in this case the integer wraps around to its maximum value.

```go
package main

import "fmt"

func main() {
	var num uint8 = 0
	fmt.Println(num - 1)
}
```

Now that we have a basic understanding of bits, binary and two's compliment, we can move onto talking about bit manipulation.

## Bitwise Operations 

**Bit manipulation**: It refers to the art and science of algorithmically altering the binary bits that represent data in a computer system. 

Since computers fundamentally operate on bits (binary digits), being able to effectively manipulate these bits can be pretty useful.

All forms of data that resides on a computer, regardless of their type, or abstraction level, are ultimately stored as a series bits.

This includes integers, floating-point numbers, characters (ASCII or Unicode alike), and even more complex structures like arrays or strings, since at the end of the day these structures are essentially just contigous blocks of memory each adressable and represented, ultimately, in binary.

Bit manipulation utilizes bitwise operations, which usually have their own mnemonics for the ISA of most computers, they are also accessible in most if not nearly every modern day programming language through bitwise operators. 

Bitwise operators are much faster than arithmetic operators, and allow for lower level control over values. They also leave a smaller memory footprint, depending on what you're doing, meaning that if you have a limited amount of memory bitwise operators can save memory. One example is packing a lot of flags within a single integer to represent lots of information at once, arithmetic operators wouldn't really be useful in this case since you're going to be reading and manipulating individual bits as opposed to decimal numbers.

Really this only matters if you're doing something intense like *game development* or *bare metal programming*.

The use of bit manipulation can also be found in graphics, encryption algorithms, data compression algorithms, low level device control (like microcontrollers) and in the optimization of algorithms in general.

## Bitwise Operators

1. **Binary AND (&)**: This operation compares two bits and returns 1 only if both bits are 1. Otherwise, it returns 0. For example, `1 & 1` results in `1`, but `1 & 0` or `0 & 1` results in `0`.

Another way to think about AND is how its used in boolean algebra. Boolean algebra, created by a George Boole, uses binary values, often represented as on/off, or true/false, to conduct logical operations, and really is the foundation behind logic gates, and digital logic in general.

If you've ever written a boolean expression you know that the AND operator returns true only if both operands are true, and this is exactly how the AND operation works in the context of Bit Manipulation.

2. **Binary OR (|)**: This operation compares two bits and returns 1 if either of the bits is 1. It only returns 0 when both bits are 0. For instance, `1 | 0` or `0 | 1` results in `1`, and `0 | 0` results in `0`.

Once again we can think about OR in the way its used within boolean alegbra.

If we write a boolean expression with an OR operator, it returns false only if both operands are false, otherwise if either operand is true, it returns true, same is true in the context of Bit Manipulation.

3. **Binary XOR (^)**: This operation compares two bits and returns 1 if the bits are different, and 0 if they are the same. For example, `1 ^ 0` or `0 ^ 1` results in `1`, but `0 ^ 0` or `1 ^ 1` results in `0`.

4. **Binary NOT (~)**: This operation inverts all the bits of its operand. If the bit is 1, it becomes 0, and if it's 0, it becomes 1. For example, `~1` results in all bits of `1` being inverted, becoming either `-2`, or `254` depending on if its signed.

In boolean algebra, the NOT operator works identically, it produces the opposite logical value of the expression, and we can think of NOT in the same way within the context of Bit Manipulation.

5. **Binary Left Shift (<<)**: This operation shifts the bits of the first operand to the left by the number of positions specified by the second operand. It is equivalent to multiplying the first operand by 2 raised to the power of the second operand. For example, `1 << 2` results in `4`.

6. **Binary Right Shift (>>)**: This operation shifts the bits of the first operand to the right by the number of positions specified by the second operand. It is equivalent to dividing the first operand by 2 raised to the power of the second operand. For example, `4 >> 2` results in `1`.

Just like in binary arithemtic, these operations are done from the LSB (Least Significant Bit) to the MSB (Most Significant Bit), from right to left. Then each bit is compared and evualated individually.

## Techniques

### Zeroing Registers

In assembly the **XOR** mnemonic is used for zeroing registers since any two bits that are the same in each of the operands will result in a `0`, this means any number that is XORed by itself becomes `0`, making it possible to zero out a value or register if the first and second operands are the same.

`Linux 32-bit assembly (Intel)`
```x86asm
global _start
_start:
	mov eax, 1 ; syscall number (sys_exit)
	mov ebx, 69 ; loading 69 into ebx
	xor ebx, ebx ; zeroing out register
	int 0x80 ; performs syscall to exit with 0
```
Heres also a simpler example in C:
```c
int main() {
	int n = 69;
	n = n^n; //zeroes out n
	return n; //exits with 0
}
```
### Bit Masking

Bit masking is a technique used to manipulate specific bits within a binary number. It involves using a mask - a binary number where certain bits are set to 1 (to select or affect those bits) or 0 (to ignore or leave those bits unchanged). 

Bit masking is used for operations like setting, clearing, toggling, or checking the value of particular bits within a number. It's a fundamental concept in low-level programming and is used across various applications, including hardware control, data compression, encryption, and more.

Bit masks can be used to combine or isolate multiple bits from a number simultaneously, which is useful in applications where multiple flags or settings are stored within a single integer.

```
0000
^
```
You can create a mask in a lot of different ways, but the simplest way is simply to do 
```go
mask := (1<<n)
```
This is Go code that produces a binary value with `1` left shifted to the value of `n`, this means you can make the `n`th bit a value of `1`, targeting it for various reasons. *Do note that this is zero indexed since we're left shifting by the value of `n`, shifting once would give us a value of `2` or `0010` in binary.*
```
0000
^ ^
```
What about targeting multiple bits?
Well, we can do that too!
```go
mask := (1<<1) | (1<<3)
```
This Go code now targets the **2nd** and **4th** bit in the number. By using the **OR** operator on the values of `10` or `2^1` and `1000` or `2^3` you've effectively set two bits to `1`. The reason for this is quite simple, an **OR** operation will compare two bits and evaluate to `1` if any operand contains a `1`, which makes it especially useful for combining bits, and creating masks that target multiple bits in the way we've done above.

You can always use the `0b` prefix to write explicit binary values directly, which you can use for creating a bit masks aswell.

### Reading Bits

```go
func getBit(n, i int) int {
	mask := (1 << i) //creates bit mask with value: 100 aka 1<<2 (2^2)

	if (n & mask) > 0 { //101 & 100 results in 100 or 4 in decimal which is larger than 0
		return 1
	}
	return 0
}

func main() {
	n := 5 // the number: 101 (binary)
	i := 2 // create index of 2
	fmt.Println(getBit(n, i)) // prints the value 1, because the 3rd bit in 5 is 1
}

```
To check the value of specific bits, you can use the bitwise AND operator with a mask where the targeted bits are `1`. If `x & (1 << n)` is nonzero, then the nth bit of `x` is set. 

The reason is simple, if the targeted bit is `1` then the AND operation results in `1`, and the number produced will be larger than `1`, otherwise the **AND** operation will cause all bits to be `0` including the bit thats targeted since its already `0`, this results in you zeroing out the number, and you'll find out that the targeted bit is a `1`.

The above code illustrates this, its a simple program that takes in an argument from the commandline, and then creates a bit mask based on the value passed in as an argument, it then uses the **AND** operator on the mask and checks if the result is more than zero, because if it is, then that means that the bit set or equal to `1`, therefore you should return `1`, otherwise you can return `0` because the bit is a zero, this allows you to find the value of any bit in a number at any position.

### Clearing Bits

```cpp
#include <iostream>
using namespace std;
void clearBit(int &n, int i)
{
    int mask = ~(1 << i);
    n = n & mask;
}

int main()
{
    int n = 13; // 1101
    clearBit(n, 3); //results in 0101
    cout << n << '\n'; //prints 5 or (0100)
    return 0;
}
```

To clear specific bits (set them to `0`), use the bitwise **AND** operator `(&)` with a mask where the bits to be cleared are `0`. This will cause all other bits to be unaffected since if the bits were `0` they remain `0` and if they were `1` they remain `1`, however the targeted bit will change to `0` since the result of the **AND** operation is `0` if any of the two operands is a `0`.

Creating a mask like this is easy, you just take a regular mask, that uses `1`'s to target bits, and you perform a **NOT** operation on it, flipping all of the bits, and now `0` will target bits for clearing.

The above code illustrates this once again. I used C++ because passing by reference makes things easier, allowing us to mutate `n` because we are passing in a reference to it, and the nagation operator in Go is kind of weird and confusing so I just wanted to keep things simpler.

### Clearing a Range of Bits
```cpp
#include <iostream>
using namespace std;

void clearBitsInRange(int &n, int i, int j)
{
    int x = (~0) << (j + 1);
    int y = (1 << i) - 1;

    int mask = x | y;
    n = n & mask;
}

int main()
{
    int n = 31; // 11111
    clearBitsInRange(n, 1, 3); //clearing bits 1 to 3 (zero indexed)
    cout << n << endl; // 17 printed (10001)
    return 0;
}
```
To clear a range of bits and set them to `0`, we use the bitwise **AND** operator `(&)` with a mask where the bits to be cleared are `0`.

The only difference between clearing a single bit and a range of bits, is the bit mask. If more than one bit in the mask is a `0`, then every bit with a `0` value will be cleared after doing an **AND** operation, while all bits in the mask that are `1`, will be unaffected.

### Replacing Bits
```cpp
#include <iostream>
using namespace std;

void clearBitsInRange(int &n, int i, int j)
{
    int x = (~0) << (j + 1);
    int y = (1 << i) - 1;

    int mask = x | y;
    n = n & mask;
}
void replaceBits(int &n, int i, int j, int m)
{
    clearBitsInRange(n, i, j);
    int mask = (m << i);
    n = n | mask;
}
int main()
{
    int n = 0b1001001; // 73
    int i = 2;
    int j = 4;
    int m = 5;
    replaceBits(n, i, j, m);
    cout << n << endl; // 85
    return 0;
}
```
To replace a segment of bits in a number, you can first clear the segment of bits, then perform a bitwise **OR** operation with the bits you want replaced.  This is pretty easy to understand because when all bits are `0`, performing an **OR** operation with any bit will fill in the bits exactly, since the **OR** operation results in `1` if either operand is a `1`.

The code first clears a range of bits, then  left shifts by the amount of the starting index for the bits its going to replace, so that it can replace the correct bits, then using OR to combne its bits with the zeroed bits, which replaces them exactly because the OR operation will not affect an operand if the other operand is all zeros.

### Setting Bits
```cpp
#include <iostream>
using namespace std;
void SetBit(int &n, int i)
{
    int mask = (1 << i);
    n = (n | mask);
}
int main()
{
    int n = 0b100111;
    SetBit(n, 4);
    cout << n << endl;
    return 0;
}
```
To set specific bits to `1`, use the bitwise **OR** operator `(|)` with a mask where the bits to be set are `1`. This makes sense because regardless of wether or not the bit is set, the result is always `1`, which means using the **OR** operation will ensure a bit is set to `1`.

### Toggling bits
```cpp
#include <iostream>
using namespace std;
void toggleBit(int &n, int i)
{
    int mask = (1 << i);
    n = (n ^ mask);
}
int main()
{
    int n = 0b110111;
    toggleBit(n, 4);
    cout << n << endl;
    return 0;
}
```
To toggle specific bits, use the bitwise **XOR** operator `(^)` with a mask where the bits to be toggled are `1`. Using an **XOR** works because **XOR** always produces `0` if both bits are the same, this means that when the mask is all `1`'s any `0` bit in the original number will be converted to `1`, and if its a `1`, it'll be converted to a `0`.

### Counting Bits
```go
func countBits(n int) int {
	count := 0
	for n > 0 {
		last_bit := n & 1
		count += last_bit
		n = n >> 1
	}
	return count
}

func main() {
	n := 31                   // 11111
	fmt.Println(countBits(n)) // should print 5
}
```
```go
func countBitsTwo(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func main() {
	n := 31 // 11111
	go fmt.Println(countBitsTwo(n)) // should print 5
}
```
You can count the number of bits in a number by iterating through each bit and adding them all together, if the bit is a 0, it won't increase the value of count, and if its a 1, it will increment the count, you iterate for aslong as the number is more than 0, because that means that there are more bits in the number to be iterated over. This is what is done in the first block of code. `n & 1` will return the last bit in a number, so once you have the last bit you can add it to a `count` to keep track of the amount, then do right shift to drop off the last bit, so that you can continue on with the next bit until you've iterated over every bit and right shifted to the number 0.

That is the naive approach, but there is a much faster and more optimal solution.

The second implementation `countBitsTwo` works slightly dfferently.Its known as Brian Kernighan's algorithm. It keeps track of a count the same way, but performs a bitwise operation that removes the last set bit from the number `n = n & (n -1)` then increments the count each time this is done. This means for every bit that exists the count is incremented, and this will only iterate for as many set bits that exist, making it faster than the first solution.

### Converting from Decimal to Binary
```go
func convertToBinary(n int) int {
	bin := 0
	p := 1
	for n > 0 {
		last_bit := n & 1
		bin += (p * last_bit)
		p = p * 10
		n = n >> 1
	}
	return bin
}

func main() {
	n := 11 // 1011
	fmt.Println(convertBinary(n))
}
```
The last thing I want to go over in this article is how to convert a decimal number to its binary representation but in decimal format. The way we'll do this, will take us right back to the beggining of this article. Remember how I said that every digit is just an exponent of its base, and the further left you go the higher the expoential value? Well this is true, but we can also use this information to represent binary numbers in decimal format, by essentially just multiplying by our base.

That is what the convertToBinary function does. The function starts by initializing a value of 0 in a variable `bin`, and it will add to this variable for the duration of its implementation. First it loops for aslong as n is more than 0, which means that there are still bits left in the number. Just like in the counting bits function, the first implementation, we will get the last bit in the number by using `n & 1`. At the end of each iteration we will right shift n to drop off the last bit. This is actually very similar to the count bits function but with some differences. The p variable of 1 is used to keep track of the placement of the digit, where its multiplied by 10 each iteration to move left once. Each iteration the value of the last bit is multiplied by the current exponential value of 10 (p) for its placement, then gets added to the value of bin, and when we have iterated through the bits, we get a number that looks like the binary representation, but in decimal format, by multiplying by the base, more specifically base 10.


That about covers the basics of bit manipulation, and binary too for the most part, hopefully you learned something. Once you understand how these operations work, you can get pretty creative and of and do some pretty interesting things, anyways thanks for reading.

