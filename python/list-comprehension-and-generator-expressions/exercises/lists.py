"""List comprehension exercises"""


def get_vowel_names(names):
    """Return a list containing all names given that start with a vowel."""
    return [
        name 
        for name in names
        if name.lower()[0] in ['a', 'e', 'i', 'o', 'u']
    ]


def power_list(nums):
    """Return a list that contains each number raised to the i-th power."""
    return [
        num**i
        for i, num in enumerate(nums)
    ]


def flatten():
    """Return a flattened version of the given 2-D matrix (list-of-lists)."""


def reverse_difference():
    """Return list subtracted from the reverse of itself."""


def matrix_add():
    """Add corresponding numbers in given 2-D matrices."""


def transpose():
    """Return a transposed version of given list of lists."""


def get_factors():
    """Return a list of all factors of the given number."""


def triples():
    """Return list of Pythagorean triples less than input num."""
