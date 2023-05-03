"""
trick:
(x ^ y) < 0 if not same sign
index % (len(arr) - 1) circle arr, length must be 2^?
remove last bit n  ^ (n - 1)
"""


arr = [1, 2, 3, 4]
index = 0

# while index < 100:
#     print(arr[index & (len(arr) - 1)])
#     index += 1

# reverse

index = 3
while index > -100:
    print(arr[index &(len(arr) - 1)])
    index -= 1