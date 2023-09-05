package main
import "fmt"
type teststruct struct {
	On bool
	Ammo int
	Power int
  }
  func (s *teststruct)Shoot() bool{
	res := false
	if s.On {
	  if s.Ammo <= 0{
	  res = false
	}else{
	  res = true
	  s.Ammo --
	}
  }
	return res
  }
  func (s *teststruct)RideBike() bool{
	res := false
	if s.On {
	  if s.Power <= 0{
	  res = false
	}else{
	  res = true
	  s.Power --
	}
  }
	return res
  }
  func main(){
	testStruct :=&teststruct{}
	fmt.Print(testStruct)
  }  