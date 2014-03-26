// Copyright (c) 2014 Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Inital yacc source generated by ebnf2y[1]
// at 2013-10-04 23:10:47.861401015 +0200 CEST
//
//  $ ebnf2y -o ql.y -oe ql.ebnf -start StatementList -pkg ql -p _
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
//
//   [1]: http://github.com/cznic/ebnf2y

package ql

import __yyfmt__ "fmt"

import (
	"fmt"
)

type yySymType struct {
	yys  int
	line int
	col  int
	item interface{}
	list []interface{}
}

const add = 57346
const alter = 57347
const and = 57348
const andand = 57349
const andnot = 57350
const as = 57351
const asc = 57352
const begin = 57353
const between = 57354
const bigIntType = 57355
const bigRatType = 57356
const blobType = 57357
const boolType = 57358
const by = 57359
const byteType = 57360
const column = 57361
const commit = 57362
const complex128Type = 57363
const complex64Type = 57364
const create = 57365
const deleteKwd = 57366
const desc = 57367
const distinct = 57368
const drop = 57369
const durationType = 57370
const eq = 57371
const falseKwd = 57372
const floatType = 57373
const float32Type = 57374
const float64Type = 57375
const floatLit = 57376
const from = 57377
const ge = 57378
const group = 57379
const identifier = 57380
const imaginaryLit = 57381
const in = 57382
const index = 57383
const insert = 57384
const intType = 57385
const int16Type = 57386
const int32Type = 57387
const int64Type = 57388
const int8Type = 57389
const is = 57390
const into = 57391
const intLit = 57392
const le = 57393
const lsh = 57394
const neq = 57395
const not = 57396
const null = 57397
const on = 57398
const order = 57399
const oror = 57400
const qlParam = 57401
const rollback = 57402
const rsh = 57403
const runeType = 57404
const selectKwd = 57405
const stringType = 57406
const stringLit = 57407
const tableKwd = 57408
const timeType = 57409
const transaction = 57410
const trueKwd = 57411
const truncate = 57412
const uintType = 57413
const uint16Type = 57414
const uint32Type = 57415
const uint64Type = 57416
const uint8Type = 57417
const unique = 57418
const update = 57419
const values = 57420
const where = 57421

var yyToknames = []string{
	"add",
	"alter",
	"and",
	"andand",
	"andnot",
	"as",
	"asc",
	"begin",
	"between",
	"bigIntType",
	"bigRatType",
	"blobType",
	"boolType",
	"by",
	"byteType",
	"column",
	"commit",
	"complex128Type",
	"complex64Type",
	"create",
	"deleteKwd",
	"desc",
	"distinct",
	"drop",
	"durationType",
	"eq",
	"falseKwd",
	"floatType",
	"float32Type",
	"float64Type",
	"floatLit",
	"from",
	"ge",
	"group",
	"identifier",
	"imaginaryLit",
	"in",
	"index",
	"insert",
	"intType",
	"int16Type",
	"int32Type",
	"int64Type",
	"int8Type",
	"is",
	"into",
	"intLit",
	"le",
	"lsh",
	"neq",
	"not",
	"null",
	"on",
	"order",
	"oror",
	"qlParam",
	"rollback",
	"rsh",
	"runeType",
	"selectKwd",
	"stringType",
	"stringLit",
	"tableKwd",
	"timeType",
	"transaction",
	"trueKwd",
	"truncate",
	"uintType",
	"uint16Type",
	"uint32Type",
	"uint64Type",
	"uint8Type",
	"unique",
	"update",
	"values",
	"where",
}
var yyStatenames = []string{}

const yyEOFCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 185
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 787

