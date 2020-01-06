%{
#include <stdio.h>
#include <stdlib.h>

int AM_MINS = 0;
int PM_MINS = 12 * 60;
extern int yylex();
int yyerror(char *ch);
%}

%union {
    int i;
}

%token <i> NUM AM PM END_TOKEN
%type <i> time <i> hour minute apm spec

%%

spec: time END_TOKEN
    { 
        $$ = $1;
        if($$ >= 24 * 60)
        {
            yyerror("Time is too big");
        }
        printf("%d minutes since midnight\n", $$);
        exit(0);
    }
;

time: hour ':' minute { $$ = $1 * 60 + $3; }
    | hour ':' minute apm { $$ = $1 * 60 + $3 + $4; }
    | hour apm { $$ = $1 * 60 + $2; }
    ;

hour: NUM  { $$ = $1; }
    | NUM NUM { $$ = $1 * 10 + $2; if($$ > 24) yyerror("hour is too big"); }
    ;

minute: NUM NUM { $$ = $1 * 10 + $2; if($$ > 59) yyerror("minute is too big"); }
    ;

apm: AM { $$ = AM_MINS; }
    | PM { $$ = PM_MINS; }
%%

int yyerror(char *ch) {
    printf("Error %s\n", ch);
    exit(-1);
}

int main() {
    yyparse();
    return 0;
}