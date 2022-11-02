package sqlutil

//var placeholder, _ = regexp.Compile("^:\\d$")
//
//func replace(statement string) string {
//	//n := strings.Count(statement, "?")
//	//for i := 0; i < n; i++ {
//	//	statement = strings.Replace(statement, "?", fmt.Sprintf(":%d", i+1), 1)
//	//}
//
//	var results = make([]byte, 0, len(statement))
//	var i int
//	for _, v := range statement {
//		if v == '?' {
//			results = append(results, []byte(fmt.Sprintf(":%d", i+1))...)
//			i++
//		} else {
//			results = append(results, byte(v))
//		}
//	}
//	return string(results)
//}
