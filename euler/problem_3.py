# https://projecteuler.net/problem=3
# TODO: optimize for suggested input: 600851475143 with divisions without calculating the primes.

n = 13195

# lpf: Largest Prime Factor.
def lpf(n):
    largest = 2
    for i in range(2, n):
        if is_prime(i) and n % i == 0 : largest = i
    return largest

def is_prime(n):
    for i in range(2, n):
        if n % i == 0: return False
    else:
        return True


print(lpf(n))
