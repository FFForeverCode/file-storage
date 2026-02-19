package main

import (
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"file-storage/p2p"
	"io"
	"sync"
)

const defaultRootFolderName = "default"

func init() {
	gob.Register(MessageStoreFile{})
	gob.Register(MessageGetFile{})
}

type MessageGetFile struct {
	ID  string
	Key string
}
type MessageStoreFile struct {
	ID   string
	Key  string
	Size int64
}
type FileServerOpts struct {
	ID                string            //主键ID
	EncKey            []byte            //key
	StorageRoot       string            //文件路径
	PathTransformFunc PathTransformFunc //路径转换方法
	Transport         p2p.Transport
	BootStrapNodes    []string
}

// FileServer 文件服务端
type FileServer struct {
	FileServerOpts
	peerLock sync.Mutex

	peers map[string]p2p.Peer

	store  *Store
	quitch chan struct{}
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storeOpts := StoreOpts{
		Root:              opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}
	if len(opts.ID) == 0 {
		opts.ID = generateID()
	}

	return &FileServer{
		FileServerOpts: opts,
		store:          NewStore(storeOpts),
		quitch:         make(chan struct{}),
		peers:          make(map[string]p2p.Peer),
	}
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = defaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		Filename: key,
	}
}

func generateID() string {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	return hex.EncodeToString(buf)

}

// todo 新增FileServer-GET方法、Store方法
