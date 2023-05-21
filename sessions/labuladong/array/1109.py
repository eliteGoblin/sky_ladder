


class Solution:
    def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
        # init 2D diff array
        diff = [0] * n
        
        for book in bookings:
            first, last, count = book

            diff[first-1] += count
            if last -1 + 1 < n:
                diff[last] -= count

        res = [0] * n
        res[0] = diff[0]

        for i in range(1, n):
            res[i] = res[i-1] + diff[i]
        return res