// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Compiler is the name of the compiler toolchain that built the
// running binary. Known toolchains are:
//
//	gc      Also known as cmd/compile.
//	gccgo   The gccgo front end, part of the GCC compiler suite.
//

// Compiler 是构建运行的二进制文件的编译器工具链的名称。已知的工具链有:
//
//	gc      也称为 cmd/compile.
//	gccgo   gccgo 前端，GCC编译器套件的一部分。
//
const Compiler = "gc"
