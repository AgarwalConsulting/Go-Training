package main

type NetworkError string

func (e NetworkError) Error() string {
	return string(e)
}
