package util

func RunMultiple(fl []func() error) error {
	for _, f := range fl {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
