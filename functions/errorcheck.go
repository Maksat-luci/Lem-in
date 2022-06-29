package functions

// ErrorCheck - returns error
func ErrorCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}
