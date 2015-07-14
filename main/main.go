package main

import (
  "fmt"
  "sort"
  "strings"
  "bufio"
  "os"
  "permutation-finder/clean_strings"
)

func IsAnagram(str1, str2 string) bool{

  clean_strings := clean_strings.CleanStrings(str1, str2)
  
  if len(clean_strings) != 2 {
    return false
  }

  str1 = clean_strings[0]
  str2 = clean_strings[1]

  if len(str1) == len(str2) {
    // downcase the strings
    str1 = strings.ToLower(str1)
    str2 = strings.ToLower(str2)

    // convert strings to slices
    string_one := strings.Split(str1, "")
    string_two := strings.Split(str2, "")

    // Sort the slices
    sort.Strings(string_one)
    sort.Strings(string_two)

    // join the slices back into strings
    sorted_one := strings.Join(string_one, "")
    sorted_two := strings.Join(string_two, "")

    // compare the strings to check equality
    if sorted_one == sorted_two {
      return true
    }
  } 
  return false
}

func main() {
  fmt.Println("Welcome to Anagram Finder. Find out if two strings are anagrams of each other.")
  
  input1 := bufio.NewReader(os.Stdin)
  fmt.Print("Enter the first string: ")
  str1, _ := input1.ReadString('\n')

  input2 := bufio.NewReader(os.Stdin)
  fmt.Print("Enter the second string: ")
  str2, _ := input2.ReadString('\n')

  fmt.Println(IsAnagram(str1, str2))

}













