using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var (tokens, interpreter) = Bootstrap(string.Join("", args));
            interpreter.Calculate(tokens);
        }


        private static (IEnumerable<Token> tokens, Interpreter Interpreter) Bootstrap(string input)
        {
            if(File.Exists(input))
                input = File.ReadAllText(input);

            return
            (
                new Tokeniser().GetTokens(input),
                new Interpreter()
            );
        }
    }
}
