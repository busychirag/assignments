# Assignment 1

## Problem Statement

You are given a JSON input string.\
Your task is to:

-   Iterate over the input
-   Print the **type** and **value** of each entity
-   If the entity is a **data structure** (map/object or array/slice),
    go inside it
-   Continue this process **recursively** until you reach primitive data
    types
-   Implement the solution using **Go reflection**

------------------------------------------------------------------------

## High-Level Approach

This solution follows four main steps:

1.  **Unmarshal JSON into `interface{}`**
2.  **Use reflection to inspect runtime types**
3.  **Use recursion to traverse nested structures**
4.  **Stop recursion at primitive data types**

------------------------------------------------------------------------

## Why `interface{}` Is Used

The structure of the JSON is **not known in advance**.\
So the JSON is unmarshalled into an empty interface:

``` go
var data interface{}
json.Unmarshal([]byte(input), &data)
```

After unmarshalling:
```

  JSON Type   Go Type
  ----------- --------------------------
  Object      `map[string]interface{}`
  Array       `[]interface{}`
  String      `string`
  Number      `float64`
  Boolean     `bool`
```
This allows the program to handle **any JSON shape dynamically**.

------------------------------------------------------------------------

## Core Function: `recursiveIterator`

This function is responsible for: - Identifying the type of the value at
runtime - Printing the type and value - Recursively going deeper if the
value is a map or slice

------------------------------------------------------------------------

## Step-by-Step Explanation

### 1. Convert to Reflection Value

``` go
val := reflect.ValueOf(data)
```

-   Wraps the runtime value in a `reflect.Value`
-   Enables inspection of its **kind** and **contents**

------------------------------------------------------------------------

### 2. Handle Invalid / Nil Values

``` go
if !val.IsValid() {
    fmt.Println("Type: nil, Value: nil")
    return
}
```

-   Prevents runtime errors
-   Acts as a base case for recursion

------------------------------------------------------------------------

### 3. Print Type and Value

``` go
fmt.Printf("Type: %v, Value: %v\n", val.Kind(), val)
```

-   `val.Kind()` returns the category (`map`, `slice`, `string`, etc.)
-   Useful for understanding the structure dynamically

------------------------------------------------------------------------

### 4. Switch Based on Kind

``` go
switch val.Kind() {
```

This determines **how to process the value**.

------------------------------------------------------------------------

## Handling JSON Objects (`reflect.Map`)

``` go
case reflect.Map:
    for _, key := range val.MapKeys() {
        fmt.Printf("Key: %v -> \n", key)
        recursiveIterator(val.MapIndex(key).Interface())
    }
```

-   Iterates through all keys
-   Extracts each value
-   Converts it back to `interface{}`
-   Calls the function recursively

------------------------------------------------------------------------

## Handling JSON Arrays (`reflect.Slice`, `reflect.Array`)

``` go
case reflect.Slice, reflect.Array:
    for i := 0; i < val.Len(); i++ {
        fmt.Printf("Index: %d -> \n", i)
        recursiveIterator(val.Index(i).Interface())
    }
```

-   Iterates through each element
-   Recursively inspects nested values

------------------------------------------------------------------------

## Primitive Data Types (Base Case)

If the value is: - `string` - `float64` - `bool`

Then: - It gets printed - No further recursion occurs

This naturally **terminates recursion**.

------------------------------------------------------------------------

## Why `.Interface()` Is Required

Reflection works with `reflect.Value`, not normal Go values.

``` go
val.MapIndex(key).Interface()
```

-   Converts a reflection value back into a usable Go value
-   Required to pass data into recursive calls
-   Bridges reflection-based code with normal Go code

------------------------------------------------------------------------

## Key Concepts Demonstrated

  Concept                 Usage
  ----------------------- ------------------------------
  Reflection              Runtime type inspection
  Recursion               Traversing nested structures
  `interface{}`           Handling unknown JSON schema
  `reflect.Kind()`        Type-based logic
  Defensive programming   Nil checks

------------------------------------------------------------------------

## Final Summary

This solution:

-   Works for **any JSON structure**
-   Requires **no predefined schema**
-   Uses **reflection + recursion** to deeply inspect data
-   Stops automatically at primitive values

It is a **generic JSON walker**, commonly used in: - Parsers -
Validators - Debugging tools - Framework internals

------------------------------------------------------------------------

