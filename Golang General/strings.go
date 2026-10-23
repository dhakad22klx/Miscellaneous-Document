
// Comparision without copy creation
strings.EqualFold(s1, s2)       // Case insensitive comparision , "go" and "Go" are same.
strings.EqualFold("Go", "go")   //true
strings.EqualFold("Go", "Java") //false

s := "GoLang"
fmt.Println(strings.ToLower(s)) // "golang"
fmt.Println(strings.ToUpper(s)) // "GOLANG"
fmt.Println(s)                  // "GoLang" — unchanged

// Checking substring existance
strings.Contains("GoLang", "Lang") // true
strings.Contains("GoLang", "java") // false
strings.Contains("GoLang", "lang") // false — use EqualFold for case-insensitive