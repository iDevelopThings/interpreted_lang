parser grammar SimpleLangParser;


options {
	superClass = SimpleLangParserBase;
	tokenVocab = SimpleLangLexer;
    language = Go;
}

program:
(imports += importStatement)*
(
    objectDeclaration |
    funcDeclaration |
    httpRoute |
    httpServerConfig
)* EOF;

importStatement: IMPORT importPath=string SEMICOLON?;

typedIdentifier: name=identifier typeName=type;

objectDeclaration: 'object' name=ID objectBody;
objectBody: LBRACE (objectFieldDeclaration SEMICOLON?)* RBRACE;
objectFieldDeclaration: typedIdentifier;
objectFieldAssignment: name=ID COLON val=expression COMMA?;

dict: LBRACE dictFieldAssignment* RBRACE;
dictFieldKey: ID | string;
dictFieldAssignment: key=dictFieldKey COLON val=expression COMMA?;

list: LBRACE listElement* RBRACE;
listElement: val=expression COMMA?;

// We can instantiate an object(struct) with the following syntax:
// var z MyObject* = &MyObject{}
objectInstantiation: AMPERSAND? name=ID
LBRACE (fields += objectFieldAssignment)* RBRACE
;

string: DOUBLE_QUOUTE_STRING | SINGLE_QUOUTE_STRING | BACKTICK_STRING;
int: VALUE_INTEGER;
float: VALUE_FLOAT;
bool: VALUE_BOOL;
null: VALUE_NULL;
value
    : int
    | float
    | bool
    | null
    | objectInstantiation
    | ID
    | string
    | dict
    | list
    ;

     /*typeName=TYPE_I8
   | typeName=TYPE_I16
   | typeName=TYPE_I32
   | typeName=TYPE_I64
   | typeName=TYPE_INTEGER
   | typeName=VALUE_NULL
   | typeName=TYPE_BOOLEAN
   | typeName=TYPE_FLOAT
   | typeName=TYPE_STRING
   |*/
type
    : typeName=identifier (isPointer=ASTERISK)*?                #simpleTypeIdentifier
    | LBRACK RBRACK (isPointer=ASTERISK)*? typeName=identifier  #arrayTypeIdentifier
    ;

blockBody: LBRACE (statements += statement)* RBRACE;

funcDeclaration
    : FUNCTION
    (LPAREN receiver=typedIdentifier RPAREN)?
    name=ID
    arguments=argumentDeclarationList
    (returnType=type)?
    blockBody
    ;

argumentDeclarationList: LPAREN (typedIdentifier (COMMA typedIdentifier)*)? RPAREN;

variableDeclaration
    : VAR
    ( typedIdentifier (EQUALS val=expression)? SEMICOLON
    | name=ID EQUALS val=expression SEMICOLON
    );

loopStatement
    : FOR cond=expression ((AS as=identifier)? (STEP step=expression)? | (STEP step=expression)? (AS as=identifier)?) blockBody
    | FOR cond=expression (STEP step=expression)? blockBody
    | FOR cond=expression (AS as=identifier)? blockBody
    | FOR cond=expression blockBody
    | FOR blockBody
    ;

baseStatement
    : loopStatement
    | expression SEMICOLON?
    | returnStmt
    | breakStmt
    | variableDeclaration
    | ifStmt
    | deleteStmt
    ;
statement: baseStatement;
httpStatement
    : baseStatement
    | httpResponse
    ;

deleteStmt: DELETE expression SEMICOLON;

elseStmt
    : ELSE blockBody #elseBlock
    | ELSE ifStmt    #elseIfBlock
    ;
ifStmt
    : IF cond=expression blockBody (elseStmt)?
    ;


returnStmt: RETURN expression SEMICOLON;
breakStmt: BREAK SEMICOLON;

httpRoute:
    ROUTE
    method=HTTP_METHOD
    path=string
    (LPAREN injectionParameters += typedIdentifier (COMMA injectionParameters += typedIdentifier)* RPAREN)?
    body=httpRouteBody;
