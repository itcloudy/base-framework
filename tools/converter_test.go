// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package tools

import (
	"fmt"
	"testing"
)

func TestIntToStr(t *testing.T) {
	var a int
	a = 10
	aStr := fmt.Sprintf("%d", a)
	result := IntToStr(a)
	if aStr != result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%d excepted %s ,got %s", a, aStr, result)
	}
}
func TestStrToInt64(t *testing.T) {
	var aStr string
	aStr = "10"
	var aInt64 int64 = 10

	result := StrToInt64(aStr)
	if aInt64 != result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%s excepted %d ,got %d", aStr, aInt64, result)
	}
}
func TestBytesToInt64(t *testing.T) {
	var aStr string
	aStr = "10"
	var aInt64 int64 = 10

	result := BytesToInt64([]byte(aStr))
	if aInt64 != result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%s excepted %d ,got %d", aStr, aInt64, result)
	}
}

func TestStrToUint64(t *testing.T) {
	var aStr string
	aStr = "10"
	var aInt64 uint64 = 10

	result := StrToUint64(aStr)
	if aInt64 != result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%s excepted %d ,got %d", aStr, aInt64, result)
	}
	aStr = "-10"

	result = StrToUint64(aStr)
	if aInt64 == result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%s excepted %d ,got %d", aStr, 0, result)
	}
}

func TestFloat64ToStr(t *testing.T) {
	var aFloat float64 = 12.33
	except := fmt.Sprintf("%f0000000", aFloat)
	result := Float64ToStr(aFloat)
	if except != result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%f excepted %s ,got %s", aFloat, except, result)

	}
}
func TestStrToInt(t *testing.T) {
	var aStr string
	aStr = "10"
	var aInt int = 10

	result := StrToInt(aStr)
	if aInt != result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%s excepted %d ,got %d", aStr, aInt, result)
	}
	aStr = "234qwe"
	result = StrToInt(aStr)
	if aInt == result {
		t.Errorf("TestIntToStr failed to achieve the expected function，%s excepted %d ,got %d", aStr, 0, result)
	}
}
