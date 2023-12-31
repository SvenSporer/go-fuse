// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

import (
	"syscall"
)

const (
	ENOATTR = Status(syscall.ENODATA) // On Linux, ENOATTR is an alias for ENODATA.

	// EREMOTEIO Remote I/O error
	EREMOTEIO = Status(syscall.EREMOTEIO)
)

// To be set in InitIn/InitOut.Flags.
//
// This flags conflict with https://github.com/macfuse/library/blob/master/include/fuse_common.h
// and should be used only on Linux.
const (
	CAP_NO_OPENDIR_SUPPORT  = (1 << 24)
	CAP_EXPLICIT_INVAL_DATA = (1 << 25)
	CAP_MAP_ALIGNMENT       = (1 << 26)
	CAP_SUBMOUNTS           = (1 << 27)
	CAP_HANDLE_KILLPRIV_V2  = (1 << 28)
	CAP_SETXATTR_EXT        = (1 << 29)
	CAP_INIT_EXT            = (1 << 30)
	CAP_INIT_RESERVED       = (1 << 31)

	// CAP_RENAME_SWAP only exists on OSX.
	CAP_RENAME_SWAP = 0x0
)

type Attr struct {
	Ino  uint64
	Size uint64

	// Blocks is the number of 512-byte blocks that the file occupies on disk.
	Blocks    uint64
	Atime     uint64
	Mtime     uint64
	Ctime     uint64
	Atimensec uint32
	Mtimensec uint32
	Ctimensec uint32
	Mode      uint32
	Nlink     uint32
	Owner
	Rdev uint32

	// Blksize is the preferred size for file system operations.
	Blksize uint32
	Padding uint32
}

type SetAttrIn struct {
	SetAttrInCommon
}

const (
	// Mask for GetAttrIn.Flags. If set, GetAttrIn has a file handle set.
	FUSE_GETATTR_FH = (1 << 0)
)

type GetAttrIn struct {
	InHeader

	Flags_ uint32
	Dummy  uint32
	Fh_    uint64
}

// Flags accesses the flags. This is a method, because OSXFuse does not
// have GetAttrIn flags.
func (g *GetAttrIn) Flags() uint32 {
	return g.Flags_
}

// Fh accesses the file handle. This is a method, because OSXFuse does not
// have GetAttrIn flags.
func (g *GetAttrIn) Fh() uint64 {
	return g.Fh_
}

type CreateIn struct {
	InHeader
	Flags uint32

	// Mode for the new file; already takes Umask into account.
	Mode uint32

	// Umask used for this create call.
	Umask   uint32
	Padding uint32
}

type MknodIn struct {
	InHeader

	// Mode to use, including the Umask value
	Mode    uint32
	Rdev    uint32
	Umask   uint32
	Padding uint32
}

type ReadIn struct {
	InHeader
	Fh        uint64
	Offset    uint64
	Size      uint32
	ReadFlags uint32
	LockOwner uint64
	Flags     uint32
	Padding   uint32
}

type WriteIn struct {
	InHeader
	Fh         uint64
	Offset     uint64
	Size       uint32
	WriteFlags uint32
	LockOwner  uint64
	Flags      uint32
	Padding    uint32
}

type SetXAttrIn struct {
	InHeader
	Size  uint32
	Flags uint32
}

type GetXAttrIn struct {
	InHeader
	Size    uint32
	Padding uint32
}

func (s *StatfsOut) FromStatfsT(statfs *syscall.Statfs_t) {
	s.Blocks = statfs.Blocks
	s.Bsize = uint32(statfs.Bsize)
	s.Bfree = statfs.Bfree
	s.Bavail = statfs.Bavail
	s.Files = statfs.Files
	s.Ffree = statfs.Ffree
	s.Frsize = uint32(statfs.Frsize)
	s.NameLen = uint32(statfs.Namelen)
}
