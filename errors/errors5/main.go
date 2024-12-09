package main

import "errors"

func validateStatus(status string) error {
	if len(status) == 0 {
		err := errors.New("status cannot be empty")
		return err
	}  
	
	if len(status) > 140{
		err2 := errors.New("status exceeds 140 characters")
		return err2

	}

	return nil
}
