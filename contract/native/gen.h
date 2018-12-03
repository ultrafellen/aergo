/**
 * @file    gen.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _GEN_H
#define _GEN_H

#include "common.h"

#include "ast.h"
#include "dsgmt.h"
#include "binaryen-c.h"

typedef struct gen_s {
    flag_t flag;
    char path[PATH_MAX_LEN + 5];

    ast_blk_t *root;
    BinaryenModuleRef module;

    dsgmt_t *dsgmt;
    int id_idx;

    int local_cnt;
    BinaryenType *locals;

    int buf_size;
    char *buf;
} gen_t;

void gen(ast_t *ast, flag_t flag, char *path);

#endif /* ! _GEN_H */
