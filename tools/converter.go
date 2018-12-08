// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package tools

import "strconv"

// IntToStr converts integer to string
func IntToStr(num int) string {
	return strconv.Itoa(num)
}

// StrToInt64 converts string to int64
func StrToInt64(s string) int64 {
	int64, _ := strconv.ParseInt(s, 10, 64)
	return int64
}

// BytesToInt64 converts []bytes to int64
func BytesToInt64(s []byte) int64 {
	int64, _ := strconv.ParseInt(string(s), 10, 64)
	return int64
}

// StrToUint64 converts string to the unsinged int64
func StrToUint64(s string) uint64 {
	ret, _ := strconv.ParseUint(s, 10, 64)
	return ret
}

// Float64ToStr converts float64 to string
func Float64ToStr(f float64, prec ...string) string {
	return strconv.FormatFloat(f, 'f', 13, 64)
}

// StrToInt converts string to integer
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// StrToFloat64 converts string to float64
func StrToFloat64(s string) float64 {
	Float64, _ := strconv.ParseFloat(s, 64)
	return Float64
}

// BytesToFloat64 converts []byte to float64
func BytesToFloat64(s []byte) float64 {
	Float64, _ := strconv.ParseFloat(string(s), 64)
	return Float64
}

// BytesToInt converts []byte to integer
func BytesToInt(s []byte) int {
	i, _ := strconv.Atoi(string(s))
	return i
}