var yyAct = []int{

	206, 162, 169, 270, 58, 207, 248, 107, 161, 222,
	172, 72, 14, 59, 81, 82, 83, 84, 56, 85,
	55, 120, 86, 87, 53, 138, 253, 28, 60, 88,
	156, 73, 89, 90, 91, 76, 293, 149, 223, 80,
	77, 148, 108, 54, 92, 93, 94, 95, 96, 284,
	232, 78, 134, 135, 136, 137, 74, 120, 120, 279,
	69, 233, 174, 97, 275, 98, 79, 274, 99, 140,
	75, 120, 100, 101, 102, 103, 104, 154, 141, 264,
	263, 267, 174, 71, 288, 114, 277, 260, 258, 62,
	120, 64, 65, 256, 265, 254, 238, 236, 219, 286,
	63, 230, 228, 115, 218, 115, 175, 221, 262, 139,
	142, 143, 144, 216, 245, 208, 180, 155, 117, 170,
	113, 166, 165, 242, 212, 54, 175, 119, 160, 168,
	115, 183, 25, 186, 187, 188, 189, 190, 191, 61,
	30, 177, 179, 39, 176, 158, 29, 167, 192, 193,
	194, 195, 25, 272, 203, 34, 120, 31, 122, 185,
	184, 209, 112, 231, 36, 215, 213, 196, 197, 198,
	199, 200, 201, 202, 214, 134, 135, 136, 137, 109,
	35, 227, 32, 134, 135, 136, 137, 229, 226, 182,
	251, 240, 210, 178, 164, 41, 47, 44, 250, 118,
	40, 33, 150, 151, 152, 153, 290, 120, 38, 2,
	235, 163, 282, 225, 108, 273, 123, 181, 110, 243,
	239, 291, 252, 285, 261, 241, 244, 211, 126, 255,
	43, 247, 246, 45, 46, 257, 48, 49, 42, 259,
	105, 111, 171, 157, 16, 133, 15, 268, 1, 146,
	50, 266, 128, 37, 269, 13, 124, 224, 173, 70,
	289, 276, 271, 66, 127, 68, 280, 130, 278, 132,
	125, 116, 12, 281, 145, 170, 283, 81, 82, 83,
	84, 249, 85, 287, 52, 86, 87, 121, 57, 292,
	237, 3, 88, 11, 73, 89, 90, 91, 76, 10,
	129, 131, 80, 77, 9, 217, 8, 92, 93, 94,
	95, 96, 7, 67, 78, 6, 220, 205, 147, 74,
	5, 159, 106, 69, 4, 0, 97, 0, 98, 79,
	0, 99, 0, 75, 0, 100, 101, 102, 103, 104,
	0, 0, 0, 0, 0, 0, 71, 0, 0, 0,
	0, 234, 62, 0, 64, 65, 81, 82, 83, 84,
	0, 85, 0, 63, 86, 87, 0, 0, 0, 0,
	0, 88, 0, 73, 89, 90, 91, 76, 0, 0,
	0, 80, 77, 0, 0, 0, 92, 93, 94, 95,
	96, 0, 0, 78, 0, 0, 0, 0, 74, 0,
	0, 0, 69, 0, 0, 97, 0, 98, 79, 0,
	99, 0, 75, 0, 100, 101, 102, 103, 104, 0,
	0, 0, 0, 0, 0, 71, 0, 0, 0, 0,
	0, 62, 0, 64, 65, 81, 82, 83, 84, 0,
	85, 204, 63, 86, 87, 0, 0, 0, 0, 0,
	88, 0, 73, 89, 90, 91, 76, 0, 0, 0,
	80, 77, 0, 0, 0, 92, 93, 94, 95, 96,
	0, 0, 78, 0, 0, 0, 0, 74, 0, 0,
	0, 69, 0, 0, 97, 0, 98, 79, 0, 99,
	0, 75, 0, 100, 101, 102, 103, 104, 0, 0,
	0, 0, 0, 0, 71, 0, 0, 0, 0, 0,
	62, 0, 64, 65, 0, 0, 0, 51, 0, 0,
	0, 63, 81, 82, 83, 84, 0, 85, 0, 0,
	86, 87, 0, 0, 0, 0, 0, 88, 0, 73,
	89, 90, 91, 76, 0, 0, 0, 80, 77, 0,
	0, 0, 92, 93, 94, 95, 96, 0, 0, 78,
	0, 0, 0, 0, 74, 0, 0, 0, 69, 0,
	0, 97, 0, 98, 79, 0, 99, 0, 75, 0,
	100, 101, 102, 103, 104, 0, 0, 0, 0, 0,
	0, 71, 0, 0, 0, 0, 0, 62, 0, 64,
	65, 81, 82, 83, 84, 0, 85, 0, 63, 86,
	87, 0, 0, 0, 0, 0, 88, 0, 73, 89,
	90, 91, 76, 0, 0, 0, 80, 77, 0, 0,
	0, 92, 93, 94, 95, 96, 0, 0, 78, 0,
	0, 0, 0, 74, 0, 0, 0, 69, 0, 0,
	97, 0, 98, 79, 0, 99, 0, 75, 0, 100,
	101, 102, 103, 104, 81, 82, 83, 84, 0, 85,
	71, 0, 86, 87, 0, 0, 0, 0, 0, 88,
	0, 0, 89, 90, 91, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 92, 93, 94, 95, 96, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 97, 17, 98, 0, 0, 99, 0,
	18, 0, 100, 101, 102, 103, 104, 0, 0, 19,
	0, 0, 20, 21, 0, 0, 22, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 24,
	0, 0, 25, 0, 0, 0, 0, 0, 0, 26,
	0, 0, 0, 0, 0, 0, 27,
}
var yyPact = []int{

	709, -70, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 80, 72, -1000,
	116, 166, 114, 115, -1000, 182, 77, 157, 709, 157,
	-1000, 159, 157, 157, 158, 157, 157, 422, -1000, 157,
	141, -1000, -1000, 214, 106, 38, 51, -1000, -1000, 36,
	164, -1000, 46, -1000, 149, 209, -1000, 216, -36, 17,
	-1000, -45, 588, 588, 588, 588, -1000, -1000, -1000, -1000,
	-1000, 509, 35, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-66, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 51, -1000, 48, -1000,
	141, 192, 156, 141, -1000, 509, 69, 141, 44, 509,
	509, -1000, 155, 509, 34, 177, 509, 105, 509, 509,
	509, 509, 509, 509, 509, 509, 509, 509, 509, 509,
	509, 509, 509, 509, 509, -1000, -1000, -1000, 343, 509,
	-45, -45, -45, -45, 32, 509, 154, -1000, -1000, 43,
	509, -1000, 651, 141, 31, -1000, 98, 22, -1000, 15,
	-1000, 26, -1000, 204, -1000, 89, -1000, 209, -1000, -1000,
	509, 20, 509, 95, -1000, 108, -36, -36, -36, -36,
	-36, -36, 17, 17, 17, 17, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -37, 264, 14, -1000, 98, -1000, 13,
	-1000, -1000, 141, 98, -1000, -1000, 153, 42, 509, -1000,
	33, 24, 161, -1000, -1000, 152, -71, 12, 509, 87,
	509, -1000, -1000, 1, -1000, 0, -1000, 27, -1000, -1000,
	-3, 11, 141, -2, -1000, 141, -1000, 161, 96, -1000,
	198, -1000, -16, -1000, -1000, -19, 509, -36, -1000, -1,
	-1000, -1000, 509, -1000, -24, -1000, -1000, -1000, -1000, 96,
	-1000, -1000, 195, 141, -1000, -1000, -36, -1000, 98, -34,
	18, -1000, 509, -1000, -1000, -1000, 2, 196, 509, -1000,
	-1000, -1000, -47, -1000,
}
var yyPgo = []int{

	0, 324, 7, 322, 321, 320, 318, 317, 8, 1,
	2, 316, 315, 313, 312, 306, 305, 304, 299, 293,
	291, 5, 0, 290, 18, 288, 24, 287, 284, 281,
	274, 272, 271, 266, 265, 263, 262, 260, 259, 139,
	4, 13, 10, 258, 257, 255, 12, 253, 250, 9,
	6, 3, 249, 209, 248, 200, 20, 246, 11, 28,
	244, 243, 38, 242, 227, 226, 225, 224, 223, 222,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 4, 4, 64, 64, 5,
	6, 7, 7, 8, 9, 10, 11, 11, 65, 65,
	12, 13, 14, 14, 15, 16, 16, 66, 66, 17,
	17, 18, 19, 20, 21, 21, 22, 23, 23, 67,
	67, 24, 24, 24, 24, 24, 24, 24, 25, 25,
	25, 25, 25, 25, 25, 26, 27, 27, 28, 28,
	29, 30, 31, 31, 32, 32, 33, 33, 68, 68,
	34, 34, 34, 34, 34, 34, 34, 35, 35, 35,
	35, 36, 37, 37, 37, 39, 39, 39, 39, 39,
	40, 40, 40, 40, 40, 41, 41, 41, 41, 41,
	41, 41, 41, 38, 38, 42, 43, 43, 69, 69,
	44, 44, 63, 63, 45, 46, 46, 47, 47, 48,
	48, 48, 49, 49, 50, 50, 51, 51, 52, 52,
	52, 52, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 53, 53, 53, 53, 53, 54, 54, 55, 56,
	56, 57, 58, 58, 58, 58, 58, 58, 58, 58,
	58, 58, 58, 58, 58, 58, 58, 58, 58, 58,
	58, 58, 58, 58, 58, 58, 60, 61, 61, 59,
	59, 59, 59, 59, 62,
}
var yyR2 = []int{

	0, 5, 6, 3, 3, 0, 3, 0, 1, 2,
	3, 0, 1, 2, 1, 3, 0, 3, 0, 1,
	1, 4, 8, 10, 8, 0, 3, 0, 1, 3,
	4, 3, 3, 0, 1, 3, 3, 0, 3, 0,
	1, 1, 5, 6, 5, 6, 3, 4, 1, 3,
	3, 3, 3, 3, 3, 2, 0, 2, 1, 3,
	3, 3, 10, 5, 0, 3, 0, 5, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 4, 0, 1, 1, 1, 1, 2, 2, 2,
	1, 3, 3, 3, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 1, 3, 2, 1, 4, 0, 1,
	0, 2, 1, 3, 1, 8, 9, 0, 1, 1,
	1, 2, 0, 1, 0, 1, 0, 1, 3, 4,
	4, 5, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 1,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 0, 1, 1,
	2, 2, 2, 2, 2,
}
var yyChk = []int{

	-1000, -54, -53, -20, -1, -5, -12, -14, -15, -17,
	-18, -19, -31, -45, -46, -57, -60, 5, 11, 20,
	23, 24, 27, 42, 60, 63, 70, 77, 97, 66,
	68, 41, 66, 35, 41, 66, 49, -47, 26, 66,
	-55, 38, -53, -55, 38, -55, -55, 38, -55, -55,
	-48, 95, -28, -26, -21, -56, -24, -25, -40, -41,
	-59, -39, 88, 99, 90, 91, -35, -13, -34, 59,
	-38, 82, -58, 30, 55, 69, 34, 39, 50, 65,
	38, 13, 14, 15, 16, 18, 21, 22, 28, 31,
	32, 33, 43, 44, 45, 46, 47, 62, 64, 67,
	71, 72, 73, 74, 75, -55, -3, -2, -9, 38,
	4, 27, 56, 82, -62, 79, -32, 82, 35, 81,
	58, -27, 9, 7, 40, 54, 12, 48, 36, 84,
	51, 85, 53, 29, 88, 89, 90, 91, 8, 92,
	52, 61, 93, 94, 95, -30, -52, -6, 86, 82,
	-39, -39, -39, -39, -21, 82, 96, -61, -62, -4,
	80, -8, -9, 19, 38, -8, -21, 78, -46, -10,
	-9, -63, -42, -43, 38, 82, -26, -56, 38, -24,
	82, 40, 12, -40, 55, 54, -40, -40, -40, -40,
	-40, -40, -41, -41, -41, -41, -59, -59, -59, -59,
	-59, -59, -59, -21, 98, -7, -22, -21, 83, -21,
	38, -64, 81, -21, -58, -9, 82, -16, 82, 83,
	-11, 81, -49, -62, -44, 9, -46, -22, 82, -40,
	6, 55, 87, 98, 87, -21, 83, -23, 83, -2,
	38, -66, 81, -22, -65, 81, -42, -49, -50, -29,
	37, 38, -69, 97, 83, -22, 6, -40, 87, -21,
	87, -67, 81, 83, 82, 83, -8, 83, -9, -50,
	-51, -36, 57, 17, 83, 83, -40, 87, -21, 83,
	-33, -51, 17, -10, 83, -68, 81, -22, 82, -37,
	10, 25, -22, 83,
}
var yyDef = []int{

	33, -2, 146, 132, 133, 134, 135, 136, 137, 138,
	139, 140, 141, 142, 143, 144, 145, 0, 0, 20,
	0, 0, 0, 0, 114, 117, 0, 0, 33, 0,
	9, 0, 0, 0, 0, 0, 0, 0, 118, 0,
	0, 148, 147, 0, 0, 0, 29, 31, 32, 64,
	0, 119, 120, 58, 56, 34, 149, 41, 48, 90,
	95, 179, 0, 0, 0, 0, 85, 86, 77, 78,
	79, 0, 0, 70, 71, 72, 73, 74, 75, 76,
	103, 152, 153, 154, 155, 156, 157, 158, 159, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 175, 151, 177, 5, 0, 14,
	0, 0, 0, 0, 30, 0, 0, 0, 0, 121,
	0, 55, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 87, 88, 89, 0, 11,
	180, 181, 182, 183, 0, 0, 0, 176, 178, 7,
	0, 1, 0, 0, 0, 25, 184, 0, 63, 0,
	16, 122, 112, 110, 106, 0, 59, 35, 57, 150,
	0, 0, 0, 0, 46, 0, 49, 50, 51, 52,
	53, 54, 91, 92, 93, 94, 96, 97, 98, 99,
	100, 101, 102, 0, 0, 0, 12, 37, 80, 0,
	104, 4, 8, 3, 13, 2, 0, 27, 0, 65,
	18, 122, 124, 123, 105, 0, 108, 0, 0, 0,
	0, 47, 61, 0, 128, 0, 10, 39, 21, 6,
	0, 0, 28, 0, 15, 19, 113, 124, 126, 125,
	0, 111, 0, 109, 42, 0, 0, 44, 130, 0,
	129, 36, 40, 22, 0, 24, 26, 66, 17, 126,
	115, 127, 0, 0, 107, 43, 45, 131, 38, 0,
	68, 116, 0, 60, 23, 62, 69, 82, 0, 81,
	83, 84, 0, 67,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 99, 3, 3, 3, 93, 92, 3,
	82, 83, 95, 91, 81, 90, 96, 94, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 98, 97,
	85, 80, 84, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 86, 3, 87, 88, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 89,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79,
}
var yyTok3 = []int{
	0,
}

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEOFCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:

		{
			yyVAL.item = &alterTableAddStmt{tableName: yyS[yypt-2].item.(string), c: yyS[yypt-0].item.(*col)}
		}
	case 2:

		{
			yyVAL.item = &alterTableDropColumnStmt{tableName: yyS[yypt-3].item.(string), colName: yyS[yypt-0].item.(string)}
		}
	case 3:

		{
			yyVAL.item = assignment{colName: yyS[yypt-2].item.(string), expr: yyS[yypt-0].item.(expression)}
		}
	case 4:

		{
			yyVAL.item = append([]assignment{yyS[yypt-2].item.(assignment)}, yyS[yypt-1].item.([]assignment)...)
		}
	case 5:

		{
			yyVAL.item = []assignment{}
		}
	case 6:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]assignment), yyS[yypt-0].item.(assignment))
		}
	case 9:

		{
			yyVAL.item = beginTransactionStmt{}
		}
	case 10:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 11:

		{
			yyVAL.item = []expression{}
		}
	case 12:
		yyVAL.item = yyS[yypt-0].item
	case 13:

		{
			yyVAL.item = &col{name: yyS[yypt-1].item.(string), typ: yyS[yypt-0].item.(int)}
		}
	case 14:
		yyVAL.item = yyS[yypt-0].item
	case 15:

		{
			yyVAL.item = append([]string{yyS[yypt-2].item.(string)}, yyS[yypt-1].item.([]string)...)
		}
	case 16:

		{
			yyVAL.item = []string{}
		}
	case 17:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]string), yyS[yypt-0].item.(string))
		}
	case 20:

		{
			yyVAL.item = commitStmt{}
		}
	case 21:

		{
			yyVAL.item = &conversion{typ: yyS[yypt-3].item.(int), val: yyS[yypt-1].item.(expression)}
		}
	case 22:

		{
			yyVAL.item = &createIndexStmt{yyS[yypt-5].item.(string), yyS[yypt-3].item.(string), yyS[yypt-1].item.(string)}
		}
	case 23:

		{
			yyVAL.item = &createIndexStmt{yyS[yypt-7].item.(string), yyS[yypt-5].item.(string), yyS[yypt-3].item.(string) + "()"}
			if yyS[yypt-3].item.(string) != "id" {
				yylex.(*lexer).err("only the built-in id() can be indexed on")
				goto ret1
			}
		}
	case 24:

		{
			yyVAL.item = &createTableStmt{tableName: yyS[yypt-5].item.(string), cols: append([]*col{yyS[yypt-3].item.(*col)}, yyS[yypt-2].item.([]*col)...)}
		}
	case 25:

		{
			yyVAL.item = []*col{}
		}
	case 26:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]*col), yyS[yypt-0].item.(*col))
		}
	case 29:

		{
			yyVAL.item = &truncateTableStmt{yyS[yypt-0].item.(string)}
		}
	case 30:

		{
			yyVAL.item = &deleteStmt{tableName: yyS[yypt-1].item.(string), where: yyS[yypt-0].item.(*whereRset).expr}
		}
	case 31:

		{
			yyVAL.item = &dropIndexStmt{indexName: yyS[yypt-0].item.(string)}
		}
	case 32:

		{
			yyVAL.item = &dropTableStmt{tableName: yyS[yypt-0].item.(string)}
		}
	case 33:

		{
			yyVAL.item = nil
		}
	case 34:
		yyVAL.item = yyS[yypt-0].item
	case 35:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(oror, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 36:

		{
			yyVAL.item = append([]expression{yyS[yypt-2].item.(expression)}, yyS[yypt-1].item.([]expression)...)
		}
	case 37:

		{
			yyVAL.item = []expression(nil)
		}
	case 38:

		{
			yyVAL.item = append(yyS[yypt-2].item.([]expression), yyS[yypt-0].item.(expression))
		}
	case 41:
		yyVAL.item = yyS[yypt-0].item
	case 42:

		{
			yyVAL.item = &pIn{expr: yyS[yypt-4].item.(expression), list: yyS[yypt-1].item.([]expression)}
		}
	case 43:

		{
			yyVAL.item = &pIn{expr: yyS[yypt-5].item.(expression), not: true, list: yyS[yypt-1].item.([]expression)}
		}
	case 44:

		{
			yyVAL.item = &pBetween{expr: yyS[yypt-4].item.(expression), l: yyS[yypt-2].item.(expression), h: yyS[yypt-0].item.(expression)}
		}
	case 45:

		{
			yyVAL.item = &pBetween{expr: yyS[yypt-5].item.(expression), not: true, l: yyS[yypt-2].item.(expression), h: yyS[yypt-0].item.(expression)}
		}
	case 46:

		{
			yyVAL.item = &isNull{expr: yyS[yypt-2].item.(expression)}
		}
	case 47:

		{
			yyVAL.item = &isNull{expr: yyS[yypt-3].item.(expression), not: true}
		}
	case 48:
		yyVAL.item = yyS[yypt-0].item
	case 49:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(ge, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 50:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('>', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 51:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(le, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 52:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('<', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 53:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(neq, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 54:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(eq, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 55:

		{
			expr, name := yyS[yypt-1].item.(expression), yyS[yypt-0].item.(string)
			if name == "" {
				s, ok := expr.(*ident)
				if ok {
					name = s.s
				}
			}
			yyVAL.item = &fld{expr: expr, name: name}
		}
	case 56:

		{
			yyVAL.item = ""
		}
	case 57:

		{
			yyVAL.item = yyS[yypt-0].item
		}
	case 58:

		{
			yyVAL.item = []*fld{yyS[yypt-0].item.(*fld)}
		}
	case 59:

		{
			l, f := yyS[yypt-2].item.([]*fld), yyS[yypt-0].item.(*fld)
			if f.name != "" {
				if f := findFld(l, f.name); f != nil {
					yylex.(*lexer).err("duplicate field name %q", f.name)
					goto ret1
				}
			}

			yyVAL.item = append(yyS[yypt-2].item.([]*fld), yyS[yypt-0].item.(*fld))
		}
	case 60:

		{
			yyVAL.item = &groupByRset{colNames: yyS[yypt-0].item.([]string)}
		}
	case 61:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 62:

		{
			yyVAL.item = &insertIntoStmt{tableName: yyS[yypt-7].item.(string), colNames: yyS[yypt-6].item.([]string), lists: append([][]expression{yyS[yypt-3].item.([]expression)}, yyS[yypt-1].item.([][]expression)...)}
		}
	case 63:

		{
			yyVAL.item = &insertIntoStmt{tableName: yyS[yypt-2].item.(string), colNames: yyS[yypt-1].item.([]string), sel: yyS[yypt-0].item.(*selectStmt)}
		}
	case 64:

		{
			yyVAL.item = []string{}
		}
	case 65:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 66:

		{
			yyVAL.item = [][]expression{}
		}
	case 67:

		{
			yyVAL.item = append(yyS[yypt-4].item.([][]expression), yyS[yypt-1].item.([]expression))
		}
	case 70:
		yyVAL.item = yyS[yypt-0].item
	case 71:
		yyVAL.item = yyS[yypt-0].item
	case 72:
		yyVAL.item = yyS[yypt-0].item
	case 73:
		yyVAL.item = yyS[yypt-0].item
	case 74:
		yyVAL.item = yyS[yypt-0].item
	case 75:
		yyVAL.item = yyS[yypt-0].item
	case 76:
		yyVAL.item = yyS[yypt-0].item
	case 77:

		{
			yyVAL.item = value{yyS[yypt-0].item}
		}
	case 78:

		{
			yyVAL.item = parameter{yyS[yypt-0].item.(int)}
		}
	case 79:

		{
			yyVAL.item = &ident{yyS[yypt-0].item.(string)}
		}
	case 80:

		{
			yyVAL.item = &pexpr{expr: yyS[yypt-1].item.(expression)}
		}
	case 81:

		{
			yyVAL.item = &orderByRset{by: yyS[yypt-1].item.([]expression), asc: yyS[yypt-0].item.(bool)}
		}
	case 82:

		{
			yyVAL.item = true // ASC by default
		}
	case 83:

		{
			yyVAL.item = true
		}
	case 84:

		{
			yyVAL.item = false
		}
	case 85:
		yyVAL.item = yyS[yypt-0].item
	case 86:
		yyVAL.item = yyS[yypt-0].item
	case 87:

		{
			var err error
			if yyVAL.item, err = newIndex(yyS[yypt-1].item.(expression), yyS[yypt-0].item.(expression)); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 88:

		{
			var err error
			s := yyS[yypt-0].item.([2]*expression)
			if yyVAL.item, err = newSlice(yyS[yypt-1].item.(expression), s[0], s[1]); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 89:

		{
			x := yylex.(*lexer)
			f, ok := yyS[yypt-1].item.(*ident)
			if !ok {
				x.err("expected identifier or qualified identifier")
				goto ret1
			}

			var err error
			var agg bool
			if yyVAL.item, agg, err = newCall(f.s, yyS[yypt-0].item.([]expression)); err != nil {
				x.err("%v", err)
				goto ret1
			}
			if n := len(x.agg); n > 0 {
				x.agg[n-1] = x.agg[n-1] || agg
			}
		}
	case 90:
		yyVAL.item = yyS[yypt-0].item
	case 91:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('^', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 92:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('|', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 93:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation('-', yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 94:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('+', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 95:
		yyVAL.item = yyS[yypt-0].item
	case 96:

		{
			var err error
			yyVAL.item, err = newBinaryOperation(andnot, yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 97:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('&', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 98:

		{
			var err error
			yyVAL.item, err = newBinaryOperation(lsh, yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 99:

		{
			var err error
			yyVAL.item, err = newBinaryOperation(rsh, yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 100:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('%', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 101:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('/', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 102:

		{
			var err error
			yyVAL.item, err = newBinaryOperation('*', yyS[yypt-2].item, yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 103:
		yyVAL.item = yyS[yypt-0].item
	case 104:

		{
			yyVAL.item = fmt.Sprintf("%s.%s", yyS[yypt-2].item.(string), yyS[yypt-0].item.(string))
		}
	case 105:

		{
			yyVAL.item = []interface{}{yyS[yypt-1].item, yyS[yypt-0].item}
		}
	case 106:
		yyVAL.item = yyS[yypt-0].item
	case 107:

		{
			yyVAL.item = yyS[yypt-2].item
		}
	case 110:

		{
			yyVAL.item = ""
		}
	case 111:

		{
			yyVAL.item = yyS[yypt-0].item
		}
	case 112:

		{
			yyVAL.list = []interface{}{yyS[yypt-0].item}
		}
	case 113:

		{
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].item)
		}
	case 114:

		{
			yyVAL.item = rollbackStmt{}
		}
	case 115:

		{
			x := yylex.(*lexer)
			n := len(x.agg)
			yyVAL.item = &selectStmt{
				distinct:      yyS[yypt-6].item.(bool),
				flds:          yyS[yypt-5].item.([]*fld),
				from:          &crossJoinRset{sources: yyS[yypt-3].list},
				hasAggregates: x.agg[n-1],
				where:         yyS[yypt-2].item.(*whereRset),
				group:         yyS[yypt-1].item.(*groupByRset),
				order:         yyS[yypt-0].item.(*orderByRset),
			}
			x.agg = x.agg[:n-1]
		}
	case 116:

		{
			x := yylex.(*lexer)
			n := len(x.agg)
			yyVAL.item = &selectStmt{
				distinct:      yyS[yypt-7].item.(bool),
				flds:          yyS[yypt-6].item.([]*fld),
				from:          &crossJoinRset{sources: yyS[yypt-4].list},
				hasAggregates: x.agg[n-1],
				where:         yyS[yypt-2].item.(*whereRset),
				group:         yyS[yypt-1].item.(*groupByRset),
				order:         yyS[yypt-0].item.(*orderByRset),
			}
			x.agg = x.agg[:n-1]
		}
	case 117:

		{
			yyVAL.item = false
		}
	case 118:

		{
			yyVAL.item = true
		}
	case 119:

		{
			yyVAL.item = []*fld{}
		}
	case 120:

		{
			yyVAL.item = yyS[yypt-0].item
		}
	case 121:

		{
			yyVAL.item = yyS[yypt-1].item
		}
	case 122:

		{
			yyVAL.item = (*whereRset)(nil)
		}
	case 123:
		yyVAL.item = yyS[yypt-0].item
	case 124:

		{
			yyVAL.item = (*groupByRset)(nil)
		}
	case 125:
		yyVAL.item = yyS[yypt-0].item
	case 126:

		{
			yyVAL.item = (*orderByRset)(nil)
		}
	case 127:
		yyVAL.item = yyS[yypt-0].item
	case 128:

		{
			yyVAL.item = [2]*expression{nil, nil}
		}
	case 129:

		{
			hi := yyS[yypt-1].item.(expression)
			yyVAL.item = [2]*expression{nil, &hi}
		}
	case 130:

		{
			lo := yyS[yypt-2].item.(expression)
			yyVAL.item = [2]*expression{&lo, nil}
		}
	case 131:

		{
			lo := yyS[yypt-3].item.(expression)
			hi := yyS[yypt-1].item.(expression)
			yyVAL.item = [2]*expression{&lo, &hi}
		}
	case 132:
		yyVAL.item = yyS[yypt-0].item
	case 133:
		yyVAL.item = yyS[yypt-0].item
	case 134:
		yyVAL.item = yyS[yypt-0].item
	case 135:
		yyVAL.item = yyS[yypt-0].item
	case 136:
		yyVAL.item = yyS[yypt-0].item
	case 137:
		yyVAL.item = yyS[yypt-0].item
	case 138:
		yyVAL.item = yyS[yypt-0].item
	case 139:
		yyVAL.item = yyS[yypt-0].item
	case 140:
		yyVAL.item = yyS[yypt-0].item
	case 141:
		yyVAL.item = yyS[yypt-0].item
	case 142:
		yyVAL.item = yyS[yypt-0].item
	case 143:
		yyVAL.item = yyS[yypt-0].item
	case 144:
		yyVAL.item = yyS[yypt-0].item
	case 145:
		yyVAL.item = yyS[yypt-0].item
	case 146:

		{
			if yyS[yypt-0].item != nil {
				yylex.(*lexer).list = []stmt{yyS[yypt-0].item.(stmt)}
			}
		}
	case 147:

		{
			if yyS[yypt-0].item != nil {
				yylex.(*lexer).list = append(yylex.(*lexer).list, yyS[yypt-0].item.(stmt))
			}
		}
	case 148:
		yyVAL.item = yyS[yypt-0].item
	case 149:
		yyVAL.item = yyS[yypt-0].item
	case 150:

		{
			var err error
			if yyVAL.item, err = newBinaryOperation(andand, yyS[yypt-2].item, yyS[yypt-0].item); err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 151:

		{
			yyVAL.item = &truncateTableStmt{tableName: yyS[yypt-0].item.(string)}
		}
	case 152:
		yyVAL.item = yyS[yypt-0].item
	case 153:
		yyVAL.item = yyS[yypt-0].item
	case 154:
		yyVAL.item = yyS[yypt-0].item
	case 155:
		yyVAL.item = yyS[yypt-0].item
	case 156:
		yyVAL.item = yyS[yypt-0].item
	case 157:
		yyVAL.item = yyS[yypt-0].item
	case 158:
		yyVAL.item = yyS[yypt-0].item
	case 159:
		yyVAL.item = yyS[yypt-0].item
	case 160:
		yyVAL.item = yyS[yypt-0].item
	case 161:
		yyVAL.item = yyS[yypt-0].item
	case 162:
		yyVAL.item = yyS[yypt-0].item
	case 163:
		yyVAL.item = yyS[yypt-0].item
	case 164:
		yyVAL.item = yyS[yypt-0].item
	case 165:
		yyVAL.item = yyS[yypt-0].item
	case 166:
		yyVAL.item = yyS[yypt-0].item
	case 167:
		yyVAL.item = yyS[yypt-0].item
	case 168:
		yyVAL.item = yyS[yypt-0].item
	case 169:
		yyVAL.item = yyS[yypt-0].item
	case 170:
		yyVAL.item = yyS[yypt-0].item
	case 171:
		yyVAL.item = yyS[yypt-0].item
	case 172:
		yyVAL.item = yyS[yypt-0].item
	case 173:
		yyVAL.item = yyS[yypt-0].item
	case 174:
		yyVAL.item = yyS[yypt-0].item
	case 175:
		yyVAL.item = yyS[yypt-0].item
	case 176:

		{
			yyVAL.item = &updateStmt{tableName: yyS[yypt-2].item.(string), list: yyS[yypt-1].item.([]assignment), where: yyS[yypt-0].item.(*whereRset).expr}
		}
	case 177:

		{
			yyVAL.item = nowhere
		}
	case 178:
		yyVAL.item = yyS[yypt-0].item
	case 179:
		yyVAL.item = yyS[yypt-0].item
	case 180:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('^', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 181:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('!', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 182:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('-', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 183:

		{
			var err error
			yyVAL.item, err = newUnaryOperation('+', yyS[yypt-0].item)
			if err != nil {
				yylex.(*lexer).err("%v", err)
				goto ret1
			}
		}
	case 184:

		{
			yyVAL.item = &whereRset{expr: yyS[yypt-0].item.(expression)}
		}
	}
	goto yystack /* stack new state and value */
}
