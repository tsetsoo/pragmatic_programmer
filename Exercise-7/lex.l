%{
#include <stdio.h>
#include <stdlib.h>
#include "bison.tab.h"
%}

%option noyywrap

%%

[0-9] {yylval.i = atoi(yytext); return NUM;}
[:] {return yytext[0];}
[ap]m {if(yytext[0] == 'a'){
    return AM;
}else {
    return PM;
}}
[\n] {return END_TOKEN;}

%%
