# Introduction to R

## Intro

- Assign a value to a variable:
```r
val_1 <- 42
```
- Create a vector of values:
```r
vector_1 <- c(50, 5)
```
- Create a vector using variable names:
```r
vector_1 <- c(value_1, value_2)
```

### Built-in functions

- Average of values in a vector:
```r
mean(vector_1)
```
- Smallest value in a vector:
```r 
min(vector_1)
```
- Largest value in a vector:
```r 
max(vector_1)
```
- Total number of elements in a vector:
```r
length(vector_1) 
```
- Sum of elements in a vector:
```r 
sum(vector_1)
```

### Resources
 
- [Notes on naming variables in R](https://www.r-bloggers.com/consistent-naming-conventions-in-r/0)
- [Documentation on vectors in R](https://cran.r-project.org/doc/manuals/r-release/R-lang.html#Vector-objects)

--- 

## Working with Vectors

### Indexing vectors by position

- Extract a single element:
```r 
vector[1]
```
- Extract a range of elements:
```r
vector[3:7]
```
- Extract multiple elements:
```r
vector[c(2,5,7)]
```

### Displaying data types

- Display the data type of a vector:
```r 
typeof(vector)
```

### Naming vector elements

- Assign name attributes to a vector:
```r 
names(vector) <- name_vector
```

### Indexing vectors by name

- Extract a single element:
```r 
vector["name_2"]
```
- Extract multiple elements:
```r 
vector[c("name_1", "name_2)]
```

### Logical operators

- Less than: ```vector_1 < vector_2```
- Greater than: ```vector_1 > vector_2```
- Less than or equal to: ```vector_1 <= vector_2```
- Greater than or equal to: ```vector_1 >= vector_2```
- Equal to: ```vector_1 == vector_2```
- Not equal to: ```vector_1 != vector_2```

### Logical indexing

- Indexing into a numeric vector using a logical vector:
```r 
numeric_vector[logical_vector]
```

### Performing arithmetic on vectors

- Add, divide, or multiply vectors:
```r 
vector_1 + vector_2
vector_1 / vector_2
vector_1 * vector_2
```

### Concepts

- R recognizes different data types:
   - Numeric (3, 5.66, 199, 6)
   - Character ("math", "%", "&", "chem+math")
   - Logical (TRUE, FALSE)
   
- R is a _1-indexed_ programming language, which means that the first element in a vector is assigned a position of one.
- When performing operations on vectors of unequal length, R "recycles" values of the shorter vector until the two vectors are the same length.

### Resources
- [Documentation on indexing vectors in R](https://cran.r-project.org/doc/manuals/r-release/R-lang.html#Indexing)
- [Documentation on R's "recycling rule"](https://cran.r-project.org/doc/manuals/r-release/R-intro.html#The-recycling-rule)

---

## Working with matrices

### Naming matrix and row columns

- Assign name attributes to rows of a matrix:
```r 
rownames(matrix)
```

- Assign name attributes to columns of a matrix:
```r 
colnames(matrix)
```

### Matrix Operations

- Finding Matrix Dimensions
```r
dim(math_chemistry)
```

- Combining Vectors or Matrices by Row
```r
rbind(matrix_1, matrix_2)
rbind(vector_1, vector_2)
rbind(vector_1, matrix_1)
```

- Combining Vectors or Matrices by Column
```r 
cbind(matrix_1, matrix_2)
cbind(vector_1, vector_2)
cbind(vector_1, matrix_1)
```

### Indexing matrices by element

- Extract a single element:
```r 
matrix[2,5] 
matrix["Stanford","patents"]
```

- Extract multiple elements:
```r 
matrix[c(1,2),c(1,3)] 
matrix[c("Harvard","Stanford"),c("world_rank","influence")]
```

### Indexing matrices by rows and columns

- Extract a single row:
```r 
matrix[1,]
matrix["Harvard",]
```

- Extract a single column:
```r 
matrix[,2] 
matrix[,"quality_of_education"]
```

- Extract multiple rows or columns:
```r 
matrix[,c("quality_of_education","influence","broad_impact")]    
matrix[,c("2,3,4")]
```

### Rank values of a vector or a subset of a matrix

- Rank values of a vector:
```r 
rank(vector)
```

- Rank values of a matrix:
```r 
rank(matrix[,"column"])
rank(matrix["row",])
```

### Calculate the sum sum of values in a vector or matrix

- Sum of values in a vector:
```r 
sum(vector)
```

- Sum of values in a matrix:
```r 
sum(matrix[,"column"])
sum(matrix["row",])
```

### Concepts

- Like vectors, matrices only contain one data type. Unlike vectors, they are two-dimensional.
- When adding a vector to a matrix, it's good practice to make sure the new vector is the same length as the number of rows or columns in the matrix.

### Resources
- [Documentation on indexing matrices in R](https://cran.r-project.org/doc/manuals/r-release/R-lang.html#Indexing-matrices-and-arrays)
