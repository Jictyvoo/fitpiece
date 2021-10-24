# Bumpingheart

The core library that contains wrappers for some standard libraries from Go

## Router

## DAO

First thing first you need to load DatabaseConnection module from modules loader. After that, use `database.GenerateDAO` function giving table name as parameter.

All DAO Classes have a predefined functions, there is:

- `Insert(values map[string]interface{})` => Used to insert a new line into the table
- `Update(values map[string]interface{}, where string)` => Used to update a table line
- `Delete(condition string)` => Used to delete a table line
- `Select(columns []string, where string)` => Used to search lines in table
- `SelectAll(where string)` => used to search all columns in all lines of table

To use insert and update functions, can use them with associative arrays, or just values in a normal array. If use a normal array the columns used was all less first if the values's array length was less than columns existing into table. And for every value less, less columns will be used in the command.
