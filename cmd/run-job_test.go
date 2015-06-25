package cmd

import (
       "testing"
)

func TestParseOptions(t *testing.T){
        //given
	options := "option1=option1,option2=option2"
	expected := "-option1 option1 -option2 option2" 

        //when
	found := parseArguments(options)

        //then
	if expected != found {
		t.Errorf("Expected/Found\n'%s'\n'%s'", expected,found)
	}
}
