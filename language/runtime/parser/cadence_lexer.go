// Code generated from parser/Cadence.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 81, 598,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3,
	54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60,
	3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3,
	62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64,
	3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3,
	66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68,
	3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 7, 69, 457,
	10, 69, 12, 69, 14, 69, 460, 11, 69, 3, 70, 5, 70, 463, 10, 70, 3, 71,
	3, 71, 5, 71, 467, 10, 71, 3, 72, 3, 72, 7, 72, 471, 10, 72, 12, 72, 14,
	72, 474, 11, 72, 3, 72, 5, 72, 477, 10, 72, 3, 72, 3, 72, 3, 72, 7, 72,
	482, 10, 72, 12, 72, 14, 72, 485, 11, 72, 3, 72, 5, 72, 488, 10, 72, 3,
	73, 3, 73, 7, 73, 492, 10, 73, 12, 73, 14, 73, 495, 11, 73, 3, 74, 3, 74,
	3, 74, 3, 74, 6, 74, 501, 10, 74, 13, 74, 14, 74, 502, 3, 75, 3, 75, 3,
	75, 3, 75, 6, 75, 509, 10, 75, 13, 75, 14, 75, 510, 3, 76, 3, 76, 3, 76,
	3, 76, 6, 76, 517, 10, 76, 13, 76, 14, 76, 518, 3, 77, 3, 77, 3, 77, 7,
	77, 524, 10, 77, 12, 77, 14, 77, 527, 11, 77, 3, 78, 3, 78, 7, 78, 531,
	10, 78, 12, 78, 14, 78, 534, 11, 78, 3, 78, 3, 78, 3, 79, 3, 79, 5, 79,
	540, 10, 79, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 6, 80, 549,
	10, 80, 13, 80, 14, 80, 550, 3, 80, 3, 80, 5, 80, 555, 10, 80, 3, 81, 3,
	81, 3, 82, 6, 82, 560, 10, 82, 13, 82, 14, 82, 561, 3, 82, 3, 82, 3, 83,
	6, 83, 567, 10, 83, 13, 83, 14, 83, 568, 3, 83, 3, 83, 3, 84, 3, 84, 3,
	84, 3, 84, 3, 84, 7, 84, 578, 10, 84, 12, 84, 14, 84, 581, 11, 84, 3, 84,
	3, 84, 3, 84, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3, 85, 7, 85, 592, 10,
	85, 12, 85, 14, 85, 595, 11, 85, 3, 85, 3, 85, 3, 579, 2, 86, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129,
	66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 2, 141, 2, 143, 71, 145, 72,
	147, 73, 149, 74, 151, 75, 153, 76, 155, 77, 157, 2, 159, 2, 161, 2, 163,
	78, 165, 79, 167, 80, 169, 81, 3, 2, 16, 5, 2, 67, 92, 97, 97, 99, 124,
	3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 4, 2, 50, 51, 97, 97, 4, 2, 50, 57,
	97, 97, 6, 2, 50, 59, 67, 72, 97, 97, 99, 104, 4, 2, 67, 92, 99, 124, 6,
	2, 50, 59, 67, 92, 97, 97, 99, 124, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94,
	9, 2, 36, 36, 41, 41, 50, 50, 94, 94, 112, 112, 116, 116, 118, 118, 5,
	2, 50, 59, 67, 72, 99, 104, 6, 2, 2, 2, 11, 11, 13, 14, 34, 34, 5, 2, 12,
	12, 15, 15, 8234, 8235, 4, 2, 12, 12, 15, 15, 2, 612, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89,
	3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2,
	97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2,
	2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111,
	3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2,
	2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3,
	2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2,
	133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 143, 3, 2,
	2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151,
	3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2,
	2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 3, 171, 3,
	2, 2, 2, 5, 173, 3, 2, 2, 2, 7, 175, 3, 2, 2, 2, 9, 177, 3, 2, 2, 2, 11,
	179, 3, 2, 2, 2, 13, 181, 3, 2, 2, 2, 15, 183, 3, 2, 2, 2, 17, 185, 3,
	2, 2, 2, 19, 187, 3, 2, 2, 2, 21, 191, 3, 2, 2, 2, 23, 193, 3, 2, 2, 2,
	25, 196, 3, 2, 2, 2, 27, 199, 3, 2, 2, 2, 29, 202, 3, 2, 2, 2, 31, 205,
	3, 2, 2, 2, 33, 207, 3, 2, 2, 2, 35, 209, 3, 2, 2, 2, 37, 212, 3, 2, 2,
	2, 39, 215, 3, 2, 2, 2, 41, 217, 3, 2, 2, 2, 43, 219, 3, 2, 2, 2, 45, 221,
	3, 2, 2, 2, 47, 223, 3, 2, 2, 2, 49, 225, 3, 2, 2, 2, 51, 230, 3, 2, 2,
	2, 53, 232, 3, 2, 2, 2, 55, 234, 3, 2, 2, 2, 57, 237, 3, 2, 2, 2, 59, 241,
	3, 2, 2, 2, 61, 243, 3, 2, 2, 2, 63, 247, 3, 2, 2, 2, 65, 250, 3, 2, 2,
	2, 67, 254, 3, 2, 2, 2, 69, 256, 3, 2, 2, 2, 71, 258, 3, 2, 2, 2, 73, 260,
	3, 2, 2, 2, 75, 272, 3, 2, 2, 2, 77, 279, 3, 2, 2, 2, 79, 288, 3, 2, 2,
	2, 81, 297, 3, 2, 2, 2, 83, 307, 3, 2, 2, 2, 85, 311, 3, 2, 2, 2, 87, 317,
	3, 2, 2, 2, 89, 322, 3, 2, 2, 2, 91, 326, 3, 2, 2, 2, 93, 331, 3, 2, 2,
	2, 95, 336, 3, 2, 2, 2, 97, 340, 3, 2, 2, 2, 99, 344, 3, 2, 2, 2, 101,
	351, 3, 2, 2, 2, 103, 355, 3, 2, 2, 2, 105, 360, 3, 2, 2, 2, 107, 368,
	3, 2, 2, 2, 109, 375, 3, 2, 2, 2, 111, 381, 3, 2, 2, 2, 113, 390, 3, 2,
	2, 2, 115, 394, 3, 2, 2, 2, 117, 398, 3, 2, 2, 2, 119, 401, 3, 2, 2, 2,
	121, 406, 3, 2, 2, 2, 123, 412, 3, 2, 2, 2, 125, 417, 3, 2, 2, 2, 127,
	423, 3, 2, 2, 2, 129, 427, 3, 2, 2, 2, 131, 434, 3, 2, 2, 2, 133, 439,
	3, 2, 2, 2, 135, 446, 3, 2, 2, 2, 137, 454, 3, 2, 2, 2, 139, 462, 3, 2,
	2, 2, 141, 466, 3, 2, 2, 2, 143, 468, 3, 2, 2, 2, 145, 489, 3, 2, 2, 2,
	147, 496, 3, 2, 2, 2, 149, 504, 3, 2, 2, 2, 151, 512, 3, 2, 2, 2, 153,
	520, 3, 2, 2, 2, 155, 528, 3, 2, 2, 2, 157, 539, 3, 2, 2, 2, 159, 554,
	3, 2, 2, 2, 161, 556, 3, 2, 2, 2, 163, 559, 3, 2, 2, 2, 165, 566, 3, 2,
	2, 2, 167, 572, 3, 2, 2, 2, 169, 587, 3, 2, 2, 2, 171, 172, 7, 61, 2, 2,
	172, 4, 3, 2, 2, 2, 173, 174, 7, 125, 2, 2, 174, 6, 3, 2, 2, 2, 175, 176,
	7, 127, 2, 2, 176, 8, 3, 2, 2, 2, 177, 178, 7, 46, 2, 2, 178, 10, 3, 2,
	2, 2, 179, 180, 7, 60, 2, 2, 180, 12, 3, 2, 2, 2, 181, 182, 7, 48, 2, 2,
	182, 14, 3, 2, 2, 2, 183, 184, 7, 93, 2, 2, 184, 16, 3, 2, 2, 2, 185, 186,
	7, 95, 2, 2, 186, 18, 3, 2, 2, 2, 187, 188, 7, 62, 2, 2, 188, 189, 7, 47,
	2, 2, 189, 190, 7, 64, 2, 2, 190, 20, 3, 2, 2, 2, 191, 192, 7, 63, 2, 2,
	192, 22, 3, 2, 2, 2, 193, 194, 7, 126, 2, 2, 194, 195, 7, 126, 2, 2, 195,
	24, 3, 2, 2, 2, 196, 197, 7, 40, 2, 2, 197, 198, 7, 40, 2, 2, 198, 26,
	3, 2, 2, 2, 199, 200, 7, 63, 2, 2, 200, 201, 7, 63, 2, 2, 201, 28, 3, 2,
	2, 2, 202, 203, 7, 35, 2, 2, 203, 204, 7, 63, 2, 2, 204, 30, 3, 2, 2, 2,
	205, 206, 7, 62, 2, 2, 206, 32, 3, 2, 2, 2, 207, 208, 7, 64, 2, 2, 208,
	34, 3, 2, 2, 2, 209, 210, 7, 62, 2, 2, 210, 211, 7, 63, 2, 2, 211, 36,
	3, 2, 2, 2, 212, 213, 7, 64, 2, 2, 213, 214, 7, 63, 2, 2, 214, 38, 3, 2,
	2, 2, 215, 216, 7, 45, 2, 2, 216, 40, 3, 2, 2, 2, 217, 218, 7, 47, 2, 2,
	218, 42, 3, 2, 2, 2, 219, 220, 7, 44, 2, 2, 220, 44, 3, 2, 2, 2, 221, 222,
	7, 49, 2, 2, 222, 46, 3, 2, 2, 2, 223, 224, 7, 39, 2, 2, 224, 48, 3, 2,
	2, 2, 225, 226, 7, 99, 2, 2, 226, 227, 7, 119, 2, 2, 227, 228, 7, 118,
	2, 2, 228, 229, 7, 106, 2, 2, 229, 50, 3, 2, 2, 2, 230, 231, 7, 40, 2,
	2, 231, 52, 3, 2, 2, 2, 232, 233, 7, 35, 2, 2, 233, 54, 3, 2, 2, 2, 234,
	235, 7, 62, 2, 2, 235, 236, 7, 47, 2, 2, 236, 56, 3, 2, 2, 2, 237, 238,
	7, 62, 2, 2, 238, 239, 7, 47, 2, 2, 239, 240, 7, 35, 2, 2, 240, 58, 3,
	2, 2, 2, 241, 242, 7, 65, 2, 2, 242, 60, 3, 2, 2, 2, 243, 244, 5, 163,
	82, 2, 244, 245, 7, 65, 2, 2, 245, 246, 7, 65, 2, 2, 246, 62, 3, 2, 2,
	2, 247, 248, 7, 99, 2, 2, 248, 249, 7, 117, 2, 2, 249, 64, 3, 2, 2, 2,
	250, 251, 7, 99, 2, 2, 251, 252, 7, 117, 2, 2, 252, 253, 7, 65, 2, 2, 253,
	66, 3, 2, 2, 2, 254, 255, 7, 66, 2, 2, 255, 68, 3, 2, 2, 2, 256, 257, 7,
	42, 2, 2, 257, 70, 3, 2, 2, 2, 258, 259, 7, 43, 2, 2, 259, 72, 3, 2, 2,
	2, 260, 261, 7, 118, 2, 2, 261, 262, 7, 116, 2, 2, 262, 263, 7, 99, 2,
	2, 263, 264, 7, 112, 2, 2, 264, 265, 7, 117, 2, 2, 265, 266, 7, 99, 2,
	2, 266, 267, 7, 101, 2, 2, 267, 268, 7, 118, 2, 2, 268, 269, 7, 107, 2,
	2, 269, 270, 7, 113, 2, 2, 270, 271, 7, 112, 2, 2, 271, 74, 3, 2, 2, 2,
	272, 273, 7, 117, 2, 2, 273, 274, 7, 118, 2, 2, 274, 275, 7, 116, 2, 2,
	275, 276, 7, 119, 2, 2, 276, 277, 7, 101, 2, 2, 277, 278, 7, 118, 2, 2,
	278, 76, 3, 2, 2, 2, 279, 280, 7, 116, 2, 2, 280, 281, 7, 103, 2, 2, 281,
	282, 7, 117, 2, 2, 282, 283, 7, 113, 2, 2, 283, 284, 7, 119, 2, 2, 284,
	285, 7, 116, 2, 2, 285, 286, 7, 101, 2, 2, 286, 287, 7, 103, 2, 2, 287,
	78, 3, 2, 2, 2, 288, 289, 7, 101, 2, 2, 289, 290, 7, 113, 2, 2, 290, 291,
	7, 112, 2, 2, 291, 292, 7, 118, 2, 2, 292, 293, 7, 116, 2, 2, 293, 294,
	7, 99, 2, 2, 294, 295, 7, 101, 2, 2, 295, 296, 7, 118, 2, 2, 296, 80, 3,
	2, 2, 2, 297, 298, 7, 107, 2, 2, 298, 299, 7, 112, 2, 2, 299, 300, 7, 118,
	2, 2, 300, 301, 7, 103, 2, 2, 301, 302, 7, 116, 2, 2, 302, 303, 7, 104,
	2, 2, 303, 304, 7, 99, 2, 2, 304, 305, 7, 101, 2, 2, 305, 306, 7, 103,
	2, 2, 306, 82, 3, 2, 2, 2, 307, 308, 7, 104, 2, 2, 308, 309, 7, 119, 2,
	2, 309, 310, 7, 112, 2, 2, 310, 84, 3, 2, 2, 2, 311, 312, 7, 103, 2, 2,
	312, 313, 7, 120, 2, 2, 313, 314, 7, 103, 2, 2, 314, 315, 7, 112, 2, 2,
	315, 316, 7, 118, 2, 2, 316, 86, 3, 2, 2, 2, 317, 318, 7, 103, 2, 2, 318,
	319, 7, 111, 2, 2, 319, 320, 7, 107, 2, 2, 320, 321, 7, 118, 2, 2, 321,
	88, 3, 2, 2, 2, 322, 323, 7, 114, 2, 2, 323, 324, 7, 116, 2, 2, 324, 325,
	7, 103, 2, 2, 325, 90, 3, 2, 2, 2, 326, 327, 7, 114, 2, 2, 327, 328, 7,
	113, 2, 2, 328, 329, 7, 117, 2, 2, 329, 330, 7, 118, 2, 2, 330, 92, 3,
	2, 2, 2, 331, 332, 7, 114, 2, 2, 332, 333, 7, 116, 2, 2, 333, 334, 7, 107,
	2, 2, 334, 335, 7, 120, 2, 2, 335, 94, 3, 2, 2, 2, 336, 337, 7, 114, 2,
	2, 337, 338, 7, 119, 2, 2, 338, 339, 7, 100, 2, 2, 339, 96, 3, 2, 2, 2,
	340, 341, 7, 117, 2, 2, 341, 342, 7, 103, 2, 2, 342, 343, 7, 118, 2, 2,
	343, 98, 3, 2, 2, 2, 344, 345, 7, 99, 2, 2, 345, 346, 7, 101, 2, 2, 346,
	347, 7, 101, 2, 2, 347, 348, 7, 103, 2, 2, 348, 349, 7, 117, 2, 2, 349,
	350, 7, 117, 2, 2, 350, 100, 3, 2, 2, 2, 351, 352, 7, 99, 2, 2, 352, 353,
	7, 110, 2, 2, 353, 354, 7, 110, 2, 2, 354, 102, 3, 2, 2, 2, 355, 356, 7,
	117, 2, 2, 356, 357, 7, 103, 2, 2, 357, 358, 7, 110, 2, 2, 358, 359, 7,
	104, 2, 2, 359, 104, 3, 2, 2, 2, 360, 361, 7, 99, 2, 2, 361, 362, 7, 101,
	2, 2, 362, 363, 7, 101, 2, 2, 363, 364, 7, 113, 2, 2, 364, 365, 7, 119,
	2, 2, 365, 366, 7, 112, 2, 2, 366, 367, 7, 118, 2, 2, 367, 106, 3, 2, 2,
	2, 368, 369, 7, 116, 2, 2, 369, 370, 7, 103, 2, 2, 370, 371, 7, 118, 2,
	2, 371, 372, 7, 119, 2, 2, 372, 373, 7, 116, 2, 2, 373, 374, 7, 112, 2,
	2, 374, 108, 3, 2, 2, 2, 375, 376, 7, 100, 2, 2, 376, 377, 7, 116, 2, 2,
	377, 378, 7, 103, 2, 2, 378, 379, 7, 99, 2, 2, 379, 380, 7, 109, 2, 2,
	380, 110, 3, 2, 2, 2, 381, 382, 7, 101, 2, 2, 382, 383, 7, 113, 2, 2, 383,
	384, 7, 112, 2, 2, 384, 385, 7, 118, 2, 2, 385, 386, 7, 107, 2, 2, 386,
	387, 7, 112, 2, 2, 387, 388, 7, 119, 2, 2, 388, 389, 7, 103, 2, 2, 389,
	112, 3, 2, 2, 2, 390, 391, 7, 110, 2, 2, 391, 392, 7, 103, 2, 2, 392, 393,
	7, 118, 2, 2, 393, 114, 3, 2, 2, 2, 394, 395, 7, 120, 2, 2, 395, 396, 7,
	99, 2, 2, 396, 397, 7, 116, 2, 2, 397, 116, 3, 2, 2, 2, 398, 399, 7, 107,
	2, 2, 399, 400, 7, 104, 2, 2, 400, 118, 3, 2, 2, 2, 401, 402, 7, 103, 2,
	2, 402, 403, 7, 110, 2, 2, 403, 404, 7, 117, 2, 2, 404, 405, 7, 103, 2,
	2, 405, 120, 3, 2, 2, 2, 406, 407, 7, 121, 2, 2, 407, 408, 7, 106, 2, 2,
	408, 409, 7, 107, 2, 2, 409, 410, 7, 110, 2, 2, 410, 411, 7, 103, 2, 2,
	411, 122, 3, 2, 2, 2, 412, 413, 7, 118, 2, 2, 413, 414, 7, 116, 2, 2, 414,
	415, 7, 119, 2, 2, 415, 416, 7, 103, 2, 2, 416, 124, 3, 2, 2, 2, 417, 418,
	7, 104, 2, 2, 418, 419, 7, 99, 2, 2, 419, 420, 7, 110, 2, 2, 420, 421,
	7, 117, 2, 2, 421, 422, 7, 103, 2, 2, 422, 126, 3, 2, 2, 2, 423, 424, 7,
	112, 2, 2, 424, 425, 7, 107, 2, 2, 425, 426, 7, 110, 2, 2, 426, 128, 3,
	2, 2, 2, 427, 428, 7, 107, 2, 2, 428, 429, 7, 111, 2, 2, 429, 430, 7, 114,
	2, 2, 430, 431, 7, 113, 2, 2, 431, 432, 7, 116, 2, 2, 432, 433, 7, 118,
	2, 2, 433, 130, 3, 2, 2, 2, 434, 435, 7, 104, 2, 2, 435, 436, 7, 116, 2,
	2, 436, 437, 7, 113, 2, 2, 437, 438, 7, 111, 2, 2, 438, 132, 3, 2, 2, 2,
	439, 440, 7, 101, 2, 2, 440, 441, 7, 116, 2, 2, 441, 442, 7, 103, 2, 2,
	442, 443, 7, 99, 2, 2, 443, 444, 7, 118, 2, 2, 444, 445, 7, 103, 2, 2,
	445, 134, 3, 2, 2, 2, 446, 447, 7, 102, 2, 2, 447, 448, 7, 103, 2, 2, 448,
	449, 7, 117, 2, 2, 449, 450, 7, 118, 2, 2, 450, 451, 7, 116, 2, 2, 451,
	452, 7, 113, 2, 2, 452, 453, 7, 123, 2, 2, 453, 136, 3, 2, 2, 2, 454, 458,
	5, 139, 70, 2, 455, 457, 5, 141, 71, 2, 456, 455, 3, 2, 2, 2, 457, 460,
	3, 2, 2, 2, 458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 138, 3, 2,
	2, 2, 460, 458, 3, 2, 2, 2, 461, 463, 9, 2, 2, 2, 462, 461, 3, 2, 2, 2,
	463, 140, 3, 2, 2, 2, 464, 467, 9, 3, 2, 2, 465, 467, 5, 139, 70, 2, 466,
	464, 3, 2, 2, 2, 466, 465, 3, 2, 2, 2, 467, 142, 3, 2, 2, 2, 468, 476,
	9, 3, 2, 2, 469, 471, 9, 4, 2, 2, 470, 469, 3, 2, 2, 2, 471, 474, 3, 2,
	2, 2, 472, 470, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 475, 3, 2, 2, 2,
	474, 472, 3, 2, 2, 2, 475, 477, 9, 3, 2, 2, 476, 472, 3, 2, 2, 2, 476,
	477, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478, 479, 7, 48, 2, 2, 479, 487,
	9, 3, 2, 2, 480, 482, 9, 4, 2, 2, 481, 480, 3, 2, 2, 2, 482, 485, 3, 2,
	2, 2, 483, 481, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484, 486, 3, 2, 2, 2,
	485, 483, 3, 2, 2, 2, 486, 488, 9, 3, 2, 2, 487, 483, 3, 2, 2, 2, 487,
	488, 3, 2, 2, 2, 488, 144, 3, 2, 2, 2, 489, 493, 9, 3, 2, 2, 490, 492,
	9, 4, 2, 2, 491, 490, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493, 491, 3, 2,
	2, 2, 493, 494, 3, 2, 2, 2, 494, 146, 3, 2, 2, 2, 495, 493, 3, 2, 2, 2,
	496, 497, 7, 50, 2, 2, 497, 498, 7, 100, 2, 2, 498, 500, 3, 2, 2, 2, 499,
	501, 9, 5, 2, 2, 500, 499, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 500,
	3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 148, 3, 2, 2, 2, 504, 505, 7, 50,
	2, 2, 505, 506, 7, 113, 2, 2, 506, 508, 3, 2, 2, 2, 507, 509, 9, 6, 2,
	2, 508, 507, 3, 2, 2, 2, 509, 510, 3, 2, 2, 2, 510, 508, 3, 2, 2, 2, 510,
	511, 3, 2, 2, 2, 511, 150, 3, 2, 2, 2, 512, 513, 7, 50, 2, 2, 513, 514,
	7, 122, 2, 2, 514, 516, 3, 2, 2, 2, 515, 517, 9, 7, 2, 2, 516, 515, 3,
	2, 2, 2, 517, 518, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2, 518, 519, 3, 2, 2,
	2, 519, 152, 3, 2, 2, 2, 520, 521, 7, 50, 2, 2, 521, 525, 9, 8, 2, 2, 522,
	524, 9, 9, 2, 2, 523, 522, 3, 2, 2, 2, 524, 527, 3, 2, 2, 2, 525, 523,
	3, 2, 2, 2, 525, 526, 3, 2, 2, 2, 526, 154, 3, 2, 2, 2, 527, 525, 3, 2,
	2, 2, 528, 532, 7, 36, 2, 2, 529, 531, 5, 157, 79, 2, 530, 529, 3, 2, 2,
	2, 531, 534, 3, 2, 2, 2, 532, 530, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533,
	535, 3, 2, 2, 2, 534, 532, 3, 2, 2, 2, 535, 536, 7, 36, 2, 2, 536, 156,
	3, 2, 2, 2, 537, 540, 5, 159, 80, 2, 538, 540, 10, 10, 2, 2, 539, 537,
	3, 2, 2, 2, 539, 538, 3, 2, 2, 2, 540, 158, 3, 2, 2, 2, 541, 542, 7, 94,
	2, 2, 542, 555, 9, 11, 2, 2, 543, 544, 7, 94, 2, 2, 544, 545, 7, 119, 2,
	2, 545, 546, 3, 2, 2, 2, 546, 548, 7, 125, 2, 2, 547, 549, 5, 161, 81,
	2, 548, 547, 3, 2, 2, 2, 549, 550, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 550,
	551, 3, 2, 2, 2, 551, 552, 3, 2, 2, 2, 552, 553, 7, 127, 2, 2, 553, 555,
	3, 2, 2, 2, 554, 541, 3, 2, 2, 2, 554, 543, 3, 2, 2, 2, 555, 160, 3, 2,
	2, 2, 556, 557, 9, 12, 2, 2, 557, 162, 3, 2, 2, 2, 558, 560, 9, 13, 2,
	2, 559, 558, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 561,
	562, 3, 2, 2, 2, 562, 563, 3, 2, 2, 2, 563, 564, 8, 82, 2, 2, 564, 164,
	3, 2, 2, 2, 565, 567, 9, 14, 2, 2, 566, 565, 3, 2, 2, 2, 567, 568, 3, 2,
	2, 2, 568, 566, 3, 2, 2, 2, 568, 569, 3, 2, 2, 2, 569, 570, 3, 2, 2, 2,
	570, 571, 8, 83, 2, 2, 571, 166, 3, 2, 2, 2, 572, 573, 7, 49, 2, 2, 573,
	574, 7, 44, 2, 2, 574, 579, 3, 2, 2, 2, 575, 578, 5, 167, 84, 2, 576, 578,
	11, 2, 2, 2, 577, 575, 3, 2, 2, 2, 577, 576, 3, 2, 2, 2, 578, 581, 3, 2,
	2, 2, 579, 580, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 580, 582, 3, 2, 2, 2,
	581, 579, 3, 2, 2, 2, 582, 583, 7, 44, 2, 2, 583, 584, 7, 49, 2, 2, 584,
	585, 3, 2, 2, 2, 585, 586, 8, 84, 2, 2, 586, 168, 3, 2, 2, 2, 587, 588,
	7, 49, 2, 2, 588, 589, 7, 49, 2, 2, 589, 593, 3, 2, 2, 2, 590, 592, 10,
	15, 2, 2, 591, 590, 3, 2, 2, 2, 592, 595, 3, 2, 2, 2, 593, 591, 3, 2, 2,
	2, 593, 594, 3, 2, 2, 2, 594, 596, 3, 2, 2, 2, 595, 593, 3, 2, 2, 2, 596,
	597, 8, 85, 2, 2, 597, 170, 3, 2, 2, 2, 24, 2, 458, 462, 466, 472, 476,
	483, 487, 493, 502, 510, 518, 525, 532, 539, 550, 554, 561, 568, 577, 579,
	593, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'{'", "'}'", "','", "':'", "'.'", "'['", "']'", "'<->'", "'='",
	"'||'", "'&&'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'auth'", "'&'", "'!'", "'<-'", "'<-!'", "'?'", "",
	"'as'", "'as?'", "'@'", "'('", "')'", "'transaction'", "'struct'", "'resource'",
	"'contract'", "'interface'", "'fun'", "'event'", "'emit'", "'pre'", "'post'",
	"'priv'", "'pub'", "'set'", "'access'", "'all'", "'self'", "'account'",
	"'return'", "'break'", "'continue'", "'let'", "'var'", "'if'", "'else'",
	"'while'", "'true'", "'false'", "'nil'", "'import'", "'from'", "'create'",
	"'destroy'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Auth", "Ampersand", "Negate", "Move", "MoveForced", "Optional",
	"NilCoalescing", "Casting", "FailableCasting", "ResourceAnnotation", "OpenParen",
	"CloseParen", "Transaction", "Struct", "Resource", "Contract", "Interface",
	"Fun", "Event", "Emit", "Pre", "Post", "Priv", "Pub", "Set", "Access",
	"All", "Self", "Account", "Return", "Break", "Continue", "Let", "Var",
	"If", "Else", "While", "True", "False", "Nil", "Import", "From", "Create",
	"Destroy", "Identifier", "PositiveFixedPointLiteral", "DecimalLiteral",
	"BinaryLiteral", "OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral",
	"StringLiteral", "WS", "Terminator", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "Equal", "Unequal", "Less", "Greater", "LessEqual",
	"GreaterEqual", "Plus", "Minus", "Mul", "Div", "Mod", "Auth", "Ampersand",
	"Negate", "Move", "MoveForced", "Optional", "NilCoalescing", "Casting",
	"FailableCasting", "ResourceAnnotation", "OpenParen", "CloseParen", "Transaction",
	"Struct", "Resource", "Contract", "Interface", "Fun", "Event", "Emit",
	"Pre", "Post", "Priv", "Pub", "Set", "Access", "All", "Self", "Account",
	"Return", "Break", "Continue", "Let", "Var", "If", "Else", "While", "True",
	"False", "Nil", "Import", "From", "Create", "Destroy", "Identifier", "IdentifierHead",
	"IdentifierCharacter", "PositiveFixedPointLiteral", "DecimalLiteral", "BinaryLiteral",
	"OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral",
	"QuotedText", "EscapedCharacter", "HexadecimalDigit", "WS", "Terminator",
	"BlockComment", "LineComment",
}

type CadenceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCadenceLexer(input antlr.CharStream) *CadenceLexer {

	l := new(CadenceLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Cadence.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CadenceLexer tokens.
const (
	CadenceLexerT__0                      = 1
	CadenceLexerT__1                      = 2
	CadenceLexerT__2                      = 3
	CadenceLexerT__3                      = 4
	CadenceLexerT__4                      = 5
	CadenceLexerT__5                      = 6
	CadenceLexerT__6                      = 7
	CadenceLexerT__7                      = 8
	CadenceLexerT__8                      = 9
	CadenceLexerT__9                      = 10
	CadenceLexerT__10                     = 11
	CadenceLexerT__11                     = 12
	CadenceLexerEqual                     = 13
	CadenceLexerUnequal                   = 14
	CadenceLexerLess                      = 15
	CadenceLexerGreater                   = 16
	CadenceLexerLessEqual                 = 17
	CadenceLexerGreaterEqual              = 18
	CadenceLexerPlus                      = 19
	CadenceLexerMinus                     = 20
	CadenceLexerMul                       = 21
	CadenceLexerDiv                       = 22
	CadenceLexerMod                       = 23
	CadenceLexerAuth                      = 24
	CadenceLexerAmpersand                 = 25
	CadenceLexerNegate                    = 26
	CadenceLexerMove                      = 27
	CadenceLexerMoveForced                = 28
	CadenceLexerOptional                  = 29
	CadenceLexerNilCoalescing             = 30
	CadenceLexerCasting                   = 31
	CadenceLexerFailableCasting           = 32
	CadenceLexerResourceAnnotation        = 33
	CadenceLexerOpenParen                 = 34
	CadenceLexerCloseParen                = 35
	CadenceLexerTransaction               = 36
	CadenceLexerStruct                    = 37
	CadenceLexerResource                  = 38
	CadenceLexerContract                  = 39
	CadenceLexerInterface                 = 40
	CadenceLexerFun                       = 41
	CadenceLexerEvent                     = 42
	CadenceLexerEmit                      = 43
	CadenceLexerPre                       = 44
	CadenceLexerPost                      = 45
	CadenceLexerPriv                      = 46
	CadenceLexerPub                       = 47
	CadenceLexerSet                       = 48
	CadenceLexerAccess                    = 49
	CadenceLexerAll                       = 50
	CadenceLexerSelf                      = 51
	CadenceLexerAccount                   = 52
	CadenceLexerReturn                    = 53
	CadenceLexerBreak                     = 54
	CadenceLexerContinue                  = 55
	CadenceLexerLet                       = 56
	CadenceLexerVar                       = 57
	CadenceLexerIf                        = 58
	CadenceLexerElse                      = 59
	CadenceLexerWhile                     = 60
	CadenceLexerTrue                      = 61
	CadenceLexerFalse                     = 62
	CadenceLexerNil                       = 63
	CadenceLexerImport                    = 64
	CadenceLexerFrom                      = 65
	CadenceLexerCreate                    = 66
	CadenceLexerDestroy                   = 67
	CadenceLexerIdentifier                = 68
	CadenceLexerPositiveFixedPointLiteral = 69
	CadenceLexerDecimalLiteral            = 70
	CadenceLexerBinaryLiteral             = 71
	CadenceLexerOctalLiteral              = 72
	CadenceLexerHexadecimalLiteral        = 73
	CadenceLexerInvalidNumberLiteral      = 74
	CadenceLexerStringLiteral             = 75
	CadenceLexerWS                        = 76
	CadenceLexerTerminator                = 77
	CadenceLexerBlockComment              = 78
	CadenceLexerLineComment               = 79
)