httpRouteBody: LBRACE
(injections += httpRouteBodyInjection*)
((statements += httpStatement)* )
RBRACE;

httpRouteBodyInjection:
FROM
HTTP_ROUTE_INJECTION_TYPE
HTTP_ROUTE_AS
typedIdentifier
HTTP_ROUTE_SEMICOLON;
httpServerConfig: HTTP_SERVER dict SEMICOLON;
httpStatus: STATUS int;

httpResponseDataType: TEXT | JSON;
httpResponseData: (
    dataType=httpResponseDataType? string |
    dataType=httpResponseDataType? expression
);
httpResponse: RESPOND WITH httpResponseData? (httpStatus)? SEMICOLON?;

// Call Expressions can:
// ex: fmt::println("Hello World")
// object -> fmt, function -> println, arguments -> "Hello World"
// ex: someFunc(1, 2, 3)
// object -> null, function -> someFunc, arguments -> 1, 2, 3

argumentList: LPAREN arguments+=expression (COMMA arguments+=expression)* RPAREN
            | LPAREN RPAREN // for no arguments
            ;

expression
    : primary
    | assignmentExpression
    ;

assignmentExpression
    : nonParenExpression                      // Non-parenthesized expressions
    | lhs=primary op=(EQUALS | PLUSEQ | MINUSEQ | ASTERISKEQ | SLASHEQ) rhs=expression  // Assignment
    ;
nonParenExpression
    : logicalOrExpressionNP
    ;
logicalOrExpressionNP
    : logicalAndExpressionNP                                     // Logical AND expressions
    | lhs=logicalOrExpressionNP op=OR rhs=expression    // Logical OR
    ;
logicalAndExpressionNP
    : equalityExpressionNP                                     // Equality expressions
    | lhs=logicalAndExpressionNP op=AND rhs=expression    // Logical AND
    ;
equalityExpressionNP
    : relationalExpressionNP                                   // Relational expressions
    | lhs=equalityExpressionNP op=(EQEQ | NE) rhs=expression  // Equality and inequality
    ;
relationalExpressionNP
    : shiftExpressionNP                                          // Additive expressions
    | lhs=relationalExpressionNP op=(LT | GT | LE | GE) rhs=expression // Comparisons
    ;
shiftExpressionNP
    : additiveExpressionNP
    | lhs=shiftExpressionNP op=(LSHIFT | RSHIFT) rhs=expression
    ;
additiveExpressionNP
    : multiplicativeExpressionNP                                   // Multiplicative expressions
    | lhs=additiveExpressionNP op=(PLUS | MINUS) rhs=expression  // Addition and subtraction
    ;
multiplicativeExpressionNP
    : powerExpressionNP                                                 // Primary expressions
    | lhs=multiplicativeExpressionNP op=(ASTERISK | SLASH) rhs=expression // Multiplication and division
    ;
powerExpressionNP
    : unaryExpressionNP                                        // Unary expressions
    | lhs=powerExpressionNP op=CARET rhs=expression     // Power
    ;
unaryExpressionNP
    : primary
    | op=(MINUS | NOT) rhs=unaryExpressionNP
    ;

postFixExpression: (identifier|value) op=(PLUSPLUS | MINUSMINUS);

primary
    : functionName=identifier argumentList                              # FunctionCall
    | value                                                             # ValuePrimary
    | postFixExpression                                                 # PostfixPrimary
    | LPAREN expression RPAREN                                          # ParenExpressionPrimary
    | primary LBRACK start=expression (isSlice=COLON (end=expression)?)? RBRACK  # ArrayPrimary
    | primary DOT functionName=identifier argumentList                  # MemberFunctionCall
    | primary COLON_COLON functionName=identifier argumentList          # StaticFunctionCall
    | primary DOT identifier                                            # MemberDotAccess
    | primary COLON_COLON identifier                                    # MemberScopedAccess
    | lhs=primary DOTDOT rhs=primary                                    # RangePrimary
    ;


identifier: ID | IDENTIFIER;
