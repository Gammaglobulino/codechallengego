# codechallengego


Task 1 
Write a function/method that accepts a list of string tokens and a separate text string as an input and checks if the latter string can be represented as a concatenation of any subset of tokens from the first argument, where each token can be used multiple (including zero) times. Examples:
[ "ab", "bc", "a" ] vs. "abc" =  true
[ "a", "ab", "bc" ] vs. "ab" =  true
[ "a", "ab", "bc" ] vs. "" =  true
[ "ab", "bc" ] <-> "abc" = false
[ "ab", "bc", "c" ] <-> "abbcbc" = true

Task 2 
The input is a list of intervals (two integer points); write a function/method that merges any N intervals in the list that have common points (intervals [1, 3] and [3, 6] have a common point of 3; [4, 8] and [6, 10] have common points of 6, 7, 8) and returns the merged list (where no two intervals intersect). it is allowed to modify the input. Try to avoid allocating extra memory for the output.
What is the computational complexity of your implementation?

tests for boths the functions included.
