using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Tokeniser
    {
        const string _equalsValue = "=";


        public IEnumerable<Token> GetTokens(string input)
        {
            foreach(var line in PreparedInput())
            {
                if(line.Length > 1)
                {
                    foreach(var word in line.Split(" "))
                        yield return NewToken(word);

                    yield return NewToken(_equalsValue);
                }
            }


            string[] PreparedInput() =>
                input
                    .Replace("(", "( ")
                    .Replace(")", " )")
                    .Replace("\r", "")
                    .Split('\n')
            ;

            Token NewToken(string value) =>
                new Token(GetTokenType(value), value)
            ;

            TokenType GetTokenType(string value) =>
                value switch
                {
                    "("             => TokenType.LeftParentheses,
                    ")"             => TokenType.RightParentheses,
                    "+"             => TokenType.Addition,
                    "*"             => TokenType.Multiplication,
                    _equalsValue    => TokenType.Equals,
                    _               => TokenType.Integer
                }
            ;
        }
    }
}
