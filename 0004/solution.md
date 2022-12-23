**Largest palindrome product**

This solution first defines a function isPalindrome that takes an integer as input and returns a boolean indicating whether or not the number is a palindrome. It does this by converting the number to a string and then checking if the string is a palindrome.

In the main function, we use two nested for loops to iterate over all possible combinations of two 3-digit numbers. For each combination, we calculate the product and check if it is a palindrome using the isPalindrome function. If it is a palindrome and is larger than the current maximum palindrome, we update the maximum palindrome.

Finally, we print out the maximum palindrome that was found.