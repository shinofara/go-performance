package function_args

import (
	"testing"
	"fmt"
)

func Variable(name, password string) string {
	return fmt.Sprintf("%s/%s\n", name, password)
}

func PtrStruct(u *User) string {
	return fmt.Sprintf("%s/%s\n", u.Name, u.Password)
}

func Struct(u User) string {
	return fmt.Sprintf("%s/%s\n", u.Name, u.Password)
}

type User struct {
	Name string
	Password string
}

func BenchmarkVariable(b *testing.B) {
	name := "shinofara"
	password := "password"
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
		_ = Variable(name, password)
  }
}

func BenchmarkPtrStruct(b *testing.B) {
	u := &User{
		Name: "shinofara",
		Password: "password",
	}
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
		_ = PtrStruct(u)
  }	
}

func BenchmarkStructr(b *testing.B) {
	u := User{
		Name: "shinofara",
		Password: "password",
	}
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
		_ = Struct(u)
  }	
}
