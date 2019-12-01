package opcodes

const (
	BLOCK           = 0x02
	LOOP            = 0x03
	BR              = 0x0C
	BR_IF           = 0x0D
	END             = 0x0B
	CALL            = 0x10
	GET_LOCAL       = 0x20
	SET_LOCAL       = 0x21
	I32_STORE_8     = 0x3A
	I32_CONST       = 0x41
	F32_CONST       = 0x43
	I32_EQZ         = 0x45
	I32_EQ          = 0x46
	F32_EQ          = 0x5B
	F32_LT          = 0x5D
	F32_GT          = 0x5E
	I32_AND         = 0x71
	F32_ADD         = 0x92
	F32_SUB         = 0x93
	F32_MUL         = 0x94
	F32_DIV         = 0x95
	I32_TRUNC_F32_S = 0xA8

	// Function type
	FUNCTION    = 0x60
	EMPTY_ARRAY = 0x0
)

// Module headers
var MAGIC_MODULE_HEADER = [...]int{0x00, 0x61, 0x73, 0x6d}
var MODULE_VERSION = [...]int{0x01, 0x00, 0x00, 0x00}

var BinaryOpcode = map[string]int{
	"+":  F32_ADD,
	"-":  F32_SUB,
	"*":  F32_MUL,
	"/":  F32_DIV,
	"==": F32_EQ,
	">":  F32_GT,
	"<":  F32_LT,
	"&&": I32_AND,
}

// Export types

type ExportType = int

type exportTypeList struct {
	FUNC   ExportType
	TABLE  ExportType
	MEM    ExportType
	GLOBAL ExportType
}

var Export = &exportTypeList{
	FUNC:   0x00,
	TABLE:  0x01,
	MEM:    0x02,
	GLOBAL: 0x03,
}

// Section types

type SectionType = int

type sectionTypeList struct {
	CUSTOM  SectionType
	TYPE    SectionType
	IMPORT  SectionType
	FUNC    SectionType
	TABLE   SectionType
	MEMORY  SectionType
	GLOBAL  SectionType
	EXPORT  SectionType
	START   SectionType
	ELEMENT SectionType
	CODE    SectionType
	DATA    SectionType
}

var Section = &sectionTypeList{
	CUSTOM:  0,
	TYPE:    1,
	IMPORT:  2,
	FUNC:    3,
	TABLE:   4,
	MEMORY:  5,
	GLOBAL:  6,
	EXPORT:  7,
	START:   8,
	ELEMENT: 9,
	CODE:    10,
	DATA:    11,
}

// Value types

type ValType = int

type valTypeList struct {
	I32 ValType
	F32 ValType
}

var Val = &valTypeList{
	I32: 0x7f,
	F32: 0x7d,
}

// Block types

type BlockType = int

type blockTypeList struct {
	VOID BlockType
}

var Block = &blockTypeList{
	VOID: 0x40,
}
