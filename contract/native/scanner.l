%{

/**
 *  @file   scanner.l
 *  @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "util.h"
#include "parser.h"

#define YY_PARAM_DECL           YYSTYPE *lval, YYLTYPE *lloc, yyscan_t yyscanner
#define YY_DECL                 int yyscan(YY_PARAM_DECL)

#undef  YY_INPUT
#define YY_INPUT(buf, result, MAX)                                             \
    do {                                                                       \
        (result) = MIN((MAX), yyextra->len - yyextra->pos);                    \
        if ((result) > 0) {                                                    \
            memcpy(buf, yyextra->src + yyextra->pos, (result));                \
            yyextra->pos += (result);                                          \
        }                                                                      \
    } while (0)

#define YY_EXTRA_TYPE           yyparam_t *
#define YY_BUF                  &yyextra->buf

#define yy_update_last_col()    lloc->last.col += yyleng

#define yy_update_last_line()                                                  \
    do {                                                                       \
        lloc->last.line++;                                                     \
        lloc->last.offset += lloc->last.col;                                   \
        lloc->last.col = 1;                                                    \
    } while (0)

#define yy_update_first()       lloc->first = lloc->last

#define yy_text()               xstrndup(yytext, yyleng)
#define yy_literal()            xstrndup(yytext + 1, yyleng - 2)

#define yy_append()             strbuf_append(YY_BUF, yytext, yyleng)

#define yy_error(ec)                                                           \
    ERROR(ec, FILENAME(yyextra->path), lloc->first.line,                       \
          make_trace(yyextra->path, lloc))

#define yy_recover()                                                           \
    do {                                                                       \
        register int c;                                                        \
        for ( ; ; ) {                                                          \
            c = input(yyscanner);                                              \
            if (c == '\r' || c == '\n') {                                      \
                unput(c);                                                      \
                break;                                                         \
            }                                                                  \
            else if (c == '\0') {                                              \
                yyterminate();                                                 \
            }                                                                  \
            lloc->last.col++;                                                  \
        }                                                                      \
    } while (0)

#define yy_return(tok)                                                         \
    do {                                                                       \
        yy_update_first();                                                     \
        yy_update_last_col();                                                  \
        return (tok);                                                          \
    } while (0)

#define yy_return_buf(tok)                                                     \
    do {                                                                       \
        lval->str = xstrndup(strbuf_text(YY_BUF), strbuf_length(YY_BUF));      \
        strbuf_reset(YY_BUF);                                                  \
        return (tok);                                                          \
    } while (0)

#define yy_scan_comment()                                                      \
    do {                                                                       \
        char c, prev = '\0';                                                   \
        while ((c = input(yyscanner)) != '\0') {                               \
            if (c == '/' && prev == '*') {                                     \
                lloc->last.col++;                                              \
                break;                                                         \
            }                                                                  \
            if (c == '\n')                                                     \
                yy_update_last_line();                                         \
            else                                                               \
                lloc->last.col++;                                              \
            prev = c;                                                          \
        }                                                                      \
        if (c == '\0') {                                                       \
            yy_error(ERROR_UNTERMINATED_COMMENT);                              \
            yyterminate();                                                     \
        }                                                                      \
    } while (0)

#define yy_scan_literal()                                                      \
    do {                                                                       \
        char c, prev = '\0';                                                   \
        while ((c = input(yyscanner)) != '\0') {                               \
            if (c == '"' && prev != '\\') {                                    \
                lloc->last.col++;                                              \
                break;                                                         \
            }                                                                  \
            if (c == '\n') {                                                   \
                strbuf_reset(YY_BUF);                                          \
                unput(c);                                                      \
                yy_error(ERROR_UNTERMINATED_STRING);                           \
                break;                                                         \
            }                                                                  \
            else {                                                             \
                lloc->last.col++;                                              \
                strbuf_append(YY_BUF, &c, 1);                                  \
            }                                                                  \
            prev = c;                                                          \
        }                                                                      \
        if (c == '\0') {                                                       \
            yy_error(ERROR_UNTERMINATED_STRING);                               \
            yyterminate();                                                     \
        }                                                                      \
    } while (0)

%}

%option case-sensitive
%option never-interactive
%option debug
%option reentrant
%option noyylineno
%option noyywrap

%x DIRECTIVE

ws                          [ \t\f\v]
nl                          ([\r\n]|\r\n)

D                           [0-9]
L                           [a-zA-Z\x80-\xFF_]
H                           [a-fA-F0-9]

integer                     [+-]?{D}+
float                       [+-]?({D}*"."{D}+|{D}+"."{D}*)
hexa                        0[xX]{H}+

id                          {L}({L}|{D})*

%%

{ws}+                       yy_update_last_col();
{nl}                        yy_update_last_line();

"//"[^\r\n]*                yy_update_last_col();

"/*" {
    yy_update_first();
    yy_update_last_col();
    yy_scan_comment();
}

"\"" {
    yy_update_first();
    yy_update_last_col();
    yy_scan_literal();
    yy_return_buf(STRING);
}

{integer}                   yy_return(INT);
{float}                     yy_return(FLOAT);
{hexa}                      yy_return(HEXA);

"+="                        yy_return(OP_ADD_ASSIGN);
"-="                        yy_return(OP_SUB_ASSIGN);
"*="                        yy_return(OP_MUL_ASSIGN);
"/="                        yy_return(OP_DIV_ASSIGN);
"%="                        yy_return(OP_MOD_ASSIGN);
"&="                        yy_return(OP_AND_ASSIGN);
"^="                        yy_return(OP_XOR_ASSIGN);
"|="                        yy_return(OP_OR_ASSIGN);
">>="                       yy_return(OP_RS_ASSIGN);
"<<="                       yy_return(OP_LS_ASSIGN);
">>"                        yy_return(OP_RSHIFT);
"<<"                        yy_return(OP_LSHIFT);
"++"                        yy_return(OP_INC);
"--"                        yy_return(OP_DEC);
"&&"                        yy_return(OP_AND);
"||"                        yy_return(OP_OR);
"<="                        yy_return(OP_LE);
">="                        yy_return(OP_GE);
"=="                        yy_return(OP_EQ);
"!="                        yy_return(OP_NE);

"{"                         yy_return('{');
"}"                         yy_return('}');
"("                         yy_return('(');
")"                         yy_return(')');
"["                         yy_return('[');
"]"                         yy_return(']');
";"                         yy_return(';');
":"                         yy_return(':');
","                         yy_return(',');
"?"                         yy_return('?');
"."                         yy_return('.');
"+"                         yy_return('+');
"-"                         yy_return('-');
"*"                         yy_return('*');
"/"                         yy_return('/');
"%"                         yy_return('%');
"<"                         yy_return('<');
">"                         yy_return('>');
"&"                         yy_return('&');
"|"                         yy_return('|');
"="                         yy_return('=');
"^"                         yy_return('^');
"!"                         yy_return('!');

"account"                   yy_return(K_ACCOUNT);
"bool"                      yy_return(K_BOOL);
"break"                     yy_return(K_BREAK);
"byte"                      yy_return(K_BYTE);
"case"                      yy_return(K_CASE);
"const"                     yy_return(K_CONST);
"continue"                  yy_return(K_CONTINUE);
"contract"                  yy_return(K_CONTRACT);
"default"                   yy_return(K_DEFAULT);
"double"                    yy_return(K_DOUBLE);
"else"                      yy_return(K_ELSE);
"false"                     yy_return(K_FALSE);
"float"                     yy_return(K_FLOAT);
"for"                       yy_return(K_FOR);
"foreach"                   yy_return(K_FOREACH);
"func"                      yy_return(K_FUNC);
"if"                        yy_return(K_IF);
"in"                        yy_return(K_IN);
"index"                     yy_return(K_INDEX);
"int"                       yy_return(K_INT);
"int16"                     yy_return(K_INT16);
"int32"                     yy_return(K_INT32);
"int64"                     yy_return(K_INT64);
"local"                     yy_return(K_LOCAL);
"map"                       yy_return(K_MAP);
"new"                       yy_return(K_NEW);
"null"                      yy_return(K_NULL);
"readonly"                  yy_return(K_READONLY);
"return"                    yy_return(K_RETURN);
"shared"                    yy_return(K_SHARED);
"string"                    yy_return(K_STRING);
"struct"                    yy_return(K_STRUCT);
"switch"                    yy_return(K_SWITCH);
"table"                     yy_return(K_TABLE);
"transfer"                  yy_return(K_TRANSFER);
"true"                      yy_return(K_TRUE);
"uint"                      yy_return(K_UINT);
"uint16"                    yy_return(K_UINT16);
"uint32"                    yy_return(K_UINT32);
"uint64"                    yy_return(K_UINT64);

"create"                    yy_return(K_CREATE);
"delete"                    yy_return(K_DELETE);
"drop"                      yy_return(K_DROP);
"insert"                    yy_return(K_INSERT);
"select"                    yy_return(K_SELECT);
"update"                    yy_return(K_UPDATE);

"#file"                     BEGIN(DIRECTIVE);
<DIRECTIVE>{ws}+
<DIRECTIVE>\"[^\"\n\r]+\"   yyextra->path = yy_literal();
<DIRECTIVE>{D}+" "{D}+ {
    lloc->last.line = atoi(strtok(yy_text(), " "));
    lloc->last.offset = atoi(strtok(NULL, " "));
}
<DIRECTIVE>{nl}             BEGIN(INITIAL);

{id} {
    lval->str = yy_text();
    yy_return(ID);
}

. {
    yy_update_first();
    yy_update_last_col();
    yy_error(ERROR_UNKNOWN_CHAR);
    yy_recover();
}

<<EOF>>                     yyterminate();

%%

int
yylex(YYSTYPE *lval, YYLTYPE *lloc, void *yyscanner)
{
    int token;
    yyparam_t *param = yyget_extra(yyscanner);

    if (param->token != 0) {
        token = param->token;
        *lloc = param->lloc;
        param->token = 0;
        return token;
    }

    return yyscan(lval, lloc, yyscanner);
}

void
yylex_set_token(void *yyscanner, int token, YYLTYPE *lloc)
{
    yyparam_t *param = yyget_extra(yyscanner);

    param->token = token;
    param->lloc = *lloc;
}

/* end of scanner.l */