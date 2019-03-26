/**
 * @file    value.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _VALUE_H
#define _VALUE_H

#include "common.h"

#define is_null_val(val)            ((val)->size == 0)

#define is_bool_val(val)            ((val)->type == TYPE_BOOL)
#define is_int_val(val)             ((val)->type == TYPE_UINT128)
#define is_f64_val(val)             ((val)->type == TYPE_DOUBLE)
#define is_str_val(val)             ((val)->type == TYPE_STRING)
#define is_ptr_val(val)             ((val)->type == TYPE_OBJECT)

#define val_ptr(val)                ((val)->ptr)
#define val_size(val)               ((val)->size)

#define val_bool(val)               ((val)->b)
#define val_mpz(val)                ((val)->z)
#define val_f64(val)                ((val)->d)
#define val_str(val)                ((char *)((val)->ptr))

#define val_i64(val)                mpz_get_ui((val)->z)

#define is_zero_val(val)                                                                 \
    (is_int_val(val) ? mpz_sgn(val_mpz(val)) == 0 :                                      \
     (is_f64_val(val) ? (val)->d == 0.0f : false))

#define is_neg_val(val)             (mpz_sgn(val_mpz(val)) < 0)

#define value_init_int(val)                                                              \
    do {                                                                                 \
        (val)->type = TYPE_UINT128;                                                      \
        (val)->size = sizeof(mpz_t);                                                     \
        (val)->ptr = (val)->z;                                                           \
        mpz_init2((val)->z, 256);                                                        \
    } while (0)

#define value_set_null(val)         value_set_ptr(val, NULL, 0)

#define value_set_bool(val, v)                                                           \
    do {                                                                                 \
        (val)->type = TYPE_BOOL;                                                         \
        (val)->size = sizeof(bool);                                                      \
        (val)->ptr = &(val)->b;                                                          \
        (val)->b = (v);                                                                  \
    } while (0)

#define value_set_int(val, v)                                                            \
    do {                                                                                 \
        value_init_int(val);                                                             \
        mpz_set_ui((val)->z, (v));                                                       \
    } while (0)

#define value_set_f64(val, v)                                                            \
    do {                                                                                 \
        (val)->type = TYPE_DOUBLE;                                                       \
        (val)->size = sizeof(double);                                                    \
        (val)->ptr = &(val)->d;                                                          \
        (val)->d = (v);                                                                  \
    } while (0)

#define value_set_str(val, v)                                                            \
    do {                                                                                 \
        (val)->type = TYPE_STRING;                                                       \
        (val)->size = strlen(v);                                                         \
        (val)->ptr = (v);                                                                \
    } while (0)

#define value_set_ptr(val, v, l)                                                         \
    do {                                                                                 \
        (val)->type = TYPE_OBJECT;                                                       \
        (val)->size = l;                                                                 \
        (val)->ptr = (v);                                                                \
    } while (0)

#define value_fits_i32(val)                                                              \
    (mpz_fits_sint_p(val_mpz(val)) || mpz_fits_uint_p(val_mpz(val)))
#define value_fits_i64(val)                                                              \
    (mpz_fits_slong_p(val_mpz(val)) || mpz_fits_ulong_p(val_mpz(val)))

#ifndef _VALUE_T
#define _VALUE_T
typedef struct value_s value_t;
#endif /* ! _VALUE_T */

#ifndef _META_T
#define _META_T
typedef struct meta_s meta_t;
#endif /* ! _META_T */

typedef void (*eval_fn_t)(value_t *, value_t *, value_t *);
typedef void (*cast_fn_t)(value_t *);

struct value_s {
    type_t type;
    int size;

    void *ptr;
    union {
        bool b;
        mpz_t z;
        double d;
        char *s;
    };
};

bool value_fit(value_t *val, meta_t *meta);

int value_cmp(value_t *x, value_t *y);

void value_eval(value_t *x, op_kind_t op, value_t *y, value_t *res);
void value_cast(value_t *val, type_t type);

static inline void
value_init(value_t *val)
{
    ASSERT(val != NULL);
    memset(val, 0x00, sizeof(value_t));
}

static inline uint32_t
value_serialize(value_t *val, char *buf, meta_t *meta)
{
    switch (val->type) {
    case TYPE_BOOL:
        *(uint32_t *)buf = val_bool(val) ? 1 : 0;
        return sizeof(uint32_t);

    case TYPE_UINT128:
        ASSERT1(!is_int128_meta(meta) && !is_uint128_meta(meta) &&
                !is_int128_meta(meta) && !is_uint128_meta(meta), meta->type);

        if (is_int64_meta(meta) || is_uint64_meta(meta)) {
            *(uint64_t *)buf = val_i64(val);
            return sizeof(uint64_t);
        }

        *(uint32_t *)buf = (uint32_t)val_i64(val);
        return sizeof(uint32_t);

    case TYPE_DOUBLE:
        if (is_double_meta(meta)) {
            *(double *)buf = val_f64(val);
            return sizeof(double);
        }

        *(float *)buf = (float)val_f64(val);
        return sizeof(float);

    case TYPE_STRING:
    case TYPE_OBJECT:
        memcpy(buf, val->ptr, val->size);
        return val->size;

    default:
        ASSERT2(!"invalid value", val->type, meta->type);
    }

    return 0;
}

#endif /* ! _VALUE_H */
