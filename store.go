package main

type PathTransformFunc func(string) PathKey
type PathKey struct {
	PathName string
	Filename string
}

type Store struct {
	StoreOpts
}

type StoreOpts struct {
	Root              string
	PathTransformFunc PathTransformFunc
}
