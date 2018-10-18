/**
 * @file    enum.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "enum.h"

char *type_names_[TYPE_MAX] = {
    "undefined",
    "bool",
    "byte",
    "int8",
    "uint8",
    "int16",
    "uint16",
    "int32",
    "uint32",
    "int64",
    "uint64",
    "float",
    "double",
    "string",
    "account",
    "struct",
    "map",
    "object",
    "void",
    "tuple"
};

char *id_kinds_[ID_MAX] = {
    "variable",
    "struct",
    "enumeration",
    "function",
    "contract"
};

char *stmt_kinds_[STMT_MAX] = {
    "null",
    "exp",
    "if",
    "loop",
    "switch",
    "case",
    "continue",
    "break",
    "return",
    "goto",
    "sql",
    "block"
};

/* end of enum.c */
