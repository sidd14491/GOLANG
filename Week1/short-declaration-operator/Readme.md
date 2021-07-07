# Short-Declaration-Operator

    Terminology:
        a) Keyword:
            # These are the words that a reserved for use by a go
                  > they are sometime called reserved word
                  > you can't use keyword for anything 
        
        b) Operator
            > In "2 + 2" the "+" is the operator
            > an operator is a character that represents an action, as for example "+"
                is an arithmetic operator which perform addition

        c) Operand:
            > In "2 + 2" the "2" is an operand

        d) statement:
                In programming a statement is the smallest standalone element of a program
                that express some action to be carried out. It is an instruction that commands
                the computer to perform a specified action.
                A program is formed by sequence  of one or more statements.
    
        e) expression:
            In programming an expression is a combination of one or more explicit values,constant,variables.operator and functions that programming language interprets and 
            compute to produce anther value. For example ,2+3 is an expression which evaluates 5
    
        Code:
            https://play.golang.org/p/_e2zM-Rbe0b

# Lexical Element

        1) Comment:
            Comment serves as program documentation
            There are two types of comment declaration in golang
                a) single line comment 
                        // Hi sid
                b) Multi-line comment
                    /* Hi 
                       sid
                    */
        2) Token
            It forms the vocabulary of go language.
            There are four classes:
                a) Identifiers
                b) Keywords
                c) Operator
                d) punctuation and literals
            
        3) Identifiers:
            Identifiers name program entities such as variable and type.
            An identifier is a sequence of one or more letters and digits.
            The first character in an identifier is must be a letter.

            ** syntax: identifier = letter { letter | unicode_digit } .
            Example:
                a) a
                b) _x9
                c) ThisVariableIsExported
                d) αβ
        
        4) Keyword:
            The following keywords are reserved and not be used as identifier 
                break        default      func         interface    select
                case         defer        go           map          struct
                chan         else         goto         package      switch
                const        fallthrough  if           range        type
                continue     for          import       return       var
        
        5) Operator and punctuation
                The following character sequence represent a operator and punctuation.
                +    &     +=    &=     &&    ==    !=    (    )
                -    |     -=    |=     ||    <     <=    [    ]
                *    ^     *=    ^=     <-    >     >=    {    }
                /    <<    /=    <<=    ++    =     :=    ,    ;
                %    >>    %=    >>=    --    !     ...   .    :
                     &^          &^=
    