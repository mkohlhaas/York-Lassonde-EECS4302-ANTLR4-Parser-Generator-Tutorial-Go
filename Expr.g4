grammar Expr;

prog : (decl | expr)+ EOF          ;
decl : IDENT ':' INT_TYPE '=' NUM  ;
expr : expr '*' expr               # multiplication
     | expr '+' expr               # addition
     | IDENT                       # variable
     | NUM                         # number ;

MUL      : '*' ;
ADD      : '+' ;
IDENT    : [a-z] [a-zA-Z0-9]* ;
NUM      : '0' | '-'? [1-9] [0-9]* ;
INT_TYPE : 'INT' ;
COMMENT  : '--' ~ [\r\n]* -> skip ;
WS       : [ \t\n]+ -> skip ;
