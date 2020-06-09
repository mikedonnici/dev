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

---

## Working with Lists

### Creating Lists

- Create a list:
```r
new_list <- list("data scientist", c(50000,40000), "programming experience")
```

- Assign names to list objects:
```r
names(new_list) <- c("job title", "salaries", "requirements")
```

- Creating list with names:
```r
new_list <- list(name1 = vector_1, name2 = vector2)
```

### Indexing lists

- Return a list of selected elements:
```
new_list[1]
new_list["job title"]
new_list[c(1,3)]
```

- Return a single element:
```r 
new_list[[1]]
new_list[["job title"]]
new_list$"job title"
```

- Return a value contained in a list element:
```r 
new_list[[c(1,3)]]
```

### Manipulating lists

- Modifying List Elements
```r
new_list[[1]] <- "junior data scientist"
new_list[[c(2,1)]] <- 40000
```

- Adding Elements to Lists
```
new_list[[4]] <- c("healthcare", "vacation")
new_list[["benefits"]] <- c("healthcare", "vacation")
```

- Combining Multiple Lists
```r
new_list_2 <- c(new_list, new_list_1)
```

- Creating a List of Lists (Nested List)
```r
new_list_3 <- list(new_list, new_list_1)
```

### Concepts

- In R, lists are specialized vectors that can contain multiple objects. The objects may consist of different data 
structures, including single data elements, vectors, and matrices.
- Storing objects in lists allows you to make use of R's features for performing the same operation on each object in 
your list.
- Lists of lists contain multiple lists as objects. Each list contained in a nested list may, in turn, contain 
objects of any data structure or type.

