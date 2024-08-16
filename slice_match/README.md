+ Implement a function, or set of functions, that when given a list of integers returns the longest ascending subsequence.

   Examples:

   ```
   S = [0, 0, 0, 0, 0, 0, 2, 0, 0]
   ss = [0, 0, 0, 0, 0, 0, 2]
   ```

   ```
   S = [1, 1, 2, 7, 2, 2, 2, 2, 2, 2]
   ss = [2, 2, 2, 2, 2, 2]
   ```

   ```
   S = [1, 2, 3, 4, 5, 1, 2, 3, 4, 5]
   ss = [1, 2, 3, 4, 5]
   ```

   ```
   S = [1, 2, 5, 2, 1, 1, 4, 6, 9, 0]
   ss = [1, 1, 4, 6, 9]
   ```

+ Implement a function, or set of functions, that when given a list of integers, and an ordering, returns the longest subsequence, ordered by the provided ordering.

   Examples:

   ```
   S = [3, 0, 0, 0, 0, 0, 2, 0, 0]
   o = descending with duplicates
   ss = [3, 0, 0, 0, 0, 0]
   ```

   ```
   S = [1, 1, 1, 2, 3, 1]
   o = ascending, no duplicates
   ss = [1, 2, 3]
   ```

   ```
   S = [1, 1, 2, 3, 7, 41, 5, 9, 13, 15]
   o = greater than 2, ascending, up to 14
   ss = [5, 9, 13]
   ```

   ```
   S = [5, 1, 21, 29, 41, 44, 18, 4]
   o = odd numbers, greater than 3, ascending by greater than 2
   ss = [21, 29, 41]
   ```
