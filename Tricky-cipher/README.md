Task "A tricky cipher"

Constraints

|  Constraint        |             Value             |
| ------------------ |:-----------------------------:|
|  Time limit        | 2 seconds                     |
|  Memory limitation | 512Mb                         | 
|  Input             | standard input or input.txt   |   
|  Output            | standard output or output.txt |
  


The well-known company Tyndex is once again recruiting interns.
Taking care of the personal data of applicants, the company came up with a clever encryption algorithm:
- The number of different characters in the full name is counted (the case is important, A and a are different characters).
- The sum of the digits in the day and month of birth multiplied by 64 is taken.
- For the first (by position in the word) letter of the surname, its number in the alphabet (in 1-indexing) multiplied by 256 is determined (the letter case is not important).
- The resulting numbers are summed up.
- The result is translated into a 16-ch number system (in uppercase).
- The result retains only 3 minor digits (if there are fewer significant digits, then the cipher is supplemented up to 3 digits with leading zeros).

Your task is to help calculate for each candidate his cipher.

Input format

In the first line, enter the number N(1≤N≤10000) — the number of candidates and ciphers.
This is followed by N lines in CSV format (fj,ij,oj,dj,mj,yj) — information about candidates:
 - surname ("fj"), name ("ij") and patronymic ("oj") (1≤∣∣j∣∣,∣∣ij∣∣,∣∣oj∣∣≤15) is a string consisting of Latin letters upper and lower case;
 - birthday ("dj"), month of birth ("mj"), and year of birth ("yj") are integers that set the correct date between January 1, 1950 and December 31, 2021.


Output format

In a single line, print N lines k1, k2, ..., kN, where kj is the cipher of the jth candidate (in uppercase). Candidates are numbered from 1 to N in the order of entry.


Example

|                     Input                |     Output    |
| ---------------------------------------- | ------------- |
| 2                                        | 710 64F       |
| Volozh,Arcady,Yurievich,11,2,1964        |               |
| Segalovich,Ilya,Valentinovich,13,9,1964  |               |



Notes

Let's consider a test example.
First candidate — Volozh,Arcady,Yurievich,11,2,1964:
 - Various characters in the full name: V, o, l, z, h, A, r, c, a, d, y, Y, u, i, e, v - there are 16 of them in total.
 - The sum of the digits in the day and month of birth is equal to 1+1+2 = 4.
 - The number in the alphabet of the first letter of the surname V is 22.
 - The final value of the cipher is 16+4⋅64+22⋅256= 5904.
 - In the 16-digit number system, this number is represented as 1710.
 - We are only interested in the last 3 digits, so that remains 710.

The second candidate is Segalovich,Ilya,Valentinovich,13,9,1964:
 - Various characters in the full name: S, e, g, a, l, o, v, i, c, h, I, y, V, n, t - there are 15 of them in total.
 - The sum of the digits in the day and month of birth is equal to 1+3+9 = 13.
 - The number in the alphabet of the first letter of the surname S is 19.
 - The final value of the cipher is 15+13⋅64+19⋅256= 5711. 
 - In the 16-digit number system, this number is represented as 164F.
 - We are only interested in the last 3 digits, so 64F remains.

 
	

