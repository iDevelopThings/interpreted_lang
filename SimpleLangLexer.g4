lexer grammar SimpleLangLexer;

options {
	superClass = SimpleLangLexerBase;
    language = Go;
}


VALUE_NULL: 'null' | 'NULL';
VALUE_BOOL: 'true' | 'false';
VALUE_INTEGER : [0-9]+;
VALUE_FLOAT
    :   '-'? VALUE_FLOAT_INT '.' VALUE_FLOAT_INT EXP? 'f'   // 1.35, 1.35E-9, 0.3, -4.5
    |   '-'? VALUE_FLOAT_INT EXP 'f'           // 1e10 -3e4
    |   '-'? VALUE_FLOAT_INT 'f'               // -3, 45
    ;

fragment VALUE_FLOAT_INT : '0' | [1-9] [0-9]* ; // no leading zeros
fragment EXP : [Ee] [+\-]? VALUE_FLOAT_INT ;

// Keywords
OBJECT: 'object';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
FUNCTION: 'func';
RETURN: 'return';
BREAK: 'break';
VAR: 'var';
FOR: 'for';
AS: 'as';
STEP: 'step';
DELETE: 'delete';

// Type names
//TYPE_I8: 'i8';
//TYPE_I16: 'i16';
//TYPE_I32: 'i32';
//TYPE_I64: 'i64';
//TYPE_INTEGER: 'int';
//VALUE_NULL: 'null' | 'NULL';
//TYPE_BOOLEAN: 'bool';
//TYPE_FLOAT: 'float';
//TYPE_STRING: 'string';

// Brackets
LBRACE: '{';
RBRACE: '}';
LPAREN: '(';
RPAREN: ')';
LBRACK: '[';
RBRACK: ']';

// Operators
SEMICOLON: ';';
COMMA: ',';
EQUALS: '=';
PLUS: '+';
PLUSPLUS: '++';
PLUSEQ: '+=';
MINUS: '-';
MINUSMINUS: '--';
MINUSEQ: '-=';
ASTERISK: '*';
ASTERISKEQ: '*=';
SLASH: '/';
SLASHEQ: '/=';
RSHIFT: '>>';
LSHIFT: '<<';
CARET: '^';
COLON: ':';
COLON_COLON: '::';
AMPERSAND: '&';
DOT: '.';
DOTDOT: '..';


LT: '<';
LE: '<=';
GT: '>';
GE: '>=';
EQEQ: '==';
NE: '!=';
AND: '&&';
PIPE: '|';
OR: '||';
NOT: '!';

ROUTE: 'route';
RESPOND: 'respond';
WITH: 'with';
TEXT: 'text';
JSON: 'json';
STATUS: 'status';
FROM: 'from' -> mode(HTTP_ROUTE_MODE);

//INJECTION_TYPE: 'body' | 'path' | 'query';
//AS: 'as';
HTTP_METHOD: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH';
HTTP_SERVER: 'httpServer';

WS: [ \t\r\n]+ -> skip;
ID: LETTER (LETTER | UNICODE_DIGIT)*;

SINGLE_LINE_COMMENT: '//' ~[\r\n]* -> skip;
MULTI_LINE_COMMENT: '/*' .*? '*/' -> skip;

DOUBLE_QUOUTE_STRING: '"' (ESC|.)*? '"';
SINGLE_QUOUTE_STRING: '\'' (ESC|.)*? '\'';
BACKTICK_STRING:      '`' ( ~[`] )* '`';

fragment ESC :   '\\' (["\\/bfnrt] | UNICODE) ;

fragment UNICODE : 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT ;

fragment DECIMALS: [0-9] ('_'? [0-9])*;

fragment OCTAL_DIGIT: [0-7];

fragment HEX_DIGIT: [0-9a-fA-F];

fragment BIN_DIGIT: [01];

fragment EXPONENT: [eE] [+-]? DECIMALS;

fragment UNICODE_DIGIT: [\p{Nd}];

fragment LETTER: UNICODE_LETTER | '_';

fragment UNICODE_LETTER: [\p{L}];

mode HTTP_ROUTE_MODE;

HTTP_ROUTE_INJECTION_TYPE: 'body' | 'path' | 'query';
HTTP_ROUTE_AS: 'as';
IDENTIFIER: LETTER (LETTER | UNICODE_DIGIT)*;

HTTP_ROUTE_WS: [ \t\r\n]+ -> skip;


HTTP_ROUTE_SEMICOLON: ';' -> mode(DEFAULT_MODE);

