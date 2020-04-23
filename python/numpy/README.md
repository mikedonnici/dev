# NumPy

NumPy (NumericalPython) is a fundamental package for data science in Python.

NumPy arrays provide a fast and memory-efficient alternative to Python lists.

```python
import numpy as np

arr = np.array(range(5))
```
NumPy arrays are _homogenius_ ie must contain elements of the same type which can be show with the `.dtype` attribute.

If a NumPy array is created with mixed types as its input it will attempt to cast them to the same type.

Homogeneity allows NumPy arrays to be more efficient and faster as it eliminates the overhead of data type checking.

**NumPy array broadcasting** is a feature whereby an operations are _vectorised_ and are be performed on all elements 
of an array at once.

```python
arr = np.array([1,2,3,4])
print(arr ** 2)
# array([1, 2, 9, 16])
```

**NumPy array indexing** is easier to use for two-dimensional lists when compared with standard list indexing.

```python
nums2 = [ 
    [1, 2, 3],
    [4, 5, 6] ] 
arr = np.array(nums)
nums[0][1] # row 1, col 2
arr[0, 1]  # row 1, col 2
arr[:, 0]  # entire first column, need a list comprehension to do this with standard lists 
```

**Boolean indexing** returns a sub array based on a condition by overlaying a boolean _mask_:

```python
arr = np.array([4,5,6,7])
mask = arr > 4     # array([False, True, True, True])
print(arr[mask])   
# array([5,6,7])
```


