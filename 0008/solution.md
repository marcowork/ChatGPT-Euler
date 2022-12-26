**Largest product in a series**

The question is to large for ChatGPT. Hence I had to rephrase the question to:
*The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832. Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product? *


This program defines the given 1000-digit number as a constant string, and then it uses a nested loop to iterate through all the possible 13-digit substrings of the number. For each substring, it calculates the product of the digits and compares it to the current maximum product. If the product of the current substring is greater than the maximum product, it updates the maximum product.

After the loops have completed, the program prints the maximum product, which is the answer to the problem.