using System;
using System.Collections.Generic;
using System.IO;


namespace AoC
{
    public enum TokenTypes
    {
        Answer,
        EndOfGroup
    }


    public record Token (TokenTypes Type, string Value);


    public class Tokeniser
    {
        readonly string _inputPath = Path.Join(Directory.GetCurrentDirectory(), "input.txt");


        public IEnumerable<Token> GetTokens()
        {
            string inputPath = Path.Join(Directory.GetCurrentDirectory(), "input.txt");
            foreach(var line in File.ReadLines(inputPath))
            {
                var tokenTypes = string.IsNullOrWhiteSpace(line) ? TokenTypes.EndOfGroup : TokenTypes.Answer;
                yield return new Token(tokenTypes, line);
            }

            yield return new Token(TokenTypes.EndOfGroup, string.Empty);
        }
    }
}
