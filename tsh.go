package tsh

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"git.woa.com/oliverzyang/tsh/errors"
	"golang.org/x/sys/unix"
)

// MD5 the md5 string
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Exists check whether the file or path exists
func Exists(path string) (bool, errors.ChainedError) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.New(err)
}

// IsDIR check whether the path is a dir
func IsDIR(path string) (bool, errors.ChainedError) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, errors.New(err)
	}
	return stat.IsDir(), errors.New(err)
}

// Readable check whether the file is readable
func Readable(path string) errors.ChainedError {
	if err := unix.Access(path, unix.R_OK); err != nil {
		return errors.New(err)
	}
	return nil
}

// Writable check whether the file is writable
func Writable(path string) errors.ChainedError {
	if err := unix.Access(path, unix.W_OK); err != nil {
		return errors.New(err)
	}
	return nil
}