### Resources
- [Documentation on Lists in R](https://cran.r-project.org/doc/manuals/r-patched/R-intro.html#Lists)

---

## Working with Data Frames

### Installing and loading packages

- Install packages:
```r
install_packages("package_name")
```

- Load packages:
```r
library(package_name)
```

### Importing data into R

- Save data as a data frame (data in .csv format)
```r
new_data_frame <- read_csv("data_set.csv")
```

### Working with data frame columns

- Select data frame columns:
```r 
data_frame_2 <- data_frame %>%
select(column_2, column_4, column_6)
```

- Add a new column to a data frame:
```r
data_frame_2 <- data_frame_1 %>%
mutate(new_column = (column_2/column_4)*100)
```

### Filtering a data frame by a single column

- Numeric data
```r
data_frame_2 <- data_frame_1 %>%
  filter(column_2 < 70)
```

- Character data
```r
data_frame_2 <- data_frame_1 %>%
  filter(column_3 == "Variable Name")
```

### Filtering a data frame by a single column

- Meeting at least one criterion (the | operator):
```r
data_frame_2 <- data_frame_1 %>% 
  filter(column_6 == "Variable Name" | column_4 > 1000)
```

- Meeting multiple criteria (the & operator):
```r 
data_frame_2 <- data_frame_1 %>% 
  filter(column_6 == "Variable Name" & column_4 > 1000)
```

### Arranging data frames by variables

- Arrange by a variable from smallest to largest:
```r
data_frame_2 <- data_frame_1 %>% 
  arrange(column_2)
```

- Arrange by a variable from largest to smallest:
```r
data_frame_2 <- data_frame_1 %>% 
  arrange(desc(column_2))
```

- Arrange by multiple variables:
```
data_frame_2 <- data_frame_1 %>% 
  arrange(column_2, desc(column_4))
```

### Concepts

- In R, packages consist of user-contributed functions, code and data that extend R's capabilities.
- The `tidyverse` is a collection of packages designed to make using R for data science more effective.
- Tibbles are a specialized type of data frame. They are a feature of packages in the tidyverse family that 
have been introduced to extend R's functionality for modern data science tasks.
- The pipe operator (`%>%`) is used to write code that chains series of operations together

### Resources
- [CRAN repository, which contains R packages](https://cran.r-project.org/)
- [the tidyverse family of packages](https://www.tidyverse.org/)
- [readr package documentation](https://readr.tidyverse.org/)
- [Documentation on tibbles](https://cran.r-project.org/web/packages/tibble/vignettes/tibble.html)
- [dplyr package documentation](https://dplyr.tidyverse.org/)

--- 

## Working with Control Structures

### Selection control structures

- If statement:
```r 
if (a > b) {  
  print("a is greater than b")
}
```

- If else statement:
```r 
if (a > b) {
  print("a is greater than b")
} else {
  print("b is greater than a")
}
```

- Nested else if statements:
```r 
if (a > b) {
  print("a is greater than b")
} else if (a < b) {
  print("b is greater than a")
} else if (a == b) {
  print("a and b are the same")
}
```

### For loops

- Print elements of a sequence:
```r  
numbers <- c(1,3,5,7,10)
for (number in numbers) {
    print(number)
}
```

- Loop over rows of a data frame:
```
for (i in 1:nrow(data_frame)) { 
  print(data_frame$variable_1[i] + data_frame$variable_2[i]) 
}
```

- Store for-loop output in a vector:
```r 
output_vec <- c()
for (i in 1:nrow(data_frame)) { 
    output_vec <- c(output_vec, data_frame$variable_1[i] + data_frame$variable_2[i]) 
}
```

- For-loop with selection control structures:
```r 
for i in 1:nrow(data_frame)){
  if (a > b) {
    print("a is greater than b")
  } else {
    print("b is greater than a")
  }
}
```

- For-loop with more than two cases:
```r 
for i in 1:nrow(data_frame)){
  if (a > b) {
    print("a is greater than b")
  } else if (a < b) {
    print("b is greater than a")
  } else if (a == b) {
    print("a and b are the same")
  }
}
```

### Concepts

- THREE TYPES OF CONTROL STRUCTURES:
   - Sequence Operations: Perform operations one after the other in the order they are specified
   - Selection Operations: Enable you to write code that tests whether a specified condition returns 
   TRUE or FALSE and executes different actions depending on the test outcome
   - Repetition Operations: Allows you to repeat blocks of code a number of times depending on specified 
   conditions

### Resources

- [Wikipedia article on control structures](https://en.wikipedia.org/wiki/Control_flow)
- [Wikipedia article on for-loops](https://en.wikipedia.org/wiki/For_loop)
- [Wikipedia article on while-loops](https://en.wikipedia.org/wiki/While_loop)
- [R-bloggers introduction to for-loops](https://www.r-bloggers.com/how-to-write-the-first-for-loop-in-r/)

---

## Working With Vectorized Functions

### Writing vectorized if-else statements:

- Using the if_else() function to write an if-else statement:
```r 
if_else(vector_1 == vector_2, "condition_1", "condition_2")
```

- Nesting if_else() functions to chain if-else statements:
```r 
if_else(vector_1 > vector_2, "condition_1",
    if_else(vector_1 < vector_2, "condition_2",
        if_else(vector_1 == vector_2, "condition_3", "condition_4")))
```

### Grouping data using `dplyr::group_by()`:

- Grouping by one variable:
```r 
data_frame %>% 
     group_by(variable)
```

- Grouping by multiple variables:
```r 
data_frame %>% 
    group_by(variable_1, variable_2) %>% 
    summarize(variable_name = function(variable_1))
```

### Summarising grouped data using `dplyr::summarize()`:

- Calculating one summary:
```r 
data_frame %>% 
    group_by(variable) %>% 
    summarize(variable_name = function(variable))
```

- Calculating multiple summaries:
```r 
data_frame %>% 
    group_by(variable_1) %>% 
    summarize(variable_name_1 = function_1(variable_1), variable_name_2 = function_2(variable_1))
```

### Concepts
- In R, vectorized solutions are often faster than using loops, and the code is usually easier to understand.
- Problems that involve splitting data into groups, applying a function to each group each group, and summarizing the 
results are known as "split_apply-combine" problems in R.
- Chaining functions using the pipe operator allows you to write more efficient code and avoid cluttering 
the global environment with intermediate variables.

### Resources
- [Blog Post on Vectorization in R](http://www.noamross.net/blog/2014/4/16/vectorization-in-r--why.html)
- [Wickham et al. paper on split-apply-combine problems](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.182.5667&rep=rep1&type=pdf)
- [maggritR package documentation](https://cran.r-project.org/web/packages/magrittr/vignettes/magrittr.html)

---

