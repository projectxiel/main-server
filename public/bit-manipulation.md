## Bits in the hardware

At the lowest level computers have circuits, and through these circuits flows electricity. A bit in the context of digital computing can be thought of as representing either the absence or presence of an electrical current through a circuit.

Transistors act as switches, allowing or blocking the flow of an electrical current based on the presence or absence of an electrical signal. They are the basis behind logic gates, and really the entirety of all modern day computation. The state of a transistor (on or off) can represent a bit (0 or 1).

Transistors can be arranged in very complex ways within integrated circuits to form not only the basis for computer processors, but also memory chips, enabling the manipulation and storage of digital information in the form of bits.

CPU's have billions of transistors. A modern day i7 processor can have upward of 3 billion transistors.

Bits are of course the smallest unit of information possible on a computer. A single byte is equivalent to 8 bits.  But bits also represent digits in binary, the base 2 number system. Binary is different then the decimal base 10 number system that we're used to. All machine instructions are comprised of binary when executed. 

Another numerical system often used in computer science is hexadecimal, a base 16 number system, which is actually a more concise way of expressing binary numbers, due to the fact that 16 is a power of 2.

But perhaps we're getting a bit ahead of ourselves, we're here to talk about bit manipulation, which is all about manipulating data at the binary level.

Before we can talk about bit manipulation we do need to talk a bit more about binary. The truth is we probably could've had a different number system looking back on it, but we settled on the number 10 likely because we have 10 fingers. If you've never counted in any other base besides 10, thinking about other number systems can be quite strange at first. Before we talk about binary lets talk about the number system we're used to.

Each digit in our number system is equivalent to an exponent of its base. The further left the digit the higher the exponential value. 10 in any number system is equal to its base. In hexadecimal 10 is equivalent to 16, in binary 10 is 2, and in the decimal system 10 is well 10. The first digit following the first is equivalent to the base to the power of 1, the further left you go, it becomes the base to the power of 2, and 3 and 4, and so on. 100 is just 10 to the 2nd power, 1,000,000 is just 10 to 6th power,  which is why scientific notation is such a neat way of expressing large values. With this information we can actually calculate the value of any arrangement of digits in any number system. 

In base 2, we can only have a 1 or a 0, which means 100, is equivalent to 2 to the power of 2, because of 2 0s, which is 4. In base 16, 100 would be 16 to the power of 2, which is 256. in an 8 bit number, the highest possible value is 255, but why exactly. Well lets look at our decimal system. What is the highest value that can fit within 2 digits, well 99, which is 100 - 1, a 1 with 8 zero after it conforming to the rules that we've established where every digital exponentially increases value by the base, this number would be 2 to the power of 8, which is 256, but 1 with 8 zeros after it is actually 9 digits, so we subtract 1 and we get the value that can fit within 8 digits, which is whatever number can be created if all bits are 1. So you can now calculate the highest possible number within any number of bits, but do so by actually understanding why, well if my explanation made sense.

Now this is all well and good, but what about negative numbers, how do those get expressed? Well, we can fit 256 values within 8 bits, but those values don't necessary all need to be positive. The way we express negative numbers in binary is with something known as 2s compliment.

Now before 2s compliment, there was 1's compliment, which can found by inverting all bits in a number. 2's compliments differs slightly in that you add 1 after inverting all the bits. Now the reason for using 2s compliment instead of 1s compliment has to do with the inefficiencies of 1s compliment For one 0 can be expressed in two ways as all 0s or all 1s, but this issue doesn't exist in 2s compliment, because you would end up with the number -1, and this means that there is only 1 representation of 0 in twos compliment. Another issue is end around carry, which has something to do with overflow and the most significant bit, but I don't quite understand it. 

Point is, no one uses that crap, we use 2s compliment, which as stated before allows you to find the negative or positive representation of any number, by flipping the bits and adding 1. take a 4 bit number 0100, which is 4, if you want -4, its simply 1100, invert and add 1, and if you want the positive representation, you can repeat the process to return to the positive value. Twos compliment is actually that simple. The most significant bit is the signed bit, which will let you know if a number is negative or positive. 1 indicates that a number is negative, and 0 indicates that the number is positive.

The only thing to keep in mind is that because you have negative numbers, you can only represent half the number of positive values as opposed to unsigned numbers. In an unsigned 8 bit integer the highest number you can express is 255, but a signed 8 bit number can only express -128 to 127, this is still 256 values, but half of the values are negative.

Now lets talk a little about binary addition and subtraction.

Binary addition works bit by bit, starting from the rightmost bit (the least significant bit), this should seem familiar since its essentially the same way as demical addition with a few changes. The rules are straightforward:
- *0 + 0* gives *0*.
- *0 + 1* or *1 + 0* gives *1*.
- *1 + 1* gives *0* and produces a carry of *1* to the next higher bit.

For example, if we add 00001111 and 00000001, it results in 00010000. This is similar to decimal addition but with only two digits (0 and 1).

So what if we run out of bits, what happens? 

## Arithmetic Overflows and Underflows
Well, lets try this, in the following code, `1` has been added to a 8-bit unsigned integer with a value of `255`, and it turns the value into `0`. This is known as an overflow, because once we run out of bits, the integer wraps around to 0.
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

