package tsh

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"github.com/yang-zzhong/tsh/errors"
	"golang.org/x/sys/unix"
)

func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

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

func IsDIR(path string) (bool, errors.ChainedError) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, errors.New(err)
	}
	return stat.IsDir(), errors.New(err)
}

func Readable(path string) errors.ChainedError {
	if err := unix.Access(path, unix.R_OK); err != nil {
		return errors.New(err)
	}
	return nil
}

func Writable(path string) errors.ChainedError {
	if err := unix.Access(path, unix.W_OK); err != nil {
		return errors.New(err)
	}
	return nil
}
